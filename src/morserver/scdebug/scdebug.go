package main

/**
 * Connect to distination with addr and port for debug device
 */

import (
	"bufio"
	"flag"
	"fmt"
	"jk/jklog"
	"net"
	"os"
	"strconv"
	"time"
)

var (
	addr = flag.String("addr", "", "dial to")
	port = flag.Int("port", 12310, "port to dial")
)

type SCDebug struct {
	addr  string
	port  int
	nethd *net.UDPConn
}

func (sc *SCDebug) Connect(addr string, port int) error {
	sc.addr = addr
	sc.port = port
	str := addr + ":" + strconv.Itoa(port)
	addrto, err := net.ResolveUDPAddr("udp4", str)
	if err != nil {
		jklog.L().Errorln(err)
		return err
	}

	sc.nethd, err = net.DialUDP("udp", nil, addrto)
	if err != nil {
		jklog.L().Errorln("dial udp failed ", err)
		return err
	}

	jklog.L().Infof("connect to %s success\n", str)
	return nil
}

func (sc *SCDebug) Close() {
	sc.nethd.Close()
}

func (sc *SCDebug) Debug() {
	rdata := make([]byte, 10240)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Input > ")

		data, _, err := reader.ReadLine()
		if err != nil {
			jklog.L().Errorln("read err ", err)
			break
		}
		if string(data) == "quit" {
			break
		}
		sc.nethd.Write([]byte(data))

		sc.nethd.SetReadDeadline(time.Now().Add(5 * time.Second))
		n, err := sc.nethd.Read(rdata)
		if err != nil {
			jklog.L().Errorln("read err ", err)
			continue
		}
		jklog.L().Infof("Read out response [%d] [%s]\n", n, string(rdata[:n]))
	}
}

func main() {
	flag.Parse()

	sc := SCDebug{}

	err := sc.Connect(*addr, *port)
	if err != nil {
		return
	}

	sc.Debug()
	sc.Close()
}
