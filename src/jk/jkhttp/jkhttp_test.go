package jkhttp

import (
	"jk/jklog"
	"testing"
)

func TestKFPostData(t *testing.T) {
	testData := "{ \"MsgBody\": [ { \"PARAMS\": { \"serverIp\": \"128.199.95.213\"}}], " +
		"\"MsgHead\": { \"ServiceCode\": \"initServerPortMsg\", \"SrcSysID\": \"0001\", " +
		"\"SrcSysSign\": \"0a2883e039909e776094edbed063bb88\", \"transactionID\": \"1436232616993\"}"

	resp, err := JKPostData("api.melinkr.cn,114.215.128.113", 8085, "unite/service", []byte(testData), 50)
	if err != nil {
		jklog.L().Errorln("error: ", err)
		return
	}
	jklog.L().Infoln(string(resp))
}
