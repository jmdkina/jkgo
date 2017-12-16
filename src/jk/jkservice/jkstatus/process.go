package jkstatus

import (
	"github.com/alecthomas/log4go"
	// "github.com/buger/jsonparser"
	"jk/jkprotocol"
	"net"
)

type Process struct {
	proto *jkprotocol.JKProtoV6
	conn  net.Conn
	data  string
	SS    *ServiceStatus
}

func ParseData(data string, conn net.Conn) (*Process, error) {
	p := &Process{}
	p.data = data
	p.conn = conn

	return p, nil
}

func (p *Process) HandleMsg() bool {
	pp, err := jkprotocol.JKProtoV6Parse(p.data)
	if err != nil {
		log4go.Error("Parser fail %s\n %s\n", err, p.data)
		return false
	}
	p.proto = pp
	log4go.Debug("status process receive message: ", *p.proto)

	if p.proto.H.C == jkprotocol.JKP_V6_KEEPALIVE_NAME && p.proto.H.R {
		str, err := p.proto.JKProtoV6MakeKeepaliveResponse()
		if err != nil {
			log4go.Error("Generate Keepalive string err %s ", err)
			return false
		}
		log4go.Debug("Give Response of keepalive msg %s ", str)
		p.SS.remoteInstance[p.conn].UpdateTime()
		p.conn.Write([]byte(str))
	}
	return true
}
