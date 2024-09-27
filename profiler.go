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
	"fmt"
	"reflect"
	"sync/atomic"
	"unsafe"
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

func (p *profile) addMemoryAllocated(ptr unsafe.Pointer, size uintptr, purpose string) {
	p.blocks[uintptr(ptr)] = memoryBlock{
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

func (p *profile) addMemoryFreed(ptr unsafe.Pointer, size uintptr, purpose string) {
	if mem, ok := p.blocks[uintptr(ptr)]; ok {
		mem.free = true
		p.blocks[uintptr(ptr)] = mem
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

// CheckObjectPtr checks if the memory blocks for the given object's fields are allocated and not free.
func (p *profile) CheckObjectPtr(object any) error {
	val := reflect.ValueOf(object)
	if val.Kind() != reflect.Pointer {
		panic("should call with pointer")
	}
	val = val.Elem()
	return p.checkRecursive(val)
}

// checkRecursive checks the memory of each field recursively.
func (p *profile) checkRecursive(v reflect.Value) error {
	switch v.Kind() {
	case reflect.Ptr:
		// If it's a pointer, dereference it
		if v.IsNil() {
			return nil // Skip nil pointers
		}
		ptr := unsafe.Pointer(v.Pointer())
		address := uintptr(ptr)

		// Check if the pointer's address is in blocks and if it's free
		if err := p.checkAddress(address); err != nil {
			return err
		}

		// Recursively check the value the pointer points to
		return p.checkRecursive(v.Elem())

	case reflect.Interface:
		// If it's an interface, get the underlying value and check it
		if v.IsNil() {
			return nil // Skip nil interfaces
		}
		return p.checkRecursive(v.Elem())

	case reflect.Struct:
		structAddress := v.UnsafeAddr()
		if err := p.checkAddress(structAddress); err != nil {
			return err
		}
		// If it's a struct, iterate over its fields
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)

			// Recursively check each field
			if err := p.checkRecursive(field); err != nil {
				return err
			}
		}

	case reflect.Array, reflect.Slice:
		// If it's an array or slice, iterate over the elements
		for i := 0; i < v.Len(); i++ {
			elem := v.Index(i)

			// Recursively check each element
			if err := p.checkRecursive(elem); err != nil {
				return err
			}
		}
	case reflect.Map:
		// If it's a map, iterate over the keys and values
		for _, key := range v.MapKeys() {
			// Recursively check the key
			if err := p.checkRecursive(key); err != nil {
				return err
			}

			// Recursively check the value
			val := v.MapIndex(key)
			if err := p.checkRecursive(val); err != nil {
				return err
			}
		}
	default:
		// For all other types, get the address and check it
		if v.CanAddr() {
			ptr := unsafe.Pointer(v.UnsafeAddr())
			address := uintptr(ptr)

			// Check the address in the blocks map
			if err := p.checkAddress(address); err != nil {
				return err
			}
		}
	}
	return nil
}

// checkAddress checks if the address exists in blocks and if the block is not free.
func (p *profile) checkAddress(address uintptr) error {
	block, found := p.blocks[address]
	if !found {
		return nil
		//return fmt.Errorf("address %x not found in profile blocks", address)
	}
	if block.free {
		return fmt.Errorf("memory block at address %x is marked as free", address)
	}
	return nil
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}
