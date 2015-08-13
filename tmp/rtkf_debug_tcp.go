package main

import (
	"flag"
	"fmt"
	"io"
	"jk/jklog"
	"net"
	"strconv"
	"time"
)

func scanLine() string {
	var c byte
	var err error
	var b []byte
	for err == nil {
		_, err = fmt.Scanf("%c", &c)

		if c != '\n' {
			b = append(b, c)
		} else {
			break
		}
	}

	return string(b)
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
				continue
			}
			for {
				fmt.Print("Cmd: > ")
				senddata := scanLine()
				if len(senddata) <= 5 || senddata == "\r" {
					continue
				}
				from.Write([]byte(senddata))

				buf := make([]byte, 2<<12)
				n, err := from.Read(buf)
				if n == 0 {
					continue
				}
				if err == io.EOF {
					break
				}
				if err != nil {
					jklog.L().Infoln("read error ", err)
					break
				}
				jklog.L().Infoln("read out data ", n, " of ", string(buf[0:n]))
			}
		}
		jklog.L().Errorln("Can't be here")
	}()

}

var (
	serverPort = flag.Int("serverPort", 23433, "which port to connect")
)

func main() {

	listenLocalTcp(*serverPort)

	for {
		time.Sleep(200 * time.Millisecond)
	}
}
