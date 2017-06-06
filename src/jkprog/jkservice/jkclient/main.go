package main

import (
	"flag"
	"jk/jknetbase"
	"time"
	l4g "github.com/alecthomas/log4go"
	"jk/jkservice/jkclient"
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

	l4g.Debug("jkclient start")
	l4g.Info("start with [%s:%d]", *address, *port)

	jknetbase.InitDeamon(*bg)

	c, err := jkclient.NewDemoClient(*address, *port)
	if err != nil {
		l4g.Debug("create demo client failed ", err)
		return
	}
	for {
		l4g.Info("-------------- start to do something")
		n := c.Send("send data out demo")
        l4g.Info("send data out len %d", n)
        data, err := c.Recv()
		if err != nil {
			l4g.Error("recv error ", err)
			break
		}
		l4g.Info("recv data ", data)
		time.Sleep(time.Millisecond*500)
	}
}