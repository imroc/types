package types

import "encoding/binary"

var defaultByteOrder binary.ByteOrder = binary.BigEndian

func SetDefaultByteOrder(enc binary.ByteOrder) {
	if enc == nil {
		return
	}
	defaultByteOrder = enc
}
