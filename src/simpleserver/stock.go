package simpleserver

import (
	"encoding/base64"
	"helper"
	"jk/jklog"
	"net/http"
)

type StockIdeas struct {
	Date     string
	Analyse  string
	Remember string
	Key      string
}

type Stock struct {
	Base
}

func NewStock(path string) *Stock {
	i := &Stock{}
	i.SetPath(path)
	i.SetFunc("/stock", i)
	return i
}

func (s *Stock) Get(w http.ResponseWriter, r *http.Request) {
	sp := SimpleParse{}
	filename := s.path + "/stock.html"
	jklog.L().Debugf("Get html [%s]\n", filename)

	err := sp.Parse(w, "", filename)
	if err != nil {
		jklog.L().Errorln("Parse error ", err)
	}
}

func (s *Stock) Post(w http.ResponseWriter, r *http.Request) {
	method := r.FormValue("op")
	code := r.FormValue("code")
	jkcmd := r.FormValue("jk")
	if code != "jmdstock" || jkcmd != "stockoperation" {
		jklog.L().Errorf("Permission failed\n")
		s.WriteSerialData(w, "", 403)
		return
	}
	jklog.L().Debugf("stock operatin with method %v\n", method)
	switch method {
	case "todayall":
		break
	case "query_ideas":
		m := helper.NewMongo("mongodb://" + "localhost" + "/")
		c := m.DB("stock").C("mrjl")
		si := []StockIdeas{}
		err := c.Find(nil).Sort("-date").All(&si)
		if err != nil {
			s.WriteSerialData(w, "data error", 400)
		} else {
			s.WriteSerialData(w, si, 200)
		}
		return
	case "addnew":
		date := r.FormValue("date")
		analyse_p := r.FormValue("analyse")
		analyse_t, _ := base64.StdEncoding.DecodeString(analyse_p)
		analyse := string(analyse_t)
		remember := r.FormValue("remember")
		key := r.FormValue("key")
		m := helper.NewMongo("mongodb://" + "localhost" + "/")
		c := m.DB("stock").C("mrjl")
		si := &StockIdeas{
			Date:     date,
			Analyse:  analyse,
			Remember: remember,
			Key:      key,
		}
		jklog.L().Debugln("insert of date ", date)
		err := c.Insert(si)
		if err != nil {
			s.WriteSerialData(w, "insert error", 500)
		} else {
			s.WriteSerialData(w, "", 200)
		}
		return
	}
	s.WriteSerialData(w, "no data", 200)
}
