package main

import (
	"flag"
	"fmt"
	"jk/jklog"
	"jkbase"
	"net/http"
	ss "simpleserver"
	. "simpleserver/dbs"
	. "simpleserver/jmdkina"
	. "simpleserver/shici"
	. "simpleserver/ws"
	"strconv"
)

var (
	conf    = flag.String("conf", "etc/simpleserver.json", "Config File")
	forerun = flag.Bool("forerun", false, "Foreground run")
)

var (
	VERSION    string
	BUILD_TIME string
	GOVERSION  string
)

func main() {
	fmt.Printf("\rVERSION: %s\n\rBUILD_TIME: %s\n\rGO_VERSION: %s\n\n", VERSION, BUILD_TIME, GOVERSION)
	flag.Parse()

	if !*forerun {
		jkbase.InitDeamon(true)
	}

	ss.GlobalSetConfig(*conf)
	html_path := ss.GlobalBaseConfig().HtmlPath

	jklog.L().InitLog(ss.GlobalBaseConfig().LogFile)

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
	NewJmdkina(html_path)
	ss.NewProject(html_path)
	NewShici(html_path)
	NewShiciAdd(html_path)
	ss.NewResume(html_path)
	ss.NewResumeEn(html_path)
	ss.NewResumeSet(html_path)
	NewJmdkinaAdd(html_path)
	NewWSSimplePage(html_path)
	ss.NewManager(html_path)

	lport := ss.GlobalBaseConfig().Port

	if ss.GlobalBaseConfig().DBType == "Mongo" {
		GlobalDBSMongoCreate(ss.GlobalBaseConfig().DBUrl)
	}

	jklog.L().Infof("Listen port %d\n", lport)
	http.ListenAndServe(":"+strconv.Itoa(lport), nil)
}
