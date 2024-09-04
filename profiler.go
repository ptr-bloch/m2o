package m2o

/*
MIT License

Copyright (c) 2024

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

import "sync/atomic"

const logSize = 1024

type Profile struct {
	memoryAllocated        uint64
	memoryAllocations      [logSize]string
	memoryAllocationsIndex int32

	memoryFreed      uint64
	memoryFrees      [logSize]string
	memoryFreesIndex int32
}

func (p *Profile) GetMemoryAllocated() uint64 {
	return atomic.LoadUint64(&p.memoryAllocated)
}

func (p *Profile) GetMemoryAllocations() []string {
	index := atomic.LoadInt32(&p.memoryAllocationsIndex)
	return p.memoryAllocations[:min(logSize-1, index)]
}

func (p *Profile) addMemoryAllocated(mem uint64, purpose string) {
	atomic.AddUint64(&p.memoryAllocated, mem)
	index := atomic.AddInt32(&p.memoryAllocationsIndex, 1)
	p.memoryAllocations[(index-1)%logSize] = purpose
}

func (p *Profile) GetMemoryFreed() uint64 {
	return atomic.LoadUint64(&p.memoryFreed)
}

func (p *Profile) GetMemoryFrees() []string {
	index := atomic.LoadInt32(&p.memoryFreesIndex)
	return p.memoryFrees[:min(logSize-1, index)]
}

func (p *Profile) addMemoryFreed(mem uint64, purpose string) {
	atomic.AddUint64(&p.memoryFreed, mem)
	index := atomic.AddInt32(&p.memoryFreesIndex, 1)
	p.memoryFrees[(index-1)%logSize] = purpose
}

func min(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}
