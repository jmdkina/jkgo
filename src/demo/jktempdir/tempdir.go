package main

import (
	"io/ioutil"
	"jk/jklog"
)

func main() {
	str , _ := ioutil.TempDir("/tmp/1234", "read")
	jklog.L().Infoln("str: ", str)
}
