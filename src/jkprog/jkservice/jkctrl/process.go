package main

import (
	"jk/jkservice/jkctrl"
	l4g "github.com/alecthomas/log4go"
)

type Process struct {
	address string
	port    int
}

func (p *Process) start(address string, port int) {
	ctrl, err := jkctrl.NewServiceCtrl(address, port)
	if err != nil {
		l4g.Debug("create service ctrl failed ", err)
		return
	}
	l4g.Info("Start recv data")
	ctrl.Recv()
}