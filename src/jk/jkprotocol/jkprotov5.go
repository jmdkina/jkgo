package jkprotocol

import (
	"encoding/json"
	"errors"
	"jk/jklog"
)

// { "preheader": { "ver":1, "crypto":1 },
//   "header": { "cmd": "", "subcmd":"", "id":"", "transaction", "xxxxx", "resp":false },
//   "body": { "data":"xxx" } }

type V5PreHeader struct {
	Version float32 `json:version`
	Crypto  int     // 0: no crypto , 1 aes
}

type V5Header struct {
	Cmd         string
	SubCmd      string
	Id          string
	Transaction string // it must be uniqueue
	Resp        bool
}

type V5Body struct {
	Data interface{}
}

type V5Base struct {
	PreHeader V5PreHeader `json:preheader`
	Header    V5Header
	Body      V5Body
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

func (base *V5Base) makeCmd(cmd, subcmd string, resp bool, data string) (string, error) {
	base.base(cmd, subcmd, resp)
	base.Body.Data = data

	d, err := json.Marshal(base)
	if err != nil {
		return "", err
	}
	return string(d), nil
}

func (base *V5Base) Register(data string) (string, error) {
	return base.makeCmd("Register", "", false, data)
}

func (base *V5Base) RegisterReponse(data string) (string, error) {
	return base.makeCmd("Register", "", true, data)
}

func (base *V5Base) Keepalive(data string) (string, error) {
	return base.makeCmd("Keepalive", "", false, data)
}

func (base *V5Base) KeepaliveResponse(data string) (string, error) {
	return base.makeCmd("Keepalive", "", true, data)
}

func (base *V5Base) Leave(data string) (string, error) {
	return base.makeCmd("Leave", "", false, data)
}

func (base *V5Base) LeaveResponse(data string) (string, error) {
	return base.makeCmd("Leave", "", true, data)
}

func (base *V5Base) Parse(data string) (interface{}, int, error) {
	jklog.L().Debugf("received data [%s]\n", data)
	v5 := V5Base{}
	err := json.Unmarshal([]byte(data), &v5)
	if err == nil {
		switch v5.Header.Cmd {
		case "Register":
			return v5, JK_PROTOCOL_C_REGISTER, nil
			break
		case "Keepalive":
			return v5, JK_PROTOCOL_C_KEEPALIVE, nil
			break
		case "Leave":
			return v5, JK_PROTOCOL_C_LEAVE, nil
			break
		default:
			return nil, 0, errors.New("Unsupported command type")
			break
		}
	}
	jklog.L().Warnln("Error unmarshal ", err)
	return nil, 0, err
}
