package jkctrl

import (
	"jk/jknetbase"
	"net"
	l4g "github.com/alecthomas/log4go"
)

type ServiceCtrl struct {
	jknetbase.JKNetBaseRecv
}

func handler_msg(conn net.Conn, data string) error {
	l4g.Debug("handler msg of ctrl")
	conn.Write([]byte("ctrl service response"))
	return nil
}

func NewServiceCtrl(addr string, port int) (*ServiceCtrl, error) {
	ctrl := &ServiceCtrl{}
	err := ctrl.New(addr, port, 1)
	if (err != nil) {
		return nil, err
	}

	ctrl.SetHandlerMsg(handler_msg)
	return ctrl, nil
}
