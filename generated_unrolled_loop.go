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
	case 2:
		r = _m[K, [2]b]()
	case 4:
		r = _m[K, [4]b]()
	case 8:
		r = _m[K, [8]b]()
	case 16:
		r = _m[K, [16]b]()
	case 24:
		r = _m[K, [24]b]()
	case 32:
		r = _m[K, [32]b]()
	case 40:
		r = _m[K, [40]b]()
	case 48:
		r = _m[K, [48]b]()
	case 56:
		r = _m[K, [56]b]()
	case 64:
		r = _m[K, [64]b]()
	case 72:
		r = _m[K, [72]b]()
	case 80:
		r = _m[K, [80]b]()
	case 88:
		r = _m[K, [88]b]()
	case 96:
		r = _m[K, [96]b]()
	case 104:
		r = _m[K, [104]b]()
	case 112:
		r = _m[K, [112]b]()
	case 120:
		r = _m[K, [120]b]()
	case 128:
		r = _m[K, [128]b]()
	case 136:
		r = _m[K, [136]b]()
	case 144:
		r = _m[K, [144]b]()
	case 152:
		r = _m[K, [152]b]()
	case 160:
		r = _m[K, [160]b]()
	case 168:
		r = _m[K, [168]b]()
	case 176:
		r = _m[K, [176]b]()
	case 184:
		r = _m[K, [184]b]()
	case 192:
		r = _m[K, [192]b]()
	case 200:
		r = _m[K, [200]b]()
	case 208:
		r = _m[K, [208]b]()
	case 216:
		r = _m[K, [216]b]()
	case 224:
		r = _m[K, [224]b]()
	case 232:
		r = _m[K, [232]b]()
	case 240:
		r = _m[K, [240]b]()
	case 248:
		r = _m[K, [248]b]()
	case 256:
		r = _m[K, [256]b]()
	case 264:
		r = _m[K, [264]b]()
	case 272:
		r = _m[K, [272]b]()
	case 280:
		r = _m[K, [280]b]()
	case 288:
		r = _m[K, [288]b]()
	case 296:
		r = _m[K, [296]b]()
	case 304:
		r = _m[K, [304]b]()
	case 312:
		r = _m[K, [312]b]()
	case 320:
		r = _m[K, [320]b]()
	case 328:
		r = _m[K, [328]b]()
	case 336:
		r = _m[K, [336]b]()
	case 344:
		r = _m[K, [344]b]()
	case 352:
		r = _m[K, [352]b]()
	case 360:
		r = _m[K, [360]b]()
	case 368:
		r = _m[K, [368]b]()
	case 376:
		r = _m[K, [376]b]()
	case 384:
		r = _m[K, [384]b]()
	case 392:
		r = _m[K, [392]b]()
	case 400:
		r = _m[K, [400]b]()
	case 408:
		r = _m[K, [408]b]()
	case 416:
		r = _m[K, [416]b]()
	case 424:
		r = _m[K, [424]b]()
	case 432:
		r = _m[K, [432]b]()
	case 440:
		r = _m[K, [440]b]()
	case 448:
		r = _m[K, [448]b]()
	case 456:
		r = _m[K, [456]b]()
	case 464:
		r = _m[K, [464]b]()
	case 472:
		r = _m[K, [472]b]()
	case 480:
		r = _m[K, [480]b]()
	case 488:
		r = _m[K, [488]b]()
	case 496:
		r = _m[K, [496]b]()
	case 504:
		r = _m[K, [504]b]()
	case 512:
		r = _m[K, [512]b]()
	case 520:
		r = _m[K, [520]b]()
	case 528:
		r = _m[K, [528]b]()
	case 536:
		r = _m[K, [536]b]()
	case 544:
		r = _m[K, [544]b]()
	case 552:
		r = _m[K, [552]b]()
	case 560:
		r = _m[K, [560]b]()
	case 568:
		r = _m[K, [568]b]()
	case 576:
		r = _m[K, [576]b]()
	case 584:
		r = _m[K, [584]b]()
	case 592:
		r = _m[K, [592]b]()
	case 600:
		r = _m[K, [600]b]()
	case 608:
		r = _m[K, [608]b]()
	case 616:
		r = _m[K, [616]b]()
	case 624:
		r = _m[K, [624]b]()
	case 632:
		r = _m[K, [632]b]()
	case 640:
		r = _m[K, [640]b]()
	case 648:
		r = _m[K, [648]b]()
	case 656:
		r = _m[K, [656]b]()
	case 664:
		r = _m[K, [664]b]()
	case 672:
		r = _m[K, [672]b]()
	case 680:
		r = _m[K, [680]b]()
	case 688:
		r = _m[K, [688]b]()
	case 696:
		r = _m[K, [696]b]()
	case 704:
		r = _m[K, [704]b]()
	case 712:
		r = _m[K, [712]b]()
	case 720:
		r = _m[K, [720]b]()
	case 728:
		r = _m[K, [728]b]()
	case 736:
		r = _m[K, [736]b]()
	case 744:
		r = _m[K, [744]b]()
	case 752:
		r = _m[K, [752]b]()
	case 760:
		r = _m[K, [760]b]()
	case 768:
		r = _m[K, [768]b]()
	case 776:
		r = _m[K, [776]b]()
	case 784:
		r = _m[K, [784]b]()
	case 792:
		r = _m[K, [792]b]()
	case 800:
		r = _m[K, [800]b]()
	case 808:
		r = _m[K, [808]b]()
	case 816:
		r = _m[K, [816]b]()
	case 824:
		r = _m[K, [824]b]()
	case 832:
		r = _m[K, [832]b]()
	case 840:
		r = _m[K, [840]b]()
	case 848:
		r = _m[K, [848]b]()
	case 856:
		r = _m[K, [856]b]()
	case 864:
		r = _m[K, [864]b]()
	case 872:
		r = _m[K, [872]b]()
	case 880:
		r = _m[K, [880]b]()
	case 888:
		r = _m[K, [888]b]()
	case 896:
		r = _m[K, [896]b]()
	case 904:
		r = _m[K, [904]b]()
	case 912:
		r = _m[K, [912]b]()
	case 920:
		r = _m[K, [920]b]()
	case 928:
		r = _m[K, [928]b]()
	case 936:
		r = _m[K, [936]b]()
	case 944:
		r = _m[K, [944]b]()
	case 952:
		r = _m[K, [952]b]()
	case 960:
		r = _m[K, [960]b]()
	case 968:
		r = _m[K, [968]b]()
	case 976:
		r = _m[K, [976]b]()
	case 984:
		r = _m[K, [984]b]()
	case 992:
		r = _m[K, [992]b]()
	case 1000:
		r = _m[K, [1000]b]()
	case 1008:
		r = _m[K, [1008]b]()
	case 1016:
		r = _m[K, [1016]b]()
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
	case 2:
		r = _o[[2]b](e)
	case 4:
		r = _o[[4]b](e)
	case 8:
		r = _o[[8]b](e)
	case 16:
		r = _o[[16]b](e)
	case 24:
		r = _o[[24]b](e)
	case 32:
		r = _o[[32]b](e)
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
	case 2:
		r = _s[[2]b]
	case 4:
		r = _s[[4]b]
	case 8:
		r = _s[[8]b]
	case 16:
		r = _s[[16]b]
	case 24:
		r = _s[[24]b]
	case 32:
		r = _s[[32]b]
	case 40:
		r = _s[[40]b]
	case 48:
		r = _s[[48]b]
	case 56:
		r = _s[[56]b]
	case 64:
		r = _s[[64]b]
	case 72:
		r = _s[[72]b]
	case 80:
		r = _s[[80]b]
	case 88:
		r = _s[[88]b]
	case 96:
		r = _s[[96]b]
	case 104:
		r = _s[[104]b]
	case 112:
		r = _s[[112]b]
	case 120:
		r = _s[[120]b]
	case 128:
		r = _s[[128]b]
	case 136:
		r = _s[[136]b]
	case 144:
		r = _s[[144]b]
	case 152:
		r = _s[[152]b]
	case 160:
		r = _s[[160]b]
	case 168:
		r = _s[[168]b]
	case 176:
		r = _s[[176]b]
	case 184:
		r = _s[[184]b]
	case 192:
		r = _s[[192]b]
	case 200:
		r = _s[[200]b]
	case 208:
		r = _s[[208]b]
	case 216:
		r = _s[[216]b]
	case 224:
		r = _s[[224]b]
	case 232:
		r = _s[[232]b]
	case 240:
		r = _s[[240]b]
	case 248:
		r = _s[[248]b]
	case 256:
		r = _s[[256]b]
	case 264:
		r = _s[[264]b]
	case 272:
		r = _s[[272]b]
	case 280:
		r = _s[[280]b]
	case 288:
		r = _s[[288]b]
	case 296:
		r = _s[[296]b]
	case 304:
		r = _s[[304]b]
	case 312:
		r = _s[[312]b]
	case 320:
		r = _s[[320]b]
	case 328:
		r = _s[[328]b]
	case 336:
		r = _s[[336]b]
	case 344:
		r = _s[[344]b]
	case 352:
		r = _s[[352]b]
	case 360:
		r = _s[[360]b]
	case 368:
		r = _s[[368]b]
	case 376:
		r = _s[[376]b]
	case 384:
		r = _s[[384]b]
	case 392:
		r = _s[[392]b]
	case 400:
		r = _s[[400]b]
	case 408:
		r = _s[[408]b]
	case 416:
		r = _s[[416]b]
	case 424:
		r = _s[[424]b]
	case 432:
		r = _s[[432]b]
	case 440:
		r = _s[[440]b]
	case 448:
		r = _s[[448]b]
	case 456:
		r = _s[[456]b]
	case 464:
		r = _s[[464]b]
	case 472:
		r = _s[[472]b]
	case 480:
		r = _s[[480]b]
	case 488:
		r = _s[[488]b]
	case 496:
		r = _s[[496]b]
	case 504:
		r = _s[[504]b]
	case 512:
		r = _s[[512]b]
	case 520:
		r = _s[[520]b]
	case 528:
		r = _s[[528]b]
	case 536:
		r = _s[[536]b]
	case 544:
		r = _s[[544]b]
	case 552:
		r = _s[[552]b]
	case 560:
		r = _s[[560]b]
	case 568:
		r = _s[[568]b]
	case 576:
		r = _s[[576]b]
	case 584:
		r = _s[[584]b]
	case 592:
		r = _s[[592]b]
	case 600:
		r = _s[[600]b]
	case 608:
		r = _s[[608]b]
	case 616:
		r = _s[[616]b]
	case 624:
		r = _s[[624]b]
	case 632:
		r = _s[[632]b]
	case 640:
		r = _s[[640]b]
	case 648:
		r = _s[[648]b]
	case 656:
		r = _s[[656]b]
	case 664:
		r = _s[[664]b]
	case 672:
		r = _s[[672]b]
	case 680:
		r = _s[[680]b]
	case 688:
		r = _s[[688]b]
	case 696:
		r = _s[[696]b]
	case 704:
		r = _s[[704]b]
	case 712:
		r = _s[[712]b]
	case 720:
		r = _s[[720]b]
	case 728:
		r = _s[[728]b]
	case 736:
		r = _s[[736]b]
	case 744:
		r = _s[[744]b]
	case 752:
		r = _s[[752]b]
	case 760:
		r = _s[[760]b]
	case 768:
		r = _s[[768]b]
	case 776:
		r = _s[[776]b]
	case 784:
		r = _s[[784]b]
	case 792:
		r = _s[[792]b]
	case 800:
		r = _s[[800]b]
	case 808:
		r = _s[[808]b]
	case 816:
		r = _s[[816]b]
	case 824:
		r = _s[[824]b]
	case 832:
		r = _s[[832]b]
	case 840:
		r = _s[[840]b]
	case 848:
		r = _s[[848]b]
	case 856:
		r = _s[[856]b]
	case 864:
		r = _s[[864]b]
	case 872:
		r = _s[[872]b]
	case 880:
		r = _s[[880]b]
	case 888:
		r = _s[[888]b]
	case 896:
		r = _s[[896]b]
	case 904:
		r = _s[[904]b]
	case 912:
		r = _s[[912]b]
	case 920:
		r = _s[[920]b]
	case 928:
		r = _s[[928]b]
	case 936:
		r = _s[[936]b]
	case 944:
		r = _s[[944]b]
	case 952:
		r = _s[[952]b]
	case 960:
		r = _s[[960]b]
	case 968:
		r = _s[[968]b]
	case 976:
		r = _s[[976]b]
	case 984:
		r = _s[[984]b]
	case 992:
		r = _s[[992]b]
	case 1000:
		r = _s[[1000]b]
	case 1008:
		r = _s[[1008]b]
	case 1016:
		r = _s[[1016]b]
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
