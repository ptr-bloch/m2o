package test

import (
	"testing"

	"github.com/ptr-bloch/m2o"
)

func TestEmptyMapDecoding(t *testing.T) {
	var result BasicStruct
	decoder, err := m2o.NewDecoder(BasicStruct{})

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(map[string]interface{}{}, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if result.Name != "" || result.Age != 0 || result.Score != 0.0 {
		t.Errorf("Empty map handling failed: got %+v", result)
	}
}
