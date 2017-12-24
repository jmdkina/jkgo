package main

import (
	"flag"
	"fmt"
	"jk/jklog"
	"net/http"
	"os"
	ss "simpleserver"
	. "simpleserver/dbs"
	"strconv"
)

var (
	port = flag.Int("port", 12306, "Listen port")
	path = flag.String("htmlpath", "", "Html path")
)

var (
	VERSION    string
	BUILD_TIME string
	GOVERSION  string
)

func main() {
	fmt.Printf("VERSION: %s\nBUILD_TIME: %s\nGO_VERSION: %s\n", VERSION, BUILD_TIME, GOVERSION)
	flag.Parse()
	html_path := *path
	if len(*path) == 0 {
		curpath, _ := os.Getwd()
		html_path = curpath + "/html"
	}

	http.Handle("/css/", http.FileServer(http.Dir(html_path)))
	http.Handle("/js/", http.FileServer(http.Dir(html_path)))
	http.Handle("/addon/", http.FileServer(http.Dir(html_path)))
	http.Handle("/images/", http.FileServer(http.Dir(html_path)))
	http.Handle("/tools/", http.FileServer(http.Dir(html_path)))

	ss.NewNotFound(html_path)
	ss.NewIndex(html_path)
	ss.NewDirServer(html_path)
	ss.NewUploadServer(html_path)
	ss.NewWebGL(html_path)
	ss.NewWebSocket(html_path)
	ss.NewDBMongo(html_path)
	ss.NewStock(html_path)
	ss.NewJmdkina(html_path)
	ss.NewProject(html_path)
	ss.NewShici(html_path)
	ss.NewResume(html_path)

	lport := *port

	GlobalDBSMongoCreate("mongodb://localhost/")

	jklog.L().Infof("Listen port %d\n", lport)
	http.ListenAndServe(":"+strconv.Itoa(lport), nil)
}
