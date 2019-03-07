package jksys

import (
	"jk/jkprotocol"
	"jk/jksys"
	cbase "jkservice/jkbase"
	"time"

	l4g "github.com/alecthomas/log4go"
	"golanger.com/log"
)

type SysClient struct {
	cbase.ClientBase
	SysInfo *jksys.KFSystemInfo
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
	sys.SysInfo = jksys.NewSystemInfo()
	sys.Register()
	return sys, nil
}

func (sc *SysClient) Register() {
	str, _ := jkprotocol.JKProtoV6MakeRegister("jksys", sc.SysInfo)
	log.Debug("Register send %s", str)
	sc.Send(str)
}

func (sc *SysClient) Keepalive(interval time.Duration) {
	sc.KeepaliveCycle(interval, "jksys")
}
