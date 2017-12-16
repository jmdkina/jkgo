package jksys

import (
	l4g "github.com/alecthomas/log4go"
	cbase "jk/jkservice/jkbase"
	"time"
)

type SysClient struct {
	cbase.ClientBase
}

func NewSysClient(addr string, port int, nettype int) (*SysClient, error) {
	l4g.Debug("New service sys client [%s:%d] \n", addr, port)
	sys := &SysClient{}
	err := sys.New(addr, port, nettype)
	if err != nil {
		return nil, err
	}

	err = sys.Dial()
	if err != nil {
		return nil, err
	}
	return sys, nil
}

func (sc *SysClient) Keepalive(interval time.Duration) {
	sc.KeepaliveCycle(interval)
}
