package simpleserver

import (
	"net/http"
	"jk/jklog"
	"helper"
	"golanger.com/utils"
	"encoding/json"
)

type DBMongo struct {
	Base
	m *helper.Mongo
}

func NewDBMongo(path string) *DBMongo {
	i := &DBMongo{}
	i.SetPath(path)
	i.SetFunc("/db", i)
	return i
}

func (b *DBMongo) Get(w http.ResponseWriter, r *http.Request) {
	sp := SimpleParse{}
	filename := b.path + "/db/db.html"
	jklog.L().Debugf("Get html [%s]\n", filename)

	err := sp.Parse(w, filename, "")
	if err != nil {
		jklog.L().Errorln("Parse error ", err)
	}
}

func (b *DBMongo) Post(w http.ResponseWriter, r *http.Request) {
	cmd := r.FormValue("jk")
	res := utils.M{
		"Status":200,
	}
	jklog.L().Debugln("deal with command ", cmd)
	switch cmd {
	case "query_dbs":
		host := r.FormValue("host")
		if b.m == nil {
			b.m = helper.NewMongo("mongodb://" + host + "/")
		}
		dbs, _ := b.m.DBSession().DatabaseNames()
		res["Result"] = dbs
		d, _ := json.Marshal(res)
		w.Write(d)
		break;
	case "query_colls":
		break;
	}
}
