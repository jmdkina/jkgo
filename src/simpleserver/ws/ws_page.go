package simpleserver

import (
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

	err := sp.Parse(w, filename, "")
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
