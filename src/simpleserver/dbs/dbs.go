package simpleserver

import (
	"jkdbs"
)

type DBS struct {
	m *jkdbs.Mongo
}

var DBSInstance DBS

func GlobalDBSMongoCreate(url string) error {
	DBSInstance.m = jkdbs.NewMongo(url)
	return nil
}

func GlobalDBS() *jkdbs.Mongo {
	return DBSInstance.m
}
