package simpleserver

import (
	"net/http"
	"jk/jklog"
	"helper"
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
	jklog.L().Debugln("deal with command ", cmd)
	switch cmd {
	case "query_dbs":
		host := r.FormValue("host")
		if b.m == nil {
			b.m = helper.NewMongo("mongodb://" + host + "/")
		}
		dbs, _ := b.m.DBSession().DatabaseNames()
		b.WriteSerialData(w, dbs, 200)
		break;
	case "query_colls":
		dbname := r.FormValue("dbname")
		d := b.m.DB(dbname)
		c, _ := d.CollectionNames()
		b.WriteSerialData(w, c, 200)
		break;
	}
}
