package main

import (
	"encoding/json"
	"io/ioutil"
	"jk/jklog"
)

type NetCheckParams struct {
	SouceAddr string `json:SourceAddr`
	DstAddr   string `json:DstAddr`
	Duration  int64  `json:Duration` // how long to test
	Reverse   bool   `json:Reverse`  // If need change SourceAddr and DstAddr
	Length    int64  `json:Length`
}

func DebugParams(p *NetCheckParams) {
	jklog.L().Debugln(p)
}

func NetCheckParamsInit(file string) (*NetCheckParams, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	ncp := &NetCheckParams{}
	err = json.Unmarshal(data, ncp)
	if err != nil {
		return nil, err
	}
	DebugParams(ncp)
	return ncp, nil
}
