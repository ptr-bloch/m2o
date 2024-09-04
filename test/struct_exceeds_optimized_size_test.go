package test

import (
	"testing"

	"github.com/ptr-bloch/m2o"
	"github.com/ptr-bloch/m2o/build"
)

type NonOptimizedBySizeType struct {
	_     [build.MaxOptimizedElementSize]byte
	Field string
}

func TestStructExceedsMaxOptimizedSize(t *testing.T) {
	source := map[string]interface{}{
		"Field": "Test",
	}

	var result NonOptimizedBySizeType
	decoder, err := m2o.NewDecoder(NonOptimizedBySizeType{}, m2o.WithSquashAnonymous())

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if result.Field != "Test" {
		t.Fatalf("Field after optimized size decoded incorrectly")
	}
}
