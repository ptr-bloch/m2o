package test

import (
	"testing"

	"github.com/ptr-bloch/m2o"
)

type AliasedFieldStruct struct {
	FirstName string `mapstructure:"name,alias=name_alias"`
	Age       int
}

func TestAliasedFields(t *testing.T) {
	source := map[string]interface{}{
		"name_alias": "John",
		"Age":        30,
	}

	var result AliasedFieldStruct
	decoder, err := m2o.NewDecoder(AliasedFieldStruct{})

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if result.FirstName != "John" || result.Age != 30 {
		t.Errorf("Aliased field handling failed: got %+v", result)
	}
}
