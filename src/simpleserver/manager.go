package simpleserver

import (
	"jk/jklog"
	"net/http"
)

type Manager struct {
	Base
}

func NewManager(path string) *Manager {
	j := &Manager{}
	j.SetPath(path)
	j.SetFunc("/manager&code=jkvalid", j)
	return j
}

func (s *Manager) Get(w http.ResponseWriter, r *http.Request) {
	sp := SimpleParse{}
	filename := s.Path() + "/manager/index.html"
	jklog.L().Debugf("Get html [%s]\n", filename)

	err := sp.Parse(w, filename, "")
	if err != nil {
		jklog.L().Errorln("parse error ", err)
	}
}

