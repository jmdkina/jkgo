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

func (c *KFClient) startClient(addr string, port int) bool {

	c.handle = cli.JKNewClientNew(addr, port)

reconn:
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
		goto reconn
	}
	jklog.L().Debugln("write done of len: ", n)

	itemData := &cli.JKClientItem{}
	n, err = c.handle.Read(itemData)
	if err != nil {
		jklog.L().Errorln("read failed, ", err)

	} else {
		jklog.L().Infoln("read data: ", string(itemData.Data))
	}

	var startTime int64
	for {
		now := time.Now().Unix()

		if now-startTime < 10 {
			time.Sleep(time.Millisecond * 1000)
			continue
		}
		startTime = now

		wstr := p.GenerateControlSaveFile(*id+".log", "Hello, I'm online - "+time.Now().String())
		_, err = c.handle.Write([]byte(wstr))
		if err != nil {
			jklog.L().Errorln("write failed, ", err)
			goto reconn
		}
	}

	c.handle.Close()
	return true
}

var (
	id   = flag.String("id", "kfun", "the unique id")
	addr = flag.String("addr", "0.0.0.0", "remote addr")
	port = flag.Int("port", JK_NET_ADDRESS_PORT, "remote port")
	// savepos = flag.String("savepos", "docs", "where to save html files")
)

func main() {
	flag.Parse()

	c := &KFClient{}
	c.startClient(*addr, *port)
}
