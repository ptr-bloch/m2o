package test

import (
	"fmt"
	"testing"

	"github.com/ptr-bloch/m2o"
)

var benchmarkResultSink any

type benchmarkBasicStruct struct {
	Name    string
	Age     int
	Score   float64
	Enabled bool
}

type benchmarkNestedPayload struct {
	Name  string
	Items []int
	Meta  map[string]int
}

type benchmarkNestedStruct struct {
	ID      int
	Before  string
	Payload *benchmarkNestedPayload
	Items   []int
	Meta    map[string]int
	After   string
}

type benchmarkNamed interface {
	GetName() string
}

type benchmarkInterfaceHolder struct {
	Property benchmarkNamed
}

type benchmarkInterfaceValue struct {
	Name string
	Code int
}

func (v benchmarkInterfaceValue) GetName() string {
	return v.Name
}

type benchmarkAnyStruct struct {
	Name  *string
	Age   int
	Score float64
}

type benchmarkMapInterfaceValueOne struct {
	Name string
	Code int
}

func (v benchmarkMapInterfaceValueOne) GetName() string {
	return "one:" + v.Name
}

type benchmarkMapInterfaceValueTwo struct {
	Name string
	Code int
}

func (v benchmarkMapInterfaceValueTwo) GetName() string {
	return "two:" + v.Name
}

func BenchmarkDecodeBasicStruct(b *testing.B) {
	decoder, err := m2o.NewDecoder(benchmarkBasicStruct{})
	if err != nil {
		b.Fatalf("create decoder: %v", err)
	}

	source := map[string]interface{}{
		"Name":    "John",
		"Age":     40,
		"Score":   120.5,
		"Enabled": true,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var target benchmarkBasicStruct
		if err := decoder.Decode(source, &target); err != nil {
			b.Fatalf("decode: %v", err)
		}
		benchmarkResultSink = target
	}
}

func BenchmarkDecodeNestedStruct(b *testing.B) {
	decoder, err := m2o.NewDecoder(benchmarkNestedStruct{})
	if err != nil {
		b.Fatalf("create decoder: %v", err)
	}

	source := map[string]interface{}{
		"ID":     42,
		"Before": "left",
		"Payload": map[string]interface{}{
			"Name":  "payload",
			"Items": []interface{}{1, 2, 3, 4, 5},
			"Meta": map[string]interface{}{
				"score": 99,
				"rank":  7,
			},
		},
		"Items": []interface{}{6, 7, 8, 9, 10},
		"Meta": map[string]interface{}{
			"outer": 100,
			"limit": 200,
		},
		"After": "right",
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var target benchmarkNestedStruct
		if err := decoder.Decode(source, &target); err != nil {
			b.Fatalf("decode: %v", err)
		}
		benchmarkResultSink = target
	}
}

func BenchmarkDecodeEmptyInterfaceStruct(b *testing.B) {
	name := "initial"
	decoder, err := m2o.NewDecoder(any(benchmarkAnyStruct{
		Name:  &name,
		Age:   1,
		Score: 2,
	}))
	if err != nil {
		b.Fatalf("create decoder: %v", err)
	}

	source := map[string]interface{}{
		"Name":  "Bill",
		"Age":   40,
		"Score": 120.0,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var target any
		if err := decoder.Decode(source, &target); err != nil {
			b.Fatalf("decode: %v", err)
		}
		benchmarkResultSink = target
	}
}

func BenchmarkDecodeTypedInterfaceMultiField(b *testing.B) {
	decoder, err := m2o.NewDecoder(benchmarkInterfaceHolder{
		Property: benchmarkInterfaceValue{},
	})
	if err != nil {
		b.Fatalf("create decoder: %v", err)
	}

	source := map[string]interface{}{
		"Property": map[string]interface{}{
			"Name": "John",
			"Code": 7,
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var target benchmarkInterfaceHolder
		if err := decoder.Decode(source, &target); err != nil {
			b.Fatalf("decode: %v", err)
		}
		benchmarkResultSink = target
	}
}

func BenchmarkDecodeMapStringInt100(b *testing.B) {
	decoder, err := m2o.NewDecoder(map[string]int{})
	if err != nil {
		b.Fatalf("create decoder: %v", err)
	}

	source := make(map[string]interface{}, 100)
	for i := 0; i < 100; i++ {
		source[fmt.Sprintf("key-%03d", i)] = i
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var target map[string]int
		if err := decoder.Decode(source, &target); err != nil {
			b.Fatalf("decode: %v", err)
		}
		benchmarkResultSink = target
	}
}

func BenchmarkDecodeMapWithInterfaceValues(b *testing.B) {
	decoder, err := m2o.NewDecoder(map[string]benchmarkNamed{
		"one": benchmarkMapInterfaceValueOne{},
		"two": benchmarkMapInterfaceValueTwo{},
	})
	if err != nil {
		b.Fatalf("create decoder: %v", err)
	}

	source := map[string]interface{}{
		"one": map[string]interface{}{
			"Name": "John",
			"Code": 1,
		},
		"two": map[string]interface{}{
			"Name": "Jane",
			"Code": 2,
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var target map[string]benchmarkNamed
		if err := decoder.Decode(source, &target); err != nil {
			b.Fatalf("decode: %v", err)
		}
		benchmarkResultSink = target
	}
}

func BenchmarkProduceNestedStruct(b *testing.B) {
	decoder, err := m2o.NewDecoder(benchmarkNestedStruct{})
	if err != nil {
		b.Fatalf("create decoder: %v", err)
	}

	source := map[string]interface{}{
		"ID":     42,
		"Before": "left",
		"Payload": map[string]interface{}{
			"Name":  "payload",
			"Items": []interface{}{1, 2, 3, 4, 5},
			"Meta": map[string]interface{}{
				"score": 99,
				"rank":  7,
			},
		},
		"Items": []interface{}{6, 7, 8, 9, 10},
		"Meta": map[string]interface{}{
			"outer": 100,
			"limit": 200,
		},
		"After": "right",
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target, err := decoder.Produce(source)
		if err != nil {
			b.Fatalf("produce: %v", err)
		}
		benchmarkResultSink = target
	}
}

func BenchmarkNewDecoderNestedStruct(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		decoder, err := m2o.NewDecoder(benchmarkNestedStruct{})
		if err != nil {
			b.Fatalf("create decoder: %v", err)
		}
		benchmarkResultSink = decoder
	}
}
