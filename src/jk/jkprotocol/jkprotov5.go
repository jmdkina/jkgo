package jkprotocol

import (
	"encoding/json"
	"strings"
	"errors"
	"jk/jklog"
)

const (
    JK_PROTOCOL_V5_REGESTER = 1 << iota
	JK_PROTOCOL_V5_LEAVE
	JK_PROTOCOL_V5_KEEPALIVE
)
// { "preheader": { "ver":1, "crypto":1 },
//   "header": { "cmd": "", "subcmd":"", "id":"", "transaction", "xxxxx", "resp":false },
//   "body": { "data":"xxx" } }

type V5PreHeader struct {
	Version     float32     `json:version`
	Crypto      int       // 0: no crypto , 1 aes
}

type V5Header struct {
	Cmd         string
	SubCmd      string
	Id          string
	Transaction string   // it must be uniqueue
	Resp        bool
}

type V5Body struct {
	Data        interface{}
}

type V5Base struct {
	PreHeader   V5PreHeader    `json:preheader`
	Header      V5Header
}

type V5Register struct {
	V5Base           `json:base`
	Body        V5Body
}

const (
	V5_VERSION = 0.1
)

// Auto set crypto id  trasaction and so on...
func (base *V5Base) base(cmd, subcmd string, resp bool) {
	base.PreHeader.Crypto = 0
	base.PreHeader.Version = V5_VERSION
	base.Header.Cmd = cmd
	base.Header.SubCmd = subcmd
	// TODO: how to generate id
	base.Header.Id = "jkprotov5"
	base.Header.Resp = resp
	// TODO: how to generate transaction
	base.Header.Transaction = "jkprotov5-2017"
}

func NewV5Register(data string) (V5Register, error) {
	v5reg := V5Register{}
	v5reg.base("Register", "", false)
	v5reg.Body.Data = data

	return v5reg, nil
}

func NewV5RegisterResponse(data string) (V5Register, error) {
	v5reg := V5Register{}
	v5reg.base("Register", "", true)
	v5reg.Body.Data = data
	return v5reg, nil
}

func (reg *V5Register) String() (string, error) {
	d, err := json.Marshal(reg)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

type V5Keepalive struct {
	V5Base
	Body            V5Body
}

func NewV5Keepalive(data string) (V5Keepalive, error) {
	v5keep := V5Keepalive{}
	v5keep.base("Keepalive", "", false)

	v5keep.Body.Data = data
	return v5keep, nil
}

func NewV5KeepaliveResponse(data string) (V5Keepalive, error) {
	v5keep := V5Keepalive{}
	v5keep.base("Keepalive", "", true)

	v5keep.Body.Data = data
	return v5keep, nil
}

func (keep *V5Keepalive) String() (string, error) {
	d, err := json.Marshal(keep)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

type V5Leave struct {
	V5Base
	Body       V5Body
}

func NewV5Leave(data string) (V5Leave, error) {
	leave := V5Leave{}
	leave.base("Leave", "", false)
	leave.Body.Data = data

	return leave, nil
}

func NewV5LeaveResponse(data string) (V5Leave, error) {
	leave := V5Leave{}
	leave.base("Leave", "", true)
	leave.Body.Data = data

	return leave, nil
}

func (leave V5Leave) String() (string, error) {
	d, err := json.Marshal(leave)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func (base *V5Base) Parse(data string) (interface{}, int, error) {
	jklog.L().Debugf("received data [%s]\n", data)
	reg := V5Register{}
	err := json.Unmarshal([]byte(data), &reg)
	if err == nil {
		if strings.Compare(reg.Header.Cmd, "Register") == 0 {
			return reg, JK_PROTOCOL_V5_REGESTER, nil
		}
	} else {
		jklog.L().Warnln("Error unmarshal register ", err)
	}

	keep := V5Keepalive{}
	err = json.Unmarshal([]byte(data), &keep)
	if err == nil {
		if strings.Compare(keep.Header.Cmd, "Keepalive") == 0 {
			return keep, JK_PROTOCOL_V5_KEEPALIVE, nil
		}
	} else {
		jklog.L().Warnln("Error unmarshal keepalive ", err)
	}

	leave := V5Leave{}
	err = json.Unmarshal([]byte(data), &leave)
	if err == nil {
		if strings.Compare(keep.Header.Cmd, "Leave") == 0 {
			return leave, JK_PROTOCOL_V5_LEAVE, nil
		}
	} else {
		jklog.L().Warnln("Error unmarshal leave ", err)
	}

	return nil, 0, errors.New("Unknown command type")
}

