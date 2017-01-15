package jkprotocol

import (
	"encoding/json"
)

// { "header": { "ver": 1, "cmd": "", "subcmd":"", "crypt":1, "id":"", "transaction", "xxxxx", "resp":false },
//   "body": { "data":"xxx" } }

type V5Header struct {
	version     int
	Cmd         string
	SubCmd      string
	Crypt       int
	Id          string
	Transaction string
}

// Command list
// Register
type RegisterBody struct {
	Data string
}

type V5Register struct {
	header V5Header
	body   RegisterBody
}

const (
	V5_VERSION = 0.1
)

func V5CmdRegister(crypt int, id, transaction string, data string) (string, error) {
	h := V5Header{
		version:     V5_VERSION,
		Cmd:         "REGISTER",
		SubCmd:      "",
		Crypt:       crypt,
		Id:          id,
		Transaction: transaction,
	}
	b := RegisterBody{
		Data: data,
	}
	v5reg := V5Register{
		header: h,
		body:   b,
	}
	d, err := json.Marshal(v5reg)
	if err != nil {
		return "", err
	}
	return string(d), nil
}
