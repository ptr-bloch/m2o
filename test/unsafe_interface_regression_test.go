package test

import (
	"reflect"
	"runtime"
	"testing"

	"github.com/ptr-bloch/m2o"
)

// These tests are baseline contract tests for unsafe-heavy decoder paths.
// They intentionally cover behavior that is already expected to pass before
// adding targeted red regression tests for the unsafe interface corruption bug.

type unsafeBaselineNamed interface {
	GetName() string
}

type unsafeBaselineInterfaceHolder struct {
	Before   string
	Property unsafeBaselineNamed
	After    string
}

type unsafeBaselineMultiFieldValue struct {
	Name string
	Code int
}

func (v unsafeBaselineMultiFieldValue) GetName() string {
	return v.Name
}

type unsafeBaselineEmptyInterfaceHolder struct {
	Value any
}

type unsafeBaselineAnyStruct struct {
	Name  *string
	Age   int
	Score float64
}

type unsafeBaselineAnySingleScalarFieldStruct struct {
	Name string
}

type unsafeBaselineMapInterfaceHolder struct {
	Values map[string]unsafeBaselineNamed
}

type unsafeBaselineMapValueOne struct {
	Name string
}

func (v unsafeBaselineMapValueOne) GetName() string {
	return "one:" + v.Name
}

type unsafeBaselineMapValueTwo struct {
	Name string
}

func (v unsafeBaselineMapValueTwo) GetName() string {
	return "two:" + v.Name
}

type unsafeBaselineNestedPayload struct {
	Name  string
	Items []int
	Meta  map[string]int
}

type unsafeBaselineNestedHolder struct {
	Before string
	Value  unsafeBaselineNestedPayload
	After  string
}

func TestUnsafeBaselineTypedInterfaceMultiFieldValueDecoding(t *testing.T) {
	decoder, err := m2o.NewDecoder(unsafeBaselineInterfaceHolder{
		Property: unsafeBaselineMultiFieldValue{},
	})
	if err != nil {
		t.Fatalf("create decoder: %v", err)
	}

	var result unsafeBaselineInterfaceHolder
	if err := decoder.Decode(map[string]interface{}{
		"Before": "left",
		"Property": map[string]interface{}{
			"Name": "John",
			"Code": 7,
		},
		"After": "right",
	}, &result); err != nil {
		t.Fatalf("decode: %v", err)
	}

	if result.Before != "left" || result.After != "right" {
		t.Fatalf("neighbor fields changed: got %#v", result)
	}
	if result.Property == nil {
		t.Fatal("decoded interface property is nil")
	}
	if got := result.Property.GetName(); got != "John" {
		t.Fatalf("decoded interface property name: got %q, want %q; value=%#v", got, "John", result.Property)
	}

	decoded, ok := result.Property.(unsafeBaselineMultiFieldValue)
	if !ok {
		t.Fatalf("decoded interface concrete type: got %T, want %T", result.Property, unsafeBaselineMultiFieldValue{})
	}
	if decoded.Code != 7 {
		t.Fatalf("decoded interface payload code: got %d, want %d", decoded.Code, 7)
	}
}

func TestUnsafeBaselineTypedInterfaceStackValueIsNotMutated(t *testing.T) {
	decoder, err := m2o.NewDecoder(unsafeBaselineInterfaceHolder{
		Property: unsafeBaselineMultiFieldValue{},
	})
	if err != nil {
		t.Fatalf("create decoder: %v", err)
	}

	stackValue := unsafeBaselineMultiFieldValue{Name: "original", Code: 1}
	result := unsafeBaselineInterfaceHolder{Property: stackValue}
	if err := decoder.Decode(map[string]interface{}{
		"Property": map[string]interface{}{
			"Name": "decoded",
			"Code": 2,
		},
	}, &result); err != nil {
		t.Fatalf("decode: %v", err)
	}

	if stackValue.Name != "original" || stackValue.Code != 1 {
		t.Fatalf("decoder mutated original stack value: got %#v", stackValue)
	}
	if result.Property == nil || result.Property.GetName() != "decoded" {
		t.Fatalf("decoded interface result mismatch: got %#v", result.Property)
	}
}

func TestUnsafeBaselineEmptyInterfacePassThroughMap(t *testing.T) {
	var billet any
	decoder, err := m2o.NewDecoder(billet)
	if err != nil {
		t.Fatalf("create decoder: %v", err)
	}

	input := map[string]interface{}{
		"Field": 100,
		"Nested": map[string]interface{}{
			"Enabled": true,
		},
	}
	var result any
	if err := decoder.Decode(input, &result); err != nil {
		t.Fatalf("decode: %v", err)
	}

	if !reflect.DeepEqual(input, result) {
		t.Fatalf("decoded empty interface map mismatch: got %#v, want %#v", result, input)
	}
}

func TestUnsafeBaselineEmptyInterfaceConcreteStructDecoding(t *testing.T) {
	name := "initial"
	decoder, err := m2o.NewDecoder(any(unsafeBaselineAnyStruct{
		Name:  &name,
		Age:   1,
		Score: 2,
	}))
	if err != nil {
		t.Fatalf("create decoder: %v", err)
	}

	var result any
	if err := decoder.Decode(map[string]interface{}{
		"Name":  "Bill",
		"Age":   40,
		"Score": 120.0,
	}, &result); err != nil {
		t.Fatalf("decode: %v", err)
	}

	decoded, ok := result.(unsafeBaselineAnyStruct)
	if !ok {
		t.Fatalf("decoded result type: got %T, want %T", result, unsafeBaselineAnyStruct{})
	}
	if decoded.Name == nil || *decoded.Name != "Bill" || decoded.Age != 40 || decoded.Score != 120 {
		t.Fatalf("decoded result mismatch: got %#v", decoded)
	}
}

