package test

import (
	"testing"

	"github.com/ptr-bloch/m2o"
)

// These tests intentionally describe currently broken behavior.
// They must be red before the decoder.go unsafe interface fix and green after it.

type unsafeBugNamed interface {
	GetName() string
}

type unsafeBugInterfaceHolder struct {
	Property unsafeBugNamed
}

type unsafeBugSinglePointerField struct {
	Name *string
}

func (v unsafeBugSinglePointerField) GetName() string {
	if v.Name == nil {
		return ""
	}
	return *v.Name
}

type unsafeBugPointerConcrete struct {
	Name string
}

func (v *unsafeBugPointerConcrete) GetName() string {
	if v == nil {
		return ""
	}
	return v.Name
}

type unsafeBugFourFieldStruct struct {
	Name    *string
	NilName *string
	Age     int
	Score   float64
}

func TestUnsafeBugTypedInterfaceSinglePointerFieldDecoding(t *testing.T) {
	source := map[string]interface{}{
		"Property": map[string]interface{}{
			"Name": "John",
		},
	}

	decoder, err := m2o.NewDecoder(unsafeBugInterfaceHolder{
		Property: unsafeBugSinglePointerField{},
	})
	if err != nil {
		t.Fatalf("create decoder: %v", err)
	}

	var result unsafeBugInterfaceHolder
	if err := decoder.Decode(source, &result); err != nil {
		t.Fatalf("decode: %v", err)
	}

	if result.Property == nil {
		t.Fatal("decoded interface property is nil")
	}

	if got := result.Property.GetName(); got != "John" {
		t.Fatalf("decoded interface property has corrupted data: got %q, want %q; value=%#v", got, "John", result.Property)
	}
}

func TestUnsafeBugTypedInterfacePointerConcreteDecoding(t *testing.T) {
	source := map[string]interface{}{
		"Property": map[string]interface{}{
			"Name": "John",
		},
	}

	decoder, err := m2o.NewDecoder(unsafeBugInterfaceHolder{
		Property: &unsafeBugPointerConcrete{},
	})
	if err != nil {
		t.Fatalf("create decoder: %v", err)
	}

	var result unsafeBugInterfaceHolder
	if err := decoder.Decode(source, &result); err != nil {
		t.Fatalf("decode: %v", err)
	}

	if result.Property == nil {
		t.Fatal("decoded interface property is nil")
	}

	if got := result.Property.GetName(); got != "John" {
		t.Fatalf("decoded pointer concrete interface property has corrupted data: got %q, want %q; value=%#v", got, "John", result.Property)
	}
}

func TestUnsafeBugEmptyInterfaceStructDecodeCheckptr(t *testing.T) {
	decoder, err := m2o.NewDecoder(any(unsafeBugFourFieldStruct{}))
	if err != nil {
		t.Fatalf("create decoder: %v", err)
	}

	var result any
	if err := decoder.Decode(map[string]interface{}{
		"Name":  "John",
		"Age":   30,
		"Score": 100.0,
	}, &result); err != nil {
		t.Fatalf("decode: %v", err)
	}

	decoded, ok := result.(unsafeBugFourFieldStruct)
	if !ok {
		t.Fatalf("decoded result type: got %T, want %T", result, unsafeBugFourFieldStruct{})
	}

	if decoded.Name == nil || *decoded.Name != "John" || decoded.Age != 30 || decoded.Score != 100 {
		t.Fatalf("decoded result mismatch: got %#v", decoded)
	}
}
