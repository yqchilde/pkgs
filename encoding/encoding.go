package encoding

import (
	"encoding"
	"errors"
	"reflect"
	"strings"
)

var (
	// ErrNotAPointer 参数指针类型错误
	ErrNotAPointer = errors.New("v argument must be a pointer")
)

// Codec 编解码器，用于json或proto等编解码
type Codec interface {
	// Marshal 编码
	Marshal(v interface{}) ([]byte, error)

	// Unmarshal 解码
	Unmarshal(data []byte, v interface{}) error

	// Name 编解码器类型名字
	Name() string
}

var registeredCodecs = make(map[string]Codec)

// RegisterCodec 注册编解码器
func RegisterCodec(codec Codec) {
	if codec == nil {
		panic("cannot register a nil Codec")
	}
	if codec.Name() == "" {
		panic("cannot register a Codec with an empty Name")
	}
	contentSubtype := strings.ToLower(codec.Name())
	registeredCodecs[contentSubtype] = codec
}

// GetCodec 获取编解码器
func GetCodec(contentSubtype string) Codec {
	if contentSubtype == "" {
		return nil
	}
	return registeredCodecs[contentSubtype]
}

type Encoding interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

// Marshal encode data
func Marshal(e Encoding, v interface{}) (data []byte, err error) {
	if !isPointer(v) {
		return data, ErrNotAPointer
	}
	bm, ok := v.(encoding.BinaryMarshaler)
	if ok && e == nil {
		data, err = bm.MarshalBinary()
		return
	}

	data, err = e.Marshal(v)
	if err == nil {
		return
	}
	if ok {
		data, err = bm.MarshalBinary()
	}
	return
}

// Unmarshal decode data
func Unmarshal(e Encoding, data []byte, v interface{}) (err error) {
	if !isPointer(v) {
		return ErrNotAPointer
	}
	bm, ok := v.(encoding.BinaryUnmarshaler)
	if ok && e == nil {
		err = bm.UnmarshalBinary(data)
		return err
	}
	err = e.Unmarshal(data, v)
	if err == nil {
		return
	}
	if ok {
		return bm.UnmarshalBinary(data)
	}
	return
}

func isPointer(data interface{}) bool {
	switch reflect.ValueOf(data).Kind() {
	case reflect.Ptr, reflect.Interface:
		return true
	default:
		return false
	}
}
