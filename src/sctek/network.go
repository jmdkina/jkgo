package sctek

import (
	"jk/jklog"
	"net"
	"strconv"
)

type SctekBroadNet struct {
	listenAddr string
	listenPort int
	udpAddr    *net.UDPAddr
	udpConn    *net.UDPConn
	connected  bool
}

func NewSctekBroadNet(addr string, port int) (*SctekBroadNet, error) {
	sb := &SctekBroadNet{
		listenAddr: addr,
		listenPort: port,
	}
	var err error
	sb.udpAddr, _ = net.ResolveUDPAddr("udp", addr+":"+strconv.Itoa(port))
	sb.udpConn, err = net.ListenUDP("udp", sb.udpAddr)
	if err != nil {
		return nil, err
	}

	sb.connected = true
	return sb, nil
}

func (sb *SctekBroadNet) Recv() ([]byte, error) {
	_buff := make([]byte, 4096)
	n, addr, err := sb.udpConn.ReadFromUDP(_buff)
	if err != nil {
		return nil, err
	}
	jklog.L().Infoln(addr)
	return _buff[:n], nil
}

func (sb *SctekBroadNet) Clear() {
	sb.udpConn.Close()
	sb.connected = false
}
