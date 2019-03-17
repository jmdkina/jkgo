package simpleserver

import (
	"jk/jklog"
	"net/http"
)

type Project struct {
	Base
}

func NewProject(path string) *Project {
	j := &Project{}
	j.SetPath(path)
	j.SetFunc("/project", j)
	return j
}

func (s *Project) Get(w http.ResponseWriter, r *http.Request) {
	sp := SimpleParse{}
	filename := s.path + "/project/project.html"
	jklog.L().Debugf("Get html [%s]\n", filename)

	err := sp.Parse(w, "", filename)
	if err != nil {
		jklog.L().Errorln("Parse error ", err)
	}
}
