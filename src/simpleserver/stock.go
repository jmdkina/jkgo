package simpleserver

import (
	"jk/jklog"
	"net/http"
)

type Stock struct {
	Base
}

func NewStock(path string) *Stock {
	i := &Stock{}
	i.SetPath(path)
	i.SetFunc("/stock", i)
	return i
}

func (s *Stock) Get(w http.ResponseWriter, r *http.Request) {
	sp := SimpleParse{}
	filename := s.path + "/stock.html"
	jklog.L().Debugf("Get html [%s]\n", filename)

	err := sp.Parse(w, filename, "")
	if err != nil {
		jklog.L().Errorln("Parse error ", err)
	}
}

func (s *Stock) Post(w http.ResponseWriter, r *http.Request) {
	method := r.FormValue("op")
	code := r.FormValue("code")
	jkcmd := r.FormValue("jk")
	if code != "jmdstock" || jkcmd != "stockoperation" {
		jklog.L().Errorf("Permission failed\n")
		s.WriteSerialData(w, "", 403)
		return
	}
	jklog.L().Debugf("stock operatin with method %v\n", method)
	switch method {
	case "todayall":
		break
	}
	s.WriteSerialData(w, "no data", 200)
}
