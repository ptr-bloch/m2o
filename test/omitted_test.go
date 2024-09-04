package test

import (
	"testing"

	"github.com/ptr-bloch/m2o"
)

type OmittedFieldStruct struct {
	Name string
	Age  int
}

func TestOmittedFields(t *testing.T) {
	source := map[string]interface{}{
		"Age":          30,
		"UnknownField": "value",
	}

	var result OmittedFieldStruct
	decoder, err := m2o.NewDecoder(OmittedFieldStruct{})

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if result.Age != 30 || result.Name != "" {
		t.Errorf("Omitted field handling failed: got %+v", result)
	}
}
