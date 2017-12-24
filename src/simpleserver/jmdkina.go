package simpleserver

import (
	"golanger.com/utils"
	"jk/jklog"
	"net/http"
	. "simpleserver/dbs"
)

type Jmdkina struct {
	Base
}

func NewJmdkina(path string) *Jmdkina {
	j := &Jmdkina{}
	j.SetPath(path)
	j.SetFunc("/jmdkina", j)
	return j
}

func (s *Jmdkina) Get(w http.ResponseWriter, r *http.Request) {
	sp := SimpleParse{}
	filename := s.path + "/jmdkina/jmdkina.html"
	jklog.L().Debugf("Get html [%s]\n", filename)

	err := sp.Parse(w, filename, "")
	if err != nil {
		jklog.L().Errorln("Parse error ", err)
	}
}

func (s *Jmdkina) Post(w http.ResponseWriter, r *http.Request) {
	cmd := r.FormValue("cmd")
	jklog.L().Debugf("jmdkina page -- cmd %s\n", cmd)
	switch cmd {
	case "query_images":
		// index := r.FormValue("index")
		// length := r.FormValue("length")
		out := []utils.M{}
		GlobalDBS().Query("proj", "images", nil, &out)
		s.WriteSerialData(w, out, 200)
		// jklog.L().Debugln(out)
	case "query_images_more":

	}
}
