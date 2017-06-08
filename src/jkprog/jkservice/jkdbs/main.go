package main

import (
	l4g "github.com/alecthomas/log4go"
	"time"
	"flag"
	"jk/jknetbase"
	"jk/jkservice/jkdbs"
)

var (
	address = flag.String("address", "0.0.0.0", "Listen local address")
	port    = flag.Int("port", 20002, "Listen local port")
	bg      = flag.Bool("bg", false, "true|false")
	logfile = flag.String("logfile", "/tmp/jkdbs.log", "log file")
	logsize = flag.Int("logsize", 1024*1024*1024, "log size")
	client_addr = flag.String("client address", "0.0.0.0", "dial client address")
	client_port    = flag.Int("client port", 20101, "dial client port")
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

	go func() {
		for {
			p, err := jkdbs.NewServiceClientCtrl(*client_addr, *client_port, 1)
			if err != nil {
				l4g.Error("service client ctrl failed ", err)
				time.Sleep(time.Second * 5)
			} else {
				l4g.Info("Start service client ctrl with [%s, %d]", *client_addr, *client_port)
				p.DoRun()
			}
		}
	}()

	for {
		time.Sleep(time.Millisecond*500)
	}
}
