package main

import (
	"jk/jklog"
	"net"
	"strconv"
	"time"
    "flag"
)


// This file is use for test for udp
// one is listen udp, each receive from udp send one data out for test
// the one is send data with udp

func connect_out(addr string, port int) {
	jklog.L().Infoln("connect out with udp of ", addr, ":", port)
	addrto, err := net.ResolveUDPAddr("udp4", addr+":"+strconv.Itoa(port))
	if err != nil {
		jklog.L().Errorln("resolve udp addr failed")
		return
	}

	sendto, err := net.DialUDP("udp", nil, addrto)
	if err != nil {
		jklog.L().Errorln("dial udp failed ", err)
		return
	}
	defer sendto.Close()

	// outdata := "Here is for test"
	n, _ := sendto.Write([]byte(*content))
	jklog.L().Infoln("send out data of len ", n)
    if *recv {
        data := make([]byte, 1024)
        sendto.Read(data)
        jklog.L().Infof("Recv data [%v]\n", data)
    }
}

func listenLocal(port int) {
	service, err := net.ResolveUDPAddr("udp4", ":"+strconv.Itoa(port))
	if err != nil {
		jklog.L().Errorln("resolve udp fail: ", err)
		return
	}

	waitresp, err := net.ListenUDP("udp", service)
	// defer waitresp.Close()

	jklog.L().Infoln("Start to listen ", port)
	go func() {
		for {
			buf := make([]byte, 1024)
			n, addr, err := waitresp.ReadFromUDP(buf[0:])
			if err != nil {
				jklog.L().Errorln("accept err :", err)
				break
			}
			if n == 0 {
				time.Sleep(500 * time.Millisecond)
			} else {
				jklog.L().Infoln("received data ", n, " of ", string(buf[0:n]))

				_, err := waitresp.WriteToUDP([]byte("I got your message!"), addr)
				if err != nil {
					jklog.L().Errorln("write response fail: ", err)
					continue
				}
				jklog.L().Infoln("write response len ", n, " to ")
			}
		}
	}()
}

var (
    addr = flag.String("addr", "", "dial to")
    port = flag.Int("port", 2782, "port to dial")
    content = flag.String("content", "", "content to send")
    recv = flag.Bool("recv", true, "If recv resposne")
)

func main() {
    flag.Parse()

	//listenLocal(port)

	connect_out(*addr, *port)
	for {
		time.Sleep(200 * time.Millisecond)
	}
}
