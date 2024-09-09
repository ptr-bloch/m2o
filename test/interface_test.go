package test

import (
	"github.com/ptr-bloch/m2o"
	"testing"
)

type StringWithSlice struct {
	String string
	Images []string
}

func (o StringWithSlice) GetName() string {
	return o.String
}

func TestInterfaceValuesDecoding(t *testing.T) {
	var elBillet StringWithSlice
	var elAnyInterfBillet any = &elBillet

	decoder2, err := m2o.NewDecoder(elAnyInterfBillet)
	if err != nil {
		t.Fatalf("Error creating	 decoder: %v", err)
	}

	elInterf, err := decoder2.Produce(elMap)
	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	checkElementInInterface(t, *elInterf)

	var elTypedInterfBillet InterfaceForMap = &elBillet

	decoder3, err := m2o.NewDecoder(elTypedInterfBillet)
	if err != nil {
		t.Fatalf("Error creating	 decoder: %v", err)
	}
	elTypedInterf, err := decoder3.Produce(elMap)
	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	checkElementInInterface(t, *elTypedInterf)

	var result map[string]InterfaceForMap

	decoder4, err := m2o.NewDecoder(map[string]InterfaceForMap{
		"key2": elBillet,
		"key3": &elBillet,
	})

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder4.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if el, ok := result["key2"].(StringWithSlice); !ok {
		t.Fatal("key2 is not decoded")
	} else {
		checkElement(t, &el)
	}

	if elInterf, ok := result["key3"]; !ok {
		t.Fatalf("element not found in map")
	} else {
		checkElementInInterface(t, elInterf)
	}
}

func checkElement(t *testing.T, el *StringWithSlice) {
	if el.String != loremIpsum {
		t.Fatalf("simple decoding to element is not working")
	}
}

func checkElementInInterface(t *testing.T, elInterf any) {
	if el, ok := elInterf.(*StringWithSlice); !ok {
		t.Fatalf("decoding to interface is not working")
	} else {
		checkElement(t, el)
	}
}
