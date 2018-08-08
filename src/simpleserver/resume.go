package simpleserver

import (
	"jk/jklog"
	"net/http"
	"io/ioutil"
	"encoding/json"
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
		content, _ := ioutil.ReadFile(s.path + "/resume/template.json")
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
			s.WriteSerialData(w, "", 200);
		}
		return
	}
}

