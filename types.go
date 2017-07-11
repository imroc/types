package types

import (
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
)

// Unmarshal bytes data to struct
func Unmarshal(data []byte, v interface{}) (err error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		err = errors.New("invalid struct pointer")
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
	offset := 0
	getData := func(length int) []byte {
		b := data[offset : offset+length]
		offset += length
		return b
	}

	defer func() {
		if r := recover(); r != nil {
			err = errors.New("insufficent data")
		}
	}()

	for i := 0; i < rv.NumField(); i++ {
		f := rv.Field(i)
		switch f.Kind() {
		case reflect.Uint8:
			vv := uint8(data[offset])
			offset++
			f.Set(reflect.ValueOf(vv))
		case reflect.Int8:
			vv := int8(data[offset])
			offset++
			f.Set(reflect.ValueOf(vv))
		case reflect.Uint16:
			b := getData(2)
			vv := binary.BigEndian.Uint16(b)
			f.Set(reflect.ValueOf(vv))
		case reflect.Int16:
			b := getData(2)
			vv := int16(binary.BigEndian.Uint16(b))
			f.Set(reflect.ValueOf(vv))
		case reflect.Uint32:
			b := getData(4)
			vv := binary.BigEndian.Uint32(b)
			f.Set(reflect.ValueOf(vv))
		case reflect.Int32:
			b := getData(4)
			vv := int32(binary.BigEndian.Uint32(b))
			f.Set(reflect.ValueOf(vv))
		case reflect.Uint64:
			b := getData(8)
			vv := binary.BigEndian.Uint64(b)
			f.Set(reflect.ValueOf(vv))
		case reflect.Int64:
			b := getData(8)
			vv := int64(binary.BigEndian.Uint64(b))
			f.Set(reflect.ValueOf(vv))
		default:
			err = fmt.Errorf("unsupported field type:%s", f.Type().String())
			return
		}
	}
	return
}
