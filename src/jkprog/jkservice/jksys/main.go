package main

import (
	"flag"
	l4g "github.com/alecthomas/log4go"
	"jk/jkservice/jksys"
	"jkbase"
	"time"
)

var (
	address     = flag.String("address", "0.0.0.0", "Listen local address")
	port        = flag.Int("port", 20103, "Listen local port")
	bg          = flag.Bool("bg", false, "true|false")
	logfile     = flag.String("logfile", "/tmp/jksys.log", "log file")
	logsize     = flag.Int("logsize", 1024*1024*1024, "log size")
	client_addr = flag.String("client address", "0.0.0.0", "dial client address")
	client_port = flag.Int("client port", 20101, "dial client port")
)

func main() {
	flag.Parse()

	jkbase.InitLog(*logfile, *logsize)

	l4g.Debug("jksys start")
	l4g.Info("Listen with [%s:%d]", *address, *port)

	jkbase.InitDeamon(*bg)

	s, err := jksys.NewSysServer(*address, *port, jkbase.NetTypeBase)
	if err != nil {
		l4g.Debug("create service ctrl failed ", err)
		return
	}
	l4g.Info("Start recv data")
	s.RecvCycle()

	c, err := jksys.NewSysClient(*client_addr, *client_port, jkbase.NetTypeBase)
	if err != nil {
		l4g.Error("New Sys client error ", err)
	} else {
		c.Keepalive(30)
	}
	for {
		time.Sleep(time.Millisecond * 500)
	}
}
