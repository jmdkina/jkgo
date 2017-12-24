package jkdbs

type JKDBS interface {
	Query(dbname string, coll string, condition interface{}, result interface{}) error
	Update(dbname string, coll string, condition interface{}) error
	Add(dbname string, coll string, data interface{}) error
	Remove(dbname string, coll string, condition interface{}) error
}

type JKDBSBase struct {
}
