package types

import "bytes"

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := Encode(&buf, v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
