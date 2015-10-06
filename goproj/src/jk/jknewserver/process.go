package jknewserver

import (
	. "jk/jkcommon"
	// "jk/jklog"
)

type JKNewServer struct {
}

func JKNewServerNew() *JKNewServer {
	return &JKNewServer{}
}

type JKServerProcessItem struct {
	id         string
	remoteAddr string
	data       []byte
}

type JKServerProcess struct {
	Item []*JKServerProcessItem
}

func (newser *JKNewServer) Start(proc *JKServerProcess) bool {
	return proc.listenLocalTCP(JK_NET_ADDRESS_PORT)
}
