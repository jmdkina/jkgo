package main

import "net"

func main() {
	laddr := net.UDPAddr{
		IP:   net.IPv4(192, 168, 137, 224),
		Port: 12306,
	}
	// 这里设置接收者的IP地址为广播地址
	raddr := net.UDPAddr{
		IP:   net.IPv4(255, 255, 255, 255),
		Port: 12306,
	}
	conn, err := net.DialUDP("udp", &laddr, &raddr)
	if err != nil {
		println(err.Error())
		return
	}
	conn.Write([]byte(`hello peers`))
	conn.Close()

}
