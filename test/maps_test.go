package test

import (
	"testing"

	"github.com/ptr-bloch/m2o"
)

type MapStruct struct {
	Prop1 string
	Data  map[string]int
	Prop2 string
}

func TestMapDecoding(t *testing.T) {
	source := map[string]interface{}{
		"Data": map[string]interface{}{
			"key1": 10,
			"key2": 20,
		},
		"Prop1": "abc",
		"Prop2": "def",
	}

	var result MapStruct
	decoder, err := m2o.NewDecoder(MapStruct{})

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if len(result.Data) != 2 || result.Data["key1"] != 10 || result.Data["key2"] != 20 {
		t.Errorf("Map decoding failed: got %+v", result)
	}

	if result.Prop1 != "abc" || result.Prop2 != "def" {
		t.Errorf("possibly map decoding corrupted neighbour memory")
	}
}

type mapKeyType string
type MapTypedKeys struct {
	Data map[mapKeyType]int
}

func TestMapTypedKeysDecoding(t *testing.T) {
	source := map[string]interface{}{
		"Data": map[string]interface{}{
			"key1": 10,
			"key2": 20,
		},
	}

	var result MapTypedKeys
	decoder, err := m2o.NewDecoder(MapTypedKeys{})

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if len(result.Data) != 2 || result.Data["key1"] != 10 || result.Data["key2"] != 20 {
		t.Errorf("Map decoding failed: got %+v", result)
	}
}
