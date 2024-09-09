package test

import (
	"github.com/ptr-bloch/m2o"
	"reflect"
	"runtime"
	"testing"
	"time"
	"unsafe"
)

type PointerStruct struct {
	Name *string
	Age  *int
}

func TestPointerFields(t *testing.T) {
	age := 30
	source := map[string]interface{}{
		"Age": age,
	}

	var profile = m2o.NewProfile()
	decoder, err := m2o.NewDecoder(PointerStruct{}, m2o.WithProfile(profile))

	if err != nil {
		t.Fatalf("Error creating decoder: %v", err)
	}

	var result = new(PointerStruct)
	err = decoder.Decode(source, result)

	if err != nil {
		t.Fatalf("Error decoding: %v", err)
	}

	if result.Age == nil || *result.Age != 30 {
		t.Errorf("Pointer field decoding failed: got %+v", result)
	}

	if result.Name != nil {
		t.Errorf("Pointer field decoding failed: expected nil, got %+v", result.Name)
	}

	var before runtime.MemStats
	var after runtime.MemStats
	var size = SizeOf(result)
	runtime.GC()
	time.Sleep(time.Millisecond * 200)
	runtime.ReadMemStats(&before)
	runtime.KeepAlive(result)
	runtime.GC()
	time.Sleep(time.Millisecond * 200)
	runtime.ReadMemStats(&after)
	freed := after.Frees - before.Frees
	_ = freed
	_ = size

	//checkMemory(t, result, profile)
}

func testMemory() {

}

func SizeOf(v interface{}) uintptr {
	val := reflect.ValueOf(v)
	return sizeOfReflectValue(val)
}

func sizeOfReflectValue(val reflect.Value) uintptr {
	switch val.Kind() {
	case reflect.Ptr, reflect.UnsafePointer:
		if val.IsNil() {
			return uintptr(unsafe.Sizeof(uintptr(0)))
		}
		return uintptr(unsafe.Sizeof(uintptr(0))) + sizeOfReflectValue(val.Elem())
	case reflect.Interface:
		if val.IsNil() {
			return uintptr(unsafe.Sizeof(val.Interface()))
		}
		return uintptr(unsafe.Sizeof(val.Interface())) + sizeOfReflectValue(val.Elem())
	case reflect.Slice:
		size := uintptr(unsafe.Sizeof(val.Pointer()))
		for i := 0; i < val.Len(); i++ {
			size += sizeOfReflectValue(val.Index(i))
		}
		return size
	case reflect.Map:
		if val.IsNil() {
			return 0
		}
		size := uintptr(unsafe.Sizeof(val.Pointer()))
		for _, key := range val.MapKeys() {
			size += sizeOfReflectValue(key) + sizeOfReflectValue(val.MapIndex(key))
		}
		return size
	case reflect.Struct:
		size := uintptr(0)
		for i := 0; i < val.NumField(); i++ {
			size += sizeOfReflectValue(val.Field(i))
		}
		return size
	default:
		return val.Type().Size()
	}
}
