package jksys

import (
	// l4g "github.com/alecthomas/log4go"
	"errors"
	"jk/jklog"
	"jkbase"
	"net"
)

type SysServer struct {
	jkbase.JKNetBase
}

func (ss *SysServer) handleMsg(conn net.Conn, data string) error {
	return errors.New("Unimplement")
}

func NewSysServer(addr string, port int, nettype int) (*SysServer, error) {
	sys := &SysServer{}
	sys.HandleMsg = sys.handleMsg
	err := sys.New(addr, port, nettype)
	if err != nil {
		return nil, err
	}

	err = sys.Listen()
	if err != nil {
		return nil, err
	}
	return sys, nil
}

func Start(address string, port int, client_addr string, client_port int, recv bool, client bool) (*SysServer, *SysClient, error) {
	s, err := NewSysServer(address, port, jkbase.NetTypeBase)
	if err != nil {
		return nil, nil, err
	}
	jklog.L().Infoln("jksys Start recv data")
	if recv {
		go s.DoRecvCycle()
	}

	var c SysClient
	if client {
		c, err := NewSysClient(client_addr, client_port, jkbase.NetTypeBase)
		if err != nil {
			jklog.L().Errorln("New sys client error ", err)
		} else {
			c.Keepalive(30)
		}
	}
	return s, &c, nil
}
