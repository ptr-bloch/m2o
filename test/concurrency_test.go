package test

import (
	"github.com/ptr-bloch/m2o"
	"sync"
	"testing"
)

type ComplexStruct struct {
	IntField int
}

func TestConcurrency(t *testing.T) {
	source := map[string]interface{}{
		"Name": "John",
		"Age":  30,
	}

	var wg sync.WaitGroup
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			var result BasicStruct
			decoder, err := m2o.NewDecoder(BasicStruct{})

			if err != nil {
				t.Errorf("Error creating decoder: %v", err)
				return
			}

			err = decoder.Decode(source, &result)

			if err != nil || result.Name != "John" || result.Age != 30 {
				t.Errorf("Concurrency failed: got %+v, error: %v", result, err)
			}
		}()
	}

	wg.Wait()
}
