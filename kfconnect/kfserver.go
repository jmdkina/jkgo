package main

import (
	"jk/jklog"
	ser "jk/jknewserver"
	pt "jk/jkprotocol"
)

type KFServerItem struct {
	Id string
}

type KFServer struct {
	Item         []*KFServerItem
	ServerHandle *ser.KFServerUDP
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

		jklog.Lfile().Debugln("Start to parse data")
		pro, err := pt.KFProtocolParse(data)
		if err != nil {
			jklog.Lfile().Errorln("Parse data failed, err : ", err)
			item.Data = []byte("Error parse")
			item.SendData <- true
			continue
		}
		pro.SetResponseCode(true, 0)
		pro.SetData([]byte("I recevied the data"))
		d, err := pro.GenerateData(false)
		item.Data = d
		item.SendData <- true
		jklog.Lfile().Infoln("Start to send data of response previous command")
		errret := <-item.Error
		if errret != nil {
			jklog.Lfile().Errorln("Error send data: ", errret)
		} else {
			jklog.Lfile().Debugln("Success of the sended data")
		}
	}
}

func main() {
	s := &KFServer{}

	s.Recv(9988)
}
