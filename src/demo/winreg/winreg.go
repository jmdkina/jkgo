package main

import (
	reg "golang.org/x/sys/windows/registry"
	"jk/jklog"
)

func main() {

	testpath := "SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion\\NetworkList\\Signatures\\Unmanaged"

	k, err := reg.OpenKey(reg.LOCAL_MACHINE, testpath, reg.QUERY_VALUE)
	if err != nil {
		jklog.L().Errorln("key open failed ", err)
		return
	}
	keyi, _ := k.Stat()
	jklog.L().Infoln("key ", keyi)
}
