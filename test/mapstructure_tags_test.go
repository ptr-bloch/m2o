package test

import (
	"testing"

	"github.com/ptr-bloch/m2o"
)

type TagStruct struct {
	Name string `mapstructure:"username"`
	Age  int    `mapstructure:"user_age"`
}

func TestTagHandling(t *testing.T) {
	source := map[string]interface{}{
		"username": "John",
		"user_age": 30,
	}

	var result TagStruct
	decoder, err := m2o.NewDecoder(TagStruct{})

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if result.Name != "John" || result.Age != 30 {
		t.Errorf("Tag handling failed: got %+v", result)
	}
}

type SpecialCharStruct struct {
	Field1 string `mapstructure:"field_1"`
	Field2 int    `mapstructure:"field-2"`
}

func TestSpecialCharacterKeys(t *testing.T) {
	source := map[string]interface{}{
		"field_1": "Test",
		"field-2": 100,
	}

	var result SpecialCharStruct
	decoder, err := m2o.NewDecoder(SpecialCharStruct{})

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if result.Field1 != "Test" || result.Field2 != 100 {
		t.Errorf("Special character key handling failed: got %+v", result)
	}
}
