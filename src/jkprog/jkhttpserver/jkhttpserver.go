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

func handler(w http.ResponseWriter, r *http.Request) {

	buf := make([]byte, 1024)
	n, _ := r.Body.Read(buf)

	jklog.L().Infoln("This is message ws handler ", string(buf[:n]))

	w.Write([]byte("Response of handler."))
}

func main() {
	flag.Parse()

	http.HandleFunc("/message/ws", handler)
	http.Handle("/", http.FileServer(http.Dir(*dir)))
	inPort := *port
	if inPort[0] != ':' {
		inPort = ":" + inPort
	}
	jklog.L().Infof("Start server in %s with local %s\n", inPort, *dir)
	http.ListenAndServe(inPort, nil)

	jklog.L().Errorln("Failed listen in %s with local %s\n", inPort, *dir)
}
