package test

import (
	"github.com/ptr-bloch/m2o"
	"reflect"
	"testing"
)

func TestDecodingToEmptyInterfaceAsMap(t *testing.T) {
	var example any
	decoder, err := m2o.NewDecoder(example)
	if err != nil {
		t.Fatalf("error creating decoder")
	}

	var target any
	var input = map[string]interface{}{
		"Field": 100,
	}

	err = decoder.Decode(input, &target)

	if err != nil {
		t.Fatalf("error decoding source data: %s", err.Error())
	}

	if !reflect.DeepEqual(input, target) {
		t.Fatal("decoded incorrectly")
	}
}

type BasicStruct4Fields struct {
	Name      *string
	NilString *string
	Age       int
	Score     float64
}

// convert from float to int and vice versa
func TestDecodingStruct4FieldsToEmptyInterface(t *testing.T) {
	name := "John"
	var example any = BasicStruct4Fields{
		Name:  &name,
		Age:   30,
		Score: 100,
	}
	decoder, err := m2o.NewDecoder(example)
	if err != nil {
		t.Fatalf("error creating decoder: %s", err.Error())
	}

	var target any
	var input = map[string]interface{}{
		"Name":  "Bill",
		"Age":   40,
		"Score": 120.0,
	}

	err = decoder.Decode(input, &target)

	if err != nil {
		t.Fatalf("error decoding source data: %s", err.Error())
	}

	if target == nil {
		t.Fatal("target is nil after decoding")
	}

	if target, ok := target.(BasicStruct4Fields); !ok {
		t.Fatal("target should contain instance of BasicStruct")
	} else {
		if *target.Name != "Bill" ||
			target.Age != 40 ||
			target.Score != 120 {
			t.Fatalf("fields are decoded incorrectly")
		}
	}
}

type BasicStruct1Field struct {
	Name *string
}

// it's important to test this type of structure separately as golang assigns this type of structure with its own way
func TestDecodingStructToEmptyInterface(t *testing.T) {
	name := "John"
	var example any = BasicStruct1Field{
		Name: &name,
	}
	decoder, err := m2o.NewDecoder(example)
	if err != nil {
		t.Fatalf("error creating decoder: %s", err.Error())
	}

	var target any
	var input = map[string]interface{}{
		"Name": "Bill",
	}

	err = decoder.Decode(input, &target)

	if err != nil {
		t.Fatalf("error decoding source data: %s", err.Error())
	}

	if target == nil {
		t.Fatal("target is nil after decoding")
	}

	if target, ok := target.(BasicStruct1Field); !ok {
		t.Fatal("target should contain instance of BasicStruct")
	} else {
		if *target.Name != "Bill" {
			t.Fatalf("fields are decoded incorrectly")
		}
	}
}
