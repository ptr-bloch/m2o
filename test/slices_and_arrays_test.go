package test

import (
	"sync"
	"testing"

	"github.com/ptr-bloch/m2o"
)

type SliceStruct struct {
	Names      []string
	Numbers    []int
	Str        string
	EmptySlice []bool
}

func TestSliceDecoding(t *testing.T) {
	t.Parallel()

	source := map[string]interface{}{
		"Names":   []interface{}{"Alice", "Bob"},
		"Numbers": []interface{}{1, 2, 3},
		"Str":     "some string",
	}

	decoder, err := m2o.NewDecoder(SliceStruct{})

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	var result SliceStruct
	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if len(result.Names) != 2 || result.Names[0] != "Alice" || result.Names[1] != "Bob" {
		t.Errorf("Slice decoding failed: got %+v", result)
	}

	if len(result.Numbers) != 3 || result.Numbers[0] != 1 || result.Numbers[2] != 3 {
		t.Errorf("Slice decoding failed: got %+v", result)
	}
}

func TestSliceWithInitializeDecoding(t *testing.T) {
	t.Parallel()

	source := map[string]interface{}{
		"Names":   []interface{}{"Alice", "Bob"},
		"Numbers": []interface{}{1, 2, 3},
		"Str":     "some string",
	}

	var billet = SliceStruct{
		Names: []string{"a", "b"},
	}

	decoder, err := m2o.NewDecoder(billet, m2o.WithInitializeFromBillet())

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	var result SliceStruct
	err = decoder.Decode(source, &result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if len(result.Names) != 2 || result.Names[0] != "Alice" || result.Names[1] != "Bob" {
		t.Errorf("Slice decoding failed: got %+v", result)
	}

	if len(result.Numbers) != 3 || result.Numbers[0] != 1 || result.Numbers[2] != 3 {
		t.Errorf("Slice decoding failed: got %+v", result)
	}
}

func TestHasIncorrectRefAfterSliceArray(t *testing.T) {
	type sliceHolder struct {
		Slice []struct {
			A, B, C, D uint32
		}
	}

	var source = map[string]interface{}{
		"Slice": []map[string]interface{}{{
			"A": 1,
			"B": 2,
			"C": 3,
			"D": 4,
		}},
	}

	var start = make(chan struct{})
	var cycles = 100000
	var wg sync.WaitGroup
	decoder, err := m2o.NewDecoder(sliceHolder{})

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	wg.Add(cycles)

	for i := 0; i < cycles; i++ {
		go func() {
			defer wg.Done()
			<-start

			_ = decoder.Decode(source, &sliceHolder{})
		}()
	}

	close(start)
	wg.Wait()
}
