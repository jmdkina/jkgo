package jknewserver

import (
	"fmt"
	"io"
	. "jk/jkcommon"
	"jk/jklog"
	"net"
	"strings"
)

func (ser *JKServerProcess) addItem(item *JKServerProcessItem) {
	for k, v := range ser.Item {

		if strings.Compare(v.remoteAddr, item.remoteAddr) == 0 {
			ser.Item[k] = item
			return
		}
	}
	ser.Item = append(ser.Item, item)
}

func (ser *JKServerProcess) listenLocalTCP(port int) bool {
	str := JK_NET_ADDRESS_LOCAL + ":" + fmt.Sprintf("%d", port)
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

	go func() {
		for {
			acp, err := lis.Accept()
			if err != nil {
				jklog.L().Errorln("accept error: ", err)
				continue
			}
			servItem := &JKServerProcessItem{}

			remAddr := acp.RemoteAddr().String()
			servItem.remoteAddr = remAddr

			ser.addItem(servItem)

			var readbuf []byte
			var lenbuf int
			go func() {
				buf := make([]byte, 2<<12)
				n, err := acp.Read(buf)
				if err == io.EOF {
					jklog.L().Infoln("EOF of read.")
					return
				}
				if err != nil {
					jklog.L().Errorln("read data failed: ", err)
					return
				}

				copy(readbuf[lenbuf:lenbuf+n], buf[0:n])
				lenbuf += n

				servItem.data = readbuf

				jklog.L().Infoln("data from ", acp.RemoteAddr().String(), " with len ", n)
			}()
		}
	}()

	return true
}
