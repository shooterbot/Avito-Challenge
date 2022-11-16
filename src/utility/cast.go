package utility

import (
	"bytes"
	"encoding/binary"
)

func BytesToInt(data []byte) int {
	return int(binary.BigEndian.Uint32(data))
}

func BytesToFloat64(data []byte) float64 {
	var res float64
	buf := bytes.NewReader(data)
	_ = binary.Read(buf, binary.BigEndian, &res)
	return res
}

func BytesToString(data []byte) string {
	return string(data)
}
