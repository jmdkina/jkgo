package sctek

import (
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
	sb.udpAddr, _ = net.ResolveUDPAddr("udp", addr+":"+strconv.Itoa(port))

	return sb, nil
}

func (sb *SctekBroadNet) Recv() ([]byte, error) {
	if !sb.connected {
		var err error
		sb.udpConn, err = net.DialUDP("udp", nil, sb.udpAddr)
		if err != nil {
			return nil, err
		}
	}
	sb.connected = true

	_buff := make([]byte, 4096)
	n, err := sb.udpConn.Read(_buff)
	if err != nil {
		return nil, err
	}
	return _buff[:n], nil
}

func (sb *SctekBroadNet) Clear() {
	sb.udpConn.Close()
	sb.connected = false
}
