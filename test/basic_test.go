package test

import (
	"testing"

	"github.com/ptr-bloch/m2o"
)

type BasicStruct struct {
	Name  string
	Age   int
	Score float64
}

func TestBasicDecoding(t *testing.T) {
	source := map[string]interface{}{
		"Name":  "John",
		"Age":   30,
		"Score": 85.5,
	}

	profile := m2o.NewProfile()
	decoder, err := m2o.NewDecoder(BasicStruct{}, m2o.WithProfile(profile))

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	var result BasicStruct
	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if result.Name != "John" || result.Age != 30 || result.Score != 85.5 {
		t.Errorf("Decoding failed: got %+v", result)
	}

	checkMemoryUsage(t, profile, &result)
}

func TestBasicProducing(t *testing.T) {
	source := map[string]interface{}{
		"Name":  "John",
		"Age":   30,
		"Score": 85.5,
	}

	decoder, err := m2o.NewDecoder(BasicStruct{})

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	resultPtr, err := decoder.Produce(source)

	if resultPtr == nil {
		t.Fatalf("Error decoding: %v", err)
	}

	result := *resultPtr

	if result.Name != "John" || result.Age != 30 || result.Score != 85.5 {
		t.Errorf("Decoding failed: got %+v", result)
	}
}

func TestIncorrectTarget(t *testing.T) {
	decoder, err := m2o.NewDecoder(struct {
	}{})

	if decoder == nil {
		t.Fatalf("error creating decoder: %v", err)
	}

	err = decoder.Decode(map[string]interface{}{}, nil)
	if err == nil {
		t.Fatal("decoder should return error with incorrect target")
	}
}
