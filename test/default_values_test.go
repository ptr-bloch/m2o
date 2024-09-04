package test

import (
	"testing"

	"github.com/ptr-bloch/m2o"
)

type DefaultStruct struct {
	Name  string
	Age   int
	Score float64
}

func TestDefaultValues(t *testing.T) {
	source := map[string]interface{}{
		"Name": "John",
	}

	var result DefaultStruct
	decoder, err := m2o.NewDecoder(DefaultStruct{})

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	result.Age = 25 // Set a default value

	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if result.Name != "John" || result.Age != 25 || result.Score != 0.0 {
		t.Errorf("Default value handling failed: got %+v", result)
	}
}
