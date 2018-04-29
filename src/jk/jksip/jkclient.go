package jksip

import (
	"jk/jklog"

	"github.com/stefankopieczek/gossip/base"
	"github.com/stefankopieczek/gossip/transaction"
	"github.com/stefankopieczek/gossip/transport"
)

type JKSipClient struct {
	m  transport.Manager
	mc *transaction.Manager
}

func NewJKSipClient() (*JKSipClient, error) {

	sc := &JKSipClient{}

	sc.m, _ = transport.NewManager("udp")
	sc.mc, _ = transaction.NewManager(sc.m, "localhost:5060")
	return sc, nil
}

func (sc *JKSipClient) Recv() error {

	return nil
}

func (sc *JKSipClient) Send() error {

	var sipurl base.Uri
	var sipheader []base.SipHeader
	r := base.NewRequest(INVITE_NAME, sipurl, "SIP/2.0", sipheader, "This is a test")
	if r == nil {
		jklog.L().Errorln("error new request")
	}
	jklog.L().Infof("Create request %s\n", r.String())
	// sc.mc.Send(r, "localhost:5061")
	return nil
}
