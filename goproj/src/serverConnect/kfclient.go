package main

import (
	// "fmt"
	. "jk/jkcommon"
	"jk/jklog"
	cli "jk/jknewclient"
	"time"
	// "encoding/binary"
	"flag"
	. "jk/jkprotocol"
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

	// tmstr := fmt.Sprintf("%d", time.Now().Unix())
	p := NewJKProtocol()

	regstr := p.GenerateRegister(*id)

	n, err := c.handle.Write([]byte(regstr))
	if err != nil {
		jklog.L().Errorln("write failed: ", err)
		return false
	}
	jklog.L().Debugln("write done of len: ", n)

	itemData := &cli.JKClientItem{}
	n, err = c.handle.Read(itemData)
	if err != nil {
		jklog.L().Errorln("read failed, ", err)
		return false
	}
	jklog.L().Infoln("read data: ", string(itemData.Data))

	wstr := p.GenerateControlSaveFile("test.log", "This is a test log\n")
	_, err = c.handle.Write([]byte(wstr))
	if err != nil {
		jklog.L().Errorln("write failed, ", err)
	}

	time.Sleep(time.Millisecond * 10000)
	c.handle.Close()
	return true
}

var (
	id = flag.String("id", "kfun", "the unique id")
	// savepos = flag.String("savepos", "docs", "where to save html files")
)

func main() {
	flag.Parse()

	c := &KFClient{}
	c.startClient()
}
