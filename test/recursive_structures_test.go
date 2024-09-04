package test

import (
	"testing"

	"github.com/ptr-bloch/m2o"
)

type RecursiveStruct struct {
	Name     string
	Children []RecursiveStruct
}

func TestRecursiveStructDecoding(t *testing.T) {
	source := map[string]interface{}{
		"Name": "Parent",
		"Children": []interface{}{
			map[string]interface{}{
				"Name": "Child1",
			},
			map[string]interface{}{
				"Name": "Child2",
				"Children": []interface{}{
					map[string]interface{}{
						"Name": "Child3",
					},
				},
			},
		},
	}

	var result RecursiveStruct
	decoder, err := m2o.NewDecoder(RecursiveStruct{})

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if result.Name != "Parent" || len(result.Children) != 2 {
		t.Errorf("Recursive struct decoding failed: got %+v", result)
	}
}
