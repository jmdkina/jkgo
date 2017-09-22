package jkprotocol

import (
	"encoding/json"
	"time"
)

const (
	JKP_V6_KEEPALIVE_NAME = "Keepalive"
)

type JKProtoV6Header struct {
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
	p.H.C = cmd
	p.H.T = time
	p.H.R = resp
	p.H.ID = id
	p.B = data
	v, err := json.Marshal(p)
	return string(v), err
}

func JKProtoV6MakeKeepalive(id string) (string, error) {
	p := &JKProtoV6{}
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
