package test

import (
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/ptr-bloch/m2o"
)

type MapStruct struct {
	Prop1 string
	Data  map[string]int
	Prop2 string
}

func TestMapDecoding(t *testing.T) {
	source := map[string]interface{}{
		"Data": map[string]interface{}{
			"key1": 10,
			"key2": 20,
		},
		"Prop1": "abc",
		"Prop2": "def",
	}

	var result MapStruct
	decoder, err := m2o.NewDecoder(MapStruct{})

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if len(result.Data) != 2 || result.Data["key1"] != 10 || result.Data["key2"] != 20 {
		t.Errorf("Map decoding failed: got %+v", result)
	}

	if result.Prop1 != "abc" || result.Prop2 != "def" {
		t.Errorf("possibly map decoding corrupted neighbour memory")
	}
}

type mapKeyType string
type MapTypedKeys struct {
	Data map[mapKeyType]int
}

func TestMapTypedKeysDecoding(t *testing.T) {
	source := map[string]interface{}{
		"Data": map[string]interface{}{
			"key1": 10,
			"key2": 20,
		},
	}

	var result MapTypedKeys
	decoder, err := m2o.NewDecoder(MapTypedKeys{})

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if len(result.Data) != 2 || result.Data["key1"] != 10 || result.Data["key2"] != 20 {
		t.Errorf("Map decoding failed: got %+v", result)
	}
}

func TestEmptyBilletMapResetsToNil(t *testing.T) {
	decoder, err := m2o.NewDecoder(struct {
		Map map[string]int
	}{
		Map: map[string]int{},
	}, m2o.WithZeroOnEmpty())

	if decoder == nil {
		t.Fatalf("decoder creation error: %v", err)
	}

	p, err := decoder.Produce(map[string]interface{}{})

	if p == nil {
		t.Fatalf("decoding error: %v", err)
	}

	if p.Map != nil {
		t.Fatalf("map property should be set to nil")
	}
}

func TestNonEmptyBilletMapResetsToNil(t *testing.T) {
	decoder, err := m2o.NewDecoder(struct {
		Map map[string]int
	}{
		Map: map[string]int{
			"key": 1,
		},
	}, m2o.WithZeroOnEmpty())

	if decoder == nil {
		t.Fatalf("decoder creation error: %v", err)
	}

	p, err := decoder.Produce(map[string]interface{}{})

	if p == nil {
		t.Fatalf("decoding error: %v", err)
	}

	if p.Map != nil {
		t.Fatalf("map property should be set to nil")
	}
}

func TestMapConcurrentDecoding(t *testing.T) {
	source := map[string]interface{}{
		"Data": map[string]interface{}{
			"key": map[string]interface{}{
				"Data": map[string]int{
					"key": 10,
				},
			},
		},
	}

	decoder, err := m2o.NewDecoder(
		struct {
			Data map[string]interface{}
		}{
			Data: map[string]interface{}{
				"key": &struct {
					Data map[string]int
				}{
					Data: map[string]int{
						"key": 11,
					},
				},
			},
		},
		m2o.WithInitializeFromBillet(),
		m2o.WithZeroOnEmpty(),
	)

	runtime.GC()
	time.Sleep(time.Millisecond * 200)

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	var wg sync.WaitGroup
	c := make(chan struct{})
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			<-c
			defer wg.Done()
			o, _ := decoder.Produce(source)
			_ = o
		}()
	}
	close(c)
	wg.Wait()
}
