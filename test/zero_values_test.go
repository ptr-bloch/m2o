package test

import (
	"testing"

	"github.com/ptr-bloch/m2o"
)

type ZeroValueStruct struct {
	Name string
	Age  int
}

func TestZeroValues(t *testing.T) {
	source := map[string]interface{}{
		"Age": 0,
	}

	var result ZeroValueStruct
	result.Age = 25

	decoder, err := m2o.NewDecoder(ZeroValueStruct{})

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if result.Age != 0 {
		t.Errorf("Zero value handling failed: got %+v", result)
	}
}
