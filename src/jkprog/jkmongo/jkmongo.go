package main

import (
	"flag"
	"helper"
	. "golanger.com/middleware"
	"golanger.com/utils"
	"time"
	"jk/jklog"
	"encoding/json"
	"io/ioutil"
)

func insertItem(db, coll string, content interface{}) error {
    m := Middleware.Get("db").(*helper.Mongo)
	c := m.C(utils.M{"name":coll})

	return c.Insert(content)
}

var (
	dab = flag.String("db", "proj", "data base name")
	coll = flag.String("coll", "test", "collection name")
	filepath = flag.String("filepath", "/opt/data/proj/ConfigDemo/sql", "file path")
	filename = flag.String("filename", "zy.json", "filename")
)

func main() {
	flag.Parse()

	data,err := ioutil.ReadFile(*filepath+"/"+*filename)
	if err != nil {
		jklog.L().Error("error read data\n")
		return
	}
	i := []utils.M{}
	err = json.Unmarshal(data, &i)
	if err != nil {
		jklog.L().Errorln("error parse data")
		return
	}
	jklog.L().Infoln("len of items: ", len(i))

	m := helper.NewMongo("mongodb://127.0.0.1/" + *dab)
	defer m.Close()
	Middleware.Add("db", m)

	for j:=0; j< len(i);j++ {
		i[j]["createtime"] = time.Now().Unix()
		i[j]["updatetime"] = time.Now().Unix()
		rt := i[j]["recordtime"].(string)
		if len(rt) > 0 {
			t, err := time.Parse("2006-01-02-15-04-05 Z0700", rt)
			if err != nil {
				jklog.L().Errorln("error parse data ", rt)
				continue
			}
			i[j]["recordtime"] = t.Unix()
		}
		jklog.L().Infoln("start to insert index: ", j)

		mm := Middleware.Get("db").(*helper.Mongo)
		c := mm.C(utils.M{"name": *coll})
		if n, _ := c.Find(utils.M{"content":i[j]["content"].(string), "title":i[j]["title"].(string)}).Count(); n > 0 {
			jklog.L().Warnln("item exist. ")
			continue
		}

		err = insertItem(*dab, *coll, i[j])
		if err != nil {
			jklog.L().Errorf("insert failed %d: %v\n", j, err)
		}
		jklog.L().Infof("Insert down %d\n", j)
	}
}
