package jksys

import (
	l4g "github.com/alecthomas/log4go"
	"jkbase"
	"net"
)

type SysServer struct {
	jkbase.JKNetBase
}

func handler_msg(conn net.Conn, data string) error {
	l4g.Debug("handler msg of sys")
	conn.Write([]byte("Hello give you a response"))
	return nil
}

func NewSysServer(addr string, port int, nettype int) (*SysServer, error) {
	sys := &SysServer{}
	err := sys.New(addr, port, nettype)
	if err != nil {
		return nil, err
	}

	sys.SetHandlerMsg(handler_msg)
	sys.Listen()
	return sys, nil
}
