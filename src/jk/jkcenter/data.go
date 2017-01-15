package jkcenter

import (
	"io"
	"net"
	"strconv"
	"strings"
)

const (
	DB     = 1
	Client = 2
)

// Format of data transfer
// { "header": { "ver": 1, "cmd": "", "subcmd":"", "crypt":1, "id":"", "transaction", "xxxxx", "resp":false },
//   "body": { "data":"xxx" } }
// DBS --> register  Center
// Client --> register Center
// Device --> register Client

type header struct {
	version int
	cmd     string
	subcmd  string
	crypt   int
	id      string
}

type Cmd struct {
	header
	body interface{}
}

type ClientInfo struct {
	conn        net.Conn
	transaction string
	id          string
}

type CenterControl struct {
	lis  net.Listener
	pair map[string]*ClientInfo // transaction <--> ClientInfo
}

// nettype, 0 tcp, 1 udp
func InitCenter(laddr string, lport int, nettype int) (*CenterControl, error) {
	cc := &CenterControl{}
	addr, err := net.ResolveTCPAddr("tcp", laddr+strconv.Itoa(lport))
	if err != nil {
		return nil, err
	}
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	cc.lis = lis

	return cc, nil
}

func (cc *CenterControl) recv() error {
	conn, err := cc.lis.Accept()
	if err != nil {
		return err
	}
	go func() {
		// Recv data
		var data []byte
		for {
			rdata := make([]byte, 2<<15)
			n, err := conn.Read(rdata)
			if err == io.EOF {
				break
			}
			data = append(data, rdata)
		}
		// Parse

		// transfer to other depends on cmd

		// if no resp
		ci := &ClientInfo{
			conn:        conn,
			transaction: "1234",
			id:          "88888",
		}
		cc.pair["1234"] = ci

		// if resp
		// response to transaction
		v, ex := cc.pair["1234"]
		if ex {
			v.conn.Write("response data")
		}

	}()
}
