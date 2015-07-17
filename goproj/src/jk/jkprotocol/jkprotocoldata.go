package jkprotocol

import (
	"strconv"
)

// First query from other device
// contain information
// Response need.
type JKDeviceInfo struct {
	Name  string
	Addr  string
	ID    int
	Model string
}

func (di *JKDeviceInfo) ToByte() []byte {
	str := "Name:" + di.Name + ";" +
		"Addr:" + di.Addr + ";" +
		"ID:" + strconv.Itoa(di.ID) + ";" +
		"Model:" + di.Model + ";"
	return []byte(str)
}

// Used when need file transfer
// Use the addr and port to contact each other
// This will use tcp connect
// Response need.
type JKFileTransfer struct {
	Addr string
	Port int
}

func (ft *JKFileTransfer) ToByte() []byte {
	str := "Addr:" + ft.Addr + ";" +
		"Port:" + strconv.Itoa(ft.Port) + ";"
	return []byte(str)
}
