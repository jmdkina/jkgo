package jkdbs

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type CMMysql struct {
	db  *sql.DB
	dsn string
}

func NewCMMysql(dsn string) (*CMMysql, error) {
	ms := &CMMysql{
		dsn: dsn,
	}
	var err error
	ms.db, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return ms, nil
}

func (ms *CMMysql) Close() {
	if ms.db != nil {
		ms.db.Close()
	}
}

func (ms *CMMysql) Query(dbname string, coll string, condition interface{}, result interface{}) error {
	/*
		stat, _ := ms.db.Prepare(coll)
		defer stat.Close()
		rows, err := stat.Query()
		defer rows.Close()

		for rows.Next() {
			rows.Scan()
		}
		err = rows.Err()
	*/
	return nil
}

func (ms *CMMysql) Update(dbname string, coll string, condition interface{}) error {
	return nil
}

func (ms *CMMysql) Add(dbname string, coll string, data interface{}) error {

	return nil
}

func (ms *CMMysql) Remove(dbname string, coll string, condition interface{}) error {
	return nil
}
