package jkstatus

import (
	"jk/jknetbase"
	"net"
	l4g "github.com/alecthomas/log4go"
)

type ServiceStatus struct {
	jknetbase.JKNetBaseRecv
}

func (ctrl *ServiceStatus) handler_msg(conn net.Conn, data string) error {
	l4g.Debug("handler msg of jkstatus")
	p, _ := ParseData(data, conn)
	p.HandleMsg()
	return nil
}

func NewServiceStatus(addr string, port int) (*ServiceStatus, error) {
	ctrl := &ServiceStatus{}
	err := ctrl.New(addr, port, 1)
	if (err != nil) {
		return nil, err
	}

	ctrl.SetHandlerMsg(ctrl.handler_msg)
	return ctrl, nil
}
