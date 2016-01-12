package main

import (
	"jk/jklog"
	ser "jk/jknewserver"
	pt "jk/jkprotocol"
	"strings"
	"time"
)

type KFServerItem struct {
	Id         string
	Online     bool
	lastOnline int64 // last online time, for max offline
	offlineMax int   // how long offline make offline.
}

type KFServer struct {
	Item         []*KFServerItem
	ServerHandle *ser.KFServerUDP
}

func (s *KFServer) addNewItem(id string) *KFServerItem {
	for _, v := range s.Item {
		if strings.Compare(id, v.Id) == 0 {
			v.lastOnline = time.Now().Unix()
			return v
		}
	}
	ditem := &KFServerItem{}
	ditem.Id = id
	ditem.Online = true
	ditem.lastOnline = time.Now().Unix()
	s.Item = append(s.Item, ditem)
	return ditem
}

func (s *KFServer) dealResponse(pro *pt.KFProtocol, ditem *KFServerItem, item *ser.KFServerUDPItem) error {

	// If need response.
	pro.SetResponseCode(true, 0)
	pro.SetData([]byte("OK"))
	d, _ := pro.GenerateDataText(false)
	item.Data = d
	item.SendData <- true
	jklog.Lfile().Infoln("Start to send data of response previous command")
	errret := <-item.Error
	if errret != nil {
		jklog.Lfile().Errorln("Error send data: ", errret)
	} else {
		jklog.Lfile().Debugln("Success of the sended data")
	}
	return nil
}

func (s *KFServer) Recv(port int) error {
	h, err := ser.NewKFServerUDP(port)
	if err != nil {
		return err
	}
	s.ServerHandle = h

	for {
		jklog.Lfile().Debugln("Start to recv data ")
		// This recv will make a new go function to wait for send command out
		item, data, err := s.ServerHandle.Recv()
		if err != nil {
			return err
		}

		go func() {

			jklog.Lfile().Debugln("Start to parse data")
			pro, err := pt.KFProtocolParseText(data)
			if err != nil {
				jklog.Lfile().Errorln("Parse data failed, err : ", err)
				item.Data = []byte("Error parse")
				item.SendData <- true
				return
			}

			ditem := s.addNewItem(string(pro.Header.Id[:]))
			s.dealResponse(pro, ditem, item)
		}()
	}
}

func main() {
	s := &KFServer{}

	s.Recv(9988)
}
