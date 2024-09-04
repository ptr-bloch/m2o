package test

import (
	"testing"
	"time"

	"github.com/ptr-bloch/m2o"
)

func TestDurationDecoder(t *testing.T) {
	source := "5m30s"

	decode, err := m2o.NewDecoder(time.Duration(0), m2o.WithDurationDecoder())

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	var result time.Duration
	err = decode.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	expectedDuration, err := time.ParseDuration(source)
	if err != nil {
		t.Fatalf("Error parsing duration: %v", err)
	}

	if result != expectedDuration {
		t.Fatalf("wrong result")
	}
}
