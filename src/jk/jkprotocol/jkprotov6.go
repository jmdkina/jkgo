package jkprotocol

import (
	"encoding/json"
	"time"
)

const (
	jkp_v6_version = "1.0.0"
)

const (
	JKP_V6_REGISTER_NAME  = "Register"
	JKP_V6_KEEPALIVE_NAME = "Keepalive"
)

type JKProtoV6Header struct {
	V  string
	C  string
	ID string
	T  int64
	R  bool
}

type JKProtoV6 struct {
	H JKProtoV6Header
	B interface{}
}

func JKProtoV6Parse(data string) (*JKProtoV6, error) {
	p := &JKProtoV6{}
	err := json.Unmarshal([]byte(data), p)
	return p, err
}

func JKProtoV6Make(cmd string, time int64, resp bool, id string, data interface{}) (string, error) {
	p := &JKProtoV6{}
	p.H.V = jkp_v6_version
	p.H.C = cmd
	p.H.T = time
	p.H.R = resp
	p.H.ID = id
	p.B = data
	v, err := json.Marshal(p)
	return string(v), err
}

func JKProtoV6MakeRegister(id string, data interface{}) (string, error) {
	return JKProtoV6Make(JKP_V6_REGISTER_NAME, time.Now().UnixNano()/1000000, false, id, data)
}

func JKProtoV6MakeKeepalive(id string) (string, error) {
	p := &JKProtoV6{}
	p.H.V = jkp_v6_version
	p.H.C = JKP_V6_KEEPALIVE_NAME
	p.H.T = time.Now().UnixNano() / 1000000
	p.H.ID = id
	p.H.R = true
	p.B = ""
	v, err := json.Marshal(p)
	return string(v), err
}

func (p *JKProtoV6) JKProtoV6MakeKeepaliveResponse() (string, error) {
	p.H.T = time.Now().UnixNano() / 1000000
	p.H.R = false
	v, err := json.Marshal(p)
	return string(v), err
}

func (p *JKProtoV6) JKProtoV6MakeCommonResponse(data interface{}) (string, error) {
	p.H.T = time.Now().UnixNano() / 1000000
	p.H.R = false
	p.B = data
	v, err := json.Marshal(p)
	return string(v), err
}
