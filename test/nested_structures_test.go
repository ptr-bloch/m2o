package test

import (
	"testing"

	"github.com/ptr-bloch/m2o"
)

type NestedStruct struct {
	Inner NestedInnerStruct
}

type NestedInnerStruct struct {
	Value int
}

func TestNestedStructDecoding(t *testing.T) {
	source := map[string]interface{}{
		"Inner": map[string]interface{}{
			"Value": 42,
		},
	}

	var result NestedStruct
	decoder, err := m2o.NewDecoder(NestedStruct{})

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if result.Inner.Value != 42 {
		t.Errorf("Nested struct decoding failed: got %+v", result)
	}
}
