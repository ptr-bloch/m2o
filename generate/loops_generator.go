package main

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
	"os"
	"text/template"

	"github.com/ptr-bloch/m2o/build"
)

func rangeExclude(start, end int) []int {
	r := make([]int, end-start)
	for i := range r {
		r[i] = start + i
	}
	return r
}

func rangeInclude(start, end int) []int {
	r := make([]int, end-start+1)
	for i := range r {
		r[i] = start + i
	}
	return r
}

func div(a, b int) int {
	return a / b
}

func sub(a, b int) int {
	return a - b
}

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

const loopTemplate = `package m2o

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

// generated automatically by generate/loops_generator.go
// DO NOT EDIT
// instead edit the template

import (
	"fmt"

	// reflection is used only for decoding slices with element size bigger than {{ .MaxElementSize }} bytes
	// for example:
	// arrays: [{{ .MaxElementSize }}]byte
	// structures with fields with total size greater than {{ .MaxElementSize }}: struct { f1 byte ... f{{ .MaxElementSize }} byte }
	// slices as element don't have this problem because it's size is Sizeof(SliceHeader) 
	"reflect"
	"unsafe"
)

type up = uintptr

func unrolledLoopOptional(f []fieldDef) d {
	switch len(f) {
	case 0:
		return nil
	{{- range $index := rangeInclude 1 .LoopSize }}
	case {{$index}}:
		return _ul{{$index}}(f)
	{{- end }}
	}
	half := len(f) / 2
	part1 := unrolledLoopOptional(f[:half])
	part2 := unrolledLoopOptional(f[half:])

	return func(fromAny any, structAddress up, isOmitted bool) {
		part1(fromAny, structAddress, isOmitted)
		part2(fromAny, structAddress, isOmitted)
	}
}

type m = map[string]interface{}
{{- range $count := rangeInclude 1 .LoopSize }}
func _ul{{$count}}(f []fieldDef) d {
	var b = *(*[{{$count}}]fieldDef)(unsafe.Pointer(&f[0]))
	return func(s any, a up, isOmitted bool) {
		if from, ok := s.(m); ok {
			var v any
			var ok bool
			{{- range $index := rangeExclude 0 $count }}
			if v, ok = from[b[{{$index}}].n]; true { b[{{$index}}].u(v, a, !ok) }
			{{- end }}
			return
		}
		slowMap(b[:], s, a, true)
	}
}
{{- end }}

func unrolledLoopRequired(f []fieldDef) d {
	switch len(f) {
	case 0:
		return nil
	{{- range $index := rangeInclude 1 .LoopSize }}
	case {{$index}}:
		return _ulm{{$index}}(f)
	{{- end }}
	}
	half := len(f) / 2
	part1 := unrolledLoopRequired(f[:half])
	part2 := unrolledLoopRequired(f[half:])

	return func(fromAny any, structAddress up, isOmitted bool) {
		part1(fromAny, structAddress, isOmitted)
		part2(fromAny, structAddress, isOmitted)
	}
}

type d = decoderFunc
func _u[T map[K]V, K comparable, V any](m T, k K) V {
	if v, ok := m[k]; !ok {
		panic(fmt.Errorf("required field %v omitted", k))
	} else {
		return v
	}
}

func slowMap(f []fieldDef, source any, a up, required bool) {
	if source == nil {
		for i := 0; i < len(f); i++ {
			field := f[i]
			// if source == nil isOmitted should be false
			field.u(nil, a, false)
		}
		return
	}

	val := reflect.ValueOf(source)
	typ := val.Type()
	kind := typ.Kind()
	if kind != reflect.Map {
		panic(fmt.Errorf("expected map, got %s", kind.String()))
	}

	keyKind := typ.Key().Kind()
	if keyKind != reflect.String {
		panic(fmt.Errorf("expected map with string keys, got key %s", keyKind.String()))
	}

	for i := 0; i < len(f); i++ {
		field := f[i]
		key := reflect.ValueOf(f[i].n)
		mapVal := val.MapIndex(key)
		if !mapVal.IsValid() {
			if required {
				panic(fmt.Errorf("required field %v omitted", f[i].n))
			}
			field.u(nil, a, true)
			continue
		}
		field.u(mapVal.Interface(), a, false)
	}
}

{{- range $count := rangeInclude 1 .LoopSize }}
func _ulm{{$count}}(f []fieldDef) d {
	var b = *(*[{{$count}}]fieldDef)(unsafe.Pointer(&f[0]))
	return func(s any, a up, isOmitted bool) {
		if from, ok := s.(m); ok {
			{{- range $index := rangeExclude 0 $count }}
			b[{{$index}}].u(_u(from, b[{{$index}}].n), a, false)
			{{- end }}
			return
		}
		slowMap(b[:], s, a, true)
	}
}
{{- end }}

type b = byte

func getMapToolkitByKeyTypeAndValueReflection[K comparable](elementType reflect.Type) (r *mapToolkit) {
	elementSize := elementType.Size()
	switch elementSize {
	{{- range $index := rangeInclude 1 .MaxElementSize }}
	case {{ $index }}: r = _m[K, [{{ $index }}]b]()
	{{- end }}
	default:
	}

	return
}

func getMapToolkitForStructKeyAndValueByReflection(mapType reflect.Type) (r *mapToolkit) {
	keySize := mapType.Key().Size()
	e := mapType.Elem() 
	switch keySize {
	{{- range $index := rangeInclude 1 .MaxElementSize }}
	case {{ $index }}: r = _o[[{{ $index }}]b](e)
	{{- end }}
	default:
	}

	return
}

func _s[T any](where uintptr, capacity int) uintptr {
	var target = (*[]T)(unsafe.Pointer(where))
	*target = make([]T, capacity)
	if capacity == 0 {
		return 0
	}
	return uintptr(unsafe.Pointer(&(*target)[0]))
}

func getSliceMaker(t reflect.Type) (r func(where uintptr, capacity int) uintptr) {
	// for most cases we have fast solution - just use make function and hope that compiler will inline it
	// this solution relies on the fact that slice header doesn't contain any type specific information
	// and everything that matters is slice element size, so we can replace
	// make([]int, 0) with make([]struct{_ int}, 0)
	// this solution degrades when element size exceeds predefined sizes
	// to speed of reflect.MakeSlice
	switch t.Elem().Size() {
	{{- range $index := rangeInclude 1 .MaxElementSize }}
	case {{ $index }}:
		r = _s[[{{ $index }}]b]
	{{- end }}
	default:
		// degrade to reflect after {{ .MaxElementSize }} bytes element size
		r = func(where uintptr, capacity int) uintptr {
			e := reflect.MakeSlice(t, capacity, capacity)
			iface := e.Interface()
			interfaceHeader := (*[2]uintptr)((unsafe.Pointer)(&iface))
	
			newSliceHeaderPtr := interfaceHeader[1]
			const headerSize = unsafe.Sizeof([]int{})
			*(*[headerSize]byte)((unsafe.Pointer)(where)) = *(*[headerSize]byte)((unsafe.Pointer)(newSliceHeaderPtr))
			return e.Index(0).UnsafeAddr()
		}
	}
	return 
}`

