package test

import (
	"testing"

	"github.com/ptr-bloch/m2o"
)

type OneTypeFieldsStruct struct {
	A, B, C, D, E, F, G int
}

func TestTypedMap(t *testing.T) {
	source := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
		"G": 7,
	}

	decoder, err := m2o.NewDecoder(OneTypeFieldsStruct{})

	if err != nil {
		t.Fatalf("could not create decoder: %s", err.Error())
	}

	produced, err := decoder.Produce(source)

	if err != nil {
		t.Fatalf("could not decode: %s", err.Error())
	}

	expected := OneTypeFieldsStruct{
		A: 1,
		B: 2,
		C: 3,
		D: 4,
		E: 5,
		F: 6,
		G: 7,
	}
	if *produced != expected {
		t.Fatalf("wrong decoding")
	}

}
