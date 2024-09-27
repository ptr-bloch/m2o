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

	// reflection is used only for decoding slices with element size bigger than 1024 bytes
	// for example:
	// arrays: [1024]byte
	// structures with fields with total size greater than 1024: struct { f1 byte ... f1024 byte }
	// slices as element don't have this problem because it's size is Sizeof(SliceHeader)
	"reflect"
	"unsafe"
)

type up = unsafe.Pointer

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
			if v, ok = from[b[0].n]; true {
				b[0].u(v, a, !ok)
			}
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
			if v, ok = from[b[0].n]; true {
				b[0].u(v, a, !ok)
			}
			if v, ok = from[b[1].n]; true {
				b[1].u(v, a, !ok)
			}
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
			if v, ok = from[b[0].n]; true {
				b[0].u(v, a, !ok)
			}
			if v, ok = from[b[1].n]; true {
				b[1].u(v, a, !ok)
			}
			if v, ok = from[b[2].n]; true {
				b[2].u(v, a, !ok)
			}
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
			if v, ok = from[b[0].n]; true {
				b[0].u(v, a, !ok)
			}
			if v, ok = from[b[1].n]; true {
				b[1].u(v, a, !ok)
			}
			if v, ok = from[b[2].n]; true {
				b[2].u(v, a, !ok)
			}
			if v, ok = from[b[3].n]; true {
				b[3].u(v, a, !ok)
			}
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
			if v, ok = from[b[0].n]; true {
				b[0].u(v, a, !ok)
			}
			if v, ok = from[b[1].n]; true {
				b[1].u(v, a, !ok)
			}
			if v, ok = from[b[2].n]; true {
				b[2].u(v, a, !ok)
			}
			if v, ok = from[b[3].n]; true {
				b[3].u(v, a, !ok)
			}
			if v, ok = from[b[4].n]; true {
				b[4].u(v, a, !ok)
			}
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
			if v, ok = from[b[0].n]; true {
				b[0].u(v, a, !ok)
			}
			if v, ok = from[b[1].n]; true {
				b[1].u(v, a, !ok)
			}
			if v, ok = from[b[2].n]; true {
				b[2].u(v, a, !ok)
			}
			if v, ok = from[b[3].n]; true {
				b[3].u(v, a, !ok)
			}
			if v, ok = from[b[4].n]; true {
				b[4].u(v, a, !ok)
			}
			if v, ok = from[b[5].n]; true {
				b[5].u(v, a, !ok)
			}
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
			if v, ok = from[b[0].n]; true {
				b[0].u(v, a, !ok)
			}
			if v, ok = from[b[1].n]; true {
				b[1].u(v, a, !ok)
			}
			if v, ok = from[b[2].n]; true {
				b[2].u(v, a, !ok)
			}
			if v, ok = from[b[3].n]; true {
				b[3].u(v, a, !ok)
			}
			if v, ok = from[b[4].n]; true {
				b[4].u(v, a, !ok)
			}
			if v, ok = from[b[5].n]; true {
				b[5].u(v, a, !ok)
			}
			if v, ok = from[b[6].n]; true {
				b[6].u(v, a, !ok)
			}
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
			if v, ok = from[b[0].n]; true {
				b[0].u(v, a, !ok)
			}
			if v, ok = from[b[1].n]; true {
				b[1].u(v, a, !ok)
			}
			if v, ok = from[b[2].n]; true {
				b[2].u(v, a, !ok)
			}
			if v, ok = from[b[3].n]; true {
				b[3].u(v, a, !ok)
			}
			if v, ok = from[b[4].n]; true {
				b[4].u(v, a, !ok)
			}
			if v, ok = from[b[5].n]; true {
				b[5].u(v, a, !ok)
			}
			if v, ok = from[b[6].n]; true {
				b[6].u(v, a, !ok)
			}
			if v, ok = from[b[7].n]; true {
				b[7].u(v, a, !ok)
			}
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
			if v, ok = from[b[0].n]; true {
				b[0].u(v, a, !ok)
			}
			if v, ok = from[b[1].n]; true {
				b[1].u(v, a, !ok)
			}
			if v, ok = from[b[2].n]; true {
				b[2].u(v, a, !ok)
			}
			if v, ok = from[b[3].n]; true {
				b[3].u(v, a, !ok)
			}
			if v, ok = from[b[4].n]; true {
				b[4].u(v, a, !ok)
			}
			if v, ok = from[b[5].n]; true {
				b[5].u(v, a, !ok)
			}
			if v, ok = from[b[6].n]; true {
				b[6].u(v, a, !ok)
			}
			if v, ok = from[b[7].n]; true {
				b[7].u(v, a, !ok)
			}
			if v, ok = from[b[8].n]; true {
				b[8].u(v, a, !ok)
			}
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
			if v, ok = from[b[0].n]; true {
				b[0].u(v, a, !ok)
			}
			if v, ok = from[b[1].n]; true {
				b[1].u(v, a, !ok)
			}
			if v, ok = from[b[2].n]; true {
				b[2].u(v, a, !ok)
			}
			if v, ok = from[b[3].n]; true {
				b[3].u(v, a, !ok)
			}
			if v, ok = from[b[4].n]; true {
				b[4].u(v, a, !ok)
			}
			if v, ok = from[b[5].n]; true {
				b[5].u(v, a, !ok)
			}
			if v, ok = from[b[6].n]; true {
				b[6].u(v, a, !ok)
			}
			if v, ok = from[b[7].n]; true {
				b[7].u(v, a, !ok)
			}
			if v, ok = from[b[8].n]; true {
				b[8].u(v, a, !ok)
			}
			if v, ok = from[b[9].n]; true {
				b[9].u(v, a, !ok)
			}
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
	case 1:
		r = _m[K, [1]b]()
	case 2:
		r = _m[K, [2]b]()
	case 3:
		r = _m[K, [3]b]()
	case 4:
		r = _m[K, [4]b]()
	case 5:
		r = _m[K, [5]b]()
	case 6:
		r = _m[K, [6]b]()
	case 7:
		r = _m[K, [7]b]()
	case 8:
		r = _m[K, [8]b]()
	case 9:
		r = _m[K, [9]b]()
	case 10:
		r = _m[K, [10]b]()
	case 11:
		r = _m[K, [11]b]()
	case 12:
		r = _m[K, [12]b]()
	case 13:
		r = _m[K, [13]b]()
	case 14:
		r = _m[K, [14]b]()
	case 15:
		r = _m[K, [15]b]()
	case 16:
		r = _m[K, [16]b]()
	case 17:
		r = _m[K, [17]b]()
	case 18:
		r = _m[K, [18]b]()
	case 19:
		r = _m[K, [19]b]()
	case 20:
		r = _m[K, [20]b]()
	case 21:
		r = _m[K, [21]b]()
	case 22:
		r = _m[K, [22]b]()
	case 23:
		r = _m[K, [23]b]()
	case 24:
		r = _m[K, [24]b]()
	case 25:
		r = _m[K, [25]b]()
	case 26:
		r = _m[K, [26]b]()
	case 27:
		r = _m[K, [27]b]()
	case 28:
		r = _m[K, [28]b]()
	case 29:
		r = _m[K, [29]b]()
	case 30:
		r = _m[K, [30]b]()
	case 31:
		r = _m[K, [31]b]()
	case 32:
		r = _m[K, [32]b]()
	case 33:
		r = _m[K, [33]b]()
	case 34:
		r = _m[K, [34]b]()
	case 35:
		r = _m[K, [35]b]()
	case 36:
		r = _m[K, [36]b]()
	case 37:
		r = _m[K, [37]b]()
	case 38:
		r = _m[K, [38]b]()
	case 39:
		r = _m[K, [39]b]()
	case 40:
		r = _m[K, [40]b]()
	case 41:
		r = _m[K, [41]b]()
	case 42:
		r = _m[K, [42]b]()
	case 43:
		r = _m[K, [43]b]()
	case 44:
		r = _m[K, [44]b]()
	case 45:
		r = _m[K, [45]b]()
	case 46:
		r = _m[K, [46]b]()
	case 47:
		r = _m[K, [47]b]()
	case 48:
		r = _m[K, [48]b]()
	case 49:
		r = _m[K, [49]b]()
	case 50:
		r = _m[K, [50]b]()
	case 51:
		r = _m[K, [51]b]()
	case 52:
		r = _m[K, [52]b]()
	case 53:
		r = _m[K, [53]b]()
	case 54:
		r = _m[K, [54]b]()
	case 55:
		r = _m[K, [55]b]()
	case 56:
		r = _m[K, [56]b]()
	case 57:
		r = _m[K, [57]b]()
	case 58:
		r = _m[K, [58]b]()
	case 59:
		r = _m[K, [59]b]()
	case 60:
		r = _m[K, [60]b]()
	case 61:
		r = _m[K, [61]b]()
	case 62:
		r = _m[K, [62]b]()
	case 63:
		r = _m[K, [63]b]()
	case 64:
		r = _m[K, [64]b]()
	case 65:
		r = _m[K, [65]b]()
	case 66:
		r = _m[K, [66]b]()
	case 67:
		r = _m[K, [67]b]()
	case 68:
		r = _m[K, [68]b]()
	case 69:
		r = _m[K, [69]b]()
	case 70:
		r = _m[K, [70]b]()
	case 71:
		r = _m[K, [71]b]()
	case 72:
		r = _m[K, [72]b]()
	case 73:
		r = _m[K, [73]b]()
	case 74:
		r = _m[K, [74]b]()
	case 75:
		r = _m[K, [75]b]()
	case 76:
		r = _m[K, [76]b]()
	case 77:
		r = _m[K, [77]b]()
	case 78:
		r = _m[K, [78]b]()
	case 79:
		r = _m[K, [79]b]()
	case 80:
		r = _m[K, [80]b]()
	case 81:
		r = _m[K, [81]b]()
	case 82:
		r = _m[K, [82]b]()
	case 83:
		r = _m[K, [83]b]()
	case 84:
		r = _m[K, [84]b]()
	case 85:
		r = _m[K, [85]b]()
	case 86:
		r = _m[K, [86]b]()
	case 87:
		r = _m[K, [87]b]()
	case 88:
		r = _m[K, [88]b]()
	case 89:
		r = _m[K, [89]b]()
	case 90:
		r = _m[K, [90]b]()
	case 91:
		r = _m[K, [91]b]()
	case 92:
		r = _m[K, [92]b]()
	case 93:
		r = _m[K, [93]b]()
	case 94:
		r = _m[K, [94]b]()
	case 95:
		r = _m[K, [95]b]()
	case 96:
		r = _m[K, [96]b]()
	case 97:
		r = _m[K, [97]b]()
	case 98:
		r = _m[K, [98]b]()
	case 99:
		r = _m[K, [99]b]()
	case 100:
		r = _m[K, [100]b]()
	case 101:
		r = _m[K, [101]b]()
	case 102:
		r = _m[K, [102]b]()
	case 103:
		r = _m[K, [103]b]()
	case 104:
		r = _m[K, [104]b]()
	case 105:
		r = _m[K, [105]b]()
	case 106:
		r = _m[K, [106]b]()
	case 107:
		r = _m[K, [107]b]()
	case 108:
		r = _m[K, [108]b]()
	case 109:
		r = _m[K, [109]b]()
	case 110:
		r = _m[K, [110]b]()
	case 111:
		r = _m[K, [111]b]()
	case 112:
		r = _m[K, [112]b]()
	case 113:
		r = _m[K, [113]b]()
	case 114:
		r = _m[K, [114]b]()
	case 115:
		r = _m[K, [115]b]()
	case 116:
		r = _m[K, [116]b]()
	case 117:
		r = _m[K, [117]b]()
	case 118:
		r = _m[K, [118]b]()
	case 119:
		r = _m[K, [119]b]()
	case 120:
		r = _m[K, [120]b]()
	case 121:
		r = _m[K, [121]b]()
	case 122:
		r = _m[K, [122]b]()
	case 123:
		r = _m[K, [123]b]()
	case 124:
		r = _m[K, [124]b]()
	case 125:
		r = _m[K, [125]b]()
	case 126:
		r = _m[K, [126]b]()
	case 127:
		r = _m[K, [127]b]()
	case 128:
		r = _m[K, [128]b]()
	case 129:
		r = _m[K, [129]b]()
	case 130:
		r = _m[K, [130]b]()
	case 131:
		r = _m[K, [131]b]()
	case 132:
		r = _m[K, [132]b]()
	case 133:
		r = _m[K, [133]b]()
	case 134:
		r = _m[K, [134]b]()
	case 135:
		r = _m[K, [135]b]()
	case 136:
		r = _m[K, [136]b]()
	case 137:
		r = _m[K, [137]b]()
	case 138:
		r = _m[K, [138]b]()
	case 139:
		r = _m[K, [139]b]()
	case 140:
		r = _m[K, [140]b]()
	case 141:
		r = _m[K, [141]b]()
	case 142:
		r = _m[K, [142]b]()
	case 143:
		r = _m[K, [143]b]()
	case 144:
		r = _m[K, [144]b]()
	case 145:
		r = _m[K, [145]b]()
	case 146:
		r = _m[K, [146]b]()
	case 147:
		r = _m[K, [147]b]()
	case 148:
		r = _m[K, [148]b]()
	case 149:
		r = _m[K, [149]b]()
	case 150:
		r = _m[K, [150]b]()
	case 151:
		r = _m[K, [151]b]()
	case 152:
		r = _m[K, [152]b]()
	case 153:
		r = _m[K, [153]b]()
	case 154:
		r = _m[K, [154]b]()
	case 155:
		r = _m[K, [155]b]()
	case 156:
		r = _m[K, [156]b]()
	case 157:
		r = _m[K, [157]b]()
	case 158:
		r = _m[K, [158]b]()
	case 159:
		r = _m[K, [159]b]()
	case 160:
		r = _m[K, [160]b]()
	case 161:
		r = _m[K, [161]b]()
	case 162:
		r = _m[K, [162]b]()
	case 163:
		r = _m[K, [163]b]()
	case 164:
		r = _m[K, [164]b]()
	case 165:
		r = _m[K, [165]b]()
	case 166:
		r = _m[K, [166]b]()
	case 167:
		r = _m[K, [167]b]()
	case 168:
		r = _m[K, [168]b]()
	case 169:
		r = _m[K, [169]b]()
	case 170:
		r = _m[K, [170]b]()
	case 171:
		r = _m[K, [171]b]()
	case 172:
		r = _m[K, [172]b]()
	case 173:
		r = _m[K, [173]b]()
	case 174:
		r = _m[K, [174]b]()
	case 175:
		r = _m[K, [175]b]()
	case 176:
		r = _m[K, [176]b]()
	case 177:
		r = _m[K, [177]b]()
	case 178:
		r = _m[K, [178]b]()
	case 179:
		r = _m[K, [179]b]()
	case 180:
		r = _m[K, [180]b]()
	case 181:
		r = _m[K, [181]b]()
	case 182:
		r = _m[K, [182]b]()
	case 183:
		r = _m[K, [183]b]()
	case 184:
		r = _m[K, [184]b]()
	case 185:
		r = _m[K, [185]b]()
	case 186:
		r = _m[K, [186]b]()
	case 187:
		r = _m[K, [187]b]()
	case 188:
		r = _m[K, [188]b]()
	case 189:
		r = _m[K, [189]b]()
	case 190:
		r = _m[K, [190]b]()
	case 191:
		r = _m[K, [191]b]()
	case 192:
		r = _m[K, [192]b]()
	case 193:
		r = _m[K, [193]b]()
	case 194:
		r = _m[K, [194]b]()
	case 195:
		r = _m[K, [195]b]()
	case 196:
		r = _m[K, [196]b]()
	case 197:
		r = _m[K, [197]b]()
	case 198:
		r = _m[K, [198]b]()
	case 199:
		r = _m[K, [199]b]()
	case 200:
		r = _m[K, [200]b]()
	case 201:
		r = _m[K, [201]b]()
	case 202:
		r = _m[K, [202]b]()
	case 203:
		r = _m[K, [203]b]()
	case 204:
		r = _m[K, [204]b]()
	case 205:
		r = _m[K, [205]b]()
	case 206:
		r = _m[K, [206]b]()
	case 207:
		r = _m[K, [207]b]()
	case 208:
		r = _m[K, [208]b]()
	case 209:
		r = _m[K, [209]b]()
	case 210:
		r = _m[K, [210]b]()
	case 211:
		r = _m[K, [211]b]()
	case 212:
		r = _m[K, [212]b]()
	case 213:
		r = _m[K, [213]b]()
	case 214:
		r = _m[K, [214]b]()
	case 215:
		r = _m[K, [215]b]()
	case 216:
		r = _m[K, [216]b]()
	case 217:
		r = _m[K, [217]b]()
	case 218:
		r = _m[K, [218]b]()
	case 219:
		r = _m[K, [219]b]()
	case 220:
		r = _m[K, [220]b]()
	case 221:
		r = _m[K, [221]b]()
	case 222:
		r = _m[K, [222]b]()
	case 223:
		r = _m[K, [223]b]()
	case 224:
		r = _m[K, [224]b]()
	case 225:
		r = _m[K, [225]b]()
	case 226:
		r = _m[K, [226]b]()
	case 227:
		r = _m[K, [227]b]()
	case 228:
		r = _m[K, [228]b]()
	case 229:
		r = _m[K, [229]b]()
	case 230:
		r = _m[K, [230]b]()
	case 231:
		r = _m[K, [231]b]()
	case 232:
		r = _m[K, [232]b]()
	case 233:
		r = _m[K, [233]b]()
	case 234:
		r = _m[K, [234]b]()
	case 235:
		r = _m[K, [235]b]()
	case 236:
		r = _m[K, [236]b]()
	case 237:
		r = _m[K, [237]b]()
	case 238:
		r = _m[K, [238]b]()
	case 239:
		r = _m[K, [239]b]()
	case 240:
		r = _m[K, [240]b]()
	case 241:
		r = _m[K, [241]b]()
	case 242:
		r = _m[K, [242]b]()
	case 243:
		r = _m[K, [243]b]()
	case 244:
		r = _m[K, [244]b]()
	case 245:
		r = _m[K, [245]b]()
	case 246:
		r = _m[K, [246]b]()
	case 247:
		r = _m[K, [247]b]()
	case 248:
		r = _m[K, [248]b]()
	case 249:
		r = _m[K, [249]b]()
	case 250:
		r = _m[K, [250]b]()
	case 251:
		r = _m[K, [251]b]()
	case 252:
		r = _m[K, [252]b]()
	case 253:
		r = _m[K, [253]b]()
	case 254:
		r = _m[K, [254]b]()
	case 255:
		r = _m[K, [255]b]()
	case 256:
		r = _m[K, [256]b]()
	case 257:
		r = _m[K, [257]b]()
	case 258:
		r = _m[K, [258]b]()
	case 259:
		r = _m[K, [259]b]()
	case 260:
		r = _m[K, [260]b]()
	case 261:
		r = _m[K, [261]b]()
	case 262:
		r = _m[K, [262]b]()
	case 263:
		r = _m[K, [263]b]()
	case 264:
		r = _m[K, [264]b]()
	case 265:
		r = _m[K, [265]b]()
	case 266:
		r = _m[K, [266]b]()
	case 267:
		r = _m[K, [267]b]()
	case 268:
		r = _m[K, [268]b]()
	case 269:
		r = _m[K, [269]b]()
	case 270:
		r = _m[K, [270]b]()
	case 271:
		r = _m[K, [271]b]()
	case 272:
		r = _m[K, [272]b]()
	case 273:
		r = _m[K, [273]b]()
	case 274:
		r = _m[K, [274]b]()
	case 275:
		r = _m[K, [275]b]()
	case 276:
		r = _m[K, [276]b]()
	case 277:
		r = _m[K, [277]b]()
	case 278:
		r = _m[K, [278]b]()
	case 279:
		r = _m[K, [279]b]()
	case 280:
		r = _m[K, [280]b]()
	case 281:
		r = _m[K, [281]b]()
	case 282:
		r = _m[K, [282]b]()
	case 283:
		r = _m[K, [283]b]()
	case 284:
		r = _m[K, [284]b]()
	case 285:
		r = _m[K, [285]b]()
	case 286:
		r = _m[K, [286]b]()
	case 287:
		r = _m[K, [287]b]()
	case 288:
		r = _m[K, [288]b]()
	case 289:
		r = _m[K, [289]b]()
	case 290:
		r = _m[K, [290]b]()
	case 291:
		r = _m[K, [291]b]()
	case 292:
		r = _m[K, [292]b]()
	case 293:
		r = _m[K, [293]b]()
	case 294:
		r = _m[K, [294]b]()
	case 295:
		r = _m[K, [295]b]()
	case 296:
		r = _m[K, [296]b]()
	case 297:
		r = _m[K, [297]b]()
	case 298:
		r = _m[K, [298]b]()
	case 299:
		r = _m[K, [299]b]()
	case 300:
		r = _m[K, [300]b]()
	case 301:
		r = _m[K, [301]b]()
	case 302:
		r = _m[K, [302]b]()
	case 303:
		r = _m[K, [303]b]()
	case 304:
		r = _m[K, [304]b]()
	case 305:
		r = _m[K, [305]b]()
	case 306:
		r = _m[K, [306]b]()
	case 307:
		r = _m[K, [307]b]()
	case 308:
		r = _m[K, [308]b]()
	case 309:
		r = _m[K, [309]b]()
	case 310:
		r = _m[K, [310]b]()
	case 311:
		r = _m[K, [311]b]()
	case 312:
		r = _m[K, [312]b]()
	case 313:
		r = _m[K, [313]b]()
	case 314:
		r = _m[K, [314]b]()
	case 315:
		r = _m[K, [315]b]()
	case 316:
		r = _m[K, [316]b]()
	case 317:
		r = _m[K, [317]b]()
	case 318:
		r = _m[K, [318]b]()
	case 319:
		r = _m[K, [319]b]()
	case 320:
		r = _m[K, [320]b]()
	case 321:
		r = _m[K, [321]b]()
	case 322:
		r = _m[K, [322]b]()
	case 323:
		r = _m[K, [323]b]()
	case 324:
		r = _m[K, [324]b]()
	case 325:
		r = _m[K, [325]b]()
	case 326:
		r = _m[K, [326]b]()
	case 327:
		r = _m[K, [327]b]()
	case 328:
		r = _m[K, [328]b]()
	case 329:
		r = _m[K, [329]b]()
	case 330:
		r = _m[K, [330]b]()
	case 331:
		r = _m[K, [331]b]()
	case 332:
		r = _m[K, [332]b]()
	case 333:
		r = _m[K, [333]b]()
	case 334:
		r = _m[K, [334]b]()
	case 335:
		r = _m[K, [335]b]()
	case 336:
		r = _m[K, [336]b]()
	case 337:
		r = _m[K, [337]b]()
	case 338:
		r = _m[K, [338]b]()
	case 339:
		r = _m[K, [339]b]()
	case 340:
		r = _m[K, [340]b]()
	case 341:
		r = _m[K, [341]b]()
	case 342:
		r = _m[K, [342]b]()
	case 343:
		r = _m[K, [343]b]()
	case 344:
		r = _m[K, [344]b]()
	case 345:
		r = _m[K, [345]b]()
	case 346:
		r = _m[K, [346]b]()
	case 347:
		r = _m[K, [347]b]()
	case 348:
		r = _m[K, [348]b]()
	case 349:
		r = _m[K, [349]b]()
	case 350:
		r = _m[K, [350]b]()
	case 351:
		r = _m[K, [351]b]()
	case 352:
		r = _m[K, [352]b]()
	case 353:
		r = _m[K, [353]b]()
	case 354:
		r = _m[K, [354]b]()
	case 355:
		r = _m[K, [355]b]()
	case 356:
		r = _m[K, [356]b]()
	case 357:
		r = _m[K, [357]b]()
	case 358:
		r = _m[K, [358]b]()
	case 359:
		r = _m[K, [359]b]()
	case 360:
		r = _m[K, [360]b]()
	case 361:
		r = _m[K, [361]b]()
	case 362:
		r = _m[K, [362]b]()
	case 363:
		r = _m[K, [363]b]()
	case 364:
		r = _m[K, [364]b]()
	case 365:
		r = _m[K, [365]b]()
	case 366:
		r = _m[K, [366]b]()
	case 367:
		r = _m[K, [367]b]()
	case 368:
		r = _m[K, [368]b]()
	case 369:
		r = _m[K, [369]b]()
	case 370:
		r = _m[K, [370]b]()
	case 371:
		r = _m[K, [371]b]()
	case 372:
		r = _m[K, [372]b]()
	case 373:
		r = _m[K, [373]b]()
	case 374:
		r = _m[K, [374]b]()
	case 375:
		r = _m[K, [375]b]()
	case 376:
		r = _m[K, [376]b]()
	case 377:
		r = _m[K, [377]b]()
	case 378:
		r = _m[K, [378]b]()
	case 379:
		r = _m[K, [379]b]()
	case 380:
		r = _m[K, [380]b]()
	case 381:
		r = _m[K, [381]b]()
	case 382:
		r = _m[K, [382]b]()
	case 383:
		r = _m[K, [383]b]()
	case 384:
		r = _m[K, [384]b]()
	case 385:
		r = _m[K, [385]b]()
	case 386:
		r = _m[K, [386]b]()
	case 387:
		r = _m[K, [387]b]()
	case 388:
		r = _m[K, [388]b]()
	case 389:
		r = _m[K, [389]b]()
	case 390:
		r = _m[K, [390]b]()
	case 391:
		r = _m[K, [391]b]()
	case 392:
		r = _m[K, [392]b]()
	case 393:
		r = _m[K, [393]b]()
	case 394:
		r = _m[K, [394]b]()
	case 395:
		r = _m[K, [395]b]()
	case 396:
		r = _m[K, [396]b]()
	case 397:
		r = _m[K, [397]b]()
	case 398:
		r = _m[K, [398]b]()
	case 399:
		r = _m[K, [399]b]()
	case 400:
		r = _m[K, [400]b]()
	case 401:
		r = _m[K, [401]b]()
	case 402:
		r = _m[K, [402]b]()
	case 403:
		r = _m[K, [403]b]()
	case 404:
		r = _m[K, [404]b]()
	case 405:
		r = _m[K, [405]b]()
	case 406:
		r = _m[K, [406]b]()
	case 407:
		r = _m[K, [407]b]()
	case 408:
		r = _m[K, [408]b]()
	case 409:
		r = _m[K, [409]b]()
	case 410:
		r = _m[K, [410]b]()
	case 411:
		r = _m[K, [411]b]()
	case 412:
		r = _m[K, [412]b]()
	case 413:
		r = _m[K, [413]b]()
	case 414:
		r = _m[K, [414]b]()
	case 415:
		r = _m[K, [415]b]()
	case 416:
		r = _m[K, [416]b]()
	case 417:
		r = _m[K, [417]b]()
	case 418:
		r = _m[K, [418]b]()
	case 419:
		r = _m[K, [419]b]()
	case 420:
		r = _m[K, [420]b]()
	case 421:
		r = _m[K, [421]b]()
	case 422:
		r = _m[K, [422]b]()
	case 423:
		r = _m[K, [423]b]()
	case 424:
		r = _m[K, [424]b]()
	case 425:
		r = _m[K, [425]b]()
	case 426:
		r = _m[K, [426]b]()
	case 427:
		r = _m[K, [427]b]()
	case 428:
		r = _m[K, [428]b]()
	case 429:
		r = _m[K, [429]b]()
	case 430:
		r = _m[K, [430]b]()
	case 431:
		r = _m[K, [431]b]()
	case 432:
		r = _m[K, [432]b]()
	case 433:
		r = _m[K, [433]b]()
	case 434:
		r = _m[K, [434]b]()
	case 435:
		r = _m[K, [435]b]()
	case 436:
		r = _m[K, [436]b]()
	case 437:
		r = _m[K, [437]b]()
	case 438:
		r = _m[K, [438]b]()
	case 439:
		r = _m[K, [439]b]()
	case 440:
		r = _m[K, [440]b]()
	case 441:
		r = _m[K, [441]b]()
	case 442:
		r = _m[K, [442]b]()
	case 443:
		r = _m[K, [443]b]()
	case 444:
		r = _m[K, [444]b]()
	case 445:
		r = _m[K, [445]b]()
	case 446:
		r = _m[K, [446]b]()
	case 447:
		r = _m[K, [447]b]()
	case 448:
		r = _m[K, [448]b]()
	case 449:
		r = _m[K, [449]b]()
	case 450:
		r = _m[K, [450]b]()
	case 451:
		r = _m[K, [451]b]()
	case 452:
		r = _m[K, [452]b]()
	case 453:
		r = _m[K, [453]b]()
	case 454:
		r = _m[K, [454]b]()
	case 455:
		r = _m[K, [455]b]()
	case 456:
		r = _m[K, [456]b]()
	case 457:
		r = _m[K, [457]b]()
	case 458:
		r = _m[K, [458]b]()
	case 459:
		r = _m[K, [459]b]()
	case 460:
		r = _m[K, [460]b]()
	case 461:
		r = _m[K, [461]b]()
	case 462:
		r = _m[K, [462]b]()
	case 463:
		r = _m[K, [463]b]()
	case 464:
		r = _m[K, [464]b]()
	case 465:
		r = _m[K, [465]b]()
	case 466:
		r = _m[K, [466]b]()
	case 467:
		r = _m[K, [467]b]()
	case 468:
		r = _m[K, [468]b]()
	case 469:
		r = _m[K, [469]b]()
	case 470:
		r = _m[K, [470]b]()
	case 471:
		r = _m[K, [471]b]()
	case 472:
		r = _m[K, [472]b]()
	case 473:
		r = _m[K, [473]b]()
	case 474:
		r = _m[K, [474]b]()
	case 475:
		r = _m[K, [475]b]()
	case 476:
		r = _m[K, [476]b]()
	case 477:
		r = _m[K, [477]b]()
	case 478:
		r = _m[K, [478]b]()
	case 479:
		r = _m[K, [479]b]()
	case 480:
		r = _m[K, [480]b]()
	case 481:
		r = _m[K, [481]b]()
	case 482:
		r = _m[K, [482]b]()
	case 483:
		r = _m[K, [483]b]()
	case 484:
		r = _m[K, [484]b]()
	case 485:
		r = _m[K, [485]b]()
	case 486:
		r = _m[K, [486]b]()
	case 487:
		r = _m[K, [487]b]()
	case 488:
		r = _m[K, [488]b]()
	case 489:
		r = _m[K, [489]b]()
	case 490:
		r = _m[K, [490]b]()
	case 491:
		r = _m[K, [491]b]()
	case 492:
		r = _m[K, [492]b]()
	case 493:
		r = _m[K, [493]b]()
	case 494:
		r = _m[K, [494]b]()
	case 495:
		r = _m[K, [495]b]()
	case 496:
		r = _m[K, [496]b]()
	case 497:
		r = _m[K, [497]b]()
	case 498:
		r = _m[K, [498]b]()
	case 499:
		r = _m[K, [499]b]()
	case 500:
		r = _m[K, [500]b]()
	case 501:
		r = _m[K, [501]b]()
	case 502:
		r = _m[K, [502]b]()
	case 503:
		r = _m[K, [503]b]()
	case 504:
		r = _m[K, [504]b]()
	case 505:
		r = _m[K, [505]b]()
	case 506:
		r = _m[K, [506]b]()
	case 507:
		r = _m[K, [507]b]()
	case 508:
		r = _m[K, [508]b]()
	case 509:
		r = _m[K, [509]b]()
	case 510:
		r = _m[K, [510]b]()
	case 511:
		r = _m[K, [511]b]()
	case 512:
		r = _m[K, [512]b]()
	case 513:
		r = _m[K, [513]b]()
	case 514:
		r = _m[K, [514]b]()
	case 515:
		r = _m[K, [515]b]()
	case 516:
		r = _m[K, [516]b]()
	case 517:
		r = _m[K, [517]b]()
	case 518:
		r = _m[K, [518]b]()
	case 519:
		r = _m[K, [519]b]()
	case 520:
		r = _m[K, [520]b]()
	case 521:
		r = _m[K, [521]b]()
	case 522:
		r = _m[K, [522]b]()
	case 523:
		r = _m[K, [523]b]()
	case 524:
		r = _m[K, [524]b]()
	case 525:
		r = _m[K, [525]b]()
	case 526:
		r = _m[K, [526]b]()
	case 527:
		r = _m[K, [527]b]()
	case 528:
		r = _m[K, [528]b]()
	case 529:
		r = _m[K, [529]b]()
	case 530:
		r = _m[K, [530]b]()
	case 531:
		r = _m[K, [531]b]()
	case 532:
		r = _m[K, [532]b]()
	case 533:
		r = _m[K, [533]b]()
	case 534:
		r = _m[K, [534]b]()
	case 535:
		r = _m[K, [535]b]()
	case 536:
		r = _m[K, [536]b]()
	case 537:
		r = _m[K, [537]b]()
	case 538:
		r = _m[K, [538]b]()
	case 539:
		r = _m[K, [539]b]()
	case 540:
		r = _m[K, [540]b]()
	case 541:
		r = _m[K, [541]b]()
	case 542:
		r = _m[K, [542]b]()
	case 543:
		r = _m[K, [543]b]()
	case 544:
		r = _m[K, [544]b]()
	case 545:
		r = _m[K, [545]b]()
	case 546:
		r = _m[K, [546]b]()
	case 547:
		r = _m[K, [547]b]()
	case 548:
		r = _m[K, [548]b]()
	case 549:
		r = _m[K, [549]b]()
	case 550:
		r = _m[K, [550]b]()
	case 551:
		r = _m[K, [551]b]()
	case 552:
		r = _m[K, [552]b]()
	case 553:
		r = _m[K, [553]b]()
	case 554:
		r = _m[K, [554]b]()
	case 555:
		r = _m[K, [555]b]()
	case 556:
		r = _m[K, [556]b]()
	case 557:
		r = _m[K, [557]b]()
	case 558:
		r = _m[K, [558]b]()
	case 559:
		r = _m[K, [559]b]()
	case 560:
		r = _m[K, [560]b]()
	case 561:
		r = _m[K, [561]b]()
	case 562:
		r = _m[K, [562]b]()
	case 563:
		r = _m[K, [563]b]()
	case 564:
		r = _m[K, [564]b]()
	case 565:
		r = _m[K, [565]b]()
	case 566:
		r = _m[K, [566]b]()
	case 567:
		r = _m[K, [567]b]()
	case 568:
		r = _m[K, [568]b]()
	case 569:
		r = _m[K, [569]b]()
	case 570:
		r = _m[K, [570]b]()
	case 571:
		r = _m[K, [571]b]()
	case 572:
		r = _m[K, [572]b]()
	case 573:
		r = _m[K, [573]b]()
	case 574:
		r = _m[K, [574]b]()
	case 575:
		r = _m[K, [575]b]()
	case 576:
		r = _m[K, [576]b]()
	case 577:
		r = _m[K, [577]b]()
	case 578:
		r = _m[K, [578]b]()
	case 579:
		r = _m[K, [579]b]()
	case 580:
		r = _m[K, [580]b]()
	case 581:
		r = _m[K, [581]b]()
	case 582:
		r = _m[K, [582]b]()
	case 583:
		r = _m[K, [583]b]()
	case 584:
		r = _m[K, [584]b]()
	case 585:
		r = _m[K, [585]b]()
	case 586:
		r = _m[K, [586]b]()
	case 587:
		r = _m[K, [587]b]()
	case 588:
		r = _m[K, [588]b]()
	case 589:
		r = _m[K, [589]b]()
	case 590:
		r = _m[K, [590]b]()
	case 591:
		r = _m[K, [591]b]()
	case 592:
		r = _m[K, [592]b]()
	case 593:
		r = _m[K, [593]b]()
	case 594:
		r = _m[K, [594]b]()
	case 595:
		r = _m[K, [595]b]()
	case 596:
		r = _m[K, [596]b]()
	case 597:
		r = _m[K, [597]b]()
	case 598:
		r = _m[K, [598]b]()
	case 599:
		r = _m[K, [599]b]()
	case 600:
		r = _m[K, [600]b]()
	case 601:
		r = _m[K, [601]b]()
	case 602:
		r = _m[K, [602]b]()
	case 603:
		r = _m[K, [603]b]()
	case 604:
		r = _m[K, [604]b]()
	case 605:
		r = _m[K, [605]b]()
	case 606:
		r = _m[K, [606]b]()
	case 607:
		r = _m[K, [607]b]()
	case 608:
		r = _m[K, [608]b]()
	case 609:
		r = _m[K, [609]b]()
	case 610:
		r = _m[K, [610]b]()
	case 611:
		r = _m[K, [611]b]()
	case 612:
		r = _m[K, [612]b]()
	case 613:
		r = _m[K, [613]b]()
	case 614:
		r = _m[K, [614]b]()
	case 615:
		r = _m[K, [615]b]()
	case 616:
		r = _m[K, [616]b]()
	case 617:
		r = _m[K, [617]b]()
	case 618:
		r = _m[K, [618]b]()
	case 619:
		r = _m[K, [619]b]()
	case 620:
		r = _m[K, [620]b]()
	case 621:
		r = _m[K, [621]b]()
	case 622:
		r = _m[K, [622]b]()
	case 623:
		r = _m[K, [623]b]()
	case 624:
		r = _m[K, [624]b]()
	case 625:
		r = _m[K, [625]b]()
	case 626:
		r = _m[K, [626]b]()
	case 627:
		r = _m[K, [627]b]()
	case 628:
		r = _m[K, [628]b]()
	case 629:
		r = _m[K, [629]b]()
	case 630:
		r = _m[K, [630]b]()
	case 631:
		r = _m[K, [631]b]()
	case 632:
		r = _m[K, [632]b]()
	case 633:
		r = _m[K, [633]b]()
	case 634:
		r = _m[K, [634]b]()
	case 635:
		r = _m[K, [635]b]()
	case 636:
		r = _m[K, [636]b]()
	case 637:
		r = _m[K, [637]b]()
	case 638:
		r = _m[K, [638]b]()
	case 639:
		r = _m[K, [639]b]()
	case 640:
		r = _m[K, [640]b]()
	case 641:
		r = _m[K, [641]b]()
	case 642:
		r = _m[K, [642]b]()
	case 643:
		r = _m[K, [643]b]()
	case 644:
		r = _m[K, [644]b]()
	case 645:
		r = _m[K, [645]b]()
	case 646:
		r = _m[K, [646]b]()
	case 647:
		r = _m[K, [647]b]()
	case 648:
		r = _m[K, [648]b]()
	case 649:
		r = _m[K, [649]b]()
	case 650:
		r = _m[K, [650]b]()
	case 651:
		r = _m[K, [651]b]()
	case 652:
		r = _m[K, [652]b]()
	case 653:
		r = _m[K, [653]b]()
	case 654:
		r = _m[K, [654]b]()
	case 655:
		r = _m[K, [655]b]()
	case 656:
		r = _m[K, [656]b]()
	case 657:
		r = _m[K, [657]b]()
	case 658:
		r = _m[K, [658]b]()
	case 659:
		r = _m[K, [659]b]()
	case 660:
		r = _m[K, [660]b]()
	case 661:
		r = _m[K, [661]b]()
	case 662:
		r = _m[K, [662]b]()
	case 663:
		r = _m[K, [663]b]()
	case 664:
		r = _m[K, [664]b]()
	case 665:
		r = _m[K, [665]b]()
	case 666:
		r = _m[K, [666]b]()
	case 667:
		r = _m[K, [667]b]()
	case 668:
		r = _m[K, [668]b]()
	case 669:
		r = _m[K, [669]b]()
	case 670:
		r = _m[K, [670]b]()
	case 671:
		r = _m[K, [671]b]()
	case 672:
		r = _m[K, [672]b]()
	case 673:
		r = _m[K, [673]b]()
	case 674:
		r = _m[K, [674]b]()
	case 675:
		r = _m[K, [675]b]()
	case 676:
		r = _m[K, [676]b]()
	case 677:
		r = _m[K, [677]b]()
	case 678:
		r = _m[K, [678]b]()
	case 679:
		r = _m[K, [679]b]()
	case 680:
		r = _m[K, [680]b]()
	case 681:
		r = _m[K, [681]b]()
	case 682:
		r = _m[K, [682]b]()
	case 683:
		r = _m[K, [683]b]()
	case 684:
		r = _m[K, [684]b]()
	case 685:
		r = _m[K, [685]b]()
	case 686:
		r = _m[K, [686]b]()
	case 687:
		r = _m[K, [687]b]()
	case 688:
		r = _m[K, [688]b]()
	case 689:
		r = _m[K, [689]b]()
	case 690:
		r = _m[K, [690]b]()
	case 691:
		r = _m[K, [691]b]()
	case 692:
		r = _m[K, [692]b]()
	case 693:
		r = _m[K, [693]b]()
	case 694:
		r = _m[K, [694]b]()
	case 695:
		r = _m[K, [695]b]()
	case 696:
		r = _m[K, [696]b]()
	case 697:
		r = _m[K, [697]b]()
	case 698:
		r = _m[K, [698]b]()
	case 699:
		r = _m[K, [699]b]()
	case 700:
		r = _m[K, [700]b]()
	case 701:
		r = _m[K, [701]b]()
	case 702:
		r = _m[K, [702]b]()
	case 703:
		r = _m[K, [703]b]()
	case 704:
		r = _m[K, [704]b]()
	case 705:
		r = _m[K, [705]b]()
	case 706:
		r = _m[K, [706]b]()
	case 707:
		r = _m[K, [707]b]()
	case 708:
		r = _m[K, [708]b]()
	case 709:
		r = _m[K, [709]b]()
	case 710:
		r = _m[K, [710]b]()
	case 711:
		r = _m[K, [711]b]()
	case 712:
		r = _m[K, [712]b]()
	case 713:
		r = _m[K, [713]b]()
	case 714:
		r = _m[K, [714]b]()
	case 715:
		r = _m[K, [715]b]()
	case 716:
		r = _m[K, [716]b]()
	case 717:
		r = _m[K, [717]b]()
	case 718:
		r = _m[K, [718]b]()
	case 719:
		r = _m[K, [719]b]()
	case 720:
		r = _m[K, [720]b]()
	case 721:
		r = _m[K, [721]b]()
	case 722:
		r = _m[K, [722]b]()
	case 723:
		r = _m[K, [723]b]()
	case 724:
		r = _m[K, [724]b]()
	case 725:
		r = _m[K, [725]b]()
	case 726:
		r = _m[K, [726]b]()
	case 727:
		r = _m[K, [727]b]()
	case 728:
		r = _m[K, [728]b]()
	case 729:
		r = _m[K, [729]b]()
	case 730:
		r = _m[K, [730]b]()
	case 731:
		r = _m[K, [731]b]()
	case 732:
		r = _m[K, [732]b]()
	case 733:
		r = _m[K, [733]b]()
	case 734:
		r = _m[K, [734]b]()
	case 735:
		r = _m[K, [735]b]()
	case 736:
		r = _m[K, [736]b]()
	case 737:
		r = _m[K, [737]b]()
	case 738:
		r = _m[K, [738]b]()
	case 739:
		r = _m[K, [739]b]()
	case 740:
		r = _m[K, [740]b]()
	case 741:
		r = _m[K, [741]b]()
	case 742:
		r = _m[K, [742]b]()
	case 743:
		r = _m[K, [743]b]()
	case 744:
		r = _m[K, [744]b]()
	case 745:
		r = _m[K, [745]b]()
	case 746:
		r = _m[K, [746]b]()
	case 747:
		r = _m[K, [747]b]()
	case 748:
		r = _m[K, [748]b]()
	case 749:
		r = _m[K, [749]b]()
	case 750:
		r = _m[K, [750]b]()
	case 751:
		r = _m[K, [751]b]()
	case 752:
		r = _m[K, [752]b]()
	case 753:
		r = _m[K, [753]b]()
	case 754:
		r = _m[K, [754]b]()
	case 755:
		r = _m[K, [755]b]()
	case 756:
		r = _m[K, [756]b]()
	case 757:
		r = _m[K, [757]b]()
	case 758:
		r = _m[K, [758]b]()
	case 759:
		r = _m[K, [759]b]()
	case 760:
		r = _m[K, [760]b]()
	case 761:
		r = _m[K, [761]b]()
	case 762:
		r = _m[K, [762]b]()
	case 763:
		r = _m[K, [763]b]()
	case 764:
		r = _m[K, [764]b]()
	case 765:
		r = _m[K, [765]b]()
	case 766:
		r = _m[K, [766]b]()
	case 767:
		r = _m[K, [767]b]()
	case 768:
		r = _m[K, [768]b]()
	case 769:
		r = _m[K, [769]b]()
	case 770:
		r = _m[K, [770]b]()
	case 771:
		r = _m[K, [771]b]()
	case 772:
		r = _m[K, [772]b]()
	case 773:
		r = _m[K, [773]b]()
	case 774:
		r = _m[K, [774]b]()
	case 775:
		r = _m[K, [775]b]()
	case 776:
		r = _m[K, [776]b]()
	case 777:
		r = _m[K, [777]b]()
	case 778:
		r = _m[K, [778]b]()
	case 779:
		r = _m[K, [779]b]()
	case 780:
		r = _m[K, [780]b]()
	case 781:
		r = _m[K, [781]b]()
	case 782:
		r = _m[K, [782]b]()
	case 783:
		r = _m[K, [783]b]()
	case 784:
		r = _m[K, [784]b]()
	case 785:
		r = _m[K, [785]b]()
	case 786:
		r = _m[K, [786]b]()
	case 787:
		r = _m[K, [787]b]()
	case 788:
		r = _m[K, [788]b]()
	case 789:
		r = _m[K, [789]b]()
	case 790:
		r = _m[K, [790]b]()
	case 791:
		r = _m[K, [791]b]()
	case 792:
		r = _m[K, [792]b]()
	case 793:
		r = _m[K, [793]b]()
	case 794:
		r = _m[K, [794]b]()
	case 795:
		r = _m[K, [795]b]()
	case 796:
		r = _m[K, [796]b]()
	case 797:
		r = _m[K, [797]b]()
	case 798:
		r = _m[K, [798]b]()
	case 799:
		r = _m[K, [799]b]()
	case 800:
		r = _m[K, [800]b]()
	case 801:
		r = _m[K, [801]b]()
	case 802:
		r = _m[K, [802]b]()
	case 803:
		r = _m[K, [803]b]()
	case 804:
		r = _m[K, [804]b]()
	case 805:
		r = _m[K, [805]b]()
	case 806:
		r = _m[K, [806]b]()
	case 807:
		r = _m[K, [807]b]()
	case 808:
		r = _m[K, [808]b]()
	case 809:
		r = _m[K, [809]b]()
	case 810:
		r = _m[K, [810]b]()
	case 811:
		r = _m[K, [811]b]()
	case 812:
		r = _m[K, [812]b]()
	case 813:
		r = _m[K, [813]b]()
	case 814:
		r = _m[K, [814]b]()
	case 815:
		r = _m[K, [815]b]()
	case 816:
		r = _m[K, [816]b]()
	case 817:
		r = _m[K, [817]b]()
	case 818:
		r = _m[K, [818]b]()
	case 819:
		r = _m[K, [819]b]()
	case 820:
		r = _m[K, [820]b]()
	case 821:
		r = _m[K, [821]b]()
	case 822:
		r = _m[K, [822]b]()
	case 823:
		r = _m[K, [823]b]()
	case 824:
		r = _m[K, [824]b]()
	case 825:
		r = _m[K, [825]b]()
	case 826:
		r = _m[K, [826]b]()
	case 827:
		r = _m[K, [827]b]()
	case 828:
		r = _m[K, [828]b]()
	case 829:
		r = _m[K, [829]b]()
	case 830:
		r = _m[K, [830]b]()
	case 831:
		r = _m[K, [831]b]()
	case 832:
		r = _m[K, [832]b]()
	case 833:
		r = _m[K, [833]b]()
	case 834:
		r = _m[K, [834]b]()
	case 835:
		r = _m[K, [835]b]()
	case 836:
		r = _m[K, [836]b]()
	case 837:
		r = _m[K, [837]b]()
	case 838:
		r = _m[K, [838]b]()
	case 839:
		r = _m[K, [839]b]()
	case 840:
		r = _m[K, [840]b]()
	case 841:
		r = _m[K, [841]b]()
	case 842:
		r = _m[K, [842]b]()
	case 843:
		r = _m[K, [843]b]()
	case 844:
		r = _m[K, [844]b]()
	case 845:
		r = _m[K, [845]b]()
	case 846:
		r = _m[K, [846]b]()
	case 847:
		r = _m[K, [847]b]()
	case 848:
		r = _m[K, [848]b]()
	case 849:
		r = _m[K, [849]b]()
	case 850:
		r = _m[K, [850]b]()
	case 851:
		r = _m[K, [851]b]()
	case 852:
		r = _m[K, [852]b]()
	case 853:
		r = _m[K, [853]b]()
	case 854:
		r = _m[K, [854]b]()
	case 855:
		r = _m[K, [855]b]()
	case 856:
		r = _m[K, [856]b]()
	case 857:
		r = _m[K, [857]b]()
	case 858:
		r = _m[K, [858]b]()
	case 859:
		r = _m[K, [859]b]()
	case 860:
		r = _m[K, [860]b]()
	case 861:
		r = _m[K, [861]b]()
	case 862:
		r = _m[K, [862]b]()
	case 863:
		r = _m[K, [863]b]()
	case 864:
		r = _m[K, [864]b]()
	case 865:
		r = _m[K, [865]b]()
	case 866:
		r = _m[K, [866]b]()
	case 867:
		r = _m[K, [867]b]()
	case 868:
		r = _m[K, [868]b]()
	case 869:
		r = _m[K, [869]b]()
	case 870:
		r = _m[K, [870]b]()
	case 871:
		r = _m[K, [871]b]()
	case 872:
		r = _m[K, [872]b]()
	case 873:
		r = _m[K, [873]b]()
	case 874:
		r = _m[K, [874]b]()
	case 875:
		r = _m[K, [875]b]()
	case 876:
		r = _m[K, [876]b]()
	case 877:
		r = _m[K, [877]b]()
	case 878:
		r = _m[K, [878]b]()
	case 879:
		r = _m[K, [879]b]()
	case 880:
		r = _m[K, [880]b]()
	case 881:
		r = _m[K, [881]b]()
	case 882:
		r = _m[K, [882]b]()
	case 883:
		r = _m[K, [883]b]()
	case 884:
		r = _m[K, [884]b]()
	case 885:
		r = _m[K, [885]b]()
	case 886:
		r = _m[K, [886]b]()
	case 887:
		r = _m[K, [887]b]()
	case 888:
		r = _m[K, [888]b]()
	case 889:
		r = _m[K, [889]b]()
	case 890:
		r = _m[K, [890]b]()
	case 891:
		r = _m[K, [891]b]()
	case 892:
		r = _m[K, [892]b]()
	case 893:
		r = _m[K, [893]b]()
	case 894:
		r = _m[K, [894]b]()
	case 895:
		r = _m[K, [895]b]()
	case 896:
		r = _m[K, [896]b]()
	case 897:
		r = _m[K, [897]b]()
	case 898:
		r = _m[K, [898]b]()
	case 899:
		r = _m[K, [899]b]()
	case 900:
		r = _m[K, [900]b]()
	case 901:
		r = _m[K, [901]b]()
	case 902:
		r = _m[K, [902]b]()
	case 903:
		r = _m[K, [903]b]()
	case 904:
		r = _m[K, [904]b]()
	case 905:
		r = _m[K, [905]b]()
	case 906:
		r = _m[K, [906]b]()
	case 907:
		r = _m[K, [907]b]()
	case 908:
		r = _m[K, [908]b]()
	case 909:
		r = _m[K, [909]b]()
	case 910:
		r = _m[K, [910]b]()
	case 911:
		r = _m[K, [911]b]()
	case 912:
		r = _m[K, [912]b]()
	case 913:
		r = _m[K, [913]b]()
	case 914:
		r = _m[K, [914]b]()
	case 915:
		r = _m[K, [915]b]()
	case 916:
		r = _m[K, [916]b]()
	case 917:
		r = _m[K, [917]b]()
	case 918:
		r = _m[K, [918]b]()
	case 919:
		r = _m[K, [919]b]()
	case 920:
		r = _m[K, [920]b]()
	case 921:
		r = _m[K, [921]b]()
	case 922:
		r = _m[K, [922]b]()
	case 923:
		r = _m[K, [923]b]()
	case 924:
		r = _m[K, [924]b]()
	case 925:
		r = _m[K, [925]b]()
	case 926:
		r = _m[K, [926]b]()
	case 927:
		r = _m[K, [927]b]()
	case 928:
		r = _m[K, [928]b]()
	case 929:
		r = _m[K, [929]b]()
	case 930:
		r = _m[K, [930]b]()
	case 931:
		r = _m[K, [931]b]()
	case 932:
		r = _m[K, [932]b]()
	case 933:
		r = _m[K, [933]b]()
	case 934:
		r = _m[K, [934]b]()
	case 935:
		r = _m[K, [935]b]()
	case 936:
		r = _m[K, [936]b]()
	case 937:
		r = _m[K, [937]b]()
	case 938:
		r = _m[K, [938]b]()
	case 939:
		r = _m[K, [939]b]()
	case 940:
		r = _m[K, [940]b]()
	case 941:
		r = _m[K, [941]b]()
	case 942:
		r = _m[K, [942]b]()
	case 943:
		r = _m[K, [943]b]()
	case 944:
		r = _m[K, [944]b]()
	case 945:
		r = _m[K, [945]b]()
	case 946:
		r = _m[K, [946]b]()
	case 947:
		r = _m[K, [947]b]()
	case 948:
		r = _m[K, [948]b]()
	case 949:
		r = _m[K, [949]b]()
	case 950:
		r = _m[K, [950]b]()
	case 951:
		r = _m[K, [951]b]()
	case 952:
		r = _m[K, [952]b]()
	case 953:
		r = _m[K, [953]b]()
	case 954:
		r = _m[K, [954]b]()
	case 955:
		r = _m[K, [955]b]()
	case 956:
		r = _m[K, [956]b]()
	case 957:
		r = _m[K, [957]b]()
	case 958:
		r = _m[K, [958]b]()
	case 959:
		r = _m[K, [959]b]()
	case 960:
		r = _m[K, [960]b]()
	case 961:
		r = _m[K, [961]b]()
	case 962:
		r = _m[K, [962]b]()
	case 963:
		r = _m[K, [963]b]()
	case 964:
		r = _m[K, [964]b]()
	case 965:
		r = _m[K, [965]b]()
	case 966:
		r = _m[K, [966]b]()
	case 967:
		r = _m[K, [967]b]()
	case 968:
		r = _m[K, [968]b]()
	case 969:
		r = _m[K, [969]b]()
	case 970:
		r = _m[K, [970]b]()
	case 971:
		r = _m[K, [971]b]()
	case 972:
		r = _m[K, [972]b]()
	case 973:
		r = _m[K, [973]b]()
	case 974:
		r = _m[K, [974]b]()
	case 975:
		r = _m[K, [975]b]()
	case 976:
		r = _m[K, [976]b]()
	case 977:
		r = _m[K, [977]b]()
	case 978:
		r = _m[K, [978]b]()
	case 979:
		r = _m[K, [979]b]()
	case 980:
		r = _m[K, [980]b]()
	case 981:
		r = _m[K, [981]b]()
	case 982:
		r = _m[K, [982]b]()
	case 983:
		r = _m[K, [983]b]()
	case 984:
		r = _m[K, [984]b]()
	case 985:
		r = _m[K, [985]b]()
	case 986:
		r = _m[K, [986]b]()
	case 987:
		r = _m[K, [987]b]()
	case 988:
		r = _m[K, [988]b]()
	case 989:
		r = _m[K, [989]b]()
	case 990:
		r = _m[K, [990]b]()
	case 991:
		r = _m[K, [991]b]()
	case 992:
		r = _m[K, [992]b]()
	case 993:
		r = _m[K, [993]b]()
	case 994:
		r = _m[K, [994]b]()
	case 995:
		r = _m[K, [995]b]()
	case 996:
		r = _m[K, [996]b]()
	case 997:
		r = _m[K, [997]b]()
	case 998:
		r = _m[K, [998]b]()
	case 999:
		r = _m[K, [999]b]()
	case 1000:
		r = _m[K, [1000]b]()
	case 1001:
		r = _m[K, [1001]b]()
	case 1002:
		r = _m[K, [1002]b]()
	case 1003:
		r = _m[K, [1003]b]()
	case 1004:
		r = _m[K, [1004]b]()
	case 1005:
		r = _m[K, [1005]b]()
	case 1006:
		r = _m[K, [1006]b]()
	case 1007:
		r = _m[K, [1007]b]()
	case 1008:
		r = _m[K, [1008]b]()
	case 1009:
		r = _m[K, [1009]b]()
	case 1010:
		r = _m[K, [1010]b]()
	case 1011:
		r = _m[K, [1011]b]()
	case 1012:
		r = _m[K, [1012]b]()
	case 1013:
		r = _m[K, [1013]b]()
	case 1014:
		r = _m[K, [1014]b]()
	case 1015:
		r = _m[K, [1015]b]()
	case 1016:
		r = _m[K, [1016]b]()
	case 1017:
		r = _m[K, [1017]b]()
	case 1018:
		r = _m[K, [1018]b]()
	case 1019:
		r = _m[K, [1019]b]()
	case 1020:
		r = _m[K, [1020]b]()
	case 1021:
		r = _m[K, [1021]b]()
	case 1022:
		r = _m[K, [1022]b]()
	case 1023:
		r = _m[K, [1023]b]()
	case 1024:
		r = _m[K, [1024]b]()
	default:
	}

	return
}

