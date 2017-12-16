package jkstatus

import (
	"net"
	"time"
)

type RemoteInfo struct {
	Remote   string
	First    int64
	Last     int64
	Online   bool
	Interval int
}

type RemoteInstance struct {
	Conn net.Conn
	Info RemoteInfo
}

func NewRemoteInstance(conn net.Conn, inter int) (*RemoteInstance, error) {
	rinfo := RemoteInfo{
		Remote:   conn.RemoteAddr().String(),
		First:    time.Now().Unix(),
		Last:     time.Now().Unix(),
		Online:   true,
		Interval: inter,
	}
	ri := &RemoteInstance{
		Conn: conn,
		Info: rinfo,
	}
	return ri, nil
}

func (ri *RemoteInstance) Update() {
	now := time.Now().Unix()
	if now-ri.Info.Last > int64(ri.Info.Interval) {
		ri.Info.Online = false
	}
}

func (ri *RemoteInstance) UpdateTime() {
	ri.Info.Last = time.Now().Unix()
	ri.Info.Online = true
}
