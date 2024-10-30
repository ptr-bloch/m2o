package test

import (
	"github.com/ptr-bloch/m2o"
	"sync"
	"testing"
	"time"
)

func TestConcurrency(t *testing.T) {
	t.Parallel()

	source := map[string]interface{}{
		"Name": "John",
		"Age":  30,
	}

	var wg sync.WaitGroup
	cycles := 100000
	c := make(chan *BasicStruct, cycles)

	profile := m2o.NewProfile()
	decoder, err := m2o.NewDecoder(BasicStruct{}, m2o.WithProfile(profile))

	if err != nil {
		t.Errorf("Error creating decoder: %v", err)
		return
	}

	wg.Add(cycles)

	for i := 0; i < cycles; i++ {
		go func() {
			defer wg.Done()

			var result BasicStruct

			err = decoder.Decode(source, &result)

			if err != nil || result.Name != "John" || result.Age != 30 {
				t.Errorf("Concurrency failed: got %+v, error: %v", result, err)
			}

			time.Sleep(time.Second)
			c <- &result
		}()
	}

	wg.Wait()

	checkMemoryUsage(t, profile, func() {
		close(c)
		for true {
			if _, ok := <-c; !ok {
				break
			}
		}
	})
}