func getMapToolkitForStructKeyAndValueByReflection(mapType reflect.Type) (r *mapToolkit) {
	keySize := mapType.Key().Size()
	e := mapType.Elem()
	switch keySize {
	case 1:
		r = _o[[1]b](e)
	case 2:
		r = _o[[2]b](e)
	case 3:
		r = _o[[3]b](e)
	case 4:
		r = _o[[4]b](e)
	case 5:
		r = _o[[5]b](e)
	case 6:
		r = _o[[6]b](e)
	case 7:
		r = _o[[7]b](e)
	case 8:
		r = _o[[8]b](e)
	case 9:
		r = _o[[9]b](e)
	case 10:
		r = _o[[10]b](e)
	case 11:
		r = _o[[11]b](e)
	case 12:
		r = _o[[12]b](e)
	case 13:
		r = _o[[13]b](e)
	case 14:
		r = _o[[14]b](e)
	case 15:
		r = _o[[15]b](e)
	case 16:
		r = _o[[16]b](e)
	case 17:
		r = _o[[17]b](e)
	case 18:
		r = _o[[18]b](e)
	case 19:
		r = _o[[19]b](e)
	case 20:
		r = _o[[20]b](e)
	case 21:
		r = _o[[21]b](e)
	case 22:
		r = _o[[22]b](e)
	case 23:
		r = _o[[23]b](e)
	case 24:
		r = _o[[24]b](e)
	case 25:
		r = _o[[25]b](e)
	case 26:
		r = _o[[26]b](e)
	case 27:
		r = _o[[27]b](e)
	case 28:
		r = _o[[28]b](e)
	case 29:
		r = _o[[29]b](e)
	case 30:
		r = _o[[30]b](e)
	case 31:
		r = _o[[31]b](e)
	case 32:
		r = _o[[32]b](e)
	case 33:
		r = _o[[33]b](e)
	case 34:
		r = _o[[34]b](e)
	case 35:
		r = _o[[35]b](e)
	case 36:
		r = _o[[36]b](e)
	case 37:
		r = _o[[37]b](e)
	case 38:
		r = _o[[38]b](e)
	case 39:
		r = _o[[39]b](e)
	case 40:
		r = _o[[40]b](e)
	case 41:
		r = _o[[41]b](e)
	case 42:
		r = _o[[42]b](e)
	case 43:
		r = _o[[43]b](e)
	case 44:
		r = _o[[44]b](e)
	case 45:
		r = _o[[45]b](e)
	case 46:
		r = _o[[46]b](e)
	case 47:
		r = _o[[47]b](e)
	case 48:
		r = _o[[48]b](e)
	case 49:
		r = _o[[49]b](e)
	case 50:
		r = _o[[50]b](e)
	case 51:
		r = _o[[51]b](e)
	case 52:
		r = _o[[52]b](e)
	case 53:
		r = _o[[53]b](e)
	case 54:
		r = _o[[54]b](e)
	case 55:
		r = _o[[55]b](e)
	case 56:
		r = _o[[56]b](e)
	case 57:
		r = _o[[57]b](e)
	case 58:
		r = _o[[58]b](e)
	case 59:
		r = _o[[59]b](e)
	case 60:
		r = _o[[60]b](e)
	case 61:
		r = _o[[61]b](e)
	case 62:
		r = _o[[62]b](e)
	case 63:
		r = _o[[63]b](e)
	case 64:
		r = _o[[64]b](e)
	case 65:
		r = _o[[65]b](e)
	case 66:
		r = _o[[66]b](e)
	case 67:
		r = _o[[67]b](e)
	case 68:
		r = _o[[68]b](e)
	case 69:
		r = _o[[69]b](e)
	case 70:
		r = _o[[70]b](e)
	case 71:
		r = _o[[71]b](e)
	case 72:
		r = _o[[72]b](e)
	case 73:
		r = _o[[73]b](e)
	case 74:
		r = _o[[74]b](e)
	case 75:
		r = _o[[75]b](e)
	case 76:
		r = _o[[76]b](e)
	case 77:
		r = _o[[77]b](e)
	case 78:
		r = _o[[78]b](e)
	case 79:
		r = _o[[79]b](e)
	case 80:
		r = _o[[80]b](e)
	case 81:
		r = _o[[81]b](e)
	case 82:
		r = _o[[82]b](e)
	case 83:
		r = _o[[83]b](e)
	case 84:
		r = _o[[84]b](e)
	case 85:
		r = _o[[85]b](e)
	case 86:
		r = _o[[86]b](e)
	case 87:
		r = _o[[87]b](e)
	case 88:
		r = _o[[88]b](e)
	case 89:
		r = _o[[89]b](e)
	case 90:
		r = _o[[90]b](e)
	case 91:
		r = _o[[91]b](e)
	case 92:
		r = _o[[92]b](e)
	case 93:
		r = _o[[93]b](e)
	case 94:
		r = _o[[94]b](e)
	case 95:
		r = _o[[95]b](e)
	case 96:
		r = _o[[96]b](e)
	case 97:
		r = _o[[97]b](e)
	case 98:
		r = _o[[98]b](e)
	case 99:
		r = _o[[99]b](e)
	case 100:
		r = _o[[100]b](e)
	case 101:
		r = _o[[101]b](e)
	case 102:
		r = _o[[102]b](e)
	case 103:
		r = _o[[103]b](e)
	case 104:
		r = _o[[104]b](e)
	case 105:
		r = _o[[105]b](e)
	case 106:
		r = _o[[106]b](e)
	case 107:
		r = _o[[107]b](e)
	case 108:
		r = _o[[108]b](e)
	case 109:
		r = _o[[109]b](e)
	case 110:
		r = _o[[110]b](e)
	case 111:
		r = _o[[111]b](e)
	case 112:
		r = _o[[112]b](e)
	case 113:
		r = _o[[113]b](e)
	case 114:
		r = _o[[114]b](e)
	case 115:
		r = _o[[115]b](e)
	case 116:
		r = _o[[116]b](e)
	case 117:
		r = _o[[117]b](e)
	case 118:
		r = _o[[118]b](e)
	case 119:
		r = _o[[119]b](e)
	case 120:
		r = _o[[120]b](e)
	case 121:
		r = _o[[121]b](e)
	case 122:
		r = _o[[122]b](e)
	case 123:
		r = _o[[123]b](e)
	case 124:
		r = _o[[124]b](e)
	case 125:
		r = _o[[125]b](e)
	case 126:
		r = _o[[126]b](e)
	case 127:
		r = _o[[127]b](e)
	case 128:
		r = _o[[128]b](e)
	case 129:
		r = _o[[129]b](e)
	case 130:
		r = _o[[130]b](e)
	case 131:
		r = _o[[131]b](e)
	case 132:
		r = _o[[132]b](e)
	case 133:
		r = _o[[133]b](e)
	case 134:
		r = _o[[134]b](e)
	case 135:
		r = _o[[135]b](e)
	case 136:
		r = _o[[136]b](e)
	case 137:
		r = _o[[137]b](e)
	case 138:
		r = _o[[138]b](e)
	case 139:
		r = _o[[139]b](e)
	case 140:
		r = _o[[140]b](e)
	case 141:
		r = _o[[141]b](e)
	case 142:
		r = _o[[142]b](e)
	case 143:
		r = _o[[143]b](e)
	case 144:
		r = _o[[144]b](e)
	case 145:
		r = _o[[145]b](e)
	case 146:
		r = _o[[146]b](e)
	case 147:
		r = _o[[147]b](e)
	case 148:
		r = _o[[148]b](e)
	case 149:
		r = _o[[149]b](e)
	case 150:
		r = _o[[150]b](e)
	case 151:
		r = _o[[151]b](e)
	case 152:
		r = _o[[152]b](e)
	case 153:
		r = _o[[153]b](e)
	case 154:
		r = _o[[154]b](e)
	case 155:
		r = _o[[155]b](e)
	case 156:
		r = _o[[156]b](e)
	case 157:
		r = _o[[157]b](e)
	case 158:
		r = _o[[158]b](e)
	case 159:
		r = _o[[159]b](e)
	case 160:
		r = _o[[160]b](e)
	case 161:
		r = _o[[161]b](e)
	case 162:
		r = _o[[162]b](e)
	case 163:
		r = _o[[163]b](e)
	case 164:
		r = _o[[164]b](e)
	case 165:
		r = _o[[165]b](e)
	case 166:
		r = _o[[166]b](e)
	case 167:
		r = _o[[167]b](e)
	case 168:
		r = _o[[168]b](e)
	case 169:
		r = _o[[169]b](e)
	case 170:
		r = _o[[170]b](e)
	case 171:
		r = _o[[171]b](e)
	case 172:
		r = _o[[172]b](e)
	case 173:
		r = _o[[173]b](e)
	case 174:
		r = _o[[174]b](e)
	case 175:
		r = _o[[175]b](e)
	case 176:
		r = _o[[176]b](e)
	case 177:
		r = _o[[177]b](e)
	case 178:
		r = _o[[178]b](e)
	case 179:
		r = _o[[179]b](e)
	case 180:
		r = _o[[180]b](e)
	case 181:
		r = _o[[181]b](e)
	case 182:
		r = _o[[182]b](e)
	case 183:
		r = _o[[183]b](e)
	case 184:
		r = _o[[184]b](e)
	case 185:
		r = _o[[185]b](e)
	case 186:
		r = _o[[186]b](e)
	case 187:
		r = _o[[187]b](e)
	case 188:
		r = _o[[188]b](e)
	case 189:
		r = _o[[189]b](e)
	case 190:
		r = _o[[190]b](e)
	case 191:
		r = _o[[191]b](e)
	case 192:
		r = _o[[192]b](e)
	case 193:
		r = _o[[193]b](e)
	case 194:
		r = _o[[194]b](e)
	case 195:
		r = _o[[195]b](e)
	case 196:
		r = _o[[196]b](e)
	case 197:
		r = _o[[197]b](e)
	case 198:
		r = _o[[198]b](e)
	case 199:
		r = _o[[199]b](e)
	case 200:
		r = _o[[200]b](e)
	case 201:
		r = _o[[201]b](e)
	case 202:
		r = _o[[202]b](e)
	case 203:
		r = _o[[203]b](e)
	case 204:
		r = _o[[204]b](e)
	case 205:
		r = _o[[205]b](e)
	case 206:
		r = _o[[206]b](e)
	case 207:
		r = _o[[207]b](e)
	case 208:
		r = _o[[208]b](e)
	case 209:
		r = _o[[209]b](e)
	case 210:
		r = _o[[210]b](e)
	case 211:
		r = _o[[211]b](e)
	case 212:
		r = _o[[212]b](e)
	case 213:
		r = _o[[213]b](e)
	case 214:
		r = _o[[214]b](e)
	case 215:
		r = _o[[215]b](e)
	case 216:
		r = _o[[216]b](e)
	case 217:
		r = _o[[217]b](e)
	case 218:
		r = _o[[218]b](e)
	case 219:
		r = _o[[219]b](e)
	case 220:
		r = _o[[220]b](e)
	case 221:
		r = _o[[221]b](e)
	case 222:
		r = _o[[222]b](e)
	case 223:
		r = _o[[223]b](e)
	case 224:
		r = _o[[224]b](e)
	case 225:
		r = _o[[225]b](e)
	case 226:
		r = _o[[226]b](e)
	case 227:
		r = _o[[227]b](e)
	case 228:
		r = _o[[228]b](e)
	case 229:
		r = _o[[229]b](e)
	case 230:
		r = _o[[230]b](e)
	case 231:
		r = _o[[231]b](e)
	case 232:
		r = _o[[232]b](e)
	case 233:
		r = _o[[233]b](e)
	case 234:
		r = _o[[234]b](e)
	case 235:
		r = _o[[235]b](e)
	case 236:
		r = _o[[236]b](e)
	case 237:
		r = _o[[237]b](e)
	case 238:
		r = _o[[238]b](e)
	case 239:
		r = _o[[239]b](e)
	case 240:
		r = _o[[240]b](e)
	case 241:
		r = _o[[241]b](e)
	case 242:
		r = _o[[242]b](e)
	case 243:
		r = _o[[243]b](e)
	case 244:
		r = _o[[244]b](e)
	case 245:
		r = _o[[245]b](e)
	case 246:
		r = _o[[246]b](e)
	case 247:
		r = _o[[247]b](e)
	case 248:
		r = _o[[248]b](e)
	case 249:
		r = _o[[249]b](e)
	case 250:
		r = _o[[250]b](e)
	case 251:
		r = _o[[251]b](e)
	case 252:
		r = _o[[252]b](e)
	case 253:
		r = _o[[253]b](e)
	case 254:
		r = _o[[254]b](e)
	case 255:
		r = _o[[255]b](e)
	case 256:
		r = _o[[256]b](e)
	case 257:
		r = _o[[257]b](e)
	case 258:
		r = _o[[258]b](e)
	case 259:
		r = _o[[259]b](e)
	case 260:
		r = _o[[260]b](e)
	case 261:
		r = _o[[261]b](e)
	case 262:
		r = _o[[262]b](e)
	case 263:
		r = _o[[263]b](e)
	case 264:
		r = _o[[264]b](e)
	case 265:
		r = _o[[265]b](e)
	case 266:
		r = _o[[266]b](e)
	case 267:
		r = _o[[267]b](e)
	case 268:
		r = _o[[268]b](e)
	case 269:
		r = _o[[269]b](e)
	case 270:
		r = _o[[270]b](e)
	case 271:
		r = _o[[271]b](e)
	case 272:
		r = _o[[272]b](e)
	case 273:
		r = _o[[273]b](e)
	case 274:
		r = _o[[274]b](e)
	case 275:
		r = _o[[275]b](e)
	case 276:
		r = _o[[276]b](e)
	case 277:
		r = _o[[277]b](e)
	case 278:
		r = _o[[278]b](e)
	case 279:
		r = _o[[279]b](e)
	case 280:
		r = _o[[280]b](e)
	case 281:
		r = _o[[281]b](e)
	case 282:
		r = _o[[282]b](e)
	case 283:
		r = _o[[283]b](e)
	case 284:
		r = _o[[284]b](e)
	case 285:
		r = _o[[285]b](e)
	case 286:
		r = _o[[286]b](e)
	case 287:
		r = _o[[287]b](e)
	case 288:
		r = _o[[288]b](e)
	case 289:
		r = _o[[289]b](e)
	case 290:
		r = _o[[290]b](e)
	case 291:
		r = _o[[291]b](e)
	case 292:
		r = _o[[292]b](e)
	case 293:
		r = _o[[293]b](e)
	case 294:
		r = _o[[294]b](e)
	case 295:
		r = _o[[295]b](e)
	case 296:
		r = _o[[296]b](e)
	case 297:
		r = _o[[297]b](e)
	case 298:
		r = _o[[298]b](e)
	case 299:
		r = _o[[299]b](e)
	case 300:
		r = _o[[300]b](e)
	case 301:
		r = _o[[301]b](e)
	case 302:
		r = _o[[302]b](e)
	case 303:
		r = _o[[303]b](e)
	case 304:
		r = _o[[304]b](e)
	case 305:
		r = _o[[305]b](e)
	case 306:
		r = _o[[306]b](e)
	case 307:
		r = _o[[307]b](e)
	case 308:
		r = _o[[308]b](e)
	case 309:
		r = _o[[309]b](e)
	case 310:
		r = _o[[310]b](e)
	case 311:
		r = _o[[311]b](e)
	case 312:
		r = _o[[312]b](e)
	case 313:
		r = _o[[313]b](e)
	case 314:
		r = _o[[314]b](e)
	case 315:
		r = _o[[315]b](e)
	case 316:
		r = _o[[316]b](e)
	case 317:
		r = _o[[317]b](e)
	case 318:
		r = _o[[318]b](e)
	case 319:
		r = _o[[319]b](e)
	case 320:
		r = _o[[320]b](e)
	case 321:
		r = _o[[321]b](e)
	case 322:
		r = _o[[322]b](e)
	case 323:
		r = _o[[323]b](e)
	case 324:
		r = _o[[324]b](e)
	case 325:
		r = _o[[325]b](e)
	case 326:
		r = _o[[326]b](e)
	case 327:
		r = _o[[327]b](e)
	case 328:
		r = _o[[328]b](e)
	case 329:
		r = _o[[329]b](e)
	case 330:
		r = _o[[330]b](e)
	case 331:
		r = _o[[331]b](e)
	case 332:
		r = _o[[332]b](e)
	case 333:
		r = _o[[333]b](e)
	case 334:
		r = _o[[334]b](e)
	case 335:
		r = _o[[335]b](e)
	case 336:
		r = _o[[336]b](e)
	case 337:
		r = _o[[337]b](e)
	case 338:
		r = _o[[338]b](e)
	case 339:
		r = _o[[339]b](e)
	case 340:
		r = _o[[340]b](e)
	case 341:
		r = _o[[341]b](e)
	case 342:
		r = _o[[342]b](e)
	case 343:
		r = _o[[343]b](e)
	case 344:
		r = _o[[344]b](e)
	case 345:
		r = _o[[345]b](e)
	case 346:
		r = _o[[346]b](e)
	case 347:
		r = _o[[347]b](e)
	case 348:
		r = _o[[348]b](e)
	case 349:
		r = _o[[349]b](e)
	case 350:
		r = _o[[350]b](e)
	case 351:
		r = _o[[351]b](e)
	case 352:
		r = _o[[352]b](e)
	case 353:
		r = _o[[353]b](e)
	case 354:
		r = _o[[354]b](e)
	case 355:
		r = _o[[355]b](e)
	case 356:
		r = _o[[356]b](e)
	case 357:
		r = _o[[357]b](e)
	case 358:
		r = _o[[358]b](e)
	case 359:
		r = _o[[359]b](e)
	case 360:
		r = _o[[360]b](e)
	case 361:
		r = _o[[361]b](e)
	case 362:
		r = _o[[362]b](e)
	case 363:
		r = _o[[363]b](e)
	case 364:
		r = _o[[364]b](e)
	case 365:
		r = _o[[365]b](e)
	case 366:
		r = _o[[366]b](e)
	case 367:
		r = _o[[367]b](e)
	case 368:
		r = _o[[368]b](e)
	case 369:
		r = _o[[369]b](e)
	case 370:
		r = _o[[370]b](e)
	case 371:
		r = _o[[371]b](e)
	case 372:
		r = _o[[372]b](e)
	case 373:
		r = _o[[373]b](e)
	case 374:
		r = _o[[374]b](e)
	case 375:
		r = _o[[375]b](e)
	case 376:
		r = _o[[376]b](e)
	case 377:
		r = _o[[377]b](e)
	case 378:
		r = _o[[378]b](e)
	case 379:
		r = _o[[379]b](e)
	case 380:
		r = _o[[380]b](e)
	case 381:
		r = _o[[381]b](e)
	case 382:
		r = _o[[382]b](e)
	case 383:
		r = _o[[383]b](e)
	case 384:
		r = _o[[384]b](e)
	case 385:
		r = _o[[385]b](e)
	case 386:
		r = _o[[386]b](e)
	case 387:
		r = _o[[387]b](e)
	case 388:
		r = _o[[388]b](e)
	case 389:
		r = _o[[389]b](e)
	case 390:
		r = _o[[390]b](e)
	case 391:
		r = _o[[391]b](e)
	case 392:
		r = _o[[392]b](e)
	case 393:
		r = _o[[393]b](e)
	case 394:
		r = _o[[394]b](e)
	case 395:
		r = _o[[395]b](e)
	case 396:
		r = _o[[396]b](e)
	case 397:
		r = _o[[397]b](e)
	case 398:
		r = _o[[398]b](e)
	case 399:
		r = _o[[399]b](e)
	case 400:
		r = _o[[400]b](e)
	case 401:
		r = _o[[401]b](e)
	case 402:
		r = _o[[402]b](e)
	case 403:
		r = _o[[403]b](e)
	case 404:
		r = _o[[404]b](e)
	case 405:
		r = _o[[405]b](e)
	case 406:
		r = _o[[406]b](e)
	case 407:
		r = _o[[407]b](e)
	case 408:
		r = _o[[408]b](e)
	case 409:
		r = _o[[409]b](e)
	case 410:
		r = _o[[410]b](e)
	case 411:
		r = _o[[411]b](e)
	case 412:
		r = _o[[412]b](e)
	case 413:
		r = _o[[413]b](e)
	case 414:
		r = _o[[414]b](e)
	case 415:
		r = _o[[415]b](e)
	case 416:
		r = _o[[416]b](e)
	case 417:
		r = _o[[417]b](e)
	case 418:
		r = _o[[418]b](e)
	case 419:
		r = _o[[419]b](e)
	case 420:
		r = _o[[420]b](e)
	case 421:
		r = _o[[421]b](e)
	case 422:
		r = _o[[422]b](e)
	case 423:
		r = _o[[423]b](e)
	case 424:
		r = _o[[424]b](e)
	case 425:
		r = _o[[425]b](e)
	case 426:
		r = _o[[426]b](e)
	case 427:
		r = _o[[427]b](e)
	case 428:
		r = _o[[428]b](e)
	case 429:
		r = _o[[429]b](e)
	case 430:
		r = _o[[430]b](e)
	case 431:
		r = _o[[431]b](e)
	case 432:
		r = _o[[432]b](e)
	case 433:
		r = _o[[433]b](e)
	case 434:
		r = _o[[434]b](e)
	case 435:
		r = _o[[435]b](e)
	case 436:
		r = _o[[436]b](e)
	case 437:
		r = _o[[437]b](e)
	case 438:
		r = _o[[438]b](e)
	case 439:
		r = _o[[439]b](e)
	case 440:
		r = _o[[440]b](e)
	case 441:
		r = _o[[441]b](e)
	case 442:
		r = _o[[442]b](e)
	case 443:
		r = _o[[443]b](e)
	case 444:
		r = _o[[444]b](e)
	case 445:
		r = _o[[445]b](e)
	case 446:
		r = _o[[446]b](e)
	case 447:
		r = _o[[447]b](e)
	case 448:
		r = _o[[448]b](e)
	case 449:
		r = _o[[449]b](e)
	case 450:
		r = _o[[450]b](e)
	case 451:
		r = _o[[451]b](e)
	case 452:
		r = _o[[452]b](e)
	case 453:
		r = _o[[453]b](e)
	case 454:
		r = _o[[454]b](e)
	case 455:
		r = _o[[455]b](e)
	case 456:
		r = _o[[456]b](e)
	case 457:
		r = _o[[457]b](e)
	case 458:
		r = _o[[458]b](e)
	case 459:
		r = _o[[459]b](e)
	case 460:
		r = _o[[460]b](e)
	case 461:
		r = _o[[461]b](e)
	case 462:
		r = _o[[462]b](e)
	case 463:
		r = _o[[463]b](e)
	case 464:
		r = _o[[464]b](e)
	case 465:
		r = _o[[465]b](e)
	case 466:
		r = _o[[466]b](e)
	case 467:
		r = _o[[467]b](e)
	case 468:
		r = _o[[468]b](e)
	case 469:
		r = _o[[469]b](e)
	case 470:
		r = _o[[470]b](e)
	case 471:
		r = _o[[471]b](e)
	case 472:
		r = _o[[472]b](e)
	case 473:
		r = _o[[473]b](e)
	case 474:
		r = _o[[474]b](e)
	case 475:
		r = _o[[475]b](e)
	case 476:
		r = _o[[476]b](e)
	case 477:
		r = _o[[477]b](e)
	case 478:
		r = _o[[478]b](e)
	case 479:
		r = _o[[479]b](e)
	case 480:
		r = _o[[480]b](e)
	case 481:
		r = _o[[481]b](e)
	case 482:
		r = _o[[482]b](e)
	case 483:
		r = _o[[483]b](e)
	case 484:
		r = _o[[484]b](e)
	case 485:
		r = _o[[485]b](e)
	case 486:
		r = _o[[486]b](e)
	case 487:
		r = _o[[487]b](e)
	case 488:
		r = _o[[488]b](e)
	case 489:
		r = _o[[489]b](e)
	case 490:
		r = _o[[490]b](e)
	case 491:
		r = _o[[491]b](e)
	case 492:
		r = _o[[492]b](e)
	case 493:
		r = _o[[493]b](e)
	case 494:
		r = _o[[494]b](e)
	case 495:
		r = _o[[495]b](e)
	case 496:
		r = _o[[496]b](e)
	case 497:
		r = _o[[497]b](e)
	case 498:
		r = _o[[498]b](e)
	case 499:
		r = _o[[499]b](e)
	case 500:
		r = _o[[500]b](e)
	case 501:
		r = _o[[501]b](e)
	case 502:
		r = _o[[502]b](e)
	case 503:
		r = _o[[503]b](e)
	case 504:
		r = _o[[504]b](e)
	case 505:
		r = _o[[505]b](e)
	case 506:
		r = _o[[506]b](e)
	case 507:
		r = _o[[507]b](e)
	case 508:
		r = _o[[508]b](e)
	case 509:
		r = _o[[509]b](e)
	case 510:
		r = _o[[510]b](e)
	case 511:
		r = _o[[511]b](e)
	case 512:
		r = _o[[512]b](e)
	case 513:
		r = _o[[513]b](e)
	case 514:
		r = _o[[514]b](e)
	case 515:
		r = _o[[515]b](e)
	case 516:
		r = _o[[516]b](e)
	case 517:
		r = _o[[517]b](e)
	case 518:
		r = _o[[518]b](e)
	case 519:
		r = _o[[519]b](e)
	case 520:
		r = _o[[520]b](e)
	case 521:
		r = _o[[521]b](e)
	case 522:
		r = _o[[522]b](e)
	case 523:
		r = _o[[523]b](e)
	case 524:
		r = _o[[524]b](e)
	case 525:
		r = _o[[525]b](e)
	case 526:
		r = _o[[526]b](e)
	case 527:
		r = _o[[527]b](e)
	case 528:
		r = _o[[528]b](e)
	case 529:
		r = _o[[529]b](e)
	case 530:
		r = _o[[530]b](e)
	case 531:
		r = _o[[531]b](e)
	case 532:
		r = _o[[532]b](e)
	case 533:
		r = _o[[533]b](e)
	case 534:
		r = _o[[534]b](e)
	case 535:
		r = _o[[535]b](e)
	case 536:
		r = _o[[536]b](e)
	case 537:
		r = _o[[537]b](e)
	case 538:
		r = _o[[538]b](e)
	case 539:
		r = _o[[539]b](e)
	case 540:
		r = _o[[540]b](e)
	case 541:
		r = _o[[541]b](e)
	case 542:
		r = _o[[542]b](e)
	case 543:
		r = _o[[543]b](e)
	case 544:
		r = _o[[544]b](e)
	case 545:
		r = _o[[545]b](e)
	case 546:
		r = _o[[546]b](e)
	case 547:
		r = _o[[547]b](e)
	case 548:
		r = _o[[548]b](e)
	case 549:
		r = _o[[549]b](e)
	case 550:
		r = _o[[550]b](e)
	case 551:
		r = _o[[551]b](e)
	case 552:
		r = _o[[552]b](e)
	case 553:
		r = _o[[553]b](e)
	case 554:
		r = _o[[554]b](e)
	case 555:
		r = _o[[555]b](e)
	case 556:
		r = _o[[556]b](e)
	case 557:
		r = _o[[557]b](e)
	case 558:
		r = _o[[558]b](e)
	case 559:
		r = _o[[559]b](e)
	case 560:
		r = _o[[560]b](e)
	case 561:
		r = _o[[561]b](e)
	case 562:
		r = _o[[562]b](e)
	case 563:
		r = _o[[563]b](e)
	case 564:
		r = _o[[564]b](e)
	case 565:
		r = _o[[565]b](e)
	case 566:
		r = _o[[566]b](e)
	case 567:
		r = _o[[567]b](e)
	case 568:
		r = _o[[568]b](e)
	case 569:
		r = _o[[569]b](e)
	case 570:
		r = _o[[570]b](e)
	case 571:
		r = _o[[571]b](e)
	case 572:
		r = _o[[572]b](e)
	case 573:
		r = _o[[573]b](e)
	case 574:
		r = _o[[574]b](e)
	case 575:
		r = _o[[575]b](e)
	case 576:
		r = _o[[576]b](e)
	case 577:
		r = _o[[577]b](e)
	case 578:
		r = _o[[578]b](e)
	case 579:
		r = _o[[579]b](e)
	case 580:
		r = _o[[580]b](e)
	case 581:
		r = _o[[581]b](e)
	case 582:
		r = _o[[582]b](e)
	case 583:
		r = _o[[583]b](e)
	case 584:
		r = _o[[584]b](e)
	case 585:
		r = _o[[585]b](e)
	case 586:
		r = _o[[586]b](e)
	case 587:
		r = _o[[587]b](e)
	case 588:
		r = _o[[588]b](e)
	case 589:
		r = _o[[589]b](e)
	case 590:
		r = _o[[590]b](e)
	case 591:
		r = _o[[591]b](e)
	case 592:
		r = _o[[592]b](e)
	case 593:
		r = _o[[593]b](e)
	case 594:
		r = _o[[594]b](e)
	case 595:
		r = _o[[595]b](e)
	case 596:
		r = _o[[596]b](e)
	case 597:
		r = _o[[597]b](e)
	case 598:
		r = _o[[598]b](e)
	case 599:
		r = _o[[599]b](e)
	case 600:
		r = _o[[600]b](e)
	case 601:
		r = _o[[601]b](e)
	case 602:
		r = _o[[602]b](e)
	case 603:
		r = _o[[603]b](e)
	case 604:
		r = _o[[604]b](e)
	case 605:
		r = _o[[605]b](e)
	case 606:
		r = _o[[606]b](e)
	case 607:
		r = _o[[607]b](e)
	case 608:
		r = _o[[608]b](e)
	case 609:
		r = _o[[609]b](e)
	case 610:
		r = _o[[610]b](e)
	case 611:
		r = _o[[611]b](e)
	case 612:
		r = _o[[612]b](e)
	case 613:
		r = _o[[613]b](e)
	case 614:
		r = _o[[614]b](e)
	case 615:
		r = _o[[615]b](e)
	case 616:
		r = _o[[616]b](e)
	case 617:
		r = _o[[617]b](e)
	case 618:
		r = _o[[618]b](e)
	case 619:
		r = _o[[619]b](e)
	case 620:
		r = _o[[620]b](e)
	case 621:
		r = _o[[621]b](e)
	case 622:
		r = _o[[622]b](e)
	case 623:
		r = _o[[623]b](e)
	case 624:
		r = _o[[624]b](e)
	case 625:
		r = _o[[625]b](e)
	case 626:
		r = _o[[626]b](e)
	case 627:
		r = _o[[627]b](e)
	case 628:
		r = _o[[628]b](e)
	case 629:
		r = _o[[629]b](e)
	case 630:
		r = _o[[630]b](e)
	case 631:
		r = _o[[631]b](e)
	case 632:
		r = _o[[632]b](e)
	case 633:
		r = _o[[633]b](e)
	case 634:
		r = _o[[634]b](e)
	case 635:
		r = _o[[635]b](e)
	case 636:
		r = _o[[636]b](e)
	case 637:
		r = _o[[637]b](e)
	case 638:
		r = _o[[638]b](e)
	case 639:
		r = _o[[639]b](e)
	case 640:
		r = _o[[640]b](e)
	case 641:
		r = _o[[641]b](e)
	case 642:
		r = _o[[642]b](e)
	case 643:
		r = _o[[643]b](e)
	case 644:
		r = _o[[644]b](e)
	case 645:
		r = _o[[645]b](e)
	case 646:
		r = _o[[646]b](e)
	case 647:
		r = _o[[647]b](e)
	case 648:
		r = _o[[648]b](e)
	case 649:
		r = _o[[649]b](e)
	case 650:
		r = _o[[650]b](e)
	case 651:
		r = _o[[651]b](e)
	case 652:
		r = _o[[652]b](e)
	case 653:
		r = _o[[653]b](e)
	case 654:
		r = _o[[654]b](e)
	case 655:
		r = _o[[655]b](e)
	case 656:
		r = _o[[656]b](e)
	case 657:
		r = _o[[657]b](e)
	case 658:
		r = _o[[658]b](e)
	case 659:
		r = _o[[659]b](e)
	case 660:
		r = _o[[660]b](e)
	case 661:
		r = _o[[661]b](e)
	case 662:
		r = _o[[662]b](e)
	case 663:
		r = _o[[663]b](e)
	case 664:
		r = _o[[664]b](e)
	case 665:
		r = _o[[665]b](e)
	case 666:
		r = _o[[666]b](e)
	case 667:
		r = _o[[667]b](e)
	case 668:
		r = _o[[668]b](e)
	case 669:
		r = _o[[669]b](e)
	case 670:
		r = _o[[670]b](e)
	case 671:
		r = _o[[671]b](e)
	case 672:
		r = _o[[672]b](e)
	case 673:
		r = _o[[673]b](e)
	case 674:
		r = _o[[674]b](e)
	case 675:
		r = _o[[675]b](e)
	case 676:
		r = _o[[676]b](e)
	case 677:
		r = _o[[677]b](e)
	case 678:
		r = _o[[678]b](e)
	case 679:
		r = _o[[679]b](e)
	case 680:
		r = _o[[680]b](e)
	case 681:
		r = _o[[681]b](e)
	case 682:
		r = _o[[682]b](e)
	case 683:
		r = _o[[683]b](e)
	case 684:
		r = _o[[684]b](e)
	case 685:
		r = _o[[685]b](e)
	case 686:
		r = _o[[686]b](e)
	case 687:
		r = _o[[687]b](e)
	case 688:
		r = _o[[688]b](e)
	case 689:
		r = _o[[689]b](e)
	case 690:
		r = _o[[690]b](e)
	case 691:
		r = _o[[691]b](e)
	case 692:
		r = _o[[692]b](e)
	case 693:
		r = _o[[693]b](e)
	case 694:
		r = _o[[694]b](e)
	case 695:
		r = _o[[695]b](e)
	case 696:
		r = _o[[696]b](e)
	case 697:
		r = _o[[697]b](e)
	case 698:
		r = _o[[698]b](e)
	case 699:
		r = _o[[699]b](e)
	case 700:
		r = _o[[700]b](e)
	case 701:
		r = _o[[701]b](e)
	case 702:
		r = _o[[702]b](e)
	case 703:
		r = _o[[703]b](e)
	case 704:
		r = _o[[704]b](e)
	case 705:
		r = _o[[705]b](e)
	case 706:
		r = _o[[706]b](e)
	case 707:
		r = _o[[707]b](e)
	case 708:
		r = _o[[708]b](e)
	case 709:
		r = _o[[709]b](e)
	case 710:
		r = _o[[710]b](e)
	case 711:
		r = _o[[711]b](e)
	case 712:
		r = _o[[712]b](e)
	case 713:
		r = _o[[713]b](e)
	case 714:
		r = _o[[714]b](e)
	case 715:
		r = _o[[715]b](e)
	case 716:
		r = _o[[716]b](e)
	case 717:
		r = _o[[717]b](e)
	case 718:
		r = _o[[718]b](e)
	case 719:
		r = _o[[719]b](e)
	case 720:
		r = _o[[720]b](e)
	case 721:
		r = _o[[721]b](e)
	case 722:
		r = _o[[722]b](e)
	case 723:
		r = _o[[723]b](e)
	case 724:
		r = _o[[724]b](e)
	case 725:
		r = _o[[725]b](e)
	case 726:
		r = _o[[726]b](e)
	case 727:
		r = _o[[727]b](e)
	case 728:
		r = _o[[728]b](e)
	case 729:
		r = _o[[729]b](e)
	case 730:
		r = _o[[730]b](e)
	case 731:
		r = _o[[731]b](e)
	case 732:
		r = _o[[732]b](e)
	case 733:
		r = _o[[733]b](e)
	case 734:
		r = _o[[734]b](e)
	case 735:
		r = _o[[735]b](e)
	case 736:
		r = _o[[736]b](e)
	case 737:
		r = _o[[737]b](e)
	case 738:
		r = _o[[738]b](e)
	case 739:
		r = _o[[739]b](e)
	case 740:
		r = _o[[740]b](e)
	case 741:
		r = _o[[741]b](e)
	case 742:
		r = _o[[742]b](e)
	case 743:
		r = _o[[743]b](e)
	case 744:
		r = _o[[744]b](e)
	case 745:
		r = _o[[745]b](e)
	case 746:
		r = _o[[746]b](e)
	case 747:
		r = _o[[747]b](e)
	case 748:
		r = _o[[748]b](e)
	case 749:
		r = _o[[749]b](e)
	case 750:
		r = _o[[750]b](e)
	case 751:
		r = _o[[751]b](e)
	case 752:
		r = _o[[752]b](e)
	case 753:
		r = _o[[753]b](e)
	case 754:
		r = _o[[754]b](e)
	case 755:
		r = _o[[755]b](e)
	case 756:
		r = _o[[756]b](e)
	case 757:
		r = _o[[757]b](e)
	case 758:
		r = _o[[758]b](e)
	case 759:
		r = _o[[759]b](e)
	case 760:
		r = _o[[760]b](e)
	case 761:
		r = _o[[761]b](e)
	case 762:
		r = _o[[762]b](e)
	case 763:
		r = _o[[763]b](e)
	case 764:
		r = _o[[764]b](e)
	case 765:
		r = _o[[765]b](e)
	case 766:
		r = _o[[766]b](e)
	case 767:
		r = _o[[767]b](e)
	case 768:
		r = _o[[768]b](e)
	case 769:
		r = _o[[769]b](e)
	case 770:
		r = _o[[770]b](e)
	case 771:
		r = _o[[771]b](e)
	case 772:
		r = _o[[772]b](e)
	case 773:
		r = _o[[773]b](e)
	case 774:
		r = _o[[774]b](e)
	case 775:
		r = _o[[775]b](e)
	case 776:
		r = _o[[776]b](e)
	case 777:
		r = _o[[777]b](e)
	case 778:
		r = _o[[778]b](e)
	case 779:
		r = _o[[779]b](e)
	case 780:
		r = _o[[780]b](e)
	case 781:
		r = _o[[781]b](e)
	case 782:
		r = _o[[782]b](e)
	case 783:
		r = _o[[783]b](e)
	case 784:
		r = _o[[784]b](e)
	case 785:
		r = _o[[785]b](e)
	case 786:
		r = _o[[786]b](e)
	case 787:
		r = _o[[787]b](e)
	case 788:
		r = _o[[788]b](e)
	case 789:
		r = _o[[789]b](e)
	case 790:
		r = _o[[790]b](e)
	case 791:
		r = _o[[791]b](e)
	case 792:
		r = _o[[792]b](e)
	case 793:
		r = _o[[793]b](e)
	case 794:
		r = _o[[794]b](e)
	case 795:
		r = _o[[795]b](e)
	case 796:
		r = _o[[796]b](e)
	case 797:
		r = _o[[797]b](e)
	case 798:
		r = _o[[798]b](e)
	case 799:
		r = _o[[799]b](e)
	case 800:
		r = _o[[800]b](e)
	case 801:
		r = _o[[801]b](e)
	case 802:
		r = _o[[802]b](e)
	case 803:
		r = _o[[803]b](e)
	case 804:
		r = _o[[804]b](e)
	case 805:
		r = _o[[805]b](e)
	case 806:
		r = _o[[806]b](e)
	case 807:
		r = _o[[807]b](e)
	case 808:
		r = _o[[808]b](e)
	case 809:
		r = _o[[809]b](e)
	case 810:
		r = _o[[810]b](e)
	case 811:
		r = _o[[811]b](e)
	case 812:
		r = _o[[812]b](e)
	case 813:
		r = _o[[813]b](e)
	case 814:
		r = _o[[814]b](e)
	case 815:
		r = _o[[815]b](e)
	case 816:
		r = _o[[816]b](e)
	case 817:
		r = _o[[817]b](e)
	case 818:
		r = _o[[818]b](e)
	case 819:
		r = _o[[819]b](e)
	case 820:
		r = _o[[820]b](e)
	case 821:
		r = _o[[821]b](e)
	case 822:
		r = _o[[822]b](e)
	case 823:
		r = _o[[823]b](e)
	case 824:
		r = _o[[824]b](e)
	case 825:
		r = _o[[825]b](e)
	case 826:
		r = _o[[826]b](e)
	case 827:
		r = _o[[827]b](e)
	case 828:
		r = _o[[828]b](e)
	case 829:
		r = _o[[829]b](e)
	case 830:
		r = _o[[830]b](e)
	case 831:
		r = _o[[831]b](e)
	case 832:
		r = _o[[832]b](e)
	case 833:
		r = _o[[833]b](e)
	case 834:
		r = _o[[834]b](e)
	case 835:
		r = _o[[835]b](e)
	case 836:
		r = _o[[836]b](e)
	case 837:
		r = _o[[837]b](e)
	case 838:
		r = _o[[838]b](e)
	case 839:
		r = _o[[839]b](e)
	case 840:
		r = _o[[840]b](e)
	case 841:
		r = _o[[841]b](e)
	case 842:
		r = _o[[842]b](e)
	case 843:
		r = _o[[843]b](e)
	case 844:
		r = _o[[844]b](e)
	case 845:
		r = _o[[845]b](e)
	case 846:
		r = _o[[846]b](e)
	case 847:
		r = _o[[847]b](e)
	case 848:
		r = _o[[848]b](e)
	case 849:
		r = _o[[849]b](e)
	case 850:
		r = _o[[850]b](e)
	case 851:
		r = _o[[851]b](e)
	case 852:
		r = _o[[852]b](e)
	case 853:
		r = _o[[853]b](e)
	case 854:
		r = _o[[854]b](e)
	case 855:
		r = _o[[855]b](e)
	case 856:
		r = _o[[856]b](e)
	case 857:
		r = _o[[857]b](e)
	case 858:
		r = _o[[858]b](e)
	case 859:
		r = _o[[859]b](e)
	case 860:
		r = _o[[860]b](e)
	case 861:
		r = _o[[861]b](e)
	case 862:
		r = _o[[862]b](e)
	case 863:
		r = _o[[863]b](e)
	case 864:
		r = _o[[864]b](e)
	case 865:
		r = _o[[865]b](e)
	case 866:
		r = _o[[866]b](e)
	case 867:
		r = _o[[867]b](e)
	case 868:
		r = _o[[868]b](e)
	case 869:
		r = _o[[869]b](e)
	case 870:
		r = _o[[870]b](e)
	case 871:
		r = _o[[871]b](e)
	case 872:
		r = _o[[872]b](e)
	case 873:
		r = _o[[873]b](e)
	case 874:
		r = _o[[874]b](e)
	case 875:
		r = _o[[875]b](e)
	case 876:
		r = _o[[876]b](e)
	case 877:
		r = _o[[877]b](e)
	case 878:
		r = _o[[878]b](e)
	case 879:
		r = _o[[879]b](e)
	case 880:
		r = _o[[880]b](e)
	case 881:
		r = _o[[881]b](e)
	case 882:
		r = _o[[882]b](e)
	case 883:
		r = _o[[883]b](e)
	case 884:
		r = _o[[884]b](e)
	case 885:
		r = _o[[885]b](e)
	case 886:
		r = _o[[886]b](e)
	case 887:
		r = _o[[887]b](e)
	case 888:
		r = _o[[888]b](e)
	case 889:
		r = _o[[889]b](e)
	case 890:
		r = _o[[890]b](e)
	case 891:
		r = _o[[891]b](e)
	case 892:
		r = _o[[892]b](e)
	case 893:
		r = _o[[893]b](e)
	case 894:
		r = _o[[894]b](e)
	case 895:
		r = _o[[895]b](e)
	case 896:
		r = _o[[896]b](e)
	case 897:
		r = _o[[897]b](e)
	case 898:
		r = _o[[898]b](e)
	case 899:
		r = _o[[899]b](e)
	case 900:
		r = _o[[900]b](e)
	case 901:
		r = _o[[901]b](e)
	case 902:
		r = _o[[902]b](e)
	case 903:
		r = _o[[903]b](e)
	case 904:
		r = _o[[904]b](e)
	case 905:
		r = _o[[905]b](e)
	case 906:
		r = _o[[906]b](e)
	case 907:
		r = _o[[907]b](e)
	case 908:
		r = _o[[908]b](e)
	case 909:
		r = _o[[909]b](e)
	case 910:
		r = _o[[910]b](e)
	case 911:
		r = _o[[911]b](e)
	case 912:
		r = _o[[912]b](e)
	case 913:
		r = _o[[913]b](e)
	case 914:
		r = _o[[914]b](e)
	case 915:
		r = _o[[915]b](e)
	case 916:
		r = _o[[916]b](e)
	case 917:
		r = _o[[917]b](e)
	case 918:
		r = _o[[918]b](e)
	case 919:
		r = _o[[919]b](e)
	case 920:
		r = _o[[920]b](e)
	case 921:
		r = _o[[921]b](e)
	case 922:
		r = _o[[922]b](e)
	case 923:
		r = _o[[923]b](e)
	case 924:
		r = _o[[924]b](e)
	case 925:
		r = _o[[925]b](e)
	case 926:
		r = _o[[926]b](e)
	case 927:
		r = _o[[927]b](e)
	case 928:
		r = _o[[928]b](e)
	case 929:
		r = _o[[929]b](e)
	case 930:
		r = _o[[930]b](e)
	case 931:
		r = _o[[931]b](e)
	case 932:
		r = _o[[932]b](e)
	case 933:
		r = _o[[933]b](e)
	case 934:
		r = _o[[934]b](e)
	case 935:
		r = _o[[935]b](e)
	case 936:
		r = _o[[936]b](e)
	case 937:
		r = _o[[937]b](e)
	case 938:
		r = _o[[938]b](e)
	case 939:
		r = _o[[939]b](e)
	case 940:
		r = _o[[940]b](e)
	case 941:
		r = _o[[941]b](e)
	case 942:
		r = _o[[942]b](e)
	case 943:
		r = _o[[943]b](e)
	case 944:
		r = _o[[944]b](e)
	case 945:
		r = _o[[945]b](e)
	case 946:
		r = _o[[946]b](e)
	case 947:
		r = _o[[947]b](e)
	case 948:
		r = _o[[948]b](e)
	case 949:
		r = _o[[949]b](e)
	case 950:
		r = _o[[950]b](e)
	case 951:
		r = _o[[951]b](e)
	case 952:
		r = _o[[952]b](e)
	case 953:
		r = _o[[953]b](e)
	case 954:
		r = _o[[954]b](e)
	case 955:
		r = _o[[955]b](e)
	case 956:
		r = _o[[956]b](e)
	case 957:
		r = _o[[957]b](e)
	case 958:
		r = _o[[958]b](e)
	case 959:
		r = _o[[959]b](e)
	case 960:
		r = _o[[960]b](e)
	case 961:
		r = _o[[961]b](e)
	case 962:
		r = _o[[962]b](e)
	case 963:
		r = _o[[963]b](e)
	case 964:
		r = _o[[964]b](e)
	case 965:
		r = _o[[965]b](e)
	case 966:
		r = _o[[966]b](e)
	case 967:
		r = _o[[967]b](e)
	case 968:
		r = _o[[968]b](e)
	case 969:
		r = _o[[969]b](e)
	case 970:
		r = _o[[970]b](e)
	case 971:
		r = _o[[971]b](e)
	case 972:
		r = _o[[972]b](e)
	case 973:
		r = _o[[973]b](e)
	case 974:
		r = _o[[974]b](e)
	case 975:
		r = _o[[975]b](e)
	case 976:
		r = _o[[976]b](e)
	case 977:
		r = _o[[977]b](e)
	case 978:
		r = _o[[978]b](e)
	case 979:
		r = _o[[979]b](e)
	case 980:
		r = _o[[980]b](e)
	case 981:
		r = _o[[981]b](e)
	case 982:
		r = _o[[982]b](e)
	case 983:
		r = _o[[983]b](e)
	case 984:
		r = _o[[984]b](e)
	case 985:
		r = _o[[985]b](e)
	case 986:
		r = _o[[986]b](e)
	case 987:
		r = _o[[987]b](e)
	case 988:
		r = _o[[988]b](e)
	case 989:
		r = _o[[989]b](e)
	case 990:
		r = _o[[990]b](e)
	case 991:
		r = _o[[991]b](e)
	case 992:
		r = _o[[992]b](e)
	case 993:
		r = _o[[993]b](e)
	case 994:
		r = _o[[994]b](e)
	case 995:
		r = _o[[995]b](e)
	case 996:
		r = _o[[996]b](e)
	case 997:
		r = _o[[997]b](e)
	case 998:
		r = _o[[998]b](e)
	case 999:
		r = _o[[999]b](e)
	case 1000:
		r = _o[[1000]b](e)
	case 1001:
		r = _o[[1001]b](e)
	case 1002:
		r = _o[[1002]b](e)
	case 1003:
		r = _o[[1003]b](e)
	case 1004:
		r = _o[[1004]b](e)
	case 1005:
		r = _o[[1005]b](e)
	case 1006:
		r = _o[[1006]b](e)
	case 1007:
		r = _o[[1007]b](e)
	case 1008:
		r = _o[[1008]b](e)
	case 1009:
		r = _o[[1009]b](e)
	case 1010:
		r = _o[[1010]b](e)
	case 1011:
		r = _o[[1011]b](e)
	case 1012:
		r = _o[[1012]b](e)
	case 1013:
		r = _o[[1013]b](e)
	case 1014:
		r = _o[[1014]b](e)
	case 1015:
		r = _o[[1015]b](e)
	case 1016:
		r = _o[[1016]b](e)
	case 1017:
		r = _o[[1017]b](e)
	case 1018:
		r = _o[[1018]b](e)
	case 1019:
		r = _o[[1019]b](e)
	case 1020:
		r = _o[[1020]b](e)
	case 1021:
		r = _o[[1021]b](e)
	case 1022:
		r = _o[[1022]b](e)
	case 1023:
		r = _o[[1023]b](e)
	case 1024:
		r = _o[[1024]b](e)
	default:
	}

	return
}

