package main

import (
	"io"
	"jk/jklog"
	. "jk/jkserver"
	"net"
	"os"
	"strconv"
	"time"
)

const (
	jk_listen_port     = 10041
	jk_listen_port_tcp = 10044
	each_length        = 2 << 15
)

func main() {

	jklog.L().Noticeln("listen udp addr " + strconv.Itoa(jk_listen_port))
	service, err := net.ResolveUDPAddr("udp4", ":"+strconv.Itoa(jk_listen_port))
	if err != nil {
		jklog.L().Errorln("resolve udp fail: ", err)
		return
	}
	waitresp, err := net.ListenUDP("udp", service)
	defer waitresp.Close()

	addrto, err := net.ResolveUDPAddr("udp4", "255.255.255.255:"+strconv.Itoa(10044))
	if err != nil {
		jklog.L().Infoln("resolve udp error: ", err)
	}
	sendto, err := net.DialUDP("udp", nil, addrto)
	if err != nil {
		jklog.L().Infoln("dial udp error: ", err)
	}
	header := NewJKCommandHeader()
	header.GenJKCommandHeaderOneSend(JK_BVWORK, JK_BVETHERNET, JK_DEVICELIST)
	n, _ := sendto.Write([]byte(header.ToString()))
	jklog.L().Infoln("-----sned len ", n)
	// return

	receiveredDataUseTcp()
	for {
		buf := make([]byte, each_length)
		jklog.L().Infoln("read udp data .....")
		n, _, err := waitresp.ReadFromUDP(buf[0:])
		if err != nil {
			jklog.L().Errorln("accept err :", err)
			break
		}
		if n == 0 {
			time.Sleep(500 * time.Microsecond)
		} else {
			jklog.L().Infoln("out recv: ", n, " of ", string(buf[0:n]))
			writeDataToFile(buf[0:n])
			tt, err := sendto.Write([]byte("1------Yes I go ayour message"))

			jklog.L().Infoln("kkkkkkkkkk sened to ", sendto.RemoteAddr().String(), " : ", tt)
			// _, err := waitresp.WriteToUDP([]byte("2====Yes I got your message"), addr)
			if err != nil {
				jklog.L().Errorln("write response fail: ", err)
				continue
			}
			// jklog.L().Infoln("write response len ", n)
		}
	}
	jklog.L().Infoln("should not here")
}

func receiveredDataUseTcp() {
	jklog.L().Noticeln("listen tcp addr " + strconv.Itoa(jk_listen_port_tcp))
	service, err := net.Listen("tcp", ":"+strconv.Itoa(jk_listen_port_tcp))
	if err != nil {
		jklog.L().Errorln("listen tcp error:", err)
		return
	}
	go func() {

		buf := make([]byte, 2<<15)
		for {
			jklog.L().Infoln("STart to wait data come ")
			conn, err := service.Accept()
			if err != nil {
				jklog.L().Errorln("error : ", err)
				break
			}
			for {
				cnts, err := conn.Read(buf)
				if io.EOF == err {
					break
				}
				if err != nil {
					jklog.L().Errorln("read error", err)
					break
				}
				// jklog.L().Infoln("read out len ", cnts)
				jklog.L().Infof("read out %d: %s\n", cnts, string(buf[0:cnts]))
				writeDataToFile(buf[0:cnts])

				jklog.L().Infoln("write back to ", conn.RemoteAddr())
				conn.Write([]byte("I received data"))
			}
		}
	}()
}

func writeDataToFile(buf []byte) {
	filename := "c:/Users/v-jmd/Desktop/mediafrom-ios"
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModeDir)
	if err != nil {
		jklog.L().Errorln("open failed ", filename)
		return
	}
	cnts, err := f.WriteString(string(buf))
	if err != nil {
		jklog.L().Errorln("write to file error ", err)
		return
	}
	jklog.L().Infoln("Write ", cnts)
	f.Close()
}

func interface_use() {
	inetf, err := net.Interfaces()
	if err != nil {
		jklog.L().Errorln("interface exec error")
		return
	}
	for _, v := range inetf {
		jklog.L().Infof("name[%s]", v.Name)
		if v.Name == "wlan0" {
			addrs, err := v.Addrs()
			if err != nil {
				jklog.L().Errorf("get addr of [%s] ERRROR desc: [%s]", v.Name, err)
				return
			}
			for _, n := range addrs {
				jklog.L().Infof("net [%s] addr: [%s]\n", v.Name, n.String())
			}
		}
	}
}
