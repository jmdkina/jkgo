package main

import (
	"jk/jklog"
	"net"
)

func DoRecv() {
	addr, err := net.ResolveTCPAddr("tcp", "192.168.133.112:10011")
	if err != nil {
		jklog.L().Infoln("Resolve tcp addr fail, err: ", err)
		return
	}
	lis, err := net.ListenTCP("tcp", addr)
	if err != nil {
		jklog.L().Infoln("listen failed, err: ", err)
		return
	}

	for {
		jklog.L().Infoln("Start to accept: ", addr.String())
		conn, err := lis.Accept()
		if err != nil {
			jklog.L().Errorln("fail accept: ", err)
			break
		}
		jklog.L().Infoln("Recv from: ", conn.RemoteAddr().String())
		conn.Write([]byte("OK"))
	}
}

func main() {
	DoRecv()
}
