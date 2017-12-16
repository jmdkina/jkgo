package jkstatus

import (
	"net"
	"time"
)

type RemoteInstance struct {
	Conn     net.Conn
	First    int64
	Last     int64
	Online   bool
	Interval int
}

func NewRemoteInstance(conn net.Conn, inter int) (*RemoteInstance, error) {
	ri := &RemoteInstance{
		Conn:     conn,
		First:    time.Now().Unix(),
		Last:     time.Now().Unix(),
		Online:   true,
		Interval: inter,
	}
	return ri, nil
}

func (ri *RemoteInstance) Update() {
	now := time.Now().Unix()
	if now-ri.Last > int64(ri.Interval) {
		ri.Online = false
	}
}

func (ri *RemoteInstance) UpdateTime() {
	ri.Last = time.Now().Unix()
	ri.Online = true
}
