package jknewclient

import (
	"errors"
	"jk/jklog"
	"net"
	"strconv"
	"time"
)

type KFClient struct {
	addr  string
	port  int
	naddr *net.UDPAddr
	nconn net.Conn
}

func KFClientNew(addr string, port int) (*KFClient, error) {
	if port < 1000 || port > 65535 {
		return nil, errors.New("port is invalid")
	}
	c := &KFClient{
		addr: addr,
		port: port,
	}
	naddr, err := net.ResolveUDPAddr("udp", addr+":"+strconv.Itoa(port))
	if err != nil {
		return nil, errors.New("resolve udp address failed")
	}
	c.naddr = naddr
	return c, nil
}

func (s *KFClient) Send(data []byte, timeout int) error {
	jklog.Lfile().Infoln("start to dial ", s.naddr.String())

	send := false
	if s.nconn != nil {
		n, err := s.nconn.Write(data)
		if err == nil {
			jklog.Lfile().Debugln("write down len: ", n)
			send = true
		} else {
			jklog.Lfile().Errorln("write failed, reconnect try one")
			s.nconn.Close()
		}
	}

	if !send {
		waitresp, err := net.DialTimeout("udp", s.naddr.String(), time.Millisecond*time.Duration(timeout))
		if err != nil {
			return err
		}
		s.nconn = waitresp
		n, err := s.nconn.Write(data)
		if err != nil {
			return err
		}
		jklog.Lfile().Debugln("write down len: ", n)
	}

	return nil
}

func (c *KFClient) Recv() ([]byte, error) {
	data := make([]byte, 1<<12)
	n, err := c.nconn.Read(data)
	if err != nil {
		return nil, err
	}
	retData := data[:n]
	return retData, nil
}

func (c *KFClient) Close() {
	c.nconn.Close()
}
