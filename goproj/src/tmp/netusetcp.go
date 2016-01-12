package main

import (
	"io"
	"jk/jklog"
	"net"
	"strconv"
	"time"
)

func connect_out_tcp(addr string, port int) {
	jklog.L().Infoln("Will connect to ", addr, ":", port)
	addrto, err := net.ResolveTCPAddr("tcp", addr+":"+strconv.Itoa(port))
	if err != nil {
		jklog.L().Infoln("error resolve tcp addr ", err)
		return
	}

	sendto, err := net.DialTCP("tcp", nil, addrto)
	if err != nil {
		jklog.L().Infoln("erro dial tcp ", err)
		return
	}
	defer sendto.Close()

	outdata := "Send to data with tcp for test"
	n, _ := sendto.Write([]byte(outdata))
	jklog.L().Infoln("Write out data len ", n)
}

func listenLocalTcp(port int) {
	addrto, err := net.ResolveTCPAddr("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		jklog.L().Errorln("error reolve tcp addr ", err)
		return
	}

	listen, err := net.ListenTCP("tcp", addrto)
	if err != nil {
		jklog.L().Errorln("list local with tcp error ", err)
		return
	}

	jklog.L().Infoln("Start to listen with tcp of ", port)
	go func() {
		for {
			jklog.L().Infoln("start to accept from remote ...")
			from, err := listen.Accept()

			if err != nil {
				jklog.L().Errorln("accept error ", err)
				return
			}
			go func() {
				for {
					buf := make([]byte, 2<<12)
					n, err := from.Read(buf)
					if err == io.EOF {
						break
					}
					if err != nil {
						jklog.L().Infoln("read error ", err)
						break
					}
					jklog.L().Infoln("read out data ", n, " of ", string(buf[0:n]))

					jklog.L().Infoln("read out data from ", from.RemoteAddr().String())
					from.Write([]byte("I have received your data."))
				}
			}()
		}
		jklog.L().Errorln("Can't be here")
	}()

}

func main() {
	jk_connect_out := false

	listenLocalTcp(0xade)

	time.Sleep(5000 * time.Millisecond)
	// jk_connect_out = true
	if jk_connect_out {
		connect_out_tcp("192.168.0.153", 10044)
	}

	for {
		time.Sleep(200 * time.Millisecond)
	}
}
