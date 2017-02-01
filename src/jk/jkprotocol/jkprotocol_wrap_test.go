package jkprotocol

import (
	"testing"
	"strings"
)

func TestJKProtocolWrap_Register(t *testing.T) {
	wrap, err := NewJKProtocolWrap(JK_PROTOCOL_VERSION_5)
	if err != nil {
		t.Fatal("Error: ", err)
	}
	str, err := wrap.Register("WrapRegister")
	if err != nil {
		t.Fatal("Error: ", err)
	}
	exceptstr := "{\"PreHeader\":{\"Version\":0.1,\"Crypto\":0}," +
		"\"Header\":{\"Cmd\":\"Register\",\"SubCmd\":\"\",\"Id\":\"jkprotov5\"," +
		"\"Transaction\":\"jkprotov5-2017\",\"Resp\":false}," +
		"\"Body\":{\"Data\":\"WrapRegister\"}}"
	if strings.Compare(str, exceptstr) != 0 {
		t.Fatalf("Error except str [%s], but real is [ %s]\n", exceptstr, str)
	}
}