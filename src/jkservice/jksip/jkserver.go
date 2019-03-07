package jksip

import (
	"jk/jklog"

	"github.com/stefankopieczek/gossip/base"
	"github.com/stefankopieczek/gossip/transaction"
	"github.com/stefankopieczek/gossip/transport"
)

type JKSipServer struct {
	m  transport.Manager
	mc *transaction.Manager
}

func NewJKSipServer() (*JKSipServer, error) {
	ss := &JKSipServer{}

	ss.m, _ = transport.NewManager("udp")
	ss.mc, _ = transaction.NewManager(ss.m, "localhost:5061")
	return ss, nil
}

func (ss *JKSipServer) Recv() error {
	x := <-ss.mc.Requests()
	d := x.Destination()
	jklog.L().Infoln("d: ", d)

	y := x.Origin()
	jklog.L().Infoln("origin: ", y)
	var msg base.SipMessage
	x.Receive(msg)

	return nil
}

func (ss *JKSipServer) Send() error {

	return nil
}
