package jkstatus

import (
	"jk/jklog"
	"jkbase"
	"net"
	"time"
)

const (
	jk_status_remote_interval = 30
)

type ServiceStatus struct {
	jkbase.JKNetBase
	remoteInstance map[net.Conn]*RemoteInstance
}

func (ctrl *ServiceStatus) RemoteInstances() []RemoteInstance {
	var ris []RemoteInstance
	for _, v := range ctrl.remoteInstance {
		ris = append(ris, *v)
	}
	return ris
}

func (ctrl *ServiceStatus) findRemoteInstance(conn net.Conn) *RemoteInstance {
	return ctrl.remoteInstance[conn]
}

func (ctrl *ServiceStatus) handleMsg(conn net.Conn, data string) error {
	jklog.L().Debugln("handler msg of jkstatus from ", conn.RemoteAddr().String())

	// Check if client has connected
	// Update device info if connected, create new one if not
	ri := ctrl.findRemoteInstance(conn)
	if ri == nil {
		rii, _ := NewRemoteInstance(conn, jk_status_remote_interval)
		ctrl.remoteInstance[conn] = rii
	}
	p, _ := ParseData(data, conn)
	p.SS = ctrl

	// Parse msg, and give response if need
	// Do other process also.
	p.HandleMsg()
	return nil
}

func (ctrl *ServiceStatus) updateRemoteInstance() {
	jklog.L().Debugln("go func for update remote instance start")
	for {
		// Traverse all connected clients and call update to check if client still connected
		// Mark disconnected if does, set update time if online.
		for _, v := range ctrl.remoteInstance {
			v.Update()
		}
		time.Sleep(time.Millisecond * 1000)
	}
	jklog.L().Debugln("go func for update remote instance end ")
}

func NewServiceStatus(addr string, port int) (*ServiceStatus, error) {
	ctrl := &ServiceStatus{}

	// Set call back functions
	ctrl.HandleMsg = ctrl.handleMsg

	// Create handle
	err := ctrl.New(addr, port, 1)
	if err != nil {
		return nil, err
	}

	ctrl.remoteInstance = make(map[net.Conn]*RemoteInstance)

	// Start listen, wait client connect
	ctrl.Listen()
	go ctrl.updateRemoteInstance()
	return ctrl, nil
}

func Start(addr string, port int, recv bool) (*ServiceStatus, error) {
	st, err := NewServiceStatus(addr, port)
	if err != nil {
		return nil, err
	}
	if recv {
		go st.DoRecvCycle()
	}
	/**
	 * Next provide http server for get status with http
	 */
	shttp, err := NewStatusHttp(12307)
	if err != nil {
		return st, err
	}
	shttp.AddLinkStatus(st)
	return st, nil
}
