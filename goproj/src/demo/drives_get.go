package main

import (
	"jk/jklog"
	. "strconv"
	"syscall"
)

func GetLogicalDrives() []string {
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	GetLogicalDrives := kernel32.MustFindProc("GetLogicalDrives")
	n, _, _ := GetLogicalDrives.Call()
	drivesString := FormatInt(int64(n), 2)
	lastDriver := 64 + len(drivesString)
	var drives []string
	for i := len(drivesString) - 1; i >= 0; i-- {
		if drivesString[i] == 49 {
			drives = append(drives, string(lastDriver-i)+":")
		}
	}
	return drives
}

func main() {
	drives := GetLogicalDrives()
	for _, drive := range drives {
		jklog.L().Infoln("The drives ", drive)
	}
}
