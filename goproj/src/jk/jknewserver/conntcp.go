package jknewserver

import (
	"fmt"
	// "io"
	. "jk/jkcommon"
	"jk/jklog"
	"net"
	"strings"
)

func (ser *JKServerProcess) addItem(item *JKServerProcessItem) {
	for k, v := range ser.Item {

		if strings.Compare(v.RemoteAddr, item.RemoteAddr) == 0 {
			ser.Item[k] = item
			return
		}
	}
	ser.Item = append(ser.Item, item)
}

func (ser *JKServerProcess) listenLocalTCP(port int) bool {
	str := JK_NET_ADDRESS_LOCAL + ":" + fmt.Sprintf("%d", port)
	jklog.L().Debugln("start to listen : ", str)
	nt, err := net.ResolveTCPAddr("tcp", str)
	if err != nil {
		jklog.L().Errorln("error resolve: ", err)
		return false
	}

	lis, err := net.ListenTCP("tcp", nt)
	if err != nil {
		jklog.L().Errorln("error listen: ", err)
		return false
	}
	ser.Listen = lis

	return true
}
