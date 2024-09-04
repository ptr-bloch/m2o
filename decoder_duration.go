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
	"time"
)

func DurationDecoder() CustomDecoder {
	return durationDecoder{}
}

type durationDecoder struct{}

var _ CustomDecoder = durationDecoder{}

func (d durationDecoder) GetType() any {
	return time.Duration(0)
}

func (d durationDecoder) Decode(raw any) any {
	if stringDuration, ok := raw.(string); !ok {
		panic(fmt.Errorf("wrong format for type time.Duration, should be a string"))
	} else {
		duration, err := time.ParseDuration(stringDuration)
		if err != nil {
			panic(err)
		}
		return duration
	}
}
