package test

import (
	"testing"

	"github.com/ptr-bloch/m2o"
)

type CaseSensitivityStruct struct {
	Field string
}

func TestCaseSensitivity(t *testing.T) {
	t.Parallel()

	source := map[string]interface{}{
		"Field": "Test",
		"field": "Ignored",
	}

	var result CaseSensitivityStruct
	profile := m2o.NewProfile()
	decoder, err := m2o.NewDecoder(CaseSensitivityStruct{}, m2o.WithProfile(profile))

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if result.Field != "Test" {
		t.Errorf("Case sensitivity handling failed: got %+v", result)
	}
}
