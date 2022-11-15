package utility

import (
	"bytes"
	"encoding/binary"
)

func BytesToFloat64(data []byte) float64 {
	var res float64
	buf := bytes.NewReader(data)
	_ = binary.Read(buf, binary.BigEndian, &res)
	return res
}
