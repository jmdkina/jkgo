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
	filename := s.Path() + "/shici/shici.html"
	jklog.L().Debugf("Get html [%s]\n", filename)

	err := sp.Parse(w, filename, "")
	if err != nil {
		jklog.L().Errorln("Parse error ", err)
	}
}

func (s *Shici) Post(w http.ResponseWriter, r *http.Request) {
	cmd := r.FormValue("cmd")
	switch cmd {
	case "query_shicis":
		out := s.queryItems(r)
		s.WriteSerialData(w, out, 200)
		return
	}
}

func (s *Shici) queryItems(r *http.Request) []utils.M {
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
	GlobalDBS().Query("proj", "shici", mc, &out)
	return out
}
