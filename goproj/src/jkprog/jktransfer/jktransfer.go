package main

import (
	"flag"
	"jk/jklog"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	style    = flag.String("type", "client", "client/server")
	dst      = flag.String("addr", "", "addr[client use]")
	filename = flag.String("filename", "", "filename[client use]")
	save     = flag.String("save", "", "save postion[serveruse]")
)

const (
	JK_TYPE_CLIENT = 0
	JK_TYPE_SERVER = 1
)

const (
	JK_CONNECT_SEND   = 9955
	JK_CONNECT_LISTEN = 9955
)

const (
	onceLen = 2 << 20
)

var execType int

func main() {
	flag.Parse()
	switch *style {
	case "client":
		execType = JK_TYPE_CLIENT
		break
	case "server":
		execType = JK_TYPE_SERVER
		break
	default:
		jklog.L().Errorln("I'm not support ", *style)
		return
	}
	if execType == JK_TYPE_CLIENT {
		if *dst == "" {
			jklog.L().Errorln("You use client, so give me addr to connect ")
			return
		}
	}

	// connect out
	switch execType {
	case JK_TYPE_CLIENT:
		c, err := connectOut(*dst, JK_CONNECT_SEND)
		if err != nil {
			jklog.L().Infoln("connect err ==>", err)
			break
		}
		sendDataOut(c, *filename)
		jklog.L().Infoln("wait a minute")
		time.Sleep(6000 * time.Millisecond)
		jklog.L().Infoln("send data down.")
		break
	case JK_TYPE_SERVER:
		listenData(JK_CONNECT_LISTEN, *save)
		jklog.L().Infoln("Should not here, until down.")
		break
	default:
		jklog.L().Errorln("Unkown type ", execType)
		break
	}
}

func connectOut(addr string, port int) (net.Conn, error) {
	connstr := addr + ":" + strconv.Itoa(port)
	conn, err := net.Dial("tcp", connstr)
	return conn, err
}

func sendDataOut(c net.Conn, name string) {
	if name == "" {
		jklog.L().Errorln("give me some name [fullpath]")
		return
	}
	f, err := os.OpenFile(name, os.O_RDONLY, os.ModePerm)
	if err != nil {
		jklog.L().Errorln("err open file ", name)
		return
	}

	sendFilename := name
	i := strings.LastIndex(name, "/")
	if i >= 0 {
		sendFilename = name[i+1 : len(name)]
	}
	c.Write([]byte("filename\n" + sendFilename))
	jklog.L().Infoln("send out ==> filename:", sendFilename)

	max := 0

	go func() {
		lastvalue := 0
		tCnts := 1
		for {
			if max == -1 {
				jklog.L().Infoln("The av send speed ", lastvalue/tCnts/1000000, "Mb/s")
				break
			}

			jklog.L().Infoln("Send Data Out", max, "(", (max-lastvalue)/5/1000000, "Mb/s)")
			lastvalue = max
			time.Sleep(5000 * time.Millisecond)
			tCnts += 5
		}
	}()

	buf := make([]byte, onceLen)
	for {
		cnts, err := f.Read(buf)
		if err != nil {
			jklog.L().Errorln("read err ==>", err)
			break
		}
		// jklog.L().Infoln("read out len ", cnts)
		wlen, err := c.Write(buf[0:cnts])
		if err != nil {
			jklog.L().Errorln("write err ==>", err)
			break
		}
		max += wlen
		// jklog.L().Infoln("write down ", wlen)
	}
	f.Close()
	jklog.L().Infoln("transfered len ", max, "of file", name)
	max = -1
}

func listenData(port int, save string) {
	if save == "" {
		jklog.L().Errorln("Give me a place to write ")
		return
	}

	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		jklog.L().Errorln("listen err", err)
		return
	}
	jklog.L().Infoln("start to listen ", ":"+strconv.Itoa(port))
	for {
		conn, err := ln.Accept()
		if err != nil {
			jklog.L().Errorln("accept err ==> ", err)
			break
		}
		go receiverData(conn, save)
	}
}

func receiverData(c net.Conn, save string) {

	var f *os.File
	max := 0

	go func() {
		lastvalue := 0
		iCnts := 1
		for {
			if max == -1 {
				jklog.L().Infoln("The av send speed ", lastvalue/iCnts/1000, "Kb/s")
				break
			}
			jklog.L().Infoln("Recived data", max, "(", (max-lastvalue)/5/1000, "Kb/s)")
			lastvalue = max
			time.Sleep(5000 * time.Millisecond)
			iCnts += 5
		}
	}()

	buf := make([]byte, onceLen)
	for {
		cnts, err := c.Read(buf)
		if err != nil {
			jklog.L().Errorln("read err ==> ", err)
			break
		}
		// If it is filename:xxx means the filename
		str := string(buf)
		if strings.HasPrefix(str, "filename") {
			lastname := strings.Split(str[0:cnts], "\n")[1]
			i := strings.LastIndex(lastname, "/")

			if i >= 0 {
				lastname = lastname[i+1 : len(lastname)]
			} else {
				// maybe it is windows
				i = strings.Index(lastname, "\\")
				if i >= 0 {
					lastname = lastname[i+1 : len(lastname)]
				}
			}
			saveplace := save + "/" + lastname

			// open the read filename
			jklog.L().Infoln("open ", saveplace, " now")
			// saveplace = "./netuse.go"

			f, err = os.Create(saveplace)
			if err != nil {
				jklog.L().Infoln("open err =>", err)
				break
			} else {
				continue
			}
		}

		cnts, err = f.Write(buf[0:cnts])
		if err != nil {
			jklog.L().Errorln("write err ==>", cnts)
			break
		}
		max += cnts
		// jklog.L().Infoln("write len down ", cnts)
	}
	if f != nil {
		f.Close()
	}
	jklog.L().Infoln("received len ", max)
	max = -1
}
