package test

import (
	"testing"

	"github.com/ptr-bloch/m2o"
)

type Interface interface {
	GetName() string
}

type StructWithInterfaceProperty struct {
	Property Interface
}

type interfaceObject struct {
	Name string
}

func (s interfaceObject) GetName() string {
	return s.Name
}

func TestInterfacePropertyDecoding(t *testing.T) {
	source := map[string]interface{}{
		"Property": map[string]interface{}{
			"Name": "John",
		},
	}

	decoder, err := m2o.NewDecoder(StructWithInterfaceProperty{
		Property: interfaceObject{},
	})

	var result StructWithInterfaceProperty

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if result.Property == nil || result.Property.GetName() != "John" {
		t.Errorf("Decoding failed: got %+v", result)
	}
}

func TestInterfacePropertyNotReallocatedWhenOnStack(t *testing.T) {
	source := map[string]interface{}{
		"Property": map[string]interface{}{
			"Name": "John",
		},
	}

	decoder, err := m2o.NewDecoder(StructWithInterfaceProperty{
		Property: interfaceObject{},
	})

	var concreteObjectOnStack = interfaceObject{}
	var result = StructWithInterfaceProperty{
		Property: concreteObjectOnStack,
	}

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if result.Property == nil {
		t.Errorf("Decoding failed: got %+v", result)
	}

	// library MUST allocate new object and MUST NOT use concreteObjectOnStack because it is located on stack
	if concreteObjectOnStack.Name != "" {
		t.Fatal("allocated new object instead of using old one")
	}
}

func TestInterfacePropertyReallocatedWhenOnHeap(t *testing.T) {
	source := map[string]interface{}{
		"Property": map[string]interface{}{
			"Name": "John",
		},
	}

	decoder, err := m2o.NewDecoder(StructWithInterfaceProperty{
		Property: interfaceObject{},
	})

	var concreteObjectOnHeap = new(interfaceObject)
	var result = StructWithInterfaceProperty{
		Property: concreteObjectOnHeap,
	}

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if result.Property == nil {
		t.Errorf("Decoding failed: got %+v", result)
	}

	// library MUST NOT allocate new object and MUST use concreteObjectOnHeap
	// if it allocates new one, concreteObjectOnHeap.Name remains empty
	if concreteObjectOnHeap.Name != "John" {
		t.Fatal("allocated new object instead of using old one")
	}
}
