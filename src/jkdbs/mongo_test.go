package jkdbs

import (
	"testing"

	"golanger.com/utils"
)

func TestMongo(t *testing.T) {
	m := NewMongo("mongodb://localhost/")
	out := []utils.M{}
	mc := MongoCondition{
		Limit: 9,
		Skip:  3,
		Order: false,
	}
	err := m.Query("proj", "images", mc, &out)
	if err != nil {
		t.Fatal("Query failed ", err)
	}
	t.Log(out)
}

func TestMongoInsert(t *testing.T) {
	m := NewMongo("mongodb://localhost/")

	i1 := utils.M{

		"author":     "jmd",
		"content":    "",
		"path":       "start",
		"createtime": "12345677",
		"updatetime": "name",
	}
	// in = append(in, i2)
	err := m.Add("proj", "images", i1)
	if err != nil {
		t.Fatal("Insert failed ", err)
	}
}
