package test

import (
	"reflect"
	"testing"

	"github.com/ptr-bloch/m2o"
	"github.com/ptr-bloch/m2o/test/objects"
)

func TestInitializeFromBillet(t *testing.T) {
	t.Parallel()

	source := map[string]interface{}{}

	var result objects.StructWithPrivateField
	billet := objects.NewStructWithPrivateField()
	decode, err := m2o.NewDecoder(billet, m2o.WithInitializeFromBillet())

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decode.Decode(source, &result)
	if !reflect.DeepEqual(billet, result) {
		t.Fatal("private property was not copied")
	}
}

func TestPartialDecoding(t *testing.T) {
	type innerStruct struct {
		A, B int
	}

	decoder, err := m2o.NewDecoder(innerStruct{
		A: 3,
		B: 4,
	}, m2o.WithInitializeFromBillet())

	if decoder == nil {
		t.Fatalf("decoder creation error: %v", err)
	}

	p, err := decoder.Produce(map[string]interface{}{
		"B": 5,
	})

	if p == nil {
		t.Fatalf("decoding error: %v", err)
	}

	if p.A != 3 {
		t.Fatalf("key2.A should be 3")
	}

	if p.B != 5 {
		t.Fatalf("key2.B should be 5")
	}
}

func TestBilletIsNotChangedWithInitializeFromBilletFlag(t *testing.T) {
	type innerStruct struct {
		A int
	}

	type outerStruct struct {
		Inner *innerStruct
	}

	billet := outerStruct{
		Inner: &innerStruct{A: 100},
	}

	decoder, err := m2o.NewDecoder(billet, m2o.WithInitializeFromBillet())

	if decoder == nil {
		t.Fatalf("decoder creation error: %v", err)
	}

	p, _ := decoder.Produce(map[string]interface{}{
		"Inner": map[string]interface{}{
			"A": 150,
		},
	})
	_ = p

	if billet.Inner.A != 100 {
		t.Fatalf("billet should not be changed")
	}
}
