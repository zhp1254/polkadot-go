package _type

import (
	"github.com/zhex/polkadot-go/codec"
	"reflect"
)

type Extrinsic struct {
}

func EncodeExtrinsic(ex interface{}) ([]byte, error) {
	// todo
	return []byte("mock"), nil
}

func DecodeExtrinsic(b []byte, target reflect.Value) (*codec.ByteInfo, error) {
	// todo
	return nil, nil
}
