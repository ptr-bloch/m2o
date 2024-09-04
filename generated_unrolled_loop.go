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

// generated automatically by generate/loops_generator.go
// DO NOT EDIT
// instead edit the template

import (
	"fmt"

	// reflection is used only for decoding slices with element size bigger than 64 bytes
	// for example:
	// arrays: [64]byte
	// structures with fields with total size greater than 64: struct { f1 byte ... f64 byte }
	// slices as element don't have this problem because it's size is Sizeof(SliceHeader) 
	"reflect"
	"unsafe"
)

type up = uintptr

func unrolledLoopOptional(f []fieldDef) d {
	switch len(f) {
	case 0:
		return nil
	case 1:
		return _ul1(f)
	case 2:
		return _ul2(f)
	case 3:
		return _ul3(f)
	case 4:
		return _ul4(f)
	case 5:
		return _ul5(f)
	case 6:
		return _ul6(f)
	case 7:
		return _ul7(f)
	case 8:
		return _ul8(f)
	case 9:
		return _ul9(f)
	case 10:
		return _ul10(f)
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
func _ul1(f []fieldDef) d {
	var b = *(*[1]fieldDef)(unsafe.Pointer(&f[0]))
	return func(s any, a up, isOmitted bool) {
		if from, ok := s.(m); ok {
			var v any
			var ok bool
			if v, ok = from[b[0].n]; true { b[0].u(v, a, !ok) }
			return
		}
		slowMap(b[:], s, a, true)
	}
}
func _ul2(f []fieldDef) d {
	var b = *(*[2]fieldDef)(unsafe.Pointer(&f[0]))
	return func(s any, a up, isOmitted bool) {
		if from, ok := s.(m); ok {
			var v any
			var ok bool
			if v, ok = from[b[0].n]; true { b[0].u(v, a, !ok) }
			if v, ok = from[b[1].n]; true { b[1].u(v, a, !ok) }
			return
		}
		slowMap(b[:], s, a, true)
	}
}
func _ul3(f []fieldDef) d {
	var b = *(*[3]fieldDef)(unsafe.Pointer(&f[0]))
	return func(s any, a up, isOmitted bool) {
		if from, ok := s.(m); ok {
			var v any
			var ok bool
			if v, ok = from[b[0].n]; true { b[0].u(v, a, !ok) }
			if v, ok = from[b[1].n]; true { b[1].u(v, a, !ok) }
			if v, ok = from[b[2].n]; true { b[2].u(v, a, !ok) }
			return
		}
		slowMap(b[:], s, a, true)
	}
}
func _ul4(f []fieldDef) d {
	var b = *(*[4]fieldDef)(unsafe.Pointer(&f[0]))
	return func(s any, a up, isOmitted bool) {
		if from, ok := s.(m); ok {
			var v any
			var ok bool
			if v, ok = from[b[0].n]; true { b[0].u(v, a, !ok) }
			if v, ok = from[b[1].n]; true { b[1].u(v, a, !ok) }
			if v, ok = from[b[2].n]; true { b[2].u(v, a, !ok) }
			if v, ok = from[b[3].n]; true { b[3].u(v, a, !ok) }
			return
		}
		slowMap(b[:], s, a, true)
	}
}
func _ul5(f []fieldDef) d {
	var b = *(*[5]fieldDef)(unsafe.Pointer(&f[0]))
	return func(s any, a up, isOmitted bool) {
		if from, ok := s.(m); ok {
			var v any
			var ok bool
			if v, ok = from[b[0].n]; true { b[0].u(v, a, !ok) }
			if v, ok = from[b[1].n]; true { b[1].u(v, a, !ok) }
			if v, ok = from[b[2].n]; true { b[2].u(v, a, !ok) }
			if v, ok = from[b[3].n]; true { b[3].u(v, a, !ok) }
			if v, ok = from[b[4].n]; true { b[4].u(v, a, !ok) }
			return
		}
		slowMap(b[:], s, a, true)
	}
}
func _ul6(f []fieldDef) d {
	var b = *(*[6]fieldDef)(unsafe.Pointer(&f[0]))
	return func(s any, a up, isOmitted bool) {
		if from, ok := s.(m); ok {
			var v any
			var ok bool
			if v, ok = from[b[0].n]; true { b[0].u(v, a, !ok) }
			if v, ok = from[b[1].n]; true { b[1].u(v, a, !ok) }
			if v, ok = from[b[2].n]; true { b[2].u(v, a, !ok) }
			if v, ok = from[b[3].n]; true { b[3].u(v, a, !ok) }
			if v, ok = from[b[4].n]; true { b[4].u(v, a, !ok) }
			if v, ok = from[b[5].n]; true { b[5].u(v, a, !ok) }
			return
		}
		slowMap(b[:], s, a, true)
	}
}
func _ul7(f []fieldDef) d {
	var b = *(*[7]fieldDef)(unsafe.Pointer(&f[0]))
	return func(s any, a up, isOmitted bool) {
		if from, ok := s.(m); ok {
			var v any
			var ok bool
			if v, ok = from[b[0].n]; true { b[0].u(v, a, !ok) }
			if v, ok = from[b[1].n]; true { b[1].u(v, a, !ok) }
			if v, ok = from[b[2].n]; true { b[2].u(v, a, !ok) }
			if v, ok = from[b[3].n]; true { b[3].u(v, a, !ok) }
			if v, ok = from[b[4].n]; true { b[4].u(v, a, !ok) }
			if v, ok = from[b[5].n]; true { b[5].u(v, a, !ok) }
			if v, ok = from[b[6].n]; true { b[6].u(v, a, !ok) }
			return
		}
		slowMap(b[:], s, a, true)
	}
}
func _ul8(f []fieldDef) d {
	var b = *(*[8]fieldDef)(unsafe.Pointer(&f[0]))
	return func(s any, a up, isOmitted bool) {
		if from, ok := s.(m); ok {
			var v any
			var ok bool
			if v, ok = from[b[0].n]; true { b[0].u(v, a, !ok) }
			if v, ok = from[b[1].n]; true { b[1].u(v, a, !ok) }
			if v, ok = from[b[2].n]; true { b[2].u(v, a, !ok) }
			if v, ok = from[b[3].n]; true { b[3].u(v, a, !ok) }
			if v, ok = from[b[4].n]; true { b[4].u(v, a, !ok) }
			if v, ok = from[b[5].n]; true { b[5].u(v, a, !ok) }
			if v, ok = from[b[6].n]; true { b[6].u(v, a, !ok) }
			if v, ok = from[b[7].n]; true { b[7].u(v, a, !ok) }
			return
		}
		slowMap(b[:], s, a, true)
	}
}
func _ul9(f []fieldDef) d {
	var b = *(*[9]fieldDef)(unsafe.Pointer(&f[0]))
	return func(s any, a up, isOmitted bool) {
		if from, ok := s.(m); ok {
			var v any
			var ok bool
			if v, ok = from[b[0].n]; true { b[0].u(v, a, !ok) }
			if v, ok = from[b[1].n]; true { b[1].u(v, a, !ok) }
			if v, ok = from[b[2].n]; true { b[2].u(v, a, !ok) }
			if v, ok = from[b[3].n]; true { b[3].u(v, a, !ok) }
			if v, ok = from[b[4].n]; true { b[4].u(v, a, !ok) }
			if v, ok = from[b[5].n]; true { b[5].u(v, a, !ok) }
			if v, ok = from[b[6].n]; true { b[6].u(v, a, !ok) }
			if v, ok = from[b[7].n]; true { b[7].u(v, a, !ok) }
			if v, ok = from[b[8].n]; true { b[8].u(v, a, !ok) }
			return
		}
		slowMap(b[:], s, a, true)
	}
}
func _ul10(f []fieldDef) d {
	var b = *(*[10]fieldDef)(unsafe.Pointer(&f[0]))
	return func(s any, a up, isOmitted bool) {
		if from, ok := s.(m); ok {
			var v any
			var ok bool
			if v, ok = from[b[0].n]; true { b[0].u(v, a, !ok) }
			if v, ok = from[b[1].n]; true { b[1].u(v, a, !ok) }
			if v, ok = from[b[2].n]; true { b[2].u(v, a, !ok) }
			if v, ok = from[b[3].n]; true { b[3].u(v, a, !ok) }
			if v, ok = from[b[4].n]; true { b[4].u(v, a, !ok) }
			if v, ok = from[b[5].n]; true { b[5].u(v, a, !ok) }
			if v, ok = from[b[6].n]; true { b[6].u(v, a, !ok) }
			if v, ok = from[b[7].n]; true { b[7].u(v, a, !ok) }
			if v, ok = from[b[8].n]; true { b[8].u(v, a, !ok) }
			if v, ok = from[b[9].n]; true { b[9].u(v, a, !ok) }
			return
		}
		slowMap(b[:], s, a, true)
	}
}

func unrolledLoopRequired(f []fieldDef) d {
	switch len(f) {
	case 0:
		return nil
	case 1:
		return _ulm1(f)
	case 2:
		return _ulm2(f)
	case 3:
		return _ulm3(f)
	case 4:
		return _ulm4(f)
	case 5:
		return _ulm5(f)
	case 6:
		return _ulm6(f)
	case 7:
		return _ulm7(f)
	case 8:
		return _ulm8(f)
	case 9:
		return _ulm9(f)
	case 10:
		return _ulm10(f)
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
func _ulm1(f []fieldDef) d {
	var b = *(*[1]fieldDef)(unsafe.Pointer(&f[0]))
	return func(s any, a up, isOmitted bool) {
		if from, ok := s.(m); ok {
			b[0].u(_u(from, b[0].n), a, false)
			return
		}
		slowMap(b[:], s, a, true)
	}
}
func _ulm2(f []fieldDef) d {
	var b = *(*[2]fieldDef)(unsafe.Pointer(&f[0]))
	return func(s any, a up, isOmitted bool) {
		if from, ok := s.(m); ok {
			b[0].u(_u(from, b[0].n), a, false)
			b[1].u(_u(from, b[1].n), a, false)
			return
		}
		slowMap(b[:], s, a, true)
	}
}
func _ulm3(f []fieldDef) d {
	var b = *(*[3]fieldDef)(unsafe.Pointer(&f[0]))
	return func(s any, a up, isOmitted bool) {
		if from, ok := s.(m); ok {
			b[0].u(_u(from, b[0].n), a, false)
			b[1].u(_u(from, b[1].n), a, false)
			b[2].u(_u(from, b[2].n), a, false)
			return
		}
		slowMap(b[:], s, a, true)
	}
}
func _ulm4(f []fieldDef) d {
	var b = *(*[4]fieldDef)(unsafe.Pointer(&f[0]))
	return func(s any, a up, isOmitted bool) {
		if from, ok := s.(m); ok {
			b[0].u(_u(from, b[0].n), a, false)
			b[1].u(_u(from, b[1].n), a, false)
			b[2].u(_u(from, b[2].n), a, false)
			b[3].u(_u(from, b[3].n), a, false)
			return
		}
		slowMap(b[:], s, a, true)
	}
}
func _ulm5(f []fieldDef) d {
	var b = *(*[5]fieldDef)(unsafe.Pointer(&f[0]))
	return func(s any, a up, isOmitted bool) {
		if from, ok := s.(m); ok {
			b[0].u(_u(from, b[0].n), a, false)
			b[1].u(_u(from, b[1].n), a, false)
			b[2].u(_u(from, b[2].n), a, false)
			b[3].u(_u(from, b[3].n), a, false)
			b[4].u(_u(from, b[4].n), a, false)
			return
		}
		slowMap(b[:], s, a, true)
	}
}
func _ulm6(f []fieldDef) d {
	var b = *(*[6]fieldDef)(unsafe.Pointer(&f[0]))
	return func(s any, a up, isOmitted bool) {
		if from, ok := s.(m); ok {
			b[0].u(_u(from, b[0].n), a, false)
			b[1].u(_u(from, b[1].n), a, false)
			b[2].u(_u(from, b[2].n), a, false)
			b[3].u(_u(from, b[3].n), a, false)
			b[4].u(_u(from, b[4].n), a, false)
			b[5].u(_u(from, b[5].n), a, false)
			return
		}
		slowMap(b[:], s, a, true)
	}
}
func _ulm7(f []fieldDef) d {
	var b = *(*[7]fieldDef)(unsafe.Pointer(&f[0]))
	return func(s any, a up, isOmitted bool) {
		if from, ok := s.(m); ok {
			b[0].u(_u(from, b[0].n), a, false)
			b[1].u(_u(from, b[1].n), a, false)
			b[2].u(_u(from, b[2].n), a, false)
			b[3].u(_u(from, b[3].n), a, false)
			b[4].u(_u(from, b[4].n), a, false)
			b[5].u(_u(from, b[5].n), a, false)
			b[6].u(_u(from, b[6].n), a, false)
			return
		}
		slowMap(b[:], s, a, true)
	}
}
func _ulm8(f []fieldDef) d {
	var b = *(*[8]fieldDef)(unsafe.Pointer(&f[0]))
	return func(s any, a up, isOmitted bool) {
		if from, ok := s.(m); ok {
			b[0].u(_u(from, b[0].n), a, false)
			b[1].u(_u(from, b[1].n), a, false)
			b[2].u(_u(from, b[2].n), a, false)
			b[3].u(_u(from, b[3].n), a, false)
			b[4].u(_u(from, b[4].n), a, false)
			b[5].u(_u(from, b[5].n), a, false)
			b[6].u(_u(from, b[6].n), a, false)
			b[7].u(_u(from, b[7].n), a, false)
			return
		}
		slowMap(b[:], s, a, true)
	}
}
func _ulm9(f []fieldDef) d {
	var b = *(*[9]fieldDef)(unsafe.Pointer(&f[0]))
	return func(s any, a up, isOmitted bool) {
		if from, ok := s.(m); ok {
			b[0].u(_u(from, b[0].n), a, false)
			b[1].u(_u(from, b[1].n), a, false)
			b[2].u(_u(from, b[2].n), a, false)
			b[3].u(_u(from, b[3].n), a, false)
			b[4].u(_u(from, b[4].n), a, false)
			b[5].u(_u(from, b[5].n), a, false)
			b[6].u(_u(from, b[6].n), a, false)
			b[7].u(_u(from, b[7].n), a, false)
			b[8].u(_u(from, b[8].n), a, false)
			return
		}
		slowMap(b[:], s, a, true)
	}
}
func _ulm10(f []fieldDef) d {
	var b = *(*[10]fieldDef)(unsafe.Pointer(&f[0]))
	return func(s any, a up, isOmitted bool) {
		if from, ok := s.(m); ok {
			b[0].u(_u(from, b[0].n), a, false)
			b[1].u(_u(from, b[1].n), a, false)
			b[2].u(_u(from, b[2].n), a, false)
			b[3].u(_u(from, b[3].n), a, false)
			b[4].u(_u(from, b[4].n), a, false)
			b[5].u(_u(from, b[5].n), a, false)
			b[6].u(_u(from, b[6].n), a, false)
			b[7].u(_u(from, b[7].n), a, false)
			b[8].u(_u(from, b[8].n), a, false)
			b[9].u(_u(from, b[9].n), a, false)
			return
		}
		slowMap(b[:], s, a, true)
	}
}

type b = byte

func getMapToolkitByKeyTypeAndValueReflection[K comparable](elementType reflect.Type) (r *mapToolkit) {
	elementSize := elementType.Size()
	switch elementSize {
	case 1: r = _m[K, [1]b]()
	case 2: r = _m[K, [2]b]()
	case 3: r = _m[K, [3]b]()
	case 4: r = _m[K, [4]b]()
	case 5: r = _m[K, [5]b]()
	case 6: r = _m[K, [6]b]()
	case 7: r = _m[K, [7]b]()
	case 8: r = _m[K, [8]b]()
	case 9: r = _m[K, [9]b]()
	case 10: r = _m[K, [10]b]()
	case 11: r = _m[K, [11]b]()
	case 12: r = _m[K, [12]b]()
	case 13: r = _m[K, [13]b]()
	case 14: r = _m[K, [14]b]()
	case 15: r = _m[K, [15]b]()
	case 16: r = _m[K, [16]b]()
	case 17: r = _m[K, [17]b]()
	case 18: r = _m[K, [18]b]()
	case 19: r = _m[K, [19]b]()
	case 20: r = _m[K, [20]b]()
	case 21: r = _m[K, [21]b]()
	case 22: r = _m[K, [22]b]()
	case 23: r = _m[K, [23]b]()
	case 24: r = _m[K, [24]b]()
	case 25: r = _m[K, [25]b]()
	case 26: r = _m[K, [26]b]()
	case 27: r = _m[K, [27]b]()
	case 28: r = _m[K, [28]b]()
	case 29: r = _m[K, [29]b]()
	case 30: r = _m[K, [30]b]()
	case 31: r = _m[K, [31]b]()
	case 32: r = _m[K, [32]b]()
	case 33: r = _m[K, [33]b]()
	case 34: r = _m[K, [34]b]()
	case 35: r = _m[K, [35]b]()
	case 36: r = _m[K, [36]b]()
	case 37: r = _m[K, [37]b]()
	case 38: r = _m[K, [38]b]()
	case 39: r = _m[K, [39]b]()
	case 40: r = _m[K, [40]b]()
	case 41: r = _m[K, [41]b]()
	case 42: r = _m[K, [42]b]()
	case 43: r = _m[K, [43]b]()
	case 44: r = _m[K, [44]b]()
	case 45: r = _m[K, [45]b]()
	case 46: r = _m[K, [46]b]()
	case 47: r = _m[K, [47]b]()
	case 48: r = _m[K, [48]b]()
	case 49: r = _m[K, [49]b]()
	case 50: r = _m[K, [50]b]()
	case 51: r = _m[K, [51]b]()
	case 52: r = _m[K, [52]b]()
	case 53: r = _m[K, [53]b]()
	case 54: r = _m[K, [54]b]()
	case 55: r = _m[K, [55]b]()
	case 56: r = _m[K, [56]b]()
	case 57: r = _m[K, [57]b]()
	case 58: r = _m[K, [58]b]()
	case 59: r = _m[K, [59]b]()
	case 60: r = _m[K, [60]b]()
	case 61: r = _m[K, [61]b]()
	case 62: r = _m[K, [62]b]()
	case 63: r = _m[K, [63]b]()
	case 64: r = _m[K, [64]b]()
	default:
	}

	return
}

func getMapToolkitForStructKeyAndValueByReflection(mapType reflect.Type) (r *mapToolkit) {
	keySize := mapType.Key().Size()
	e := mapType.Elem() 
	switch keySize {
	case 1: r = _o[[1]b](e)
	case 2: r = _o[[2]b](e)
	case 3: r = _o[[3]b](e)
	case 4: r = _o[[4]b](e)
	case 5: r = _o[[5]b](e)
	case 6: r = _o[[6]b](e)
	case 7: r = _o[[7]b](e)
	case 8: r = _o[[8]b](e)
	case 9: r = _o[[9]b](e)
	case 10: r = _o[[10]b](e)
	case 11: r = _o[[11]b](e)
	case 12: r = _o[[12]b](e)
	case 13: r = _o[[13]b](e)
	case 14: r = _o[[14]b](e)
	case 15: r = _o[[15]b](e)
	case 16: r = _o[[16]b](e)
	case 17: r = _o[[17]b](e)
	case 18: r = _o[[18]b](e)
	case 19: r = _o[[19]b](e)
	case 20: r = _o[[20]b](e)
	case 21: r = _o[[21]b](e)
	case 22: r = _o[[22]b](e)
	case 23: r = _o[[23]b](e)
	case 24: r = _o[[24]b](e)
	case 25: r = _o[[25]b](e)
	case 26: r = _o[[26]b](e)
	case 27: r = _o[[27]b](e)
	case 28: r = _o[[28]b](e)
	case 29: r = _o[[29]b](e)
	case 30: r = _o[[30]b](e)
	case 31: r = _o[[31]b](e)
	case 32: r = _o[[32]b](e)
	case 33: r = _o[[33]b](e)
	case 34: r = _o[[34]b](e)
	case 35: r = _o[[35]b](e)
	case 36: r = _o[[36]b](e)
	case 37: r = _o[[37]b](e)
	case 38: r = _o[[38]b](e)
	case 39: r = _o[[39]b](e)
	case 40: r = _o[[40]b](e)
	case 41: r = _o[[41]b](e)
	case 42: r = _o[[42]b](e)
	case 43: r = _o[[43]b](e)
	case 44: r = _o[[44]b](e)
	case 45: r = _o[[45]b](e)
	case 46: r = _o[[46]b](e)
	case 47: r = _o[[47]b](e)
	case 48: r = _o[[48]b](e)
	case 49: r = _o[[49]b](e)
	case 50: r = _o[[50]b](e)
	case 51: r = _o[[51]b](e)
	case 52: r = _o[[52]b](e)
	case 53: r = _o[[53]b](e)
	case 54: r = _o[[54]b](e)
	case 55: r = _o[[55]b](e)
	case 56: r = _o[[56]b](e)
	case 57: r = _o[[57]b](e)
	case 58: r = _o[[58]b](e)
	case 59: r = _o[[59]b](e)
	case 60: r = _o[[60]b](e)
	case 61: r = _o[[61]b](e)
	case 62: r = _o[[62]b](e)
	case 63: r = _o[[63]b](e)
	case 64: r = _o[[64]b](e)
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
	case 1:
		r = _s[[1]b]
	case 2:
		r = _s[[2]b]
	case 3:
		r = _s[[3]b]
	case 4:
		r = _s[[4]b]
	case 5:
		r = _s[[5]b]
	case 6:
		r = _s[[6]b]
	case 7:
		r = _s[[7]b]
	case 8:
		r = _s[[8]b]
	case 9:
		r = _s[[9]b]
	case 10:
		r = _s[[10]b]
	case 11:
		r = _s[[11]b]
	case 12:
		r = _s[[12]b]
	case 13:
		r = _s[[13]b]
	case 14:
		r = _s[[14]b]
	case 15:
		r = _s[[15]b]
	case 16:
		r = _s[[16]b]
	case 17:
		r = _s[[17]b]
	case 18:
		r = _s[[18]b]
	case 19:
		r = _s[[19]b]
	case 20:
		r = _s[[20]b]
	case 21:
		r = _s[[21]b]
	case 22:
		r = _s[[22]b]
	case 23:
		r = _s[[23]b]
	case 24:
		r = _s[[24]b]
	case 25:
		r = _s[[25]b]
	case 26:
		r = _s[[26]b]
	case 27:
		r = _s[[27]b]
	case 28:
		r = _s[[28]b]
	case 29:
		r = _s[[29]b]
	case 30:
		r = _s[[30]b]
	case 31:
		r = _s[[31]b]
	case 32:
		r = _s[[32]b]
	case 33:
		r = _s[[33]b]
	case 34:
		r = _s[[34]b]
	case 35:
		r = _s[[35]b]
	case 36:
		r = _s[[36]b]
	case 37:
		r = _s[[37]b]
	case 38:
		r = _s[[38]b]
	case 39:
		r = _s[[39]b]
	case 40:
		r = _s[[40]b]
	case 41:
		r = _s[[41]b]
	case 42:
		r = _s[[42]b]
	case 43:
		r = _s[[43]b]
	case 44:
		r = _s[[44]b]
	case 45:
		r = _s[[45]b]
	case 46:
		r = _s[[46]b]
	case 47:
		r = _s[[47]b]
	case 48:
		r = _s[[48]b]
	case 49:
		r = _s[[49]b]
	case 50:
		r = _s[[50]b]
	case 51:
		r = _s[[51]b]
	case 52:
		r = _s[[52]b]
	case 53:
		r = _s[[53]b]
	case 54:
		r = _s[[54]b]
	case 55:
		r = _s[[55]b]
	case 56:
		r = _s[[56]b]
	case 57:
		r = _s[[57]b]
	case 58:
		r = _s[[58]b]
	case 59:
		r = _s[[59]b]
	case 60:
		r = _s[[60]b]
	case 61:
		r = _s[[61]b]
	case 62:
		r = _s[[62]b]
	case 63:
		r = _s[[63]b]
	case 64:
		r = _s[[64]b]
	default:
		// degrade to reflect after 64 bytes element size
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
}