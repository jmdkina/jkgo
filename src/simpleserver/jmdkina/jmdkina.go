package simpleserver

import (
	"golanger.com/utils"
	"jk/jklog"
	"jkdbs"
	"net/http"
	. "simpleserver"
	. "simpleserver/dbs"
	"strconv"
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
	filename := s.Path() + "/jmdkina/jmdkina.html"
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
		out := s.queryImages(r)
		s.WriteSerialData(w, out, 200)
		// jklog.L().Debugln(out)
	case "query_images_more":
		out := s.queryImages(r)
		s.WriteSerialData(w, out, 200)
	}
}

func (s *Jmdkina) queryImages(r *http.Request) []utils.M {
	index := r.FormValue("index")
	length := r.FormValue("length")
	out := []utils.M{}
	limit, _ := strconv.Atoi(length)
	skip, _ := strconv.Atoi(index)
	mc := jkdbs.MongoCondition{
		Limit: limit,
		Skip:  skip,
		Order: false,
	}
	GlobalDBS().Query("proj", "images", mc, &out)
	return out
}
