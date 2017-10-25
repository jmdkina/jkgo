package jkstatus

import (
	l4g "github.com/alecthomas/log4go"
	"jk/jknetbase"
	"net"
	"time"
)

const (
	jk_status_remote_interval = 30
)

type ServiceStatus struct {
	jknetbase.JKNetBaseRecv
	remoteInstance map[net.Conn]RemoteInstance
}

func (ctrl *ServiceStatus) findRemoteInstance(conn net.Conn) *RemoteInstance {
	return ctrl.remoteInstance[conn]
}

func (ctrl *ServiceStatus) handler_msg(conn net.Conn, data string) error {
	l4g.Debug("handler msg of jkstatus")
	ri := ctrl.findRemoteInstance(conn)
	if ri == nil {
		rii := NewRemoteInstance(conn, jk_status_remote_interval)
		ctrl.remoteInstance[conn] = rii
	}
	p, _ := ParseData(data, conn)
	p.SS = ctrl
	p.HandleMsg()
	return nil
}

func (ctrl *ServiceStatus) updateRemoteInstance() {
	l4g.Debug("go func for update remote instance start")
	for {
		for k, v := range ctrl.remoteInstance {
			v.Update()
		}
		time.Sleep(time.Duration * 1000)
	}
	l4g.Debug("go func for update remote instance end ")
}

func NewServiceStatus(addr string, port int) (*ServiceStatus, error) {
	ctrl := &ServiceStatus{}
	err := ctrl.New(addr, port, 1)
	if err != nil {
		return nil, err
	}

	ctrl.SetHandlerMsg(ctrl.handler_msg)
	go updateRemoteInstance()
	return ctrl, nil
}
