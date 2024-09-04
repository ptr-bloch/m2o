package test

import (
	"testing"
	"time"

	"github.com/ptr-bloch/m2o"
)

func TestCustomDecoder(t *testing.T) {
	source := "1985-04-12T23:20:50Z"

	decode, err := m2o.NewDecoder(time.Time{}, m2o.WithTimeDecoder(time.RFC3339))

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	var result time.Time
	err = decode.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if source != result.Format(time.RFC3339) {
		t.Fatalf("time.Time parsed incorrectly")
	}
}
