package jkbase

import (
	"errors"
	"github.com/alecthomas/log4go"
	"io"
	"net"
	"strconv"
)

var NetTypeBase = 1

type HandlerMsgCallback func(conn net.Conn, data string) error

type JKNetBase struct {
	address string
	port    int
	nettype int

	// As client dial to
	conn     net.Conn
	recvdata string

	// As server listen
	listener net.Listener
	// return recv data to handle
	handler_msg HandlerMsgCallback
	items       map[string](JKNetBaseItem)
}

type JKNetBaseItem struct {
	conn       net.Conn
	remoteaddr string
	data       string
}

func (nb *JKNetBase) GetAddress() (string, int) {
	return nb.address, nb.port
}

func (nb *JKNetBase) GetNettype() int {
	return nb.nettype
}

func (nb *JKNetBase) New(addr string, port int, nettype int) error {
	nb.address = addr
	nb.port = port
	nb.nettype = nettype
	nb.items = make(map[string]JKNetBaseItem)
	return nil
}

func (nb *JKNetBase) Listen() error {
	if nb.conn != nil {
		return errors.New("Has used for dial")
	}
	var err error
	nb.listener, err = net.Listen("tcp", nb.address+":"+strconv.Itoa(nb.port))
	if err != nil {
		return err
	}
	return nil
}

func (nb *JKNetBase) Dial() error {
	if nb.listener != nil {
		return errors.New("Has used for listen")
	}
	conn, err := net.Dial("tcp", nb.address+":"+strconv.Itoa(nb.port))
	if err != nil {
		return err
	}
	nb.conn = conn
	return nil
}

func (nb *JKNetBase) Send(data string) int {
	n, err := nb.conn.Write([]byte(data))
	if err != nil {
		return -1
	}
	return n
}

func (nb *JKNetBase) RecvCycle() error {
	go func() {
		for {
			log4go.Debug("Start to accept ...")
			conn, err := nb.listener.Accept()
			if err != nil {
				log4go.Error("accept failed ", err)
				return
			}
			log4go.Info("accept one client %s", conn.RemoteAddr().String())
			go func() {
				for {
					rdata := make([]byte, 2<<15)
					n, err := conn.Read(rdata)
					if err == io.EOF {
						log4go.Error("IO EOF exit")
						break
					}
					if err != nil {
						log4go.Error("Read error %s", err)
						break
					}

					log4go.Debug("Recv data of length %d, from %s", n, conn.RemoteAddr().String())
					item := JKNetBaseItem{}
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

func (nb *JKNetBase) SetHandlerMsg(callback HandlerMsgCallback) {
	nb.handler_msg = callback
}

func (nb *JKNetBase) RecvClient() ([]byte, error) {
	rdata := make([]byte, 2<<15)
	n, err := nb.conn.Read(rdata)
	if err != nil {
		return nil, err
	}
	nb.recvdata = string(rdata[0:n])
	return rdata[0:n], nil
}
