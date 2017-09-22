//
// +build windows

package main

import (
	"flag"
	l4g "github.com/alecthomas/log4go"
	"jk/jkcenter"
	"time"
)

var (
	address = flag.String("local_address", "0.0.0.0", "Listen local address")
	port    = flag.Int("local_port", 24433, "Listen local port")
	bg      = flag.Bool("bg", false, "true|false")
)

func main() {
	flag.Parse()

	lw := l4g.NewFileLogWriter("/tmp/jkcenter.log", false)
	if lw != nil {
		lw.SetRotateSize(1024 * 1024 * 1024)
		l4g.AddFilter("file", l4g.FINE, lw)
	}
	l4g.Debug("jkcenter start")

	cc, err := jkcenter.InitCenter(*address, *port, 0)
	if err != nil {
		l4g.Error("Error init ", err)
		return
	}

	// thread doing
	err = cc.DoCycle()
	if err != nil {
		l4g.Error("do cycle failed ", err)
		return
	}

	// thread doing
	err = cc.Recv()
	if err != nil {
		l4g.Error("recv failed ", err)
		return
	}

	for {
		time.Sleep(500 * time.Millisecond)
	}
}
