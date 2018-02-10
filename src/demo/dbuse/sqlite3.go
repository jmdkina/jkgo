package main

import (
	"jk/jklog"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"database/sql"
	"flag"
)

func main() {
	dbname := flag.String("dbname", "", "db name")
	flag.Parse()
	db, err := sql.Open("sqlite3", *dbname)
	if err != nil {
		jklog.L().Errorln(err)
		return
	}
	defer db.Close()

	sqlStmt := `select * from collection_data`

	rows, err := db.Query(sqlStmt)
	if err != nil {
		jklog.L().Errorln(err)
	}
	defer rows.Close()
	for rows.Next() {
		var collection_name int
		var record_id string
		err = rows.Scan(&collection_name, &record_id)
		if err != nil {
			jklog.L().Errorln(err)
		}
		fmt.Println(collection_name, record_id)
	}
}
