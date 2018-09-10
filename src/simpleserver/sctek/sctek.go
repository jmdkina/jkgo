package simpleserver

import (
	"jk/jklog"
	"net/http"
	. "simpleserver"
)

type PageSctek struct {
	Base
}

func NewPageSctek(path string) *PageSctek {
	j := &PageSctek{}
	j.SetPath(path)
	j.SetFunc("/sctek", j)
	return j
}

func (s *PageSctek) Get(w http.ResponseWriter, r *http.Request) {
	sp := SimpleParse{}
	filename := s.Path() + "/sctek/sctek.html"
	jklog.L().Debugf("Get html [%s]\n", filename)

	err := sp.Parse(w, filename, "")
	if err != nil {
		jklog.L().Errorln("Parse error ", err)
	}
}

func (s *PageSctek) Post(w http.ResponseWriter, r *http.Request) {
	cmd := r.FormValue("cmd")
	jklog.L().Infof("Recv command of [%s]\n", cmd)
	switch cmd {

	}
}
