package jksys

import (
	// l4g "github.com/alecthomas/log4go"
	"errors"
	"jkbase"
	"net"
)

type SysServer struct {
	jkbase.JKNetBase
}

func (ss *SysServer) HandleMsg(conn net.Conn, data string) error {
	return errors.New("Unimplement")
}

func NewSysServer(addr string, port int, nettype int) (*SysServer, error) {
	sys := &SysServer{}
	err := sys.New(addr, port, nettype)
	if err != nil {
		return nil, err
	}

	sys.Listen()
	return sys, nil
}

func (sys *SysServer) RecvCycle() {
	sys.DoRecvCycle(sys)
}