func _s[T any](where unsafe.Pointer, capacity int) unsafe.Pointer {
	var target = (*[]T)(where)
	*target = make([]T, capacity)
	if capacity == 0 {
		return nil
	}
	return unsafe.Pointer(&(*target)[0])
}

func getSliceMaker(t reflect.Type) (r func(where unsafe.Pointer, capacity int) unsafe.Pointer) {
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
	case 65:
		r = _s[[65]b]
	case 66:
		r = _s[[66]b]
	case 67:
		r = _s[[67]b]
	case 68:
		r = _s[[68]b]
	case 69:
		r = _s[[69]b]
	case 70:
		r = _s[[70]b]
	case 71:
		r = _s[[71]b]
	case 72:
		r = _s[[72]b]
	case 73:
		r = _s[[73]b]
	case 74:
		r = _s[[74]b]
	case 75:
		r = _s[[75]b]
	case 76:
		r = _s[[76]b]
	case 77:
		r = _s[[77]b]
	case 78:
		r = _s[[78]b]
	case 79:
		r = _s[[79]b]
	case 80:
		r = _s[[80]b]
	case 81:
		r = _s[[81]b]
	case 82:
		r = _s[[82]b]
	case 83:
		r = _s[[83]b]
	case 84:
		r = _s[[84]b]
	case 85:
		r = _s[[85]b]
	case 86:
		r = _s[[86]b]
	case 87:
		r = _s[[87]b]
	case 88:
		r = _s[[88]b]
	case 89:
		r = _s[[89]b]
	case 90:
		r = _s[[90]b]
	case 91:
		r = _s[[91]b]
	case 92:
		r = _s[[92]b]
	case 93:
		r = _s[[93]b]
	case 94:
		r = _s[[94]b]
	case 95:
		r = _s[[95]b]
	case 96:
		r = _s[[96]b]
	case 97:
		r = _s[[97]b]
	case 98:
		r = _s[[98]b]
	case 99:
		r = _s[[99]b]
	case 100:
		r = _s[[100]b]
	case 101:
		r = _s[[101]b]
	case 102:
		r = _s[[102]b]
	case 103:
		r = _s[[103]b]
	case 104:
		r = _s[[104]b]
	case 105:
		r = _s[[105]b]
	case 106:
		r = _s[[106]b]
	case 107:
		r = _s[[107]b]
	case 108:
		r = _s[[108]b]
	case 109:
		r = _s[[109]b]
	case 110:
		r = _s[[110]b]
	case 111:
		r = _s[[111]b]
	case 112:
		r = _s[[112]b]
	case 113:
		r = _s[[113]b]
	case 114:
		r = _s[[114]b]
	case 115:
		r = _s[[115]b]
	case 116:
		r = _s[[116]b]
	case 117:
		r = _s[[117]b]
	case 118:
		r = _s[[118]b]
	case 119:
		r = _s[[119]b]
	case 120:
		r = _s[[120]b]
	case 121:
		r = _s[[121]b]
	case 122:
		r = _s[[122]b]
	case 123:
		r = _s[[123]b]
	case 124:
		r = _s[[124]b]
	case 125:
		r = _s[[125]b]
	case 126:
		r = _s[[126]b]
	case 127:
		r = _s[[127]b]
	case 128:
		r = _s[[128]b]
	case 129:
		r = _s[[129]b]
	case 130:
		r = _s[[130]b]
	case 131:
		r = _s[[131]b]
	case 132:
		r = _s[[132]b]
	case 133:
		r = _s[[133]b]
	case 134:
		r = _s[[134]b]
	case 135:
		r = _s[[135]b]
	case 136:
		r = _s[[136]b]
	case 137:
		r = _s[[137]b]
	case 138:
		r = _s[[138]b]
	case 139:
		r = _s[[139]b]
	case 140:
		r = _s[[140]b]
	case 141:
		r = _s[[141]b]
	case 142:
		r = _s[[142]b]
	case 143:
		r = _s[[143]b]
	case 144:
		r = _s[[144]b]
	case 145:
		r = _s[[145]b]
	case 146:
		r = _s[[146]b]
	case 147:
		r = _s[[147]b]
	case 148:
		r = _s[[148]b]
	case 149:
		r = _s[[149]b]
	case 150:
		r = _s[[150]b]
	case 151:
		r = _s[[151]b]
	case 152:
		r = _s[[152]b]
	case 153:
		r = _s[[153]b]
	case 154:
		r = _s[[154]b]
	case 155:
		r = _s[[155]b]
	case 156:
		r = _s[[156]b]
	case 157:
		r = _s[[157]b]
	case 158:
		r = _s[[158]b]
	case 159:
		r = _s[[159]b]
	case 160:
		r = _s[[160]b]
	case 161:
		r = _s[[161]b]
	case 162:
		r = _s[[162]b]
	case 163:
		r = _s[[163]b]
	case 164:
		r = _s[[164]b]
	case 165:
		r = _s[[165]b]
	case 166:
		r = _s[[166]b]
	case 167:
		r = _s[[167]b]
	case 168:
		r = _s[[168]b]
	case 169:
		r = _s[[169]b]
	case 170:
		r = _s[[170]b]
	case 171:
		r = _s[[171]b]
	case 172:
		r = _s[[172]b]
	case 173:
		r = _s[[173]b]
	case 174:
		r = _s[[174]b]
	case 175:
		r = _s[[175]b]
	case 176:
		r = _s[[176]b]
	case 177:
		r = _s[[177]b]
	case 178:
		r = _s[[178]b]
	case 179:
		r = _s[[179]b]
	case 180:
		r = _s[[180]b]
	case 181:
		r = _s[[181]b]
	case 182:
		r = _s[[182]b]
	case 183:
		r = _s[[183]b]
	case 184:
		r = _s[[184]b]
	case 185:
		r = _s[[185]b]
	case 186:
		r = _s[[186]b]
	case 187:
		r = _s[[187]b]
	case 188:
		r = _s[[188]b]
	case 189:
		r = _s[[189]b]
	case 190:
		r = _s[[190]b]
	case 191:
		r = _s[[191]b]
	case 192:
		r = _s[[192]b]
	case 193:
		r = _s[[193]b]
	case 194:
		r = _s[[194]b]
	case 195:
		r = _s[[195]b]
	case 196:
		r = _s[[196]b]
	case 197:
		r = _s[[197]b]
	case 198:
		r = _s[[198]b]
	case 199:
		r = _s[[199]b]
	case 200:
		r = _s[[200]b]
	case 201:
		r = _s[[201]b]
	case 202:
		r = _s[[202]b]
	case 203:
		r = _s[[203]b]
	case 204:
		r = _s[[204]b]
	case 205:
		r = _s[[205]b]
	case 206:
		r = _s[[206]b]
	case 207:
		r = _s[[207]b]
	case 208:
		r = _s[[208]b]
	case 209:
		r = _s[[209]b]
	case 210:
		r = _s[[210]b]
	case 211:
		r = _s[[211]b]
	case 212:
		r = _s[[212]b]
	case 213:
		r = _s[[213]b]
	case 214:
		r = _s[[214]b]
	case 215:
		r = _s[[215]b]
	case 216:
		r = _s[[216]b]
	case 217:
		r = _s[[217]b]
	case 218:
		r = _s[[218]b]
	case 219:
		r = _s[[219]b]
	case 220:
		r = _s[[220]b]
	case 221:
		r = _s[[221]b]
	case 222:
		r = _s[[222]b]
	case 223:
		r = _s[[223]b]
	case 224:
		r = _s[[224]b]
	case 225:
		r = _s[[225]b]
	case 226:
		r = _s[[226]b]
	case 227:
		r = _s[[227]b]
	case 228:
		r = _s[[228]b]
	case 229:
		r = _s[[229]b]
	case 230:
		r = _s[[230]b]
	case 231:
		r = _s[[231]b]
	case 232:
		r = _s[[232]b]
	case 233:
		r = _s[[233]b]
	case 234:
		r = _s[[234]b]
	case 235:
		r = _s[[235]b]
	case 236:
		r = _s[[236]b]
	case 237:
		r = _s[[237]b]
	case 238:
		r = _s[[238]b]
	case 239:
		r = _s[[239]b]
	case 240:
		r = _s[[240]b]
	case 241:
		r = _s[[241]b]
	case 242:
		r = _s[[242]b]
	case 243:
		r = _s[[243]b]
	case 244:
		r = _s[[244]b]
	case 245:
		r = _s[[245]b]
	case 246:
		r = _s[[246]b]
	case 247:
		r = _s[[247]b]
	case 248:
		r = _s[[248]b]
	case 249:
		r = _s[[249]b]
	case 250:
		r = _s[[250]b]
	case 251:
		r = _s[[251]b]
	case 252:
		r = _s[[252]b]
	case 253:
		r = _s[[253]b]
	case 254:
		r = _s[[254]b]
	case 255:
		r = _s[[255]b]
	case 256:
		r = _s[[256]b]
	case 257:
		r = _s[[257]b]
	case 258:
		r = _s[[258]b]
	case 259:
		r = _s[[259]b]
	case 260:
		r = _s[[260]b]
	case 261:
		r = _s[[261]b]
	case 262:
		r = _s[[262]b]
	case 263:
		r = _s[[263]b]
	case 264:
		r = _s[[264]b]
	case 265:
		r = _s[[265]b]
	case 266:
		r = _s[[266]b]
	case 267:
		r = _s[[267]b]
	case 268:
		r = _s[[268]b]
	case 269:
		r = _s[[269]b]
	case 270:
		r = _s[[270]b]
	case 271:
		r = _s[[271]b]
	case 272:
		r = _s[[272]b]
	case 273:
		r = _s[[273]b]
	case 274:
		r = _s[[274]b]
	case 275:
		r = _s[[275]b]
	case 276:
		r = _s[[276]b]
	case 277:
		r = _s[[277]b]
	case 278:
		r = _s[[278]b]
	case 279:
		r = _s[[279]b]
	case 280:
		r = _s[[280]b]
	case 281:
		r = _s[[281]b]
	case 282:
		r = _s[[282]b]
	case 283:
		r = _s[[283]b]
	case 284:
		r = _s[[284]b]
	case 285:
		r = _s[[285]b]
	case 286:
		r = _s[[286]b]
	case 287:
		r = _s[[287]b]
	case 288:
		r = _s[[288]b]
	case 289:
		r = _s[[289]b]
	case 290:
		r = _s[[290]b]
	case 291:
		r = _s[[291]b]
	case 292:
		r = _s[[292]b]
	case 293:
		r = _s[[293]b]
	case 294:
		r = _s[[294]b]
	case 295:
		r = _s[[295]b]
	case 296:
		r = _s[[296]b]
	case 297:
		r = _s[[297]b]
	case 298:
		r = _s[[298]b]
	case 299:
		r = _s[[299]b]
	case 300:
		r = _s[[300]b]
	case 301:
		r = _s[[301]b]
	case 302:
		r = _s[[302]b]
	case 303:
		r = _s[[303]b]
	case 304:
		r = _s[[304]b]
	case 305:
		r = _s[[305]b]
	case 306:
		r = _s[[306]b]
	case 307:
		r = _s[[307]b]
	case 308:
		r = _s[[308]b]
	case 309:
		r = _s[[309]b]
	case 310:
		r = _s[[310]b]
	case 311:
		r = _s[[311]b]
	case 312:
		r = _s[[312]b]
	case 313:
		r = _s[[313]b]
	case 314:
		r = _s[[314]b]
	case 315:
		r = _s[[315]b]
	case 316:
		r = _s[[316]b]
	case 317:
		r = _s[[317]b]
	case 318:
		r = _s[[318]b]
	case 319:
		r = _s[[319]b]
	case 320:
		r = _s[[320]b]
	case 321:
		r = _s[[321]b]
	case 322:
		r = _s[[322]b]
	case 323:
		r = _s[[323]b]
	case 324:
		r = _s[[324]b]
	case 325:
		r = _s[[325]b]
	case 326:
		r = _s[[326]b]
	case 327:
		r = _s[[327]b]
	case 328:
		r = _s[[328]b]
	case 329:
		r = _s[[329]b]
	case 330:
		r = _s[[330]b]
	case 331:
		r = _s[[331]b]
	case 332:
		r = _s[[332]b]
	case 333:
		r = _s[[333]b]
	case 334:
		r = _s[[334]b]
	case 335:
		r = _s[[335]b]
	case 336:
		r = _s[[336]b]
	case 337:
		r = _s[[337]b]
	case 338:
		r = _s[[338]b]
	case 339:
		r = _s[[339]b]
	case 340:
		r = _s[[340]b]
	case 341:
		r = _s[[341]b]
	case 342:
		r = _s[[342]b]
	case 343:
		r = _s[[343]b]
	case 344:
		r = _s[[344]b]
	case 345:
		r = _s[[345]b]
	case 346:
		r = _s[[346]b]
	case 347:
		r = _s[[347]b]
	case 348:
		r = _s[[348]b]
	case 349:
		r = _s[[349]b]
	case 350:
		r = _s[[350]b]
	case 351:
		r = _s[[351]b]
	case 352:
		r = _s[[352]b]
	case 353:
		r = _s[[353]b]
	case 354:
		r = _s[[354]b]
	case 355:
		r = _s[[355]b]
	case 356:
		r = _s[[356]b]
	case 357:
		r = _s[[357]b]
	case 358:
		r = _s[[358]b]
	case 359:
		r = _s[[359]b]
	case 360:
		r = _s[[360]b]
	case 361:
		r = _s[[361]b]
	case 362:
		r = _s[[362]b]
	case 363:
		r = _s[[363]b]
	case 364:
		r = _s[[364]b]
	case 365:
		r = _s[[365]b]
	case 366:
		r = _s[[366]b]
	case 367:
		r = _s[[367]b]
	case 368:
		r = _s[[368]b]
	case 369:
		r = _s[[369]b]
	case 370:
		r = _s[[370]b]
	case 371:
		r = _s[[371]b]
	case 372:
		r = _s[[372]b]
	case 373:
		r = _s[[373]b]
	case 374:
		r = _s[[374]b]
	case 375:
		r = _s[[375]b]
	case 376:
		r = _s[[376]b]
	case 377:
		r = _s[[377]b]
	case 378:
		r = _s[[378]b]
	case 379:
		r = _s[[379]b]
	case 380:
		r = _s[[380]b]
	case 381:
		r = _s[[381]b]
	case 382:
		r = _s[[382]b]
	case 383:
		r = _s[[383]b]
	case 384:
		r = _s[[384]b]
	case 385:
		r = _s[[385]b]
	case 386:
		r = _s[[386]b]
	case 387:
		r = _s[[387]b]
	case 388:
		r = _s[[388]b]
	case 389:
		r = _s[[389]b]
	case 390:
		r = _s[[390]b]
	case 391:
		r = _s[[391]b]
	case 392:
		r = _s[[392]b]
	case 393:
		r = _s[[393]b]
	case 394:
		r = _s[[394]b]
	case 395:
		r = _s[[395]b]
	case 396:
		r = _s[[396]b]
	case 397:
		r = _s[[397]b]
	case 398:
		r = _s[[398]b]
	case 399:
		r = _s[[399]b]
	case 400:
		r = _s[[400]b]
	case 401:
		r = _s[[401]b]
	case 402:
		r = _s[[402]b]
	case 403:
		r = _s[[403]b]
	case 404:
		r = _s[[404]b]
	case 405:
		r = _s[[405]b]
	case 406:
		r = _s[[406]b]
	case 407:
		r = _s[[407]b]
	case 408:
		r = _s[[408]b]
	case 409:
		r = _s[[409]b]
	case 410:
		r = _s[[410]b]
	case 411:
		r = _s[[411]b]
	case 412:
		r = _s[[412]b]
	case 413:
		r = _s[[413]b]
	case 414:
		r = _s[[414]b]
	case 415:
		r = _s[[415]b]
	case 416:
		r = _s[[416]b]
	case 417:
		r = _s[[417]b]
	case 418:
		r = _s[[418]b]
	case 419:
		r = _s[[419]b]
	case 420:
		r = _s[[420]b]
	case 421:
		r = _s[[421]b]
	case 422:
		r = _s[[422]b]
	case 423:
		r = _s[[423]b]
	case 424:
		r = _s[[424]b]
	case 425:
		r = _s[[425]b]
	case 426:
		r = _s[[426]b]
	case 427:
		r = _s[[427]b]
	case 428:
		r = _s[[428]b]
	case 429:
		r = _s[[429]b]
	case 430:
		r = _s[[430]b]
	case 431:
		r = _s[[431]b]
	case 432:
		r = _s[[432]b]
	case 433:
		r = _s[[433]b]
	case 434:
		r = _s[[434]b]
	case 435:
		r = _s[[435]b]
	case 436:
		r = _s[[436]b]
	case 437:
		r = _s[[437]b]
	case 438:
		r = _s[[438]b]
	case 439:
		r = _s[[439]b]
	case 440:
		r = _s[[440]b]
	case 441:
		r = _s[[441]b]
	case 442:
		r = _s[[442]b]
	case 443:
		r = _s[[443]b]
	case 444:
		r = _s[[444]b]
	case 445:
		r = _s[[445]b]
	case 446:
		r = _s[[446]b]
	case 447:
		r = _s[[447]b]
	case 448:
		r = _s[[448]b]
	case 449:
		r = _s[[449]b]
	case 450:
		r = _s[[450]b]
	case 451:
		r = _s[[451]b]
	case 452:
		r = _s[[452]b]
	case 453:
		r = _s[[453]b]
	case 454:
		r = _s[[454]b]
	case 455:
		r = _s[[455]b]
	case 456:
		r = _s[[456]b]
	case 457:
		r = _s[[457]b]
	case 458:
		r = _s[[458]b]
	case 459:
		r = _s[[459]b]
	case 460:
		r = _s[[460]b]
	case 461:
		r = _s[[461]b]
	case 462:
		r = _s[[462]b]
	case 463:
		r = _s[[463]b]
	case 464:
		r = _s[[464]b]
	case 465:
		r = _s[[465]b]
	case 466:
		r = _s[[466]b]
	case 467:
		r = _s[[467]b]
	case 468:
		r = _s[[468]b]
	case 469:
		r = _s[[469]b]
	case 470:
		r = _s[[470]b]
	case 471:
		r = _s[[471]b]
	case 472:
		r = _s[[472]b]
	case 473:
		r = _s[[473]b]
	case 474:
		r = _s[[474]b]
	case 475:
		r = _s[[475]b]
	case 476:
		r = _s[[476]b]
	case 477:
		r = _s[[477]b]
	case 478:
		r = _s[[478]b]
	case 479:
		r = _s[[479]b]
	case 480:
		r = _s[[480]b]
	case 481:
		r = _s[[481]b]
	case 482:
		r = _s[[482]b]
	case 483:
		r = _s[[483]b]
	case 484:
		r = _s[[484]b]
	case 485:
		r = _s[[485]b]
	case 486:
		r = _s[[486]b]
	case 487:
		r = _s[[487]b]
	case 488:
		r = _s[[488]b]
	case 489:
		r = _s[[489]b]
	case 490:
		r = _s[[490]b]
	case 491:
		r = _s[[491]b]
	case 492:
		r = _s[[492]b]
	case 493:
		r = _s[[493]b]
	case 494:
		r = _s[[494]b]
	case 495:
		r = _s[[495]b]
	case 496:
		r = _s[[496]b]
	case 497:
		r = _s[[497]b]
	case 498:
		r = _s[[498]b]
	case 499:
		r = _s[[499]b]
	case 500:
		r = _s[[500]b]
	case 501:
		r = _s[[501]b]
	case 502:
		r = _s[[502]b]
	case 503:
		r = _s[[503]b]
	case 504:
		r = _s[[504]b]
	case 505:
		r = _s[[505]b]
	case 506:
		r = _s[[506]b]
	case 507:
		r = _s[[507]b]
	case 508:
		r = _s[[508]b]
	case 509:
		r = _s[[509]b]
	case 510:
		r = _s[[510]b]
	case 511:
		r = _s[[511]b]
	case 512:
		r = _s[[512]b]
	case 513:
		r = _s[[513]b]
	case 514:
		r = _s[[514]b]
	case 515:
		r = _s[[515]b]
	case 516:
		r = _s[[516]b]
	case 517:
		r = _s[[517]b]
	case 518:
		r = _s[[518]b]
	case 519:
		r = _s[[519]b]
	case 520:
		r = _s[[520]b]
	case 521:
		r = _s[[521]b]
	case 522:
		r = _s[[522]b]
	case 523:
		r = _s[[523]b]
	case 524:
		r = _s[[524]b]
	case 525:
		r = _s[[525]b]
	case 526:
		r = _s[[526]b]
	case 527:
		r = _s[[527]b]
	case 528:
		r = _s[[528]b]
	case 529:
		r = _s[[529]b]
	case 530:
		r = _s[[530]b]
	case 531:
		r = _s[[531]b]
	case 532:
		r = _s[[532]b]
	case 533:
		r = _s[[533]b]
	case 534:
		r = _s[[534]b]
	case 535:
		r = _s[[535]b]
	case 536:
		r = _s[[536]b]
	case 537:
		r = _s[[537]b]
	case 538:
		r = _s[[538]b]
	case 539:
		r = _s[[539]b]
	case 540:
		r = _s[[540]b]
	case 541:
		r = _s[[541]b]
	case 542:
		r = _s[[542]b]
	case 543:
		r = _s[[543]b]
	case 544:
		r = _s[[544]b]
	case 545:
		r = _s[[545]b]
	case 546:
		r = _s[[546]b]
	case 547:
		r = _s[[547]b]
	case 548:
		r = _s[[548]b]
	case 549:
		r = _s[[549]b]
	case 550:
		r = _s[[550]b]
	case 551:
		r = _s[[551]b]
	case 552:
		r = _s[[552]b]
	case 553:
		r = _s[[553]b]
	case 554:
		r = _s[[554]b]
	case 555:
		r = _s[[555]b]
	case 556:
		r = _s[[556]b]
	case 557:
		r = _s[[557]b]
	case 558:
		r = _s[[558]b]
	case 559:
		r = _s[[559]b]
	case 560:
		r = _s[[560]b]
	case 561:
		r = _s[[561]b]
	case 562:
		r = _s[[562]b]
	case 563:
		r = _s[[563]b]
	case 564:
		r = _s[[564]b]
	case 565:
		r = _s[[565]b]
	case 566:
		r = _s[[566]b]
	case 567:
		r = _s[[567]b]
	case 568:
		r = _s[[568]b]
	case 569:
		r = _s[[569]b]
	case 570:
		r = _s[[570]b]
	case 571:
		r = _s[[571]b]
	case 572:
		r = _s[[572]b]
	case 573:
		r = _s[[573]b]
	case 574:
		r = _s[[574]b]
	case 575:
		r = _s[[575]b]
	case 576:
		r = _s[[576]b]
	case 577:
		r = _s[[577]b]
	case 578:
		r = _s[[578]b]
	case 579:
		r = _s[[579]b]
	case 580:
		r = _s[[580]b]
	case 581:
		r = _s[[581]b]
	case 582:
		r = _s[[582]b]
	case 583:
		r = _s[[583]b]
	case 584:
		r = _s[[584]b]
	case 585:
		r = _s[[585]b]
	case 586:
		r = _s[[586]b]
	case 587:
		r = _s[[587]b]
	case 588:
		r = _s[[588]b]
	case 589:
		r = _s[[589]b]
	case 590:
		r = _s[[590]b]
	case 591:
		r = _s[[591]b]
	case 592:
		r = _s[[592]b]
	case 593:
		r = _s[[593]b]
	case 594:
		r = _s[[594]b]
	case 595:
		r = _s[[595]b]
	case 596:
		r = _s[[596]b]
	case 597:
		r = _s[[597]b]
	case 598:
		r = _s[[598]b]
	case 599:
		r = _s[[599]b]
	case 600:
		r = _s[[600]b]
	case 601:
		r = _s[[601]b]
	case 602:
		r = _s[[602]b]
	case 603:
		r = _s[[603]b]
	case 604:
		r = _s[[604]b]
	case 605:
		r = _s[[605]b]
	case 606:
		r = _s[[606]b]
	case 607:
		r = _s[[607]b]
	case 608:
		r = _s[[608]b]
	case 609:
		r = _s[[609]b]
	case 610:
		r = _s[[610]b]
	case 611:
		r = _s[[611]b]
	case 612:
		r = _s[[612]b]
	case 613:
		r = _s[[613]b]
	case 614:
		r = _s[[614]b]
	case 615:
		r = _s[[615]b]
	case 616:
		r = _s[[616]b]
	case 617:
		r = _s[[617]b]
	case 618:
		r = _s[[618]b]
	case 619:
		r = _s[[619]b]
	case 620:
		r = _s[[620]b]
	case 621:
		r = _s[[621]b]
	case 622:
		r = _s[[622]b]
	case 623:
		r = _s[[623]b]
	case 624:
		r = _s[[624]b]
	case 625:
		r = _s[[625]b]
	case 626:
		r = _s[[626]b]
	case 627:
		r = _s[[627]b]
	case 628:
		r = _s[[628]b]
	case 629:
		r = _s[[629]b]
	case 630:
		r = _s[[630]b]
	case 631:
		r = _s[[631]b]
	case 632:
		r = _s[[632]b]
	case 633:
		r = _s[[633]b]
	case 634:
		r = _s[[634]b]
	case 635:
		r = _s[[635]b]
	case 636:
		r = _s[[636]b]
	case 637:
		r = _s[[637]b]
	case 638:
		r = _s[[638]b]
	case 639:
		r = _s[[639]b]
	case 640:
		r = _s[[640]b]
	case 641:
		r = _s[[641]b]
	case 642:
		r = _s[[642]b]
	case 643:
		r = _s[[643]b]
	case 644:
		r = _s[[644]b]
	case 645:
		r = _s[[645]b]
	case 646:
		r = _s[[646]b]
	case 647:
		r = _s[[647]b]
	case 648:
		r = _s[[648]b]
	case 649:
		r = _s[[649]b]
	case 650:
		r = _s[[650]b]
	case 651:
		r = _s[[651]b]
	case 652:
		r = _s[[652]b]
	case 653:
		r = _s[[653]b]
	case 654:
		r = _s[[654]b]
	case 655:
		r = _s[[655]b]
	case 656:
		r = _s[[656]b]
	case 657:
		r = _s[[657]b]
	case 658:
		r = _s[[658]b]
	case 659:
		r = _s[[659]b]
	case 660:
		r = _s[[660]b]
	case 661:
		r = _s[[661]b]
	case 662:
		r = _s[[662]b]
	case 663:
		r = _s[[663]b]
	case 664:
		r = _s[[664]b]
	case 665:
		r = _s[[665]b]
	case 666:
		r = _s[[666]b]
	case 667:
		r = _s[[667]b]
	case 668:
		r = _s[[668]b]
	case 669:
		r = _s[[669]b]
	case 670:
		r = _s[[670]b]
	case 671:
		r = _s[[671]b]
	case 672:
		r = _s[[672]b]
	case 673:
		r = _s[[673]b]
	case 674:
		r = _s[[674]b]
	case 675:
		r = _s[[675]b]
	case 676:
		r = _s[[676]b]
	case 677:
		r = _s[[677]b]
	case 678:
		r = _s[[678]b]
	case 679:
		r = _s[[679]b]
	case 680:
		r = _s[[680]b]
	case 681:
		r = _s[[681]b]
	case 682:
		r = _s[[682]b]
	case 683:
		r = _s[[683]b]
	case 684:
		r = _s[[684]b]
	case 685:
		r = _s[[685]b]
	case 686:
		r = _s[[686]b]
	case 687:
		r = _s[[687]b]
	case 688:
		r = _s[[688]b]
	case 689:
		r = _s[[689]b]
	case 690:
		r = _s[[690]b]
	case 691:
		r = _s[[691]b]
	case 692:
		r = _s[[692]b]
	case 693:
		r = _s[[693]b]
	case 694:
		r = _s[[694]b]
	case 695:
		r = _s[[695]b]
	case 696:
		r = _s[[696]b]
	case 697:
		r = _s[[697]b]
	case 698:
		r = _s[[698]b]
	case 699:
		r = _s[[699]b]
	case 700:
		r = _s[[700]b]
	case 701:
		r = _s[[701]b]
	case 702:
		r = _s[[702]b]
	case 703:
		r = _s[[703]b]
	case 704:
		r = _s[[704]b]
	case 705:
		r = _s[[705]b]
	case 706:
		r = _s[[706]b]
	case 707:
		r = _s[[707]b]
	case 708:
		r = _s[[708]b]
	case 709:
		r = _s[[709]b]
	case 710:
		r = _s[[710]b]
	case 711:
		r = _s[[711]b]
	case 712:
		r = _s[[712]b]
	case 713:
		r = _s[[713]b]
	case 714:
		r = _s[[714]b]
	case 715:
		r = _s[[715]b]
	case 716:
		r = _s[[716]b]
	case 717:
		r = _s[[717]b]
	case 718:
		r = _s[[718]b]
	case 719:
		r = _s[[719]b]
	case 720:
		r = _s[[720]b]
	case 721:
		r = _s[[721]b]
	case 722:
		r = _s[[722]b]
	case 723:
		r = _s[[723]b]
	case 724:
		r = _s[[724]b]
	case 725:
		r = _s[[725]b]
	case 726:
		r = _s[[726]b]
	case 727:
		r = _s[[727]b]
	case 728:
		r = _s[[728]b]
	case 729:
		r = _s[[729]b]
	case 730:
		r = _s[[730]b]
	case 731:
		r = _s[[731]b]
	case 732:
		r = _s[[732]b]
	case 733:
		r = _s[[733]b]
	case 734:
		r = _s[[734]b]
	case 735:
		r = _s[[735]b]
	case 736:
		r = _s[[736]b]
	case 737:
		r = _s[[737]b]
	case 738:
		r = _s[[738]b]
	case 739:
		r = _s[[739]b]
	case 740:
		r = _s[[740]b]
	case 741:
		r = _s[[741]b]
	case 742:
		r = _s[[742]b]
	case 743:
		r = _s[[743]b]
	case 744:
		r = _s[[744]b]
	case 745:
		r = _s[[745]b]
	case 746:
		r = _s[[746]b]
	case 747:
		r = _s[[747]b]
	case 748:
		r = _s[[748]b]
	case 749:
		r = _s[[749]b]
	case 750:
		r = _s[[750]b]
	case 751:
		r = _s[[751]b]
	case 752:
		r = _s[[752]b]
	case 753:
		r = _s[[753]b]
	case 754:
		r = _s[[754]b]
	case 755:
		r = _s[[755]b]
	case 756:
		r = _s[[756]b]
	case 757:
		r = _s[[757]b]
	case 758:
		r = _s[[758]b]
	case 759:
		r = _s[[759]b]
	case 760:
		r = _s[[760]b]
	case 761:
		r = _s[[761]b]
	case 762:
		r = _s[[762]b]
	case 763:
		r = _s[[763]b]
	case 764:
		r = _s[[764]b]
	case 765:
		r = _s[[765]b]
	case 766:
		r = _s[[766]b]
	case 767:
		r = _s[[767]b]
	case 768:
		r = _s[[768]b]
	case 769:
		r = _s[[769]b]
	case 770:
		r = _s[[770]b]
	case 771:
		r = _s[[771]b]
	case 772:
		r = _s[[772]b]
	case 773:
		r = _s[[773]b]
	case 774:
		r = _s[[774]b]
	case 775:
		r = _s[[775]b]
	case 776:
		r = _s[[776]b]
	case 777:
		r = _s[[777]b]
	case 778:
		r = _s[[778]b]
	case 779:
		r = _s[[779]b]
	case 780:
		r = _s[[780]b]
	case 781:
		r = _s[[781]b]
	case 782:
		r = _s[[782]b]
	case 783:
		r = _s[[783]b]
	case 784:
		r = _s[[784]b]
	case 785:
		r = _s[[785]b]
	case 786:
		r = _s[[786]b]
	case 787:
		r = _s[[787]b]
	case 788:
		r = _s[[788]b]
	case 789:
		r = _s[[789]b]
	case 790:
		r = _s[[790]b]
	case 791:
		r = _s[[791]b]
	case 792:
		r = _s[[792]b]
	case 793:
		r = _s[[793]b]
	case 794:
		r = _s[[794]b]
	case 795:
		r = _s[[795]b]
	case 796:
		r = _s[[796]b]
	case 797:
		r = _s[[797]b]
	case 798:
		r = _s[[798]b]
	case 799:
		r = _s[[799]b]
	case 800:
		r = _s[[800]b]
	case 801:
		r = _s[[801]b]
	case 802:
		r = _s[[802]b]
	case 803:
		r = _s[[803]b]
	case 804:
		r = _s[[804]b]
	case 805:
		r = _s[[805]b]
	case 806:
		r = _s[[806]b]
	case 807:
		r = _s[[807]b]
	case 808:
		r = _s[[808]b]
	case 809:
		r = _s[[809]b]
	case 810:
		r = _s[[810]b]
	case 811:
		r = _s[[811]b]
	case 812:
		r = _s[[812]b]
	case 813:
		r = _s[[813]b]
	case 814:
		r = _s[[814]b]
	case 815:
		r = _s[[815]b]
	case 816:
		r = _s[[816]b]
	case 817:
		r = _s[[817]b]
	case 818:
		r = _s[[818]b]
	case 819:
		r = _s[[819]b]
	case 820:
		r = _s[[820]b]
	case 821:
		r = _s[[821]b]
	case 822:
		r = _s[[822]b]
	case 823:
		r = _s[[823]b]
	case 824:
		r = _s[[824]b]
	case 825:
		r = _s[[825]b]
	case 826:
		r = _s[[826]b]
	case 827:
		r = _s[[827]b]
	case 828:
		r = _s[[828]b]
	case 829:
		r = _s[[829]b]
	case 830:
		r = _s[[830]b]
	case 831:
		r = _s[[831]b]
	case 832:
		r = _s[[832]b]
	case 833:
		r = _s[[833]b]
	case 834:
		r = _s[[834]b]
	case 835:
		r = _s[[835]b]
	case 836:
		r = _s[[836]b]
	case 837:
		r = _s[[837]b]
	case 838:
		r = _s[[838]b]
	case 839:
		r = _s[[839]b]
	case 840:
		r = _s[[840]b]
	case 841:
		r = _s[[841]b]
	case 842:
		r = _s[[842]b]
	case 843:
		r = _s[[843]b]
	case 844:
		r = _s[[844]b]
	case 845:
		r = _s[[845]b]
	case 846:
		r = _s[[846]b]
	case 847:
		r = _s[[847]b]
	case 848:
		r = _s[[848]b]
	case 849:
		r = _s[[849]b]
	case 850:
		r = _s[[850]b]
	case 851:
		r = _s[[851]b]
	case 852:
		r = _s[[852]b]
	case 853:
		r = _s[[853]b]
	case 854:
		r = _s[[854]b]
	case 855:
		r = _s[[855]b]
	case 856:
		r = _s[[856]b]
	case 857:
		r = _s[[857]b]
	case 858:
		r = _s[[858]b]
	case 859:
		r = _s[[859]b]
	case 860:
		r = _s[[860]b]
	case 861:
		r = _s[[861]b]
	case 862:
		r = _s[[862]b]
	case 863:
		r = _s[[863]b]
	case 864:
		r = _s[[864]b]
	case 865:
		r = _s[[865]b]
	case 866:
		r = _s[[866]b]
	case 867:
		r = _s[[867]b]
	case 868:
		r = _s[[868]b]
	case 869:
		r = _s[[869]b]
	case 870:
		r = _s[[870]b]
	case 871:
		r = _s[[871]b]
	case 872:
		r = _s[[872]b]
	case 873:
		r = _s[[873]b]
	case 874:
		r = _s[[874]b]
	case 875:
		r = _s[[875]b]
	case 876:
		r = _s[[876]b]
	case 877:
		r = _s[[877]b]
	case 878:
		r = _s[[878]b]
	case 879:
		r = _s[[879]b]
	case 880:
		r = _s[[880]b]
	case 881:
		r = _s[[881]b]
	case 882:
		r = _s[[882]b]
	case 883:
		r = _s[[883]b]
	case 884:
		r = _s[[884]b]
	case 885:
		r = _s[[885]b]
	case 886:
		r = _s[[886]b]
	case 887:
		r = _s[[887]b]
	case 888:
		r = _s[[888]b]
	case 889:
		r = _s[[889]b]
	case 890:
		r = _s[[890]b]
	case 891:
		r = _s[[891]b]
	case 892:
		r = _s[[892]b]
	case 893:
		r = _s[[893]b]
	case 894:
		r = _s[[894]b]
	case 895:
		r = _s[[895]b]
	case 896:
		r = _s[[896]b]
	case 897:
		r = _s[[897]b]
	case 898:
		r = _s[[898]b]
	case 899:
		r = _s[[899]b]
	case 900:
		r = _s[[900]b]
	case 901:
		r = _s[[901]b]
	case 902:
		r = _s[[902]b]
	case 903:
		r = _s[[903]b]
	case 904:
		r = _s[[904]b]
	case 905:
		r = _s[[905]b]
	case 906:
		r = _s[[906]b]
	case 907:
		r = _s[[907]b]
	case 908:
		r = _s[[908]b]
	case 909:
		r = _s[[909]b]
	case 910:
		r = _s[[910]b]
	case 911:
		r = _s[[911]b]
	case 912:
		r = _s[[912]b]
	case 913:
		r = _s[[913]b]
	case 914:
		r = _s[[914]b]
	case 915:
		r = _s[[915]b]
	case 916:
		r = _s[[916]b]
	case 917:
		r = _s[[917]b]
	case 918:
		r = _s[[918]b]
	case 919:
		r = _s[[919]b]
	case 920:
		r = _s[[920]b]
	case 921:
		r = _s[[921]b]
	case 922:
		r = _s[[922]b]
	case 923:
		r = _s[[923]b]
	case 924:
		r = _s[[924]b]
	case 925:
		r = _s[[925]b]
	case 926:
		r = _s[[926]b]
	case 927:
		r = _s[[927]b]
	case 928:
		r = _s[[928]b]
	case 929:
		r = _s[[929]b]
	case 930:
		r = _s[[930]b]
	case 931:
		r = _s[[931]b]
	case 932:
		r = _s[[932]b]
	case 933:
		r = _s[[933]b]
	case 934:
		r = _s[[934]b]
	case 935:
		r = _s[[935]b]
	case 936:
		r = _s[[936]b]
	case 937:
		r = _s[[937]b]
	case 938:
		r = _s[[938]b]
	case 939:
		r = _s[[939]b]
	case 940:
		r = _s[[940]b]
	case 941:
		r = _s[[941]b]
	case 942:
		r = _s[[942]b]
	case 943:
		r = _s[[943]b]
	case 944:
		r = _s[[944]b]
	case 945:
		r = _s[[945]b]
	case 946:
		r = _s[[946]b]
	case 947:
		r = _s[[947]b]
	case 948:
		r = _s[[948]b]
	case 949:
		r = _s[[949]b]
	case 950:
		r = _s[[950]b]
	case 951:
		r = _s[[951]b]
	case 952:
		r = _s[[952]b]
	case 953:
		r = _s[[953]b]
	case 954:
		r = _s[[954]b]
	case 955:
		r = _s[[955]b]
	case 956:
		r = _s[[956]b]
	case 957:
		r = _s[[957]b]
	case 958:
		r = _s[[958]b]
	case 959:
		r = _s[[959]b]
	case 960:
		r = _s[[960]b]
	case 961:
		r = _s[[961]b]
	case 962:
		r = _s[[962]b]
	case 963:
		r = _s[[963]b]
	case 964:
		r = _s[[964]b]
	case 965:
		r = _s[[965]b]
	case 966:
		r = _s[[966]b]
	case 967:
		r = _s[[967]b]
	case 968:
		r = _s[[968]b]
	case 969:
		r = _s[[969]b]
	case 970:
		r = _s[[970]b]
	case 971:
		r = _s[[971]b]
	case 972:
		r = _s[[972]b]
	case 973:
		r = _s[[973]b]
	case 974:
		r = _s[[974]b]
	case 975:
		r = _s[[975]b]
	case 976:
		r = _s[[976]b]
	case 977:
		r = _s[[977]b]
	case 978:
		r = _s[[978]b]
	case 979:
		r = _s[[979]b]
	case 980:
		r = _s[[980]b]
	case 981:
		r = _s[[981]b]
	case 982:
		r = _s[[982]b]
	case 983:
		r = _s[[983]b]
	case 984:
		r = _s[[984]b]
	case 985:
		r = _s[[985]b]
	case 986:
		r = _s[[986]b]
	case 987:
		r = _s[[987]b]
	case 988:
		r = _s[[988]b]
	case 989:
		r = _s[[989]b]
	case 990:
		r = _s[[990]b]
	case 991:
		r = _s[[991]b]
	case 992:
		r = _s[[992]b]
	case 993:
		r = _s[[993]b]
	case 994:
		r = _s[[994]b]
	case 995:
		r = _s[[995]b]
	case 996:
		r = _s[[996]b]
	case 997:
		r = _s[[997]b]
	case 998:
		r = _s[[998]b]
	case 999:
		r = _s[[999]b]
	case 1000:
		r = _s[[1000]b]
	case 1001:
		r = _s[[1001]b]
	case 1002:
		r = _s[[1002]b]
	case 1003:
		r = _s[[1003]b]
	case 1004:
		r = _s[[1004]b]
	case 1005:
		r = _s[[1005]b]
	case 1006:
		r = _s[[1006]b]
	case 1007:
		r = _s[[1007]b]
	case 1008:
		r = _s[[1008]b]
	case 1009:
		r = _s[[1009]b]
	case 1010:
		r = _s[[1010]b]
	case 1011:
		r = _s[[1011]b]
	case 1012:
		r = _s[[1012]b]
	case 1013:
		r = _s[[1013]b]
	case 1014:
		r = _s[[1014]b]
	case 1015:
		r = _s[[1015]b]
	case 1016:
		r = _s[[1016]b]
	case 1017:
		r = _s[[1017]b]
	case 1018:
		r = _s[[1018]b]
	case 1019:
		r = _s[[1019]b]
	case 1020:
		r = _s[[1020]b]
	case 1021:
		r = _s[[1021]b]
	case 1022:
		r = _s[[1022]b]
	case 1023:
		r = _s[[1023]b]
	case 1024:
		r = _s[[1024]b]
	default:
		// degrade to reflect after 1024 bytes element size
		r = func(where unsafe.Pointer, capacity int) unsafe.Pointer {
			e := reflect.MakeSlice(t, capacity, capacity)
			iface := e.Interface()
			interfaceHeader := (*[2]unsafe.Pointer)((unsafe.Pointer)(&iface))

			newSliceHeaderPtr := interfaceHeader[1]
			const headerSize = unsafe.Sizeof([]int{})
			*(*[headerSize]byte)(where) = *(*[headerSize]byte)(newSliceHeaderPtr)
			return e.Index(0).UnsafePointer()
		}
	}
	return
}
