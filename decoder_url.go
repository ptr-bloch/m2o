package m2o

import (
	"fmt"
	"net/url"
)

func URLDecoder() CustomDecoder {
	return urlDecoder{}
}

type urlDecoder struct{}

var _ CustomDecoder = urlDecoder{}

func (d urlDecoder) GetType() any {
	return url.URL{}
}

func (d urlDecoder) Decode(raw any) any {
	if stringURL, ok := raw.(string); !ok {
		panic(fmt.Errorf("wrong format for type url.URL, should be a string"))
	} else {
		parsedURL, err := url.Parse(stringURL)
		if err != nil {
			panic(err)
		}
		return *parsedURL
	}
}
