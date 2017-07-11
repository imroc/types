package types

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"reflect"
)

type Decoder struct {
	r io.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r: r,
	}
}

// Decode decode the reader to specified struct
func (d *Decoder) Decode(v interface{}) (err error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		err = errors.New("types: invalid struct pointer")
		return
	}
	for {
		rv = rv.Elem()
		if rv.Kind() == reflect.Ptr {
			continue
		} else {
			break
		}
	}

	getData := func(length int) []byte {
		b := make([]byte, length)
		l, err := d.r.Read(b)
		if err != nil {
			panic(err)
		}
		if l != length {
			panic("insufficient data")
		}
		return b
	}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("types: decode failed: %v", r)
		}
	}()

	for i := 0; i < rv.NumField(); i++ {
		f := rv.Field(i)
		var value interface{}
		switch f.Kind() {
		case reflect.Uint8:
			b := getData(1)
			value = uint8(b[0])
		case reflect.Int8:
			b := getData(1)
			value = int8(b[0])
		case reflect.Uint16:
			b := getData(2)
			value = binary.BigEndian.Uint16(b)
		case reflect.Int16:
			b := getData(2)
			value = int16(binary.BigEndian.Uint16(b))
		case reflect.Uint32:
			b := getData(4)
			value = binary.BigEndian.Uint32(b)
		case reflect.Int32:
			b := getData(4)
			value = int32(binary.BigEndian.Uint32(b))
		case reflect.Uint64:
			b := getData(8)
			value = binary.BigEndian.Uint64(b)
		case reflect.Int64:
			b := getData(8)
			value = int64(binary.BigEndian.Uint64(b))
		default:
			err = fmt.Errorf("types: unsupported field type:%s", f.Type().String())
			return
		}
		f.Set(reflect.ValueOf(value))
	}
	return
}

// Decode decode the reader to specified struct
func Decode(r io.Reader, v interface{}) (err error) {
	d := NewDecoder(r)
	return d.Decode(v)
}
