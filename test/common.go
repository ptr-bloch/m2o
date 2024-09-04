package test

import (
	"github.com/ptr-bloch/m2o"
	"runtime"
	"testing"
	"time"
)

func checkMemory(t *testing.T, obj any, profile *m2o.Profile) {
	runtime.GC()
	time.Sleep(time.Millisecond * 200)

	if profile.GetMemoryFreed() > 0 {
		t.Fatalf("freed some memory: %v", profile.GetMemoryFrees())
	}

	runtime.KeepAlive(obj)

	runtime.GC()
	time.Sleep(time.Millisecond * 200)

	if profile.GetMemoryFreed() != profile.GetMemoryAllocated() {
		t.Fatalf("not all memory freed")
	}
}
