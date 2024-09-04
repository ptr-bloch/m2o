package test

import (
	"bytes"
	"net"
	"testing"

	"github.com/ptr-bloch/m2o"
)

func TestIPDecoder(t *testing.T) {
	source := "192.168.1.1"

	decode, err := m2o.NewDecoder(net.IP{}, m2o.WithIPDecoder())

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	var result net.IP
	err = decode.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	expectedIP := net.ParseIP(source)
	if expectedIP == nil {
		t.Fatalf("Error parsing IP: %v", source)
	}

	if bytes.Compare(result, expectedIP) != 0 {
		t.Fatalf("ip parsed incorrectly")
	}
}
