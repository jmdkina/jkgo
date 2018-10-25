package main

import (
	"encoding/json"
	"jk/jklog"
	"jkbase"
	"net/http"
	"sctek"
	"time"
)

type SCServerHandle struct {
	sd *sctek.SctekDiscover
}

var schandle SCServerHandle

type WebSctekPage struct {
	jkbase.WebBaseInfo
}

func (wsp *WebSctekPage) Get(w http.ResponseWriter, r *http.Request) {
	jklog.L().Debugln("sctek get")
	filename := "./html/sctek/sctek.html"
	buffer, err := json.Marshal(schandle.sd.DevList)
	if err != nil {
		jklog.L().Errorln("parse error ", err)
	}
	jklog.L().Debugln("parse out ", string(buffer))
	err = wsp.Parse(w, filename, schandle.sd.DevList)
	if err != nil {
		jklog.L().Errorln("error parse ", err)
	}
}

func (wsp *WebSctekPage) Post(w http.ResponseWriter, r *http.Request) {
	jklog.L().Debugln("sctek post")
}

func main() {
	var err error
	schandle.sd, err = sctek.NewSctekDiscover()
	if err != nil {
		jklog.L().Errorln(err)
		return
	}
	go func() {
		schandle.sd.Discover(10)
	}()

	webh, _ := jkbase.NewWebBaseHandle(12309, "./html")
	wsp := &WebSctekPage{}
	wsp.SetFunc("/sctek", wsp)

	url := "http://127.0.0.1:12309/sctek"
	err = jkbase.JKOpenBrowser(url)
	if err != nil {
		jklog.L().Errorln("Open brower error ", err)
	} else {
		jklog.L().Infoln("Open brower success")
	}

	webh.Listen()
	for {
		time.Sleep(time.Millisecond * 500)
	}
}
