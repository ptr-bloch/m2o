package m2o

/**
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
	"runtime"
	"unsafe"
)

// buildToMapDecoder builds decoderFunc which decodes source data into any map
// possible target map key type: only scalar types
// possible target map value type: any type
//
// usage scenario 1: prefilled billet map. Use this scenario if you want to be able to define input data format in runtime
// in this scenario map contains key which can be decoded, it makes possible to decode only same keys as from billet map
// if map value type is non-empty interface value then it should be filled with concrete value to inform library to
// which concrete value it should be decoded. Example:
//
//	billet := map[string]interface{ Method() } {
//		"key1": SomeConcreteTypeWithMethod{},
//		"key2": AnotherConcreteTypeWithMethod{},
//	}
//	decoder, err := map2object.NewDecoder(billet)
//
// in this case source value for key1 will always be decoded to SomeConcreteTypeWithMethod type and
// key2 to AnotherConcreteTypeWithMethod type
//
// usage scenario 2: empty billet map
// in this case map value type cannot be non-empty interface because library could not know where to take target concrete type
// for decoding
func (b *builder) buildToMapDecoder(value reflect.Value, offsetFromParentAddress uintptr) (decoderFunc, debugInfo) {
	mapType := value.Type()
	keyType := mapType.Key()

	if keyType.Kind() == reflect.Interface || keyType.Kind() == reflect.Struct {
		panic("interface and struct keys of map are not supported")
	}

	valueType := mapType.Elem()

	var debug []string

	toolkit := b.getMapToolkit(mapType)
	make, set := toolkit.Make, toolkit.Set

	var keyMalloc = b.getMemoryAllocator(keyType, "map key")
	//проблема судя по всему здесь, для динамически собираемых структур аллоцировать нужно не для типа значения мапы, а для типа внутреннего значения
	var valMalloc = b.getMemoryAllocator(valueType, "map value")

	if value.Len() > 0 {
		mapRange := value.MapRange()

		var fields []fieldDef

		for mapRange.Next() {
			exportName := mapRange.Key().String()
			keyCopyPointer := reflect.New(keyType)
			keyCopyPointer.Elem().Set(mapRange.Key())
			keyCopy := keyCopyPointer.Elem()
			keyPtr := unsafe.Pointer(keyCopy.UnsafeAddr())
			value := mapRange.Value()

			valueUnmarshaler, info := b.buildDecoder(value, valueType, 0)
			debug = append(debug, info...)

			fields = append(fields, fieldDef{
				u: func(sourceData any, mapHeaderAddress unsafe.Pointer, isOmitted bool) {
					if isOmitted {
						return
					}

					valPtr := valMalloc()
					valueUnmarshaler(sourceData, valPtr, isOmitted)
					set(mapHeaderAddress, keyPtr, valPtr)

					ensureGCDoesNotCollect(keyCopy)
					ensureGCDoesNotCollect(valPtr)
				},

				n: exportName,
			})
		}

		var typedIterator decoderFunc

		fieldsRequiredByDefault := b.config.requiredByDefault
		if fieldsRequiredByDefault {
			typedIterator = unrolledLoopRequired(fields)
		} else {
			typedIterator = unrolledLoopOptional(fields)
		}

		if typedIterator == nil {
			return nil, nil
		}

		zeroOnEmpty := b.config.zeroOnEmpty
		copyDefaultsFromBillet := b.config.copyDefaultsFromBillet

		return func(source any, parentAddress unsafe.Pointer, isOmitted bool) {
			mapAddr := unsafe.Pointer(uintptr(parentAddress) + offsetFromParentAddress)

			if source != nil ||
				source == nil && zeroOnEmpty ||
				source == nil && copyDefaultsFromBillet {

				make(mapAddr)
				// we've made a new map, if we have to leave it zero on empty input, just return
				// we do not take into account b.config.copyDefaultsFromBillet as it should copy only if object should be decoded
				// but while source data is nil, object should not be decoded
				if source == nil && zeroOnEmpty {
					return
				}
			}

			if inputMap, ok := source.(mapStringAny); ok {
				typedIterator(inputMap, mapAddr, isOmitted)
				return
			}

			slowMap(fields, source, mapAddr, fieldsRequiredByDefault)
		}, debug
	}

	if keyType.Kind() == reflect.Interface || valueType.Kind() == reflect.Interface {
		panic(fmt.Errorf("cannot decode to map with interface keys or values. Supply decoder with concrete keys and values where it should decode, even placed within interface values"))
	}

	keyUnmarshaler, info := b.buildDecoder(reflect.New(keyType).Elem(), keyType, 0)
	debug = append(debug, info...)

	valueUnmarshaler, info := b.buildDecoder(reflect.New(valueType).Elem(), valueType, 0)
	debug = append(debug, info...)

	return func(a any, parentAddress unsafe.Pointer, isOmitted bool) {
		inputMap, ok := a.(mapStringAny)
		mapUnsafePtr := unsafe.Pointer(uintptr(parentAddress) + offsetFromParentAddress)
		make(mapUnsafePtr)

		if ok {
			for key, value := range inputMap {
				keyPtr := keyMalloc()
				keyUnmarshaler(key, keyPtr, isOmitted)

				valPtr := valMalloc()
				valueUnmarshaler(value, valPtr, isOmitted)
				set(mapUnsafePtr, keyPtr, valPtr)

				runtime.KeepAlive(keyPtr)
				runtime.KeepAlive(valPtr)
			}
			return
		}

		iterator := reflect.ValueOf(a).MapRange()
		for iterator.Next() {
			key := iterator.Key().Interface()
			value := iterator.Value().Interface()

			keyPtr := keyMalloc()
			keyUnmarshaler(key, keyPtr, false)

			valPtr := valMalloc()
			valueUnmarshaler(value, valPtr, false)
			set(unsafe.Pointer(uintptr(parentAddress)+offsetFromParentAddress), keyPtr, valPtr)
		}
	}, debug
}

type mapToolkit struct {
	Make    func(mapHeaderAddress unsafe.Pointer)
	Set     func(mapHeaderAddress, keyAddress, valueAddress unsafe.Pointer)
	Delete  func(mapHeaderAddress, keyAddress unsafe.Pointer)
	Iterate func(mapHeaderAddress unsafe.Pointer, visitor func(key, value any))
}

func (b *builder) getMapToolkit(mapType reflect.Type) (r *mapToolkit) {
	valueType := mapType.Elem()
	keyKind := mapType.Key().Kind()
	switch keyKind {
	case reflect.Int8:
		r = getMapToolkitForKnownKeyTypeAndValueReflectType[int8](valueType)
	case reflect.Int16:
		r = getMapToolkitForKnownKeyTypeAndValueReflectType[int16](valueType)
	case reflect.Int32:
		r = getMapToolkitForKnownKeyTypeAndValueReflectType[int32](valueType)
	case reflect.Int64:
		r = getMapToolkitForKnownKeyTypeAndValueReflectType[int64](valueType)
	case reflect.Uint8:
		r = getMapToolkitForKnownKeyTypeAndValueReflectType[uint8](valueType)
	case reflect.Uint16:
		r = getMapToolkitForKnownKeyTypeAndValueReflectType[uint16](valueType)
	case reflect.Uint32:
		r = getMapToolkitForKnownKeyTypeAndValueReflectType[uint32](valueType)
	case reflect.Uint64:
		r = getMapToolkitForKnownKeyTypeAndValueReflectType[uint64](valueType)
	case reflect.Float32:
		r = getMapToolkitForKnownKeyTypeAndValueReflectType[float32](valueType)
	case reflect.Float64:
		r = getMapToolkitForKnownKeyTypeAndValueReflectType[int64](valueType)
	case reflect.Complex64:
		r = getMapToolkitForKnownKeyTypeAndValueReflectType[complex64](valueType)
	case reflect.Complex128:
		r = getMapToolkitForKnownKeyTypeAndValueReflectType[complex128](valueType)
	case reflect.String:
		r = getMapToolkitForKnownKeyTypeAndValueReflectType[string](valueType)
	case reflect.Struct:
		r = getMapToolkitForStructKeyAndValueByReflection(mapType)
	default:
	}
	if r == nil {
		return b.getSlowMapMaker(mapType)
	}
	return
}

// getMapToolkitForKnownKeyTypeAndValueReflectType returns map toolkit for known type of key and value type as reflection type
func getMapToolkitForKnownKeyTypeAndValueReflectType[K comparable](valueType reflect.Type) *mapToolkit {
	return _o[K](valueType)
}

// _o see getMapToolkitForKnownKeyTypeAndValueReflectType
// short name's used because function is used from generated code widely
func _o[K comparable](valueType reflect.Type) *mapToolkit {
	switch valueType.Kind() {
	case reflect.Int:
		return getMapToolkitForKnownKeyAndValueTypes[K, int]()
	case reflect.Int8:
		return getMapToolkitForKnownKeyAndValueTypes[K, int8]()
	case reflect.Int16:
		return getMapToolkitForKnownKeyAndValueTypes[K, int16]()
	case reflect.Int32:
		return getMapToolkitForKnownKeyAndValueTypes[K, int32]()
	case reflect.Int64:
		return getMapToolkitForKnownKeyAndValueTypes[K, int64]()
	case reflect.Uint:
		return getMapToolkitForKnownKeyAndValueTypes[K, uint]()
	case reflect.Uint8:
		return getMapToolkitForKnownKeyAndValueTypes[K, uint8]()
	case reflect.Uint16:
		return getMapToolkitForKnownKeyAndValueTypes[K, uint16]()
	case reflect.Uint32:
		return getMapToolkitForKnownKeyAndValueTypes[K, uint32]()
	case reflect.Uint64:
		return getMapToolkitForKnownKeyAndValueTypes[K, uint64]()
	case reflect.Float32:
		return getMapToolkitForKnownKeyAndValueTypes[K, float32]()
	case reflect.Float64:
		return getMapToolkitForKnownKeyAndValueTypes[K, int64]()
	case reflect.Complex64:
		return getMapToolkitForKnownKeyAndValueTypes[K, complex64]()
	case reflect.Complex128:
		return getMapToolkitForKnownKeyAndValueTypes[K, complex128]()
	case reflect.String:
		return getMapToolkitForKnownKeyAndValueTypes[K, string]()
	case reflect.Interface:
		// downgrade to reflection
		if valueType.NumMethod() > 0 {
			return nil
		}
		return getMapToolkitForKnownKeyAndValueTypes[K, interface{}]()
	default:
		return getMapToolkitByKeyTypeAndValueReflection[K](valueType)
	}
}

// getMapToolkitForKnownKeyAndValueTypes returns map toolkit for known key and value types
func getMapToolkitForKnownKeyAndValueTypes[K comparable, V any]() *mapToolkit {
	return _m[K, V]()
}

// _m see getMapToolkitForKnownKeyAndValueTypes
// short name's used because function is widely used from generated code
func _m[K comparable, V any]() *mapToolkit {
	return &mapToolkit{
		Make: func(u unsafe.Pointer) {
			*(*map[K]V)(u) = make(map[K]V)
		},
		Set: func(mapAddr, keyAddr, valAddr unsafe.Pointer) {
			key := (*K)(keyAddr)
			val := (*V)(valAddr)
			Map := *(*map[K]V)(mapAddr)
			Map[*key] = *val
		},
		Delete: func(mapAddr, keyAddr unsafe.Pointer) {
			key := (*K)(keyAddr)
			Map := *(*map[K]V)(mapAddr)
			delete(Map, *key)
		},
		Iterate: func(mapAddr unsafe.Pointer, visitor func(key any, value any)) {
			Map := *(*map[K]V)(mapAddr)
			for k, v := range Map {
				visitor(k, v)
			}
		},
	}
}

func (b *builder) getSlowMapMaker(mapType reflect.Type) *mapToolkit {
	keyType := mapType.Key()
	valueType := mapType.Elem()

	b.debug = append(b.debug, "using map toolkit based on reflection")

	return &mapToolkit{
		Make: func(mapHeaderAddress unsafe.Pointer) {
			pointerToMap := reflect.NewAt(mapType, unsafe.Pointer(mapHeaderAddress))
			mapValue := reflect.MakeMap(mapType)
			pointerToMap.Elem().Set(mapValue)
		},
		Set: func(mapHeaderAddress, keyAddress, valueAddress unsafe.Pointer) {
			Map := reflect.NewAt(mapType, mapHeaderAddress).Elem()
			key := reflect.NewAt(keyType, keyAddress).Elem()
			value := reflect.NewAt(valueType, valueAddress).Elem()
			Map.SetMapIndex(key, value)
		},
		Delete: func(mapHeaderAddress, keyAddress unsafe.Pointer) {
			Map := reflect.NewAt(mapType, mapHeaderAddress).Elem()
			key := reflect.NewAt(keyType, keyAddress).Elem()
			value := reflect.Zero(valueType)
			Map.SetMapIndex(key, value)
		},
	}
}