/* ul
var b = *(*[{{$count}}]fieldDef)(unsafe.Pointer((&f[0])))
	return func(s any, a up) {
		from := s.(m)

		{{- range $index := rangeExclude 0 $count }}
		if v, ok := from[b[{{$index}}].n]; ok { b[{{$index}}].u(v, a) }
		{{- end }}
	}
*/

/* ulm
{{- range $index := rangeInclude 0 (div $count 4 ) }}
	var {{- range $offset := rangeInclude 0 (sub (min 4 (sub $count (mul $index 4))) 1) }} n{{add (mul $index 4) $offset}}, u{{add (mul $index 4) $offset}}, {{- end }} _ =
	{{- range $offset := rangeInclude 0 (sub (min 4 (sub $count (mul $index 4))) 1) }} f[{{add (mul $index 4) $offset}}].n, f[{{add (mul $index 4) $offset}}].u, {{- end }} 0
	{{- end }}
	return func(s any, a up) {
		from := s.(m)
		{{- range $index := rangeInclude 0 (div $count 4 ) }}
		var {{- range $offset := rangeInclude 0 (sub (min 4 (sub $count (mul $index 4))) 1) }} v{{add (mul $index 4) $offset}}, {{- end }} _ :=
		{{- range $offset := rangeInclude 0 (sub (min 4 (sub $count (mul $index 4))) 1) }} _u(from, n{{add (mul $index 4) $offset}}), {{- end }} 0
		{{- end }}

		{{- range $index := rangeExclude 0 $count }}
		u{{$index}}(v{{$index}}, a)
		{{- end }}
	}

*/

type loopTemplateData struct {
	LoopSize       int
	MaxElementSize int
}

func main() {
	fileName := fmt.Sprintf("generated_unrolled_loop.go")
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	count := build.MaxOptimizedElementSize
	data := loopTemplateData{
		LoopSize:       10,
		MaxElementSize: count,
	}

	tmpl, err := template.New("unrolledLoop").Funcs(template.FuncMap{
		"rangeExclude": rangeExclude,
		"rangeInclude": rangeInclude,
		"div":          div,
		"sub":          sub,
		"add":          add,
		"mul":          mul,
		"min":          min,
	}).Parse(loopTemplate)

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(file, data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Generated %s\n", fileName)
}
