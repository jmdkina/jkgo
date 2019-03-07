package jkbase

import (
	"github.com/alecthomas/log4go"
	"jk/jkprotocol"
	"jkbase"
	"time"
)

/*
 * Service port
 * status : 20101
 * dbs : 20102
 * sys : 20103
 */

type ClientBase struct {
	jkbase.JKNetBase
}

func (c *ClientBase) KeepaliveCycle(interval time.Duration, id string) error {
	for {
		str, _ := jkprotocol.JKProtoV6MakeKeepalive(id)
		n := c.Send(str)
		log4go.Debug("Send keepalive %d", n)
		if n < 0 {
			log4go.Warn("Send keepalived fail, redial")
			c.Dial()
			continue
		}
		recv, err := c.RecvClient()
		if err != nil {
			return err
		}
		log4go.Debug("Recv response of keepalive %s", string(recv))
		time.Sleep(time.Second * interval)
	}

	return nil
}
