package simpleserver

import (
	"fmt"
	"jk/jklog"
	"net/http"
	. "simpleserver"
	"strconv"
)

type WSSimplePage struct {
	Base
}

func NewWSSimplePage(path string) *WSSimplePage {
	j := &WSSimplePage{}
	j.SetPath(path)
	j.SetFunc("/wssimple", j)
	return j
}

func (s *WSSimplePage) Get(w http.ResponseWriter, r *http.Request) {
	sp := SimpleParse{}
	filename := s.Path() + "/wssimple/ws.html"
	jklog.L().Debugf("Get html [%s]\n", filename)

	err := sp.Parse(w, "", filename)
	if err != nil {
		jklog.L().Errorln("Parse error ", err)
	}
}

func (s *WSSimplePage) Post(w http.ResponseWriter, r *http.Request) {
	cmd := r.FormValue("cmd")
	jklog.L().Infof("Recv command of [%s]\n", cmd)
	switch cmd {
	case "start":
		port := r.FormValue("port")
		GlobalWSSimple().LisPort, _ = strconv.Atoi(port)
		GlobalWSSimple().Start()
		s.WriteSerialData(w, "", 200)
		return
	case "stop":
		GlobalWSSimple().Stop()
		s.WriteSerialData(w, "", 200)
		return
	case "exec":
		c := r.FormValue("do")
		GlobalWSSimple().CommonCmd(c)
		s.WriteSerialData(w, "", 200)
		return
	}
}

type WSSimplePageClient struct {
	Base
	wsclient    *WSClientSimple
	isconnected bool
}

func NewWSSimplePageClient(path string) *WSSimplePageClient {
	j := &WSSimplePageClient{}
	j.SetPath(path)
	j.SetFunc("/wsclient", j)
	return j
}

func (s *WSSimplePageClient) Get(w http.ResponseWriter, r *http.Request) {
	sp := SimpleParse{}
	filename := s.Path() + "/wssimple/wsclient.html"
	jklog.L().Debugf("Get html [%s]\n", filename)

	err := sp.Parse(w, "", filename)
	if err != nil {
		jklog.L().Errorln("Parse error ", err)
	}
}

func (s *WSSimplePageClient) Post(w http.ResponseWriter, r *http.Request) {
	cmd := r.FormValue("cmd")
	jklog.L().Infof("wsclient Recv command of [%s]\n", cmd)
	switch cmd {
	case "start":
		addr := r.FormValue("addr")
		port := r.FormValue("port")
		url := r.FormValue("url")
		var err error
		porti, _ := strconv.Atoi(port)
		s.wsclient, err = NewWSClientSimple(addr, url, porti)
		jklog.L().Debugf("wsclient start connect [%s:%s/%s]\n", addr, port, url)
		err = s.wsclient.Start()
		if err != nil {
			s.WriteSerialData(w, err.Error(), 403)
		} else {

			s.isconnected = true
			s.WriteSerialData(w, "success", 200)
		}
	case "stop":
		if s.isconnected {
			fmt.Printf("wsclient %v\n", s.wsclient)
			s.wsclient.Close()
		}
		s.isconnected = false
		jklog.L().Debugln("wsclient stop")
		s.WriteSerialData(w, "", 200)
	case "send":
		msg := r.FormValue("msg")
		jklog.L().Debugf("wsclient send message [%s]\n", msg)
		if s.isconnected {
			_, buf, err := s.wsclient.Send([]byte(msg), true)
			if err != nil {
				s.WriteSerialData(w, err.Error(), 500)
			} else {
				s.WriteSerialData(w, string(buf), 200)
			}
		} else {
			s.WriteSerialData(w, "not connected", 301)
		}
	case "recv":
		jklog.L().Debugln("wsclient recv message")
		if s.isconnected {
			_, buf := s.wsclient.Recv()
			s.WriteSerialData(w, string(buf), 200)
		} else {
			s.WriteSerialData(w, "not connected", 301)
		}
	}
}
