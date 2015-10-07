package jkprotocol

import (
	"strings"
	"testing"
)

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
