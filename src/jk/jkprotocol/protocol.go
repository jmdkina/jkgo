package jkprotocol

import (
	// "jk/jkcommon"
	"fmt"
	"jk/jklog"
	"jk/kfmd5"
	"strconv"
	"strings"
	"time"
)

const (
	jk_seperate = "\r\n"
	jk_version  = "0.0.1"
)

const (
	JK_PROTOCOL_CMD_REGISTER = 1 << iota
	JK_PROTOCOL_CMD_NOTIFY        // 2
	JK_PROTOCOL_CMD_CONTROL       // 4
)

const (
	JK_PROTOCOL_CMD_SAVEFILE = 1 << iota
)

const (
	jk_protocol_key = "ab3w-be82f231-aa-cd1b"
)

type JKProtocolHead struct {
	version    string
	command    int
	subcommand int
	id         string
	timeval    int64
	sign       string
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
	p.head.timeval = time.Now().Unix()
	firststring := p.head.version + jk_seperate + p.head.id + jk_seperate +
		strconv.Itoa(p.head.command) + jk_seperate + strconv.Itoa(p.head.subcommand) + jk_seperate +
		fmt.Sprintf("%d", p.head.timeval)
	key := kfmd5.Sum([]byte(firststring + jk_seperate + jk_protocol_key))
	return firststring + jk_seperate + fmt.Sprintf("%x", key)
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

func (p *JKProtocol) GenerateResponseOK() string {
	p.data = "OK"
	return p.toString() + jk_seperate + p.data
}

func (p *JKProtocol) GenerateResponseFail() string {
	p.data = "FAIL"
	return p.toString() + jk_seperate + p.data
}

func (p *JKProtocol) checkJKProtocolHeadValid() bool {
	firststring := p.head.version + jk_seperate + p.head.id + jk_seperate +
		strconv.Itoa(p.head.command) + jk_seperate + strconv.Itoa(p.head.subcommand) + jk_seperate +
		fmt.Sprintf("%d", p.head.timeval)
	key := kfmd5.Sum([]byte(firststring + jk_seperate + jk_protocol_key))
	if strings.Compare(p.head.sign, fmt.Sprintf("%x", key)) == 0 {
		return true
	}
	jklog.Lfile().Errorf("sign: [%s][%s]\n", p.head.sign, fmt.Sprintf("%x", key))
	jklog.Lfile().Debugln("the string is : ", firststring)
	return false
}

// Server parse files
func ParseJKProtocol(data string) *JKProtocol {
	// jklog.L().Debugln("data: ", data)
	strs := strings.Split(data, jk_seperate)

	if len(strs) < 6 {
		jklog.L().Errorln("wrong length")
		return nil
	}
	if strings.Compare(strs[0], jk_version) != 0 {
		jklog.Lfile().Errorln("wrong version")
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
	// p.head.timeval = strconv.Atoi(strs[4])
	p.head.timeval, err = strconv.ParseInt(strs[4], 10, 64)
	if err != nil {
		jklog.Lfile().Errorln("parse time value failed. ", err)
		return nil
	}
	p.head.sign = strs[5]

	if !p.checkJKProtocolHeadValid() {
		jklog.Lfile().Errorln("head valid check fail")
		return nil
	}

	headlen := len(p.head.version) + len(p.head.id) + len(strs[2]) +
		len(strs[3]) + len(jk_seperate)*6 + len(strs[4]) + len(p.head.sign)
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
			jklog.Lfile().Errorln("length error")
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
			jklog.Lfile().Errorln("length error")
			return "", ""
		}
		return strs[0], strs[1]
	}
	return "", ""
}
