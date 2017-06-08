package jkdbs

import (
	"jk/jknetbase"
	"jk/jkprotocol"
	"github.com/alecthomas/log4go"
	"time"
)

type ServiceClientCtrl struct {
	client jknetbase.JKNetBaseClient
	proto *jkprotocol.JKProtocolWrap
}

func NewServiceClientCtrl(addr string, port int, nettype int) (*ServiceClientCtrl, error) {
	c := &ServiceClientCtrl{}
	c.client = jknetbase.JKNetBaseClient{}
	err := c.client.New(addr, port, nettype)
	if err != nil {
		return nil, err
	}
	c.proto, err = jkprotocol.NewJKProtocolWrap(jkprotocol.JK_PROTOCOL_VERSION_5)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *ServiceClientCtrl) DoRun() error {
	for {
		str, err := c.proto.Keepalive("")
		if err != nil {
			return err
		}
		n := c.client.Send(str)
		log4go.Debug("Send keepalive %d", n)
		recv, err := c.client.Recv()
		if err != nil {
			return err
		}
		log4go.Debug("Recv response of keepalive %s", string(recv))
		time.Sleep(time.Second * 30)
	}

	return nil
}