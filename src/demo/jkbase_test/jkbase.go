package main

import (
	"jkbase"
	"jk/jklog"
)

func main() {
	url := "http://127.0.0.1:8080"
	err := jkbase.JKOpenBrowser(url)
	if err != nil {
		jklog.L().Infof("error open %v\n", err)
		return
	}
	jklog.L().Infoln("open success")
}