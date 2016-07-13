package jkprotocol

import (
	"bytes"
	"strings"
	"testing"
)

func TestJKProtocolV4(t *testing.T) {
	p, _ := JKProtoV4New()
	str := "doityour self"
	p.GenerateHeader(0, false, uint32(len(str)))
	p.GenerateBody(str)
	d, err := p.ToByte()
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	pn, _ := JKProtoV4New()
	err = pn.Parse(d)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
}

func TestJKProtoUp(t *testing.T) {
	p, _ := JKProtoUpNew(JK_PROTOCOL_VERSION_4, "123444")
	data, err := p.JKProtoUpRegister()
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	pn, _ := JKProtoUpNew(JK_PROTOCOL_VERSION_4, "123444")
	err = pn.JKProtoUpParse(data)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if strings.Compare(pn.Info.Id, "123444") != 0 {
		t.Fatalf("expect id 123444, but real is %s", pn.Info.Id)
	}
	if strings.Compare(pn.Info.Cmd, "Register") != 0 {
		t.Fatalf("expect cmd Register, but real is %s", pn.Info.Cmd)
	}
	if pn.Info.Seq != 1 {
		t.Fatalf("expect seq 1, but real is %d", pn.Info.Seq)
	}
}

func TestKFProtocol(t *testing.T) {
	p := NewKFProtocol()
	p.Init()
	p.SetCmd(KF_CMD_QUERY, KF_SUBCMD_KEEPALIVE, []byte("1234"))
	p.SetData([]byte("just for test"))
	gd, err := p.GenerateData(false)
	if err != nil {
		t.Fatal("error fo generate data : ", err)
	}
	pp, err := KFProtocolParse(gd)
	if err != nil {
		t.Fatal("error of parse protocol: ", err)
	}
	if bytes.Compare(pp.Body.Data, []byte("just for test")) != 0 {
		t.Fatal("error for data: ", string(pp.Body.Data))
	}
}

func TestKFProtocolText(t *testing.T) {
	p := NewKFProtocol()
	p.Init()
	p.SetCmd(KF_CMD_QUERY, KF_SUBCMD_KEEPALIVE, []byte("1234"))
	p.SetData([]byte("just for test"))
	gd, err := p.GenerateDataText(false)
	if err != nil {
		t.Fatal("error fo generate data : ", err)
	}
	pp, err := KFProtocolParseText(gd)
	if err != nil {
		t.Fatal("error of parse protocol: ", err)
	}
	if bytes.Compare(pp.Body.Data, []byte("just for test")) != 0 {
		t.Fatal("error for data: ", string(pp.Body.Data))
	}
}

func TestProtocol(t *testing.T) {
	p := NewJKProtocol()
	regstr := p.GenerateRegister("4efd")

	pt := ParseJKProtocol(regstr)
	if pt.ID() != "4efd" {
		t.Fatalf("need id 4efd, but real is %s\n", pt.ID())
	}

	if pt.Command() != JK_PROTOCOL_CMD_REGISTER {
		t.Fatalf("need command %d, but real is %d\n", JK_PROTOCOL_CMD_REGISTER, pt.Command())
	}

	filename := "/tmp/testf"
	notistr := p.GenerateNotifySaveFile(filename)
	pt = ParseJKProtocol(notistr)
	if pt.ID() != "4efd" {
		t.Fatalf("need id 4efd, but real is %s\n", pt.ID())
	}
	if pt.Command() != JK_PROTOCOL_CMD_NOTIFY {
		t.Fatalf("Need command %d, but real is %d\n", JK_PROTOCOL_CMD_NOTIFY, pt.Command())
	}
	if pt.SubCommand() != JK_PROTOCOL_CMD_SAVEFILE {
		t.Fatalf("Need subcommand %d, but real is %d\n", JK_PROTOCOL_CMD_SAVEFILE, pt.SubCommand())
	}
	outfilename := pt.ParseFilename()
	if strings.Compare(outfilename, filename) != 0 {
		t.Fatalf("Need filename %s, but real is %s\n", filename, outfilename)
	}

	filedata := "this is testf"
	contstr := p.GenerateControlSaveFile(filename, filedata)
	pt = ParseJKProtocol(contstr)
	if pt.ID() != "4efd" {
		t.Fatalf("need id 4efd, but real is %s\n", pt.ID())
	}
	if pt.Command() != JK_PROTOCOL_CMD_CONTROL {
		t.Fatalf("Need command %d, but real is %d\n", JK_PROTOCOL_CMD_CONTROL, pt.Command())
	}
	if pt.SubCommand() != JK_PROTOCOL_CMD_SAVEFILE {
		t.Fatalf("Need subcommand %d, but real is %d\n", JK_PROTOCOL_CMD_SAVEFILE, pt.SubCommand())
	}
	outfilename, outdata := pt.ParseFilenameData()
	if strings.Compare(outfilename, filename) != 0 {
		t.Fatalf("Need filename %s, but real is %s\n", filename, outfilename)
	}

	if strings.Compare(filedata, outdata) != 0 {
		t.Fatalf("Need data %s, but real is %s\n", filedata, outdata)
	}
}
