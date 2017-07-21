package main

import (
	"flag"
	"net"
	"strconv"
	"jk/jkprotocol"
	l4g "github.com/alecthomas/log4go"
	"time"
)

var (
	address = flag.String("address", "0.0.0.0", "address to connect")
	port    = flag.Int("port", 24433, "port to connect")
)

type ClientInfo struct {
    conn      net.Conn
	base      *jkprotocol.JKProtocolWrap
}

func (ci *ClientInfo) Connect(addr string, port int) error {
	conn, err := net.Dial("tcp", addr + ":" + strconv.Itoa(port))
	if err != nil {
		return err
	}
	ci.conn = conn
	return nil
}

func (ci *ClientInfo) Send(data string) (int, error) {
	n, err := ci.conn.Write([]byte(data))
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (ci *ClientInfo) Close() error {
	err := ci.conn.Close()
	return err
}

func (ci *ClientInfo) Keepalive() error {
	go func() {
		for {
			str, err := ci.base.Keepalive("Keepalive")
			if err != nil {
				l4g.Error("Generate keepalive failed ", err)
				return
			}
			_, err = ci.Send(str)
			l4g.Debug("Give keepalive response  [%s]\n", str)
			if err != nil {
				l4g.Error("Send keepalive failed ", err)
				return
			}
			time.Sleep(60000 * time.Millisecond)
		}
	}()

	return nil
}

func main() {
	flag.Parse()

	ci := ClientInfo{}
	ci.Connect(*address, *port)
	defer ci.Close()

	ci.base, _ = jkprotocol.NewJKProtocolWrap(jkprotocol.JK_PROTOCOL_VERSION_5)

    str, err := ci.base.Register("Register")
	if err != nil {
		l4g.Error("Generate Register failed ", err)
		return
	}
	n, err := ci.Send(str)
	if err != nil {
		l4g.Error("Send register failed ", err)
		return
	}
	l4g.Debug("send register success %d\n", n)


	ci.Keepalive()

	for {
		time.Sleep(500*time.Millisecond)
	}

	str, err = ci.base.Leave("Leave")
	if err != nil {
		l4g.Error("Generate leave failed ", err)
		return
	}
	ci.Send(str)
	ci.Close()
}