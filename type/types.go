package _type

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

func StringToPtr(v string) *string {
	return &v
}

func IntToPtr(v int) *int {
	return &v
}

func Int32ToPtr(v int32) *int32 {
	return &v
}

func Int64ToPtr(v int64) *int64 {
	return &v
}

func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func BytesToInt(b []byte) int32 {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}

func Int64ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}

func StringToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}
