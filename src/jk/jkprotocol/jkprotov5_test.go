package jkprotocol

import (
	"testing"
	"strings"
)

func TestJKProtoV5Reg(t *testing.T) {
	v5base := V5Base{}
	str, err := v5base.Register("FromRegister")
	if err != nil {
		t.Fatalf("make register command failed, %v\n", err)
	}
	needstr := "{\"PreHeader\":{\"Version\":0.1,\"Crypto\":0}," +
	           "\"Header\":{\"Cmd\":\"Register\",\"SubCmd\":\"\",\"Id\":\"jkprotov5\"," +
	           "\"Transaction\":\"jkprotov5-2017\",\"Resp\":false}," +
	           "\"Body\":{\"Data\":\"FromRegister\"}}"
	if strings.Compare(str, needstr) != 0 {
		t.Fatalf("Except [%s], but [%s]\n", str, needstr)
	}
}
