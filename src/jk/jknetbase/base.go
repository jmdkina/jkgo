package jknetbase

import (
	"net"
	"errors"
	"strconv"
	"io"
	"github.com/alecthomas/log4go"
)

type HandlerMsgCallback func(conn net.Conn, data string) error

type JKNetBaseRecv struct {
	listener   net.Listener

	address    string
	port       int
	nettype    int

	handler_msg  HandlerMsgCallback

	items      map[string](JKNetBaseRecvItem)
}

type JKNetBaseRecvItem struct {
	conn        net.Conn
	remoteaddr  string
	data        string
}

func (nb *JKNetBaseRecv) New(addr string, port int, nettype int) error {
	nb.address = addr
	nb.port = port
	nb.nettype = nettype
	nb.items = make(map[string](JKNetBaseRecvItem))

	if nettype == 1 {
		var err error
		nb.listener, err = net.Listen("tcp", addr + ":" + strconv.Itoa(port))
		if err != nil {
			return err
		}
	} else {
		return errors.New("Unsupported nettype")
	}
	return nil
}

func (nb *JKNetBaseRecv) Recv() error {
	go func() {
		for {
			conn, err := nb.listener.Accept()
			if err != nil {
				log4go.Error("accept failed ", err)
				return
			}
			log4go.Info("accept one client %s", conn.RemoteAddr().String())
			go func() {
				for {
					rdata := make([]byte, 2 << 15)
					n, err := conn.Read(rdata)
					if err == io.EOF {
						log4go.Error("IO EOF exit")
						break
					}

					log4go.Debug("Recv data of length %d, from %s", n, conn.RemoteAddr().String())
					item := JKNetBaseRecvItem{}
					item.conn = conn
					item.remoteaddr = conn.RemoteAddr().String()
					item.data = string(rdata[0:n])
					nb.items[item.remoteaddr] = item
					nb.handler_msg(conn, item.data)
				}
			}()
		}
	}()
	return nil
}

func (nb *JKNetBaseRecv) SetHandlerMsg(callback HandlerMsgCallback) {
	nb.handler_msg = callback
}

type JKNetBaseClient struct {
	address    string
	port       int
	nettype    int

	conn net.Conn
	recvdata string
}

func (nb *JKNetBaseClient) New(addr string, port int, nettype int) error {
	if nettype == 1 {
		nb.address = addr
		nb.port = port
		nb.nettype = nettype
		conn, err := net.Dial("tcp", addr + ":" + strconv.Itoa(port))
		if err != nil {
			return err
		}
		nb.conn = conn
	} else {
		return errors.New("Unsupported nettype")
	}
	return nil
}

func (nb *JKNetBaseClient) Recv() ([]byte, error) {
	rdata := make([]byte, 2<<15)
	n, err := nb.conn.Read(rdata)
	if err != nil {
		return nil, err
	}
	nb.recvdata = string(rdata[0:n])
	return rdata[0:n], nil
}

func (nb *JKNetBaseClient) Send(data string) int {
	n, err := nb.conn.Write([]byte(data))
	if err != nil {
		return -1
	}
	return n
}
