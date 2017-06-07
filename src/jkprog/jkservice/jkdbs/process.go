package main

import (
	"jk/jkservice/jkdbs"
	l4g "github.com/alecthomas/log4go"
)

type Process struct {
	address string
	port    int
}

func (p *Process) start(addr string, port int) {
	p.address = addr
	p.port = port
	dbs, err := jkdbs.NewServiceDBS(addr, port)
	if err != nil {
		l4g.Debug("create dbs failed ", err)
		return
	}
	l4g.Info("Start recv data")
	dbs.Recv()
}