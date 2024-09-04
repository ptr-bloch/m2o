package test

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/ptr-bloch/m2o"
)

type InterfaceForMap interface {
	GetName() string
}

type interfaceObject1 struct {
	Name string
}

func (s interfaceObject1) GetName() string {
	return s.Name + "1"
}

type interfaceObject2 struct {
	Name string
}

func (s interfaceObject2) GetName() string {
	return s.Name + "2"
}

func TestMapWithInterfaceValuesDecoding(t *testing.T) {
	source := map[string]interface{}{
		"key1": map[string]interface{}{
			"Name": "John",
		},
		"key2": map[string]interface{}{
			"Name": "John",
		},
	}

	profile := &m2o.Profile{}

	decode, err := m2o.NewDecoder(map[string]InterfaceForMap{
		"key1": interfaceObject1{},
		"key2": interfaceObject2{},
	}, m2o.WithProfile(profile))

	var result map[string]InterfaceForMap

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decode.Decode(source, &result)

	runtime.GC()
	time.Sleep(time.Millisecond * 100)

	freed := profile.GetMemoryFreed()
	if uint64(46) != freed {
		msg := fmt.Sprintf("only two map values and two interface values should be freed, instead: %v", profile.GetMemoryFrees())
		t.Fatalf(msg)
	}

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if result == nil ||
		result["key1"] == nil ||
		result["key2"] == nil ||
		result["key1"].GetName() != "John1" ||
		result["key2"].GetName() != "John2" {
		t.Errorf("Decoding failed: got %+v", result)
	}

	runtime.GC()
	time.Sleep(time.Millisecond * 100)

	alloc := profile.GetMemoryAllocated()
	freed = profile.GetMemoryFreed()
	if alloc != freed {
		t.Fatal("all the memory allocated should be freed")
	}
}
