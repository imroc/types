package types

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"reflect"
)

type Encoder struct {
	w   io.Writer
	enc binary.ByteOrder
}

func NewEncoder(w io.Writer, littleEndian bool) *Encoder {
	e := &Encoder{
		w: w,
	}
	if littleEndian {
		e.enc = binary.LittleEndian
	} else {
		e.enc = binary.BigEndian
	}
	return e
}

func (e *Encoder) Encode(v interface{}) (err error) {
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
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Field(i)
		var b []byte
		switch f.Kind() {
		case reflect.Uint8:
			b = []byte{uint8(f.Uint())}
		case reflect.Int8:
			b = []byte{uint8(f.Int())}
		case reflect.Uint16:
			vv := uint16(f.Uint())
			b = make([]byte, 2)
			binary.BigEndian.PutUint16(b, vv)
		case reflect.Int16:
			vv := uint16(f.Int())
			b = make([]byte, 2)
			binary.BigEndian.PutUint16(b, vv)
		case reflect.Uint32:
			vv := uint32(f.Uint())
			b = make([]byte, 4)
			binary.BigEndian.PutUint32(b, vv)
		case reflect.Int32:
			vv := uint32(f.Int())
			b = make([]byte, 4)
			binary.BigEndian.PutUint32(b, vv)
		case reflect.Uint64:
			vv := f.Uint()
			b = make([]byte, 8)
			binary.BigEndian.PutUint64(b, vv)
		case reflect.Int64:
			vv := uint64(f.Int())
			b = make([]byte, 8)
			binary.BigEndian.PutUint64(b, vv)
		default:
			err = fmt.Errorf("types: unsupported field type:%s", f.Type().String())
			return
		}
		_, err = e.w.Write(b)
		if err != nil {
			return
		}
	}
	return
}

func Encode(w io.Writer, v interface{}) error {
	e := NewEncoder(w, false)
	return e.Encode(v)
}
