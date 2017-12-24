package jkdbs

import (
	"golanger.com/utils"
	"testing"
)

func TestMongo(t *testing.T) {
	m := NewMongo("mongodb://localhost/")
	out := []utils.M{}
	err := m.Query("stock", "bills", nil, &out)
	if err != nil {
		t.Fatal("Query failed ", err)
	}
	t.Log(out)
}
