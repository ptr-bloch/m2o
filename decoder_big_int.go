package m2o

import (
	"fmt"
	"math/big"
)

func BigIntDecoder() CustomDecoder {
	return bigIntDecoder{}
}

type bigIntDecoder struct{}

var _ CustomDecoder = bigIntDecoder{}

func (d bigIntDecoder) GetType() any {
	return *big.NewInt(0)
}

func (d bigIntDecoder) Decode(raw any) any {
	if stringBigInt, ok := raw.(string); !ok {
		panic(fmt.Errorf("wrong format for type *big.Int, should be a string"))
	} else {
		bigInt := new(big.Int)
		if _, ok := bigInt.SetString(stringBigInt, 10); !ok {
			panic(fmt.Errorf("invalid big.Int format: %s", stringBigInt))
		}
		return *bigInt
	}
}
