package main

import (
	"bytes"
	"encoding/binary"
	"jk/jklog"
)

// how many bytes to use @cnts
func BytesToInt(v int64, cnts int) []byte {
	buf := make([]byte, cnts)
	binary.PutVarint(buf, v)
	return buf
}

func IntToBytes(buf []byte) int64 {
	nbuf := bytes.NewBuffer(buf)
	n, err := binary.ReadVarint(nbuf)
	if err != nil {
		return -1
	}
	return n
}

func main() {
	v := 231
	buf := BytesToInt(int64(v), 4)
	jklog.L().Infoln("buf: ", buf)

	v1 := IntToBytes(buf)
	jklog.L().Infoln("int: ", int(v1))
}
