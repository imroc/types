package types

import (
	"fmt"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	type foo struct {
		A int8
		B uint8
		C int16
		D uint16
		E int32
		F uint32
		G int64
		H uint64
	}
	data := []byte{
		0x01,
		0x02,
		0x00, 0x03,
		0x00, 0x04,
		0x00, 0x00, 0x00, 0x05,
		0x00, 0x00, 0x00, 0x06,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x07,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08,
	}
	var f foo
	err := Unmarshal(data, &f)
	if err != nil {
		t.Errorf("Unmarshal error:%v", err)
	}
	if !(f.A == 1 && f.B == 2 && f.C == 3 && f.D == 4 && f.E == 5 && f.F == 6 && f.G == 7 && f.H == 8) {
		t.Errorf("bad Unmarshal result")
	}
	fmt.Printf("%+#v", f)
}
