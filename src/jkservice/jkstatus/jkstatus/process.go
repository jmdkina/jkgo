package jkstatus

import (
	"encoding/json"
	"jk/jklog"
	"jk/jksys"

	// "github.com/buger/jsonparser"
	"jk/jkprotocol"
	"net"
)

/**
 * This model receive message from server, then parse msg with protocol,
 * and save data to struct, give response if necessary
 */

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
		jklog.L().Errorf("Parser fail %s\n %s\n", err, p.data)
		return false
	}
	p.proto = pp
	jklog.L().Debugf("status process receive message: %v\n", *p.proto)

	switch p.proto.H.C {
	case jkprotocol.JKP_V6_KEEPALIVE_NAME:
		if p.proto.H.R {
			str, err := p.proto.JKProtoV6MakeKeepaliveResponse()
			if err != nil {
				jklog.L().Errorln("Generate Keepalive string err ", err)
				return false
			}
			jklog.L().Debugln("Give Response of keepalive msg ", str)
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
			jklog.L().Errorln("json unmarshal fail ", err)
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
