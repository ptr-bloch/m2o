package test

import (
	"testing"

	"github.com/ptr-bloch/m2o"
)

type UnexportedFieldStruct struct {
	Name string
	age  int // unexported
}

func TestUnexportedFields(t *testing.T) {
	source := map[string]interface{}{
		"Name": "John",
		"age":  30,
	}

	var result UnexportedFieldStruct
	decoder, err := m2o.NewDecoder(UnexportedFieldStruct{})

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if result.Name != "John" || result.age != 0 {
		t.Errorf("Unexported field handling failed: got %+v", result)
	}
}
