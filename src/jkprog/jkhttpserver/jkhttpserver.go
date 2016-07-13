package main

import (
	"flag"
	"jk/jklog"
	"net/http"
)

var (
	port = flag.String("port", ":8080", "Server port")
	dir  = flag.String("dir", "./", "Files position")
)

func main() {
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*dir)))
	inPort := *port
	if inPort[0] != ':' {
		inPort = ":" + inPort
	}
	jklog.L().Infof("Start server in %s with local %s\n", inPort, *dir)
	http.ListenAndServe(inPort, nil)

	jklog.L().Errorln("Failed listen in %s with local %s\n", inPort, *dir)
}
