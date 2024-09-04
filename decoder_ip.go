package m2o

import (
	"fmt"
	"net"
)

func IPDecoder() CustomDecoder {
	return ipDecoder{}
}

type ipDecoder struct{}

var _ CustomDecoder = ipDecoder{}

func (d ipDecoder) GetType() any {
	return net.IP{}
}

func (d ipDecoder) Decode(raw any) any {
	if stringIP, ok := raw.(string); !ok {
		panic(fmt.Errorf("wrong format for type net.IP, should be a string"))
	} else {
		ip := net.ParseIP(stringIP)
		if ip == nil {
			panic(fmt.Errorf("invalid IP address: %s", stringIP))
		}
		return ip
	}
}
