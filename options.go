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
)

type Option func(*builder)

// WithAllFieldsRequired - throw an error when not all fields were filled
func WithAllFieldsRequired() Option {
	return func(b *builder) {
		b.config.requiredByDefault = true
	}
}

// WithZeroOnEmpty - set fields to zero value if source data for field is nil
func WithZeroOnEmpty() Option {
	return func(b *builder) {
		b.config.zeroOnEmpty = true
	}
}

// WithInitializeFromBillet makes decoder initialize every object with fields from billet
// for example you can set billet to some initial state (including unexported fields) and expect
// every object after initialization include this state
// for example see test/initialize_from_billet_test.go
func WithInitializeFromBillet() Option {
	return func(b *builder) {
		b.config.copyDefaultsFromBillet = true
	}
}

// WithSquashAnonymous makes decoder to squash field if it's anonymous (embedded)
// To squash a struct field means behave as its fields are fields of parent structure:
// for example see test/squashing_test.go
func WithSquashAnonymous() Option {
	return func(b *builder) {
		b.config.squashAnonymousFields = true
	}
}

// CustomDecoder is any type which introduces type which it can decode and decoder itself
// should be used in conjunction with WithCustomDecoder(decoder CustomDecoder) method.
// When builder during traversing billet object encounters value of type which has its customer decoder
// it adds this custom decoder to execution plan
type CustomDecoder interface {
	// GetType should return value which decoder can decode
	// called during execution plan building phase
	GetType() any

	// Decode should return decoded value
	// Values type should be the same as GetType() returns
	// panics if return value type differs from type of value, returned by GetType()
	Decode(any) any
}

// WithTimeDecoder allows decoding to fields of type time.Time
func WithTimeDecoder(layout string) Option {
	return WithCustomDecoder(TimeDecoder(layout))
}

// WithDurationDecoder allows decoding to fields of type time.Duration
func WithDurationDecoder() Option {
	return WithCustomDecoder(DurationDecoder())
}

// WithIPDecoder enables decoding to net.IP
func WithIPDecoder() Option {
	return WithCustomDecoder(IPDecoder())
}

// WithURLDecoder enables decoding to url.URL
func WithURLDecoder() Option {
	return WithCustomDecoder(URLDecoder())
}

func WithBigIntDecoder() Option {
	return WithCustomDecoder(BigIntDecoder())
}

func WithCustomDecoder(decoder CustomDecoder) Option {
	return func(b *builder) {
		v := decoder.GetType()
		rValue := reflect.ValueOf(v)
		t := rValue.Type()
		if _, exists := b.customDecoders[t]; exists {
			var name string
			if t.PkgPath() != "" {
				name = t.PkgPath() + "/"
			}
			name += t.Name()
			panic(fmt.Errorf("multiple custom decoders for type %s", name))
		}
		b.customDecoders[t] = decoder.Decode
	}
}

func WithProfile(p *Profile) Option {
	return func(b *builder) {
		b.profile = p
	}
}

// WithMustUseAllSource - throw an error when not all input data was used
// TODO is not implemented yet
func WithMustUseAllSource() Option {
	return func(b *builder) {
		b.config.mustUseAllSource = true
	}
}
