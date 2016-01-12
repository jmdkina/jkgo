package main

import (
	// "fmt"
	. "jk/jkcommon"
	"jk/jklog"
	cli "jk/jknewclient"
	"time"
	// "encoding/binary"
	"flag"
	daemon "github.com/tyranron/daemonigo"
	. "jk/jkprotocol"
)

type KFClient struct {
	handle *cli.JKNewClient
	item   *cli.JKClientItem
}

func (c *KFClient) startClient(addr string, port int) bool {

	jklog.Lfile().Debugln("Start to connect ", addr, ": ", port)
	c.handle = cli.JKNewClientNew(addr, port)

reconn:
	jklog.Lfile().Infoln("start to connect ...")
	ret := c.handle.CliConnect(0, true)
	if !ret {
		jklog.Lfile().Errorln("connect failed.")
		return false
	}

	// tmstr := fmt.Sprintf("%d", time.Now().Unix())
	p := NewJKProtocol()

	regstr := p.GenerateRegister(*id)

	n, err := c.handle.Write([]byte(regstr))
	if err != nil {
		jklog.Lfile().Errorln("write failed: ", err)
		goto reconn
	}
	jklog.Lfile().Debugln("write done of len: ", n)

	itemData := &cli.JKClientItem{}
	n, err = c.handle.Read(itemData)
	if err != nil {
		jklog.Lfile().Errorln("read failed, ", err)

	} else {
		jklog.Lfile().Infoln("read data: ", string(itemData.Data))
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
			jklog.Lfile().Errorln("write failed, ", err)
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
	signal     = flag.String("action", "start", "{quit|stop|reload}")
	background = flag.Bool("d", false, "Background run")
)

func main() {
	flag.Parse()

	jklog.InitLog("/tmp/jkclient.log")

	jklog.Lfile().Infoln("Program start ...")
	// Daemonizing echo server application.
	if *background {
		jklog.Lfile().Infoln("background run now.")
		switch isDaemon, err := daemon.Daemonize(*signal); {
		case !isDaemon:
			return
		case err != nil:
			jklog.L().Errorln("daemon start failed : ", err.Error())
		}
	}
	c := &KFClient{}
	c.startClient(*addr, *port)
}