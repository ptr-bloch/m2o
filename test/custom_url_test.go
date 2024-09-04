package test

import (
	"net/url"
	"testing"

	"github.com/ptr-bloch/m2o"
)

func TestURLDecoder(t *testing.T) {
	source := "https://www.example.com/path?query=123"

	decode, err := m2o.NewDecoder(url.URL{}, m2o.WithURLDecoder())

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	var result url.URL
	err = decode.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	expectedURL, err := url.Parse(source)
	if err != nil {
		t.Fatalf("Error parsing URL: %v", err)
	}

	if result != *expectedURL {
		t.Fatal("url decoded incorrectly")
	}
}
