package types

import (
	"bytes"
)

// Unmarshal bytes data to struct
func Unmarshal(data []byte, v interface{}) (err error) {
	r := bytes.NewReader(data)
	d := NewDecoder(r, false)
	return d.Decode(v)
}
