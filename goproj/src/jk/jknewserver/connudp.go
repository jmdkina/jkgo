package jknewserver

import (
	"errors"
	"jk/jklog"
	"net"
	"strconv"
	"strings"
)

type KFServerUDPItem struct {
	RemoteAddr string
	Data       []byte
	c          *net.UDPAddr
	SendData   chan bool
	Error      chan error
}

type KFServerUDP struct {
	port int
	c    *net.UDPConn
	Item []*KFServerUDPItem
}

func NewKFServerUDP(port int) (*KFServerUDP, error) {
	if port < 1000 || port > 1<<16-1 {
		return nil, errors.New("Invalid port!")
	}

	s := &KFServerUDP{
		port: port,
	}

	service, err := net.ResolveUDPAddr("udp4", ":"+strconv.Itoa(port))
	if err != nil {
		return nil, err
	}

	waitresp, err := net.ListenUDP("udp", service)
	if err != nil {
		return nil, err
	}
	s.c = waitresp

	jklog.L().Infoln("listen port : ", port)

	return s, nil
}

func (s *KFServerUDP) Close() {
	s.c.Close()
}

func (s *KFServerUDP) FindItem(item *KFServerUDPItem) (*KFServerUDPItem, bool) {
	for _, v := range s.Item {
		if strings.Compare(v.RemoteAddr, item.RemoteAddr) == 0 {
			return v, true
		}
	}
	return item, false
}

func (s *KFServerUDP) Recv() (*KFServerUDPItem, []byte, error) {
	jklog.Lfile().Debugln("start to recv data")
	saveData := make([]byte, 1<<10)
	n, c, err := s.c.ReadFromUDP(saveData)
	if err != nil {
		return nil, nil, err
	}
	jklog.Lfile().Debugln("deal recv data.")
	item := &KFServerUDPItem{}
	item.RemoteAddr = c.String()
	item.Data = make([]byte, n)
	item.Data = saveData[:n]
	item.c = c
	item.SendData = make(chan bool)
	item.Error = make(chan error)

	item, ret := s.FindItem(item)
	if !ret {
		jklog.L().Infoln("New remote address is : ", item.RemoteAddr)
		s.Item = append(s.Item, item)
		go func() {
			<-item.SendData
			jklog.L().Debugln("will send data to ", item.c.String(), ", len: ", len(item.Data))
			_, err := s.c.WriteToUDP(item.Data, item.c)
			if err != nil {
				item.Error <- errors.New("Error of send data")
			} else {
				item.Error <- nil
			}
		}()
		return item, item.Data, nil
	}

	return item, item.Data, nil
}