func TestUnsafeBaselineEmptyInterfaceSingleScalarFieldStructDecoding(t *testing.T) {
	decoder, err := m2o.NewDecoder(any(unsafeBaselineAnySingleScalarFieldStruct{}))
	if err != nil {
		t.Fatalf("create decoder: %v", err)
	}

	var result any
	if err := decoder.Decode(map[string]interface{}{
		"Name": "Bill",
	}, &result); err != nil {
		t.Fatalf("decode: %v", err)
	}

	decoded, ok := result.(unsafeBaselineAnySingleScalarFieldStruct)
	if !ok {
		t.Fatalf("decoded result type: got %T, want %T", result, unsafeBaselineAnySingleScalarFieldStruct{})
	}
	if decoded.Name != "Bill" {
		t.Fatalf("decoded name: got %q, want %q", decoded.Name, "Bill")
	}
}

func TestUnsafeBaselineMapWithInterfaceValuesKeepsPerKeyConcreteType(t *testing.T) {
	decoder, err := m2o.NewDecoder(map[string]unsafeBaselineNamed{
		"one": unsafeBaselineMapValueOne{},
		"two": unsafeBaselineMapValueTwo{},
	})
	if err != nil {
		t.Fatalf("create decoder: %v", err)
	}

	var result map[string]unsafeBaselineNamed
	if err := decoder.Decode(map[string]interface{}{
		"one": map[string]interface{}{"Name": "John"},
		"two": map[string]interface{}{"Name": "Jane"},
	}, &result); err != nil {
		t.Fatalf("decode: %v", err)
	}

	if len(result) != 2 {
		t.Fatalf("decoded map size: got %d, want %d; value=%#v", len(result), 2, result)
	}
	if got := result["one"].GetName(); got != "one:John" {
		t.Fatalf("decoded key one: got %q, want %q", got, "one:John")
	}
	if got := result["two"].GetName(); got != "two:Jane" {
		t.Fatalf("decoded key two: got %q, want %q", got, "two:Jane")
	}
	if _, ok := result["one"].(unsafeBaselineMapValueOne); !ok {
		t.Fatalf("decoded key one concrete type: got %T", result["one"])
	}
	if _, ok := result["two"].(unsafeBaselineMapValueTwo); !ok {
		t.Fatalf("decoded key two concrete type: got %T", result["two"])
	}
}

func TestUnsafeBaselineRepeatedDecodeDoesNotShareTargetState(t *testing.T) {
	decoder, err := m2o.NewDecoder(unsafeBaselineNestedHolder{})
	if err != nil {
		t.Fatalf("create decoder: %v", err)
	}

	var first unsafeBaselineNestedHolder
	if err := decoder.Decode(map[string]interface{}{
		"Before": "first-before",
		"Value": map[string]interface{}{
			"Name":  "first",
			"Items": []interface{}{1, 2},
			"Meta":  map[string]interface{}{"score": 10},
		},
		"After": "first-after",
	}, &first); err != nil {
		t.Fatalf("decode first: %v", err)
	}

	var second unsafeBaselineNestedHolder
	if err := decoder.Decode(map[string]interface{}{
		"Before": "second-before",
		"Value": map[string]interface{}{
			"Name":  "second",
			"Items": []interface{}{3, 4},
			"Meta":  map[string]interface{}{"score": 20},
		},
		"After": "second-after",
	}, &second); err != nil {
		t.Fatalf("decode second: %v", err)
	}

	first.Value.Items[0] = 100
	first.Value.Meta["score"] = 1000

	if first.Before != "first-before" || first.After != "first-after" || first.Value.Name != "first" {
		t.Fatalf("first target mismatch: got %#v", first)
	}
	if second.Before != "second-before" || second.After != "second-after" || second.Value.Name != "second" {
		t.Fatalf("second target mismatch: got %#v", second)
	}
	if second.Value.Items[0] != 3 || second.Value.Meta["score"] != 20 {
		t.Fatalf("second target shares mutable state with first: first=%#v second=%#v", first, second)
	}
}

func TestUnsafeBaselineDecodedValuesSurviveGC(t *testing.T) {
	decoder, err := m2o.NewDecoder(unsafeBaselineNestedHolder{})
	if err != nil {
		t.Fatalf("create decoder: %v", err)
	}

	var result unsafeBaselineNestedHolder
	if err := decoder.Decode(map[string]interface{}{
		"Before": "left",
		"Value": map[string]interface{}{
			"Name":  "stable",
			"Items": []interface{}{1, 2, 3},
			"Meta":  map[string]interface{}{"score": 99},
		},
		"After": "right",
	}, &result); err != nil {
		t.Fatalf("decode: %v", err)
	}

	for i := 0; i < 5; i++ {
		runtime.GC()
		runtime.Gosched()
	}

	if result.Before != "left" || result.After != "right" {
		t.Fatalf("neighbor fields changed after GC: got %#v", result)
	}
	if result.Value.Name != "stable" || len(result.Value.Items) != 3 || result.Value.Items[2] != 3 || result.Value.Meta["score"] != 99 {
		t.Fatalf("decoded values changed after GC: got %#v", result)
	}
}
