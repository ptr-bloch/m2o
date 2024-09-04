package test

import (
	"reflect"
	"testing"

	"github.com/ptr-bloch/m2o"
	"github.com/ptr-bloch/m2o/test/objects"
)

func TestInitializeFromBillet(t *testing.T) {
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
