package jkstatus

import (
	"github.com/alecthomas/log4go"
	"github.com/buger/jsonparser"
	"jk/jkprotocol"
	"net"
)

type Process struct {
	proto *jkprotocol.JKProtocolWrap
	conn  net.Conn
}

func ParseData(data string, conn net.Conn) (*Process, error) {
	p := &Process{}
	var err error
	p.proto, err = jkprotocol.NewJKProtocolWrap(jkprotocol.JK_PROTOCOL_VERSION_5)
	if err != nil {
		return nil, err
	}

	_, err = p.proto.Parse(data)
	if err != nil {
		return nil, err
	}
	p.conn = conn

	return p, nil
}

func (p *Process) isKeepalive() bool {
	return p.proto.CmdType == jkprotocol.JK_PROTOCOL_C_KEEPALIVE
}

func (p *Process) HandleMsg() bool {
	if p.isKeepalive() {
		str, err := p.proto.KeepaliveResponse("")
		if err != nil {
			return false
		}
		log4go.Debug("Give Response of keepalive msg %s ", str)
		p.conn.Write([]byte(str))
		return true
	} else {

	}
	return true
}
