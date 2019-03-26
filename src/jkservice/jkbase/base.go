package jkbase

import (
	"jk/jklog"
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
		jklog.L().Debugf("Send keepalive %d\n", n)
		if n < 0 {
			jklog.L().Warnln("Send keepalived fail, redial")
			c.Dial()
			continue
		}
		recv, err := c.RecvClient()
		if err != nil {
			return err
		}
		jklog.L().Debugf("Recv response of keepalive %s\n", string(recv))
		time.Sleep(time.Second * interval)
	}

	return nil
}
