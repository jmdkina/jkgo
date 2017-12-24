package jkdbs

import (
	"golanger.com/utils"
	"testing"
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
