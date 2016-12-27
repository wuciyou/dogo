package common

import (
// "bytes"
)

// 使用大端模式将int64转成byte数组
func Uint64Tobytes(d uint64) []byte {
	data := make([]byte, 8)
	offset := 7
	for i := 7; i >= 0; i-- {
		data[offset-i] = byte((d >> uint(i*8)) & 0xff)
	}
	return data
}

// 使用大端模式将byte数组转int64
func BytesToUint64(d []byte) uint64 {
	var data uint64
	for i := 7; i >= 0; i-- {
		data |= uint64(d[7-i]) << uint(i*8)
	}
	return data
}
