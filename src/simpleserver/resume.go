package simpleserver

import (
	"jk/jklog"
	"net/http"
	"jkdbs"
	"io/ioutil"
	"encoding/json"
	. "simpleserver/dbs"
	"os"
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

type ResumeEn struct{
    Base
}

func NewResumeEn(path string) *ResumeEn {
	j := &ResumeEn{}
	j.SetPath(path)
	j.SetFunc("/resume_en", j)
	return j
}

func (s *ResumeEn) Get(w http.ResponseWriter, r *http.Request) {
	sp := SimpleParse{}
	filename := s.path + "/resume/resume_en.html"
	jklog.L().Debugf("Get html [%s]\n", filename)

	err := sp.Parse(w, filename, "")
	if err != nil {
		jklog.L().Errorln("Parse error ", err)
	}
}

type ResumeSet struct {
	Base
}

func NewResumeSet(path string) *ResumeSet {
	j := &ResumeSet{}
	j.SetPath(path)
	j.SetFunc("/resume_set", j)
	return j
}

func (s *ResumeSet) Get(w http.ResponseWriter, r *http.Request) {
	sp := SimpleParse{}
	filename := s.path + "/resume/resume_set.html"
	jklog.L().Debugf("Get htmle [%s]\n", filename)

	err := sp.Parse(w, filename, "")
	if err != nil {
		jklog.L().Errorln("Parse error ", err)
	}
}

type ResumeBaseInfoBase struct {
}

type ResumeBaseInfo struct {
	BaseInfo        ResumeBaseInfoBase
}

func (s *ResumeSet) Post(w http.ResponseWriter, r *http.Request) {
	cmd := r.FormValue("cmd")
	switch cmd {
	case "query_info":
		var out []interface{}
	    mc := jkdbs.MongoCondition{
			Limit: 10,
			Skip:  0,
			Order: false,
	    }
		var content []byte
		err := GlobalDBS().Query("proj", "resume", mc, &out)
		if err != nil || len(out) == 0 {
			content, _ = ioutil.ReadFile(s.path + "/resume/template.json")
			jklog.L().Infoln("Get from fiel")
		} else {
			content, err = json.Marshal(out[0])
			if err != nil {
				jklog.L().Errorln("Error parse from database ", err)
			} else {
				jklog.L().Infoln("Get from DB")
			}
		}
		s.WriteSerialData(w, string(content), 200)
		return
	case "change_new":
		code := r.FormValue("code")
		if code != "jkresumeset" {
			s.WriteSerialData(w, "No Permission", 403)
			return
		}
		data := r.FormValue("content")
		var out interface{}
	    err := json.Unmarshal([]byte(data), &out)
		if err != nil {
			s.WriteSerialData(w, "Invalid Str", 400);
			jklog.L().Errorln("error parse ", err)
		} else {
			ioutil.WriteFile(s.path + "/resume/template.json", []byte(data), os.ModePerm)
			GlobalDBS().Remove("proj", "resume", nil)
			err := GlobalDBS().Add("proj", "resume", out)
			if err != nil {
				jklog.L().Errorln("write data base failed ", err)
				s.WriteSerialData(w, "write db fail", 401)
			} else {
				s.WriteSerialData(w, "", 200);
		    }
		}
		return
	}
}

