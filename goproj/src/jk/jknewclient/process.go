package jknewclient

import (
	// "fmt"
	"jk/jklog"
	"net"
	"strconv"
	"time"
)

type JKClientItem struct {
	Data []byte
}

type JKNewClient struct {
	Addr string
	Port int
	Conn *net.TCPConn
}

func JKNewClientNew(addr string, port int) *JKNewClient {
	cli := &JKNewClient{
		Addr: addr,
		Port: port,
	}
	return cli
}

func (cli *JKNewClient) CliConnect(connTimes int, cycle bool) bool {
	if connTimes < 0 {
		connTimes = 0
	}

	connStr := cli.Addr + ":" + strconv.Itoa(cli.Port)
	jklog.L().Debugln("goto connect to : ", connStr)
	resv, err := net.ResolveTCPAddr("tcp", connStr)
	if err != nil {
		jklog.L().Errorln("failed resolve: ", err)
		return false
	}

	conncnts := 0
reconn:

	if connTimes > 0 && conncnts > connTimes {
		return false
	}

	conn, err := net.DialTCP("tcp", nil, resv)
	if err != nil {
		jklog.L().Errorln("connect failed: ", err)
		if cycle {
			time.Sleep(time.Millisecond * 1000)
			conncnts = conncnts + 1
			goto reconn
		} else {
			return false
		}
	}
	jklog.L().Debugln("connect success with ", cli.Addr, ", port: ", cli.Port)

	cli.Conn = conn

	return true
}

func (cli *JKNewClient) Write(data []byte) (int, error) {
	return cli.Conn.Write(data)
}

func (cli *JKNewClient) Read(item *JKClientItem) (int, error) {
	item.Data = make([]byte, 2<<12)
	// TODO: change to read cycle.
	return cli.Conn.Read(item.Data)
}

func (cli *JKNewClient) Close() bool {
	cli.Conn.Close()
	return true
}
