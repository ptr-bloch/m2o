package test

import (
	"runtime"
	"testing"
	"time"
)

func checkMemoryUsage(t *testing.T, profiler interface{ HasFreedBlocks() bool }, object any) {
	runtime.GC()
	time.Sleep(200 * time.Millisecond)
	if profiler.HasFreedBlocks() {
		t.Fatal("has premature freed memory")
	}
	runtime.KeepAlive(object)
}
