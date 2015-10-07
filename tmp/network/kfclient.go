package main

import (
	"fmt"
	. "jk/jkcommon"
	"jk/jklog"
	cli "jk/jknewclient"
	"time"
	// "encoding/binary"
)

type KFClient struct {
	handle *cli.JKNewClient
	item   *cli.JKClientItem
}

func (c *KFClient) startClient() bool {
	c.handle = cli.JKNewClientNew(JK_NET_ADDRESS_LOCAL, JK_NET_ADDRESS_PORT)

	ret := c.handle.CliConnect(0, true)
	if !ret {
		return false
	}

	tmstr := fmt.Sprintf("%d", time.Now().Unix())

	n, err := c.handle.Write([]byte("KFClientTest-" + tmstr))
	if err != nil {
		jklog.L().Errorln("write failed: ", err)
		return false
	}
	jklog.L().Debugln("write done of len: ", n)

	time.Sleep(time.Millisecond * 10000)
	c.handle.Write([]byte("Second data..."))
	time.Sleep(time.Millisecond * 10000)
	c.handle.Close()
	return true
}

func main() {
	c := &KFClient{}
	c.startClient()
}
