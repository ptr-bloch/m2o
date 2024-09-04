package todo

import (
	"testing"

	"github.com/ptr-bloch/m2o"
)

type UnusedSourceStruct struct {
	Name  string
	Age   int
	Score float64
}

func TestUnusedSourceDecoding(t *testing.T) {
	source := map[string]interface{}{}

	var result UnusedSourceStruct
	decode, err := m2o.NewDecoder(UnusedSourceStruct{}, m2o.WithMustUseAllSource())

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decode.Decode(source, &result)
	if err == nil {
		t.Fatalf("should return error when trying to skip mandatory field")
	}
}
