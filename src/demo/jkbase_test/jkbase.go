package main

import (
	"jk/jklog"
	"jkbase"
	"net/http"
	"time"
)

type DemoPage struct {
	jkbase.WebBaseInfo
}

func (dp *DemoPage) Get(w http.ResponseWriter, r *http.Request) {
	jklog.L().Debugln("get in demopage")
	filename := "./html/demopage/demo.html"
	err := dp.Parse(w, filename, "")
	if err != nil {
		jklog.L().Errorln("error ", err)
	}
}

type DemoPageTwo struct {
	jkbase.WebBaseInfo
}

func (dp *DemoPageTwo) Get(w http.ResponseWriter, r *http.Request) {
	jklog.L().Debugln("get in demopage two")
	filename := "./html/demopage/demotwo.html"
	err := dp.Parse(w, filename, "")
	if err != nil {
		jklog.L().Errorln("error ", err)
	}
}

func main() {

	webh, _ := jkbase.NewWebBaseHandle(8089, "html")
	dp := &DemoPage{}
	dp.SetFunc("/demopage", dp)
	dp2 := &DemoPageTwo{}
	dp2.SetFunc("/demopagetwo", dp2)

	go func() {
		webh.Listen()
		jklog.L().Infoln("Error list return")
	}()

	url := "http://127.0.0.1:8089/demopage"
	err := jkbase.JKOpenBrowser(url)
	if err != nil {
		jklog.L().Infof("error open %v\n", err)
		return
	}
	jklog.L().Infoln("open success")
	for {
		time.Sleep(time.Second * 1)
	}
}
