package test

import (
	"testing"

	"github.com/ptr-bloch/m2o"
)

// TestZeroScalarFieldOnEmpty - if int field is reset to 0 after decoding
// TestZeroPointerFieldOnEmpty - if pointer field is reset to nil
// TestZeroEmptyInterfaceFieldOnEmpty - if interface{} field is reset to nil
// TestZeroTypedInterfaceFieldOnEmpty - if interface{ Method() } field is reset to nil
// TestClearMapOnEmptyMap - if structured map is empty on empty input
// TestDeleteMapKeyOnEmptyKey - if structured map one key deleted and one remains correctly decoded

type StructureToCheckZero[T any] struct {
	NotEmpty1 uint64
	Field     T
	NotEmpty2 uint64
}

func TestZeroScalarFieldOnEmpty(t *testing.T) {
	notEmpty := uint64(0xffffffffffffffff)
	source := map[string]interface{}{
		"NotEmpty1": notEmpty,
		"NotEmpty2": notEmpty,
	}

	decode, err := m2o.NewDecoder(StructureToCheckZero[int]{}, m2o.WithZeroOnEmpty())

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	var result = StructureToCheckZero[int]{
		Field: 10, // should be reset to nil
	}

	err = decode.Decode(source, &result)
	if err != nil {
		t.Fatalf("error decoding: %s", err.Error())
	}

	if result.Field != 0 {
		t.Fatalf("properties should be emptied")
	}

	if result.NotEmpty1 != notEmpty || result.NotEmpty2 != notEmpty {
		t.Fatalf("not all fields should be emptied")
	}
}

func TestZeroPointerFieldOnEmpty(t *testing.T) {
	notEmpty := uint64(0xffffffffffffffff)
	source := map[string]interface{}{
		"NotEmpty1": notEmpty,
		"NotEmpty2": notEmpty,
	}

	decode, err := m2o.NewDecoder(StructureToCheckZero[*int]{}, m2o.WithZeroOnEmpty())

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	var result = StructureToCheckZero[*int]{
		Field: new(int), // should be reset to nil
	}

	err = decode.Decode(source, &result)
	if err != nil {
		t.Fatalf("error decoding: %s", err.Error())
	}

	if result.Field != nil {
		t.Fatalf("properties should be emptied")
	}

	if result.NotEmpty1 != notEmpty || result.NotEmpty2 != notEmpty {
		t.Fatalf("not all fields should be emptied")
	}
}

func TestZeroEmptyInterfaceFieldOnEmpty(t *testing.T) {
	notEmpty := uint64(0xffffffffffffffff)
	source := map[string]interface{}{
		"NotEmpty1": notEmpty,
		"NotEmpty2": notEmpty,
	}

	decode, err := m2o.NewDecoder(StructureToCheckZero[any]{
		Field: 1,
	}, m2o.WithZeroOnEmpty())

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	var result = StructureToCheckZero[any]{
		Field: 1, // should be reset to nil
	}

	err = decode.Decode(source, &result)
	if err != nil {
		t.Fatalf("error decoding: %s", err.Error())
	}

	if result.Field != nil {
		t.Fatalf("properties should be emptied")
	}

	if result.NotEmpty1 != notEmpty || result.NotEmpty2 != notEmpty {
		t.Fatalf("not all fields should be emptied")
	}
}

type IntWithMethod int

func (IntWithMethod) Method() {}

func TestZeroTypedInterfaceFieldOnEmpty(t *testing.T) {
	notEmpty := uint64(0xffffffffffffffff)
	source := map[string]interface{}{
		"NotEmpty1": notEmpty,
		"NotEmpty2": notEmpty,
	}

	decode, err := m2o.NewDecoder(StructureToCheckZero[interface{ Method() }]{
		Field: IntWithMethod(1),
	}, m2o.WithZeroOnEmpty())

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	var result = StructureToCheckZero[interface{ Method() }]{
		Field: IntWithMethod(1), // should be reset to nil
	}

	err = decode.Decode(source, &result)
	if err != nil {
		t.Fatalf("error decoding: %s", err.Error())
	}

	if result.Field != nil {
		t.Fatalf("properties should be emptied")
	}

	if result.NotEmpty1 != notEmpty || result.NotEmpty2 != notEmpty {
		t.Fatalf("not all fields should be emptied")
	}
}

func TestClearMapOnEmptyMap(t *testing.T) {
	source := map[string]interface{}{
		"Prop1": "abc",
		"Prop2": "def",
	}

	decoder, err := m2o.NewDecoder(source, m2o.WithZeroOnEmpty())

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	var result = map[string]interface{}{
		"Prop1": "123",
		"Prop2": "456",
	}

	err = decoder.Decode(nil, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if len(result) > 0 {
		t.Fatal("map should be empty")
	}
}

func TestDeleteMapKeyOnEmptyKey(t *testing.T) {
	source := map[string]interface{}{
		"Prop1": "abc",
		"Prop2": "def",
	}

	decoder, err := m2o.NewDecoder(source, m2o.WithZeroOnEmpty())

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	var result = map[string]interface{}{
		"Prop1": "123",
		"Prop2": "456",
	}

	err = decoder.Decode(map[string]interface{}{
		"Prop2": "xwz",
	}, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if _, ok := result["Prop1"]; ok {
		t.Fatal("key Prop1 should be unset")
	}

	if Prop2, ok := result["Prop2"]; !ok || Prop2 != "xwz" {
		t.Fatalf("key Prop2 should be set to 'xwz'")
	}
}
