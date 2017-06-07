package main

import (
	l4g "github.com/alecthomas/log4go"
	"time"
	"flag"
	"jk/jknetbase"
)

var (
	address = flag.String("address", "0.0.0.0", "Listen local address")
	port    = flag.Int("port", 20001, "Listen local port")
	bg      = flag.Bool("bg", false, "true|false")
	logfile = flag.String("logfile", "/tmp/jkctrl.log", "log file")
	logsize = flag.Int("logsize", 1024*1024*1024, "log size")
)

func main() {
	flag.Parse()

	jknetbase.InitLog(*logfile, *logsize)

	l4g.Debug("jkctrl start")
	l4g.Info("Listen with [%s:%d]", *address, *port)

	jknetbase.InitDeamon(*bg)

	p := Process{}
	p.start(*address, *port)
	for {
		time.Sleep(time.Millisecond*500)
	}
}
