package simpleserver

import (
	"jk/jklog"
	"net/http"
)

type Shici struct {
	Base
}

func NewShici(path string) *Shici {
	j := &Shici{}
	j.SetPath(path)
	j.SetFunc("/shici", j)
	return j
}

func (s *Shici) Get(w http.ResponseWriter, r *http.Request) {
	sp := SimpleParse{}
	filename := s.path + "/shici/shici.html"
	jklog.L().Debugf("Get html [%s]\n", filename)

	err := sp.Parse(w, filename, "")
	if err != nil {
		jklog.L().Errorln("Parse error ", err)
	}
}
