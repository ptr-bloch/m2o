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

import (
	"sync/atomic"
)

const logSize = 1024

type memoryBlock struct {
	size uintptr
	free bool
}

type profile struct {
	memoryAllocated        uint64
	memoryAllocations      [logSize]string
	memoryAllocationsIndex int32

	memoryFreed      uint64
	memoryFrees      [logSize]string
	memoryFreesIndex int32

	blocks map[uintptr]memoryBlock
}

func NewProfile() *profile {
	return &profile{
		blocks: make(map[uintptr]memoryBlock),
	}
}

func (p *profile) GetMemoryAllocated() uint64 {
	return atomic.LoadUint64(&p.memoryAllocated)
}

func (p *profile) GetMemoryAllocations() []string {
	index := atomic.LoadInt32(&p.memoryAllocationsIndex)
	return p.memoryAllocations[:min(logSize-1, index)]
}

func (p *profile) addMemoryAllocated(ptr, size uintptr, purpose string) {
	p.blocks[ptr] = memoryBlock{
		size: size,
	}
	atomic.AddUint64(&p.memoryAllocated, uint64(size))
	index := atomic.AddInt32(&p.memoryAllocationsIndex, 1)
	p.memoryAllocations[(index-1)%logSize] = purpose
}

func (p *profile) GetMemoryFreed() uint64 {
	return atomic.LoadUint64(&p.memoryFreed)
}

func (p *profile) GetMemoryFrees() []string {
	index := atomic.LoadInt32(&p.memoryFreesIndex)
	return p.memoryFrees[:min(logSize-1, index)]
}

func (p *profile) addMemoryFreed(ptr, size uintptr, purpose string) {
	if mem, ok := p.blocks[ptr]; ok {
		mem.free = true
		p.blocks[ptr] = mem
	} else {
		panic("freeing not allocated memory?")
	}

	atomic.AddUint64(&p.memoryFreed, uint64(size))
	index := atomic.AddInt32(&p.memoryFreesIndex, 1)
	p.memoryFrees[(index-1)%logSize] = purpose
}

func (p *profile) HasFreedBlocks() bool {
	for _, block := range p.blocks {
		if block.free {
			return true
		}
	}
	return false
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}
