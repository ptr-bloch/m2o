package todo

import (
	"github.com/ptr-bloch/m2o"
	"testing"
)

type mapKeyStructType struct {
	A, B int
}
type MapStructKeys struct {
	Data map[mapKeyStructType]int
}

func TestMapStructKeysDecoding(t *testing.T) {
	key1 := mapKeyStructType{
		A: 1,
		B: 2,
	}
	key2 := mapKeyStructType{
		A: 3,
		B: 4,
	}
	source := map[string]interface{}{
		"Data": map[mapKeyStructType]interface{}{
			key1: 10,
			key2: 20,
		},
	}

	decoder, err := m2o.NewDecoder(MapStructKeys{})

	var result MapStructKeys

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if len(result.Data) != 2 || result.Data[key1] != 10 || result.Data[key2] != 20 {
		t.Errorf("Map decoding failed: got %+v", result)
	}
}
