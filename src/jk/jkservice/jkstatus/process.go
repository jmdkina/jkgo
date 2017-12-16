package jkstatus

import (
	"encoding/json"
	"github.com/alecthomas/log4go"
	"jk/jksys"

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
	log4go.Debug("status process receive message: %v", *p.proto)

	switch p.proto.H.C {
	case jkprotocol.JKP_V6_KEEPALIVE_NAME:
		if p.proto.H.R {
			str, err := p.proto.JKProtoV6MakeKeepaliveResponse()
			if err != nil {
				log4go.Error("Generate Keepalive string err %s ", err)
				return false
			}
			log4go.Debug("Give Response of keepalive msg %s ", str)
			p.SS.remoteInstance[p.conn].UpdateTime()
			p.SS.remoteInstance[p.conn].Info.ID = p.proto.H.ID
			p.conn.Write([]byte(str))
		}
	case jkprotocol.JKP_V6_REGISTER_NAME:
		// If use interface, can't transfer back to correct interface
		// so we define a new with body of the struct
		// Why can't transfer from interface ???
		type OutSys struct {
			H jkprotocol.JKProtoV6Header
			B jksys.KFSystemInfo
		}
		var outSys OutSys
		err := json.Unmarshal([]byte(p.data), &outSys)
		if err != nil {
			log4go.Error("json unmarshal fail ", err)
			break
		}
		p.SS.remoteInstance[p.conn].SysInfo = outSys.B
		if p.proto.H.R {
			str, _ := pp.JKProtoV6MakeCommonResponse("")
			p.conn.Write([]byte(str))
		}
	}
	return true
}
