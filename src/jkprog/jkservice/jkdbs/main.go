package main

import (
	"jk/jkservice/jkdbs"
	l4g "github.com/alecthomas/log4go"
	"time"
	"flag"
	"jk/jknetbase"
)

var (
	address = flag.String("address", "0.0.0.0", "Listen local address")
	port    = flag.Int("port", 20002, "Listen local port")
	bg      = flag.Bool("bg", false, "true|false")
	logfile = flag.String("logfile", "/tmp/jkdbs.log", "log file")
	logsize = flag.Int("logsize", 1024*1024*1024, "log size")
)

func main() {
	flag.Parse()

	jknetbase.InitLog(*logfile, *logsize)

	l4g.Debug("jkdbs start")
	l4g.Info("Listen with [%s:%d]", *address, *port)

	jknetbase.InitDeamon(*bg)

	dbs, err := jkdbs.NewServiceDBS(*address, *port)
	if err != nil {
		l4g.Debug("create dbs failed ", err)
		return
	}
	l4g.Info("Start recv data")
	dbs.Recv()
	for {
		time.Sleep(time.Millisecond*500)
	}
}
