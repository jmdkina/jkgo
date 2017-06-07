package jkdbs

import (
	"jk/jknetbase"
	"net"
	l4g "github.com/alecthomas/log4go"
)

type ServiceDBS struct {
	jknetbase.JKNetBaseRecv
}

func handler_msg(conn net.Conn, data string) error {
	l4g.Debug("handler msg of dbs")
	conn.Write([]byte("Hello give you a response"))
	return nil
}

func NewServiceDBS(addr string, port int) (*ServiceDBS, error) {
	dbs := &ServiceDBS{}
	err := dbs.New(addr, port, 1)
	if (err != nil) {
		return nil, err
	}

	dbs.SetHandlerMsg(handler_msg)
	return dbs, nil
}
