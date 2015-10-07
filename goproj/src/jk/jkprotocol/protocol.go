package jkprotocol

import (
	// "jk/jkcommon"
	"jk/jklog"
	"strconv"
	"strings"
)

const (
	jk_seperate = "\r\n"
	jk_version  = "0.0.1"
)

const (
	JK_PROTOCOL_CMD_REGISTER = 1 << iota
	JK_PROTOCOL_CMD_NOTIFY
	JK_PROTOCOL_CMD_CONTROL
)

const (
	JK_PROTOCOL_CMD_SAVEFILE = 1 << iota
)

type JKProtocolHead struct {
	version    string
	command    int
	subcommand int
	id         string
}

type JKProtocol struct {
	head JKProtocolHead

	data string
}

func NewJKProtocol() *JKProtocol {
	pro := &JKProtocol{}
	pro.head.version = jk_version
	return pro
}

func (p *JKProtocol) toString() string {
	return p.head.version + jk_seperate + p.head.id + jk_seperate +
		strconv.Itoa(p.head.command) + jk_seperate + strconv.Itoa(p.head.subcommand)
}

func (p *JKProtocol) GenerateRegister(id string) string {
	p.head.command = JK_PROTOCOL_CMD_REGISTER
	p.head.id = id
	return p.toString()
}

func (p *JKProtocol) GenerateNotifySaveFile(filename string) string {
	p.head.command = JK_PROTOCOL_CMD_NOTIFY
	p.head.subcommand = JK_PROTOCOL_CMD_SAVEFILE
	return p.toString() + jk_seperate + filename
}

func (p *JKProtocol) GenerateControlSaveFile(filename string, data string) string {
	p.head.command = JK_PROTOCOL_CMD_CONTROL
	p.head.subcommand = JK_PROTOCOL_CMD_SAVEFILE
	return p.toString() + jk_seperate + filename + jk_seperate + data
}

// Server parse files
func ParseJKProtocol(data string) *JKProtocol {
	// jklog.L().Debugln("data: ", data)
	strs := strings.Split(data, jk_seperate)

	if len(strs) < 4 {
		jklog.L().Errorln("wrong length")
		return nil
	}
	if strings.Compare(strs[0], jk_version) != 0 {
		jklog.L().Errorln("wrong version")
		return nil
	}
	cmd, err := strconv.Atoi(strs[2])
	if err != nil {
		return nil
	}
	subcmd, err := strconv.Atoi(strs[3])
	if err != nil {
		return nil
	}

	p := &JKProtocol{}
	p.head.command = cmd
	p.head.subcommand = subcmd
	p.head.version = strs[0]
	p.head.id = strs[1]

	headlen := len(p.head.version) + len(p.head.id) + len(strs[2]) +
		len(strs[3]) + len(jk_seperate)*4
	// datalen := len(data) - headlen

	if headlen < len(data) {
		p.data = data[headlen:]
	}

	return p
}

func (p *JKProtocol) ID() string {
	return p.head.id
}

func (p *JKProtocol) Command() int {
	return p.head.command
}

func (p *JKProtocol) SubCommand() int {
	return p.head.subcommand
}

func (p *JKProtocol) ParseFilename() string {
	if p.head.command == JK_PROTOCOL_CMD_NOTIFY && p.head.subcommand == JK_PROTOCOL_CMD_SAVEFILE {
		strs := strings.Split(p.data, jk_seperate)
		if len(strs) < 1 {
			jklog.L().Errorln("length error")
			return ""
		}
		return strs[0]
	}
	return ""
}

func (p *JKProtocol) ParseFilenameData() (string, string) {
	if p.head.command == JK_PROTOCOL_CMD_CONTROL && p.head.subcommand == JK_PROTOCOL_CMD_SAVEFILE {
		strs := strings.Split(p.data, jk_seperate)
		if len(strs) < 2 {
			jklog.L().Errorln("length error")
			return "", ""
		}
		return strs[0], strs[1]
	}
	return "", ""
}
