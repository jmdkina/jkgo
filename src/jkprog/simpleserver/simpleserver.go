package main

import (
	ss "simpleserver"
	"net/http"
	"jk/jklog"
	"strconv"
	"os"
	"flag"
)

var (
	port    = flag.Int("port", 12306, "Listen port")
	path      = flag.String("htmlpath", "", "Html path")
)

func main() {
	flag.Parse()
	html_path := *path
	if len(*path) == 0 {
		curpath, _ := os.Getwd()
		html_path = curpath + "/html"
	}

	http.Handle("/css/", http.FileServer(http.Dir(html_path)))
	http.Handle("/js/", http.FileServer(http.Dir(html_path)))

	ss.NewNotFound(html_path)
	ss.NewIndex(html_path)

	lport := *port

	jklog.L().Debugf("Listen port %d\n", lport)
	http.ListenAndServe(":" + strconv.Itoa(lport), nil)
}