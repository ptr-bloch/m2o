package test

import (
	"github.com/ptr-bloch/m2o"
	"strings"
	"testing"
)

type StructWithBigStrings struct {
	Strings []string
}

func TestBigString(t *testing.T) {
	t.Parallel()

	var str1 = strings.Repeat("A", 1024*1024)
	var str2 = strings.Repeat("B", 1024*1024)
	var input = map[string]interface{}{
		"Strings": []interface{}{
			str1,
			str2,
		},
	}

	profile := m2o.NewProfile()
	decoder, err := m2o.NewDecoder(StructWithBigStrings{}, m2o.WithProfile(profile))

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	var result StructWithBigStrings
	err = decoder.Decode(input, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if len(result.Strings) != 2 || result.Strings[0] != str1 || result.Strings[1] != str2 {
		t.Errorf("Decoding failed: should be two strings in slice")
	}

	if result.Strings[0] != str1 || result.Strings[1] != str2 {
		t.Errorf("Decoding failed: strings decoded incorrectly")
	}

	checkMemoryUsage(t, profile, &result)
}
