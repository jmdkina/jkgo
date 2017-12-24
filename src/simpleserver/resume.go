package simpleserver

import (
	"jk/jklog"
	"net/http"
)

type Resume struct {
	Base
}

func NewResume(path string) *Resume {
	j := &Resume{}
	j.SetPath(path)
	j.SetFunc("/resume", j)
	return j
}

func (s *Resume) Get(w http.ResponseWriter, r *http.Request) {
	sp := SimpleParse{}
	filename := s.path + "/resume/resume.html"
	jklog.L().Debugf("Get html [%s]\n", filename)

	err := sp.Parse(w, filename, "")
	if err != nil {
		jklog.L().Errorln("Parse error ", err)
	}
}
