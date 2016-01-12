package jknewserver

import (
	"io"
	. "jk/jkcommon"
	"jk/jklog"
	"net"
//	"time"
)

type JKNewServer struct {
	Port int
}

func JKNewServerNew() *JKNewServer {
	return &JKNewServer{}
}

type JKServerProcessItem struct {
	Id         string
	RemoteAddr string
	Conn       net.Conn
	Data       []byte
	ReadDone   chan bool
	TimeLast   int64
}

type JKServerProcess struct {
	Listen *net.TCPListener
	Item   []*JKServerProcessItem
}

func (newser *JKNewServer) Start(proc *JKServerProcess) bool {
	return proc.listenLocalTCP(JK_NET_ADDRESS_PORT)
}

func (newser *JKNewServer) Accept(proc *JKServerProcess) net.Conn {
	acp, err := proc.Listen.Accept()
	if err != nil {
		jklog.Lfile().Errorln("accept failed. ", err)
		return nil
	}
	jklog.Lfile().Debugln("got accept from: ", acp.RemoteAddr().String())
	return acp
}

func (newser *JKNewServer) Read(proc *JKServerProcess, procItem *JKServerProcessItem) (int, error) {
	servItem := procItem

	remAddr := procItem.Conn.RemoteAddr().String()
	servItem.RemoteAddr = remAddr

	proc.addItem(servItem)

	jklog.Lfile().Debugln("first read first bytes.")
	// first read 4 bytes for length.
	buflen := make([]byte, 4)
//	procItem.Conn.SetReadDeadline(time.Unix(10000, 10000));
	_, err := procItem.Conn.Read(buflen)
	if err != nil {
		jklog.Lfile().Errorln("read failed of first read. ", err)
		return 0, err
	}
	datalen := int(BytesToInt32(buflen))
	jklog.L().Debugf("%x,%x,%x,%x\n", buflen[0], buflen[1], buflen[2], buflen[3])
	if datalen > (2000000) {
		jklog.Lfile().Errorln("Data is too long, I only receive some of them.")
		datalen = 2000000
	}
	jklog.Lfile().Debugln("Will read the data of length: ", datalen)

	readbuf := make([]byte, datalen)
	lenbuf := 0

	jklog.Lfile().Debugln("Start read data cycle from : ", remAddr)
	for {
		if lenbuf >= datalen {
			procItem.ReadDone <- true
			break
		}
		buf := make([]byte, 1024)
		n, err := procItem.Conn.Read(buf)
		if err == io.EOF {
			jklog.Lfile().Infoln("EOF of read.")
			break
		}
		// jklog.L().Debugln("The length read len :", n)
		if n > datalen {
			// More data, just cut it.
			n = datalen
		}
		if err != nil {
			jklog.Lfile().Errorln("read data failed: ", err)
			return 0, err
		}

		copy(readbuf[lenbuf:lenbuf+n], buf[0:n])
		lenbuf += n

		servItem.Data = readbuf
	}
	// procItem.ReadDone <- true
	jklog.L().Infoln("Read data done from ", procItem.RemoteAddr, " with len ", lenbuf)
	return lenbuf, nil
}

func (newser *JKNewServer) Write(procItem *JKServerProcessItem, data []byte) (int, error) {
	return procItem.Conn.Write(data)
}

func (newser *JKNewServer) Close(proc *JKServerProcess) bool {
	proc.Listen.Close()

	// close all.
	return true
}
