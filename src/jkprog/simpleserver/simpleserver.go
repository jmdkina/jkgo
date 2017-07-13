package main

import (
	ss "simpleserver"
	"net/http"
	"jk/jklog"
	"strconv"
)

func main() {
	b := ss.Base{}
	http.HandleFunc("/index", b.ServeHttp)

	port := 12306

	jklog.L().Debugf("Listen port %d\n", port)
	http.ListenAndServe(":" + strconv.Itoa(port), nil)
}