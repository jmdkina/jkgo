package main

import (
	"flag"
	"net"
	"strconv"
	"jk/jkprotocol"
	"jk/jklog"
	"time"
)

var (
	address = flag.String("local_address", "0.0.0.0", "Listen local address")
	port    = flag.Int("local_port", 24433, "Listen local port")
)

type ClientInfo struct {
    conn      net.Conn
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
			keep, err := jkprotocol.NewV5Keepalive("Keepalive")
			if err != nil {
				jklog.L().Errorln("Generate keepalive failed ", err)
				return
			}
			str, _ := keep.String()
			_, err = ci.Send(str)
			jklog.L().Debugf("Give keepalive response  [%s]\n", str)
			if err != nil {
				jklog.L().Errorln("Send keepalive failed ", err)
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

    reg, err := jkprotocol.NewV5Register("Register")
	if err != nil {
		jklog.L().Errorln("Generate Register failed ", err)
		return
	}
	str, _ := reg.String()
	n, err := ci.Send(str)
	if err != nil {
		jklog.L().Errorln("Send register failed ", err)
		return
	}
	jklog.L().Debugf("send register success %d\n", n)


	ci.Keepalive()

	for {
		time.Sleep(500*time.Millisecond)
	}

	leave, err := jkprotocol.NewV5Leave("Leave")
	if err != nil {
		jklog.L().Errorln("Generate leave failed ", err)
		return
	}
	str, _ =  leave.String()
	ci.Send(str)
	ci.Close()
}
