package types

import (
	"bytes"
	"encoding/binary"
)

// Unmarshal bytes data to struct
func Unmarshal(data []byte, v interface{}) (err error) {
	r := bytes.NewReader(data)
	d := NewDecoder(r, binary.BigEndian)
	return d.Decode(v)
}
