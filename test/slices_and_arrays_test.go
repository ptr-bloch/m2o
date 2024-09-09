package test

import (
	"testing"

	"github.com/ptr-bloch/m2o"
)

type SliceStruct struct {
	Names      []string
	Numbers    []int
	Str        string
	EmptySlice []bool
}

func TestSliceDecoding(t *testing.T) {
	source := map[string]interface{}{
		"Names":   []interface{}{"Alice", "Bob"},
		"Numbers": []interface{}{1, 2, 3},
		"Str":     "some string",
	}

	var result SliceStruct
	decoder, err := m2o.NewDecoder(SliceStruct{})

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if len(result.Names) != 2 || result.Names[0] != "Alice" || result.Names[1] != "Bob" {
		t.Errorf("Slice decoding failed: got %+v", result)
	}

	if len(result.Numbers) != 3 || result.Numbers[0] != 1 || result.Numbers[2] != 3 {
		t.Errorf("Slice decoding failed: got %+v", result)
	}
}
