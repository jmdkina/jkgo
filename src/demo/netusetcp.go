package main

import (
	"io"
	// crypto "jk/jkeasycrypto"
	// "io/ioutil"
	"jk/jklog"
	"net"
	"os"
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

	outdata := "Send to data with tcp for test\n"
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
	f, _ := os.OpenFile("/tmp/recvdata.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
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
					// plain := crypto.JKAESDecrypt([]byte("32bitstringtofor"), string(buf[0:n]))
					plain := string(buf[0:n])
					f.WriteString(plain)
					f.WriteString("\n")
					jklog.L().Infoln("read out data ", n, " of ", plain)

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

	listenLocalTcp(22224)

	time.Sleep(5000 * time.Millisecond)
	// jk_connect_out = true
	if jk_connect_out {
		connect_out_tcp("192.168.0.153", 10044)
		// connect_out_tcp("192.168.99.144", 44444)
	}

	for {
		time.Sleep(2000 * time.Millisecond)
		// connect_out_tcp("192.168.99.144", 44444)
	}
}
