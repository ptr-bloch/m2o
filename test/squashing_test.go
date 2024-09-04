package test

import (
	"testing"

	"github.com/ptr-bloch/m2o"
)

type EmbeddedStruct struct {
	InnerStruct
	OuterName string
}

type SquashedStruct struct {
	StructToSquash InnerStruct `mapstructure:",squash"`
	OuterName      string
}

type InnerStruct struct {
	InnerField int
}

func TestEmbeddedSquashedStruct(t *testing.T) {
	source := map[string]interface{}{
		"InnerField": 42,
		"OuterName":  "Test",
	}

	var result EmbeddedStruct
	decoder, err := m2o.NewDecoder(EmbeddedStruct{}, m2o.WithSquashAnonymous())

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if result.InnerField != 42 || result.OuterName != "Test" {
		t.Errorf("Embedded struct decoding failed: got %+v", result)
	}
}

func TestEmbeddedNotSquashedImplicitlyStruct(t *testing.T) {
	source := map[string]interface{}{
		"InnerField": 42,
		"OuterName":  "Test",
	}

	var result EmbeddedStruct
	decoder, err := m2o.NewDecoder(EmbeddedStruct{})

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	// by default squashing is disabled
	if result.InnerField == 42 || result.OuterName != "Test" {
		t.Errorf("Embedded struct decoding failed: got %+v", result)
	}
}

func TestSquashedStruct(t *testing.T) {
	source := map[string]interface{}{
		"InnerField": 42,
		"OuterName":  "Test",
	}

	var result SquashedStruct
	decoder, err := m2o.NewDecoder(SquashedStruct{})

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if result.StructToSquash.InnerField != 42 {
		t.Errorf("Embedded struct decoding failed: got %+v", result)
	}
}
