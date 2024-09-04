package test

import (
	"testing"

	"github.com/ptr-bloch/m2o"
)

type StructWithMandatoryField struct {
	Required string `map2object:",required"`
	Optional string
}

// TestRequiredDecoding checks:
// 1. by default fields are not required
// 2. requiring field by tag works
func TestRequiredDecoding(t *testing.T) {
	source := map[string]interface{}{
		"Optional": "optional",
	}

	var result StructWithMandatoryField
	decode, err := m2o.NewDecoder(StructWithMandatoryField{})

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decode.Decode(source, &result)

	if err == nil {
		t.Fatal("should return error when trying skip mandatory field")
	}
}

// TestAllRequiredDecoding checks that requiring fields by default works
func TestAllRequiredDecoding(t *testing.T) {
	source := map[string]interface{}{
		"Required": "required",
	}

	var result StructWithMandatoryField
	decode, err := m2o.NewDecoder(StructWithMandatoryField{}, m2o.WithAllFieldsRequired())

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decode.Decode(source, &result)

	if err == nil {
		t.Fatalf("should return error when trying skip mandatory field")
	}
}
