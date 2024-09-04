package test

import (
	"testing"

	"github.com/ptr-bloch/m2o"
)

type ErrorStruct struct {
	Age int
}

func TestErrorCase(t *testing.T) {
	source := map[string]interface{}{
		"Age": "invalid_int",
	}

	var result ErrorStruct
	decoder, err := m2o.NewDecoder(ErrorStruct{})

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(source, &result)

	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}
