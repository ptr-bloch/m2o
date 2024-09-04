package test

import (
	"math/big"
	"testing"

	"github.com/ptr-bloch/m2o"
)

func TestBigIntDecoder(t *testing.T) {
	source := "123456789012345678901234567890"

	decode, err := m2o.NewDecoder(*big.NewInt(0), m2o.WithBigIntDecoder())

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	var result big.Int
	err = decode.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	expectedBigInt := new(big.Int)
	expectedBigInt.SetString(source, 10)

	if result.String() != expectedBigInt.String() {
		t.Fatal("big int decoded incorrectly")
	}
}
