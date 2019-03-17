package simpleserver

import (
	"jk/jklog"
	"net/http"
	. "simpleserver"
	. "simpleserver/dbs"
	"time"

	"golanger.com/utils"
)

type ShiciAdd struct {
	Base
}

func NewShiciAdd(path string) *ShiciAdd {
	j := &ShiciAdd{}
	j.SetPath(path)
	j.SetFunc("/shiciadd", j)
	return j
}

func (s *ShiciAdd) Get(w http.ResponseWriter, r *http.Request) {
	sp := SimpleParse{}
	filename := s.Path() + "/shici/add.html"
	jklog.L().Debugf("Get html [%s]\n", filename)

	err := sp.Parse(w, "", filename)
	if err != nil {
		jklog.L().Errorln("Parse error ", err)
	}
}

func (s *ShiciAdd) Post(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	if code != "shiciadd" {
		return
	}
	cmd := r.FormValue("cmd")
	jklog.L().Debugf("shiciadd page -- cmd %s\n", cmd)
	switch cmd {
	case "add":
		err := s.addShici(r)
		if err == nil {
			s.WriteSerialData(w, "", 200)
		} else {
			s.WriteSerialData(w, "NoPermission", 403)
		}
		break
	}
}

func (s *ShiciAdd) addShici(r *http.Request) error {
	title := r.FormValue("title")
	subtitle := r.FormValue("subtitle")
	author := r.FormValue("author")
	content := r.FormValue("content")
	recordtime := time.Now().Unix()
	createtime := r.FormValue("createtime")

	item := utils.M{
		"title":      title,
		"subtitle":   subtitle,
		"author":     author,
		"content":    content,
		"recordtime": recordtime,
		"createtime": createtime,
	}
	GlobalDBS().Add("proj", "shici", item)
	return nil
}
