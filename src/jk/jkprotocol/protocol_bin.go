package jkprotocol

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"errors"
	"fmt"
	"jk/jklog"
	"strconv"
	"strings"
	"time"
)

//
// +-----------+-----------+----------+--------+-----------+-----------+
// | version   | direction | reseved  |   cmd  |  subcmd   |  code     |  (4bytes)
// | 4 bits    |   1bits   |  3 bits  |  4 bits|   12 bits |   8bits   |  (32 bits, 4 byte)
// +-----------+-----------+----------+--------+-----------+-----------+
// | id                                            |   reserve         |  (8bytes)
// | 48bits                                        |   16 bits         |  (96 bits, 12 byte)
// +-----------+-----------+----------+--------+-----------+-----------+
// |  transaction                                                      |  (8bytes)
// |   64bits                                                          |
// |                                                                   |  (160 bits, 20 byte)
// +-----------+-----------+----------+--------+-----------+-----------+
// |  sequence                                     |        length     |  (4bytes)
// |   16bits                                      |        16bits     |  (192 bits, 24 byte)
// +-----------+-----------+----------+--------+-----------+-----------+
// |  Sign                                                             |  (16bytes)
// |  16 bytes                                                         |
// |                                                                   |
// |                                                                   |  (~ bits, 40 byte)
// +-----------+-----------+----------+--------+-----------+-----------+
// |   Data                                                            |
// +-----------+-----------+----------+--------+-----------+-----------+
//
// (in bits)
// 1. version (4) : format 0.0.0.0        d[0] first
// 2. dir     (1) : 0 send, 1 response
// 3. none    (3) : reserved
// 4. Cmd     (4) : Main Command     d[1]   second
// 5. SubCmd  (12) : Sub Command      d[1-2]   third
// 6. Transaction (32) : Just use timestamp  d[3:6] fourth
// 7. seq     (16) : seq, restart when > 2^16  d[7:8] eighth
// 8. len     (16) : Data length, exclude the 16 bytes sign d[9:10] tenth
// id: 6bytes
// code: 1bytes means fail code, success when all 1, max support 2^8-1 codes, only effective when response
// code: first bits 1 success, 0 fail, others: fail code
//
// 9. Use same sequence if the data is continue, you must recalculate sign
//    each command even the data is continuing.
//
// We are not success with the binary define, we will use it later,
// but now we just use string for quick developer to use.
// string format like below.
// header: version-id-direction-cmd-subcmd-transaction-sequence-sign-length-ret-retcode
// direction : 0 for send command, 1 for response
// header must 120 bytes, pad 0 if not enough. ret and retcode only effective when direction is 1
//

const (
	KF_CMD_QUERY = 1 << iota
	KF_CMD_CONTROL
	KF_CMD_NOTIFY
)

const (
	KF_SUBCMD_REGISTER  = 1
	KF_SUBCMD_KEEPALIVE = 2

	KF_SUBCMD_COMMAND = 0x100 // execute command
	KF_SUBCMD_FILE    = 0x101 // Get file info (bidirectional)

	KF_SUBCMD_OFFLINE = 0x500
)

const (
	kf_protocol_key = "ab3w-be82f231-aa-cd1b"
)

type KFProtocolHeader struct {
	Cmd         uint
	SubCmd      uint
	Direction   uint8 // direction for 0 send, 1 response
	Version     uint
	Transaction uint64
	Sequence    uint16
	Length      uint16
	Sign        [16]byte
	Ret         bool
	Code        int
	Id          [16]byte
	IdStr       string
	hlen        uint
}

var (
	mixlen = 120
)

type KFProtocolBody struct {
	Data []byte
}

type KFProtocol struct {
	Header KFProtocolHeader
	Body   KFProtocolBody
}

var curSeq uint16 // max 2 ^ 14

func NewKFProtocol() *KFProtocol {
	return &KFProtocol{}
}

func (p *KFProtocol) Init() {
	p.Header.Version = 1
	p.Header.hlen = uint(mixlen)
	copy(p.Header.Id[:], []byte("NOID"))
}

func (p *KFProtocol) SetCmd(cmd, subcmd uint, id []byte) {
	p.Header.Cmd = cmd
	p.Header.SubCmd = subcmd
	copy(p.Header.Id[:], id[:6])
}

func (p *KFProtocol) SetData(data []byte) {
	p.Header.Length = uint16(len(data))
	p.Body.Data = data
}

func (p *KFProtocol) generateSign() [16]byte {
	str := fmt.Sprintf("%02d-%s-%d-%d-%d-%d-%s-%d", p.Header.Version,
		p.Header.IdStr, p.Header.Cmd,
		p.Header.SubCmd, p.Header.Transaction, p.Header.Sequence, kf_protocol_key, p.Header.Length)
	jklog.L().Debugln("sign is ", str)
	return md5.Sum([]byte(str))
}

func (p *KFProtocol) generate() {
	p.Header.Direction = 0
	p.Header.Sequence = curSeq + 1
	curSeq++
	p.Header.Transaction = uint64(time.Now().Unix())
	p.Header.Sign = p.generateSign()
}

func (p *KFProtocol) checkValid() bool {
	jklog.L().Debugf("cmd:%d,subcmd:%d, direction: %d, length: %d, sequence: %d\n", p.Header.Cmd,
		p.Header.SubCmd, p.Header.Direction, p.Header.Length, p.Header.Sequence)
	if p.Header.Cmd > (1<<4 - 1) {
		jklog.Lfile().Errorln("header command too long")
		return false
	}
	if p.Header.SubCmd > (1<<12 - 1) {
		jklog.Lfile().Errorln("header sub command too long")
		return false
	}
	if p.Header.Direction != 0 && p.Header.Direction != 1 {
		jklog.Lfile().Errorln("header direction invalid")
		return false
	}
	if p.Header.Length > (1<<16 - 1) {
		jklog.Lfile().Errorln("header length too long ")
		return false
	}
	if p.Header.Sequence > (1<<16 - 1) {
		jklog.Lfile().Errorln("header sequence too long")
		return false
	}
	return true
}

func (p *KFProtocol) SetResponseCode(result bool, code int) error {
	if code > 1<<7 {
		return errors.New("code is too long")
	}
	p.Header.Ret = result
	p.Header.Code = code
	return nil
}

func (p *KFProtocol) GenerateDataText(dir bool) ([]byte, error) {
	ret := p.checkValid()
	if !ret {
		return nil, errors.New("Invalid args")
	}
	if dir {
		p.Header.Direction = 1
	}

	p.generate()
	jklog.L().Debugf("start to generate data of len %d\n", p.Header.Length+uint16(p.Header.hlen))

	retvalue := "0"
	if p.Header.Ret {
		retvalue = "1"
	}

	//version-id-direction-cmd-subcmd-transaction-sequence-sign-length-ret-retcode
	str := fmt.Sprintf("%d", p.Header.Version) + "-" +
		p.Header.IdStr + "-" +
		fmt.Sprintf("%d", p.Header.Direction) + "-" +
		fmt.Sprintf("%d", p.Header.Cmd) + "-" +
		fmt.Sprintf("%d", p.Header.SubCmd) + "-" +
		fmt.Sprintf("%d", p.Header.Transaction) + "-" +
		fmt.Sprintf("%d", p.Header.Sequence) + "-" +
		fmt.Sprintf("%x", p.Header.Sign) + "-" +
		fmt.Sprintf("%d", p.Header.Length) + "-" +
		retvalue + "-" +
		fmt.Sprintf("%d", p.Header.Code)

	jklog.L().Debugf("header is %s\n", str)

	d := make([]byte, p.Header.Length+uint16(p.Header.hlen))
	copy(d[:], []byte(str))
	copy(d[p.Header.hlen:], p.Body.Data)

	return d, nil
}

func (p *KFProtocol) GenerateData(dir bool) ([]byte, error) {
	ret := p.checkValid()
	if !ret {
		return nil, errors.New("Invalid args")
	}

	p.generate()
	jklog.L().Debugln("start to generate data of len:", p.Header.Length+uint16(p.Header.hlen))
	d := make([]byte, p.Header.Length+uint16(p.Header.hlen))
	// version is 0.0.0.1
	var i uint

	for i = 0; i < 4; i++ {
		j := p.Header.Version & (1 << i)
		if j != 0 {
			d[0] |= 1 << (4 + i)
		}
	}
	// direction is true for send
	if dir {
		d[0] |= 1 << 3
	}
	// start from 8th bits, the 3rd byte
	// cmd is 4 bits

	for i = 0; i < 4; i++ {
		j := p.Header.Cmd & (1 << i)
		if j != 0 {
			d[1] |= 1 << (i + 4)
		}
	}

	// start from 12th bits
	// subcmd is 12 bits
	for i = 0; i < 8; i++ {
		j := p.Header.SubCmd & (1 << i)
		if j != 0 {
			d[2] |= 1 << i
		}
	}
	// first 4 bits in the seconds byte
	for i = 0; i < 4; i++ {
		j := p.Header.SubCmd & (1 << (8 + i))
		if j != 0 {
			d[1] |= 1 << i
		}
	}

	// code
	if p.Header.Ret {
		d[3] |= 1 << 7
	}
	if p.Header.Code > 1<<7 {
		return nil, errors.New("code is too long")
	}
	for i = 0; i < 7; i++ {
		j := p.Header.Code & (1 << i)
		if j != 0 {
			d[3] |= 1 << i
		}
	}

	dindex := 4
	// Set id (6 bytes)
	jklog.Lfile().Debugln("start to id")
	copy(d[dindex:dindex+6], p.Header.Id[:])

	// 2 bytes reservd

	// transcation
	dindex += 8
	jklog.Lfile().Debugln("start to timestamp :", p.Header.Transaction)

	// timestamp from the third byte
	tr := make([]byte, 8)
	binary.LittleEndian.PutUint64(tr, p.Header.Transaction)
	copy(d[dindex:dindex+8], tr[:8])

	jklog.Lfile().Debugln("start sequence")

	dindex += 8
	// seq from the eighth byte
	tr = make([]byte, 2)
	binary.LittleEndian.PutUint16(tr, p.Header.Sequence)
	copy(d[dindex:dindex+2], tr)

	jklog.Lfile().Debugln("start length")

	dindex += 2
	// len from the tenth byte
	tr = make([]byte, 2)
	binary.LittleEndian.PutUint16(tr, p.Header.Length)
	copy(d[dindex:dindex+2], tr)

	jklog.Lfile().Debugln("start sign")

	dindex += 2
	// sign  16 bytes
	nsign := p.Header.Sign[:]
	jklog.L().Debugf("len d : %d, len sign: %d, %d\n", cap(d), len(nsign), dindex)
	copy(d[dindex:dindex+16], nsign)

	dindex += 16

	jklog.L().Debugln("start data")
	jklog.L().Debugf("len start d : %d, data : %d\n", dindex, len(p.Body.Data))
	// Data
	n := copy(d[dindex:], p.Body.Data)

	jklog.L().Debugf("data string: %s, len copy: %d\n", string(d[dindex:]), n)

	return d, nil
}

func KFProtocolParseText(data []byte) (*KFProtocol, error) {
	p := NewKFProtocol()
	jklog.L().Debugln("start to parse data of len: ", len(data))

	if len(data) < mixlen {
		return nil, errors.New("Not enough data")
	}

	p.Header.hlen = uint(mixlen)

	strs := strings.Split(string(data), "-")
	//version-id-direction-cmd-subcmd-transaction-sequence-sign-length-ret-retcode
	if len(strs) < 11 {
		return nil, errors.New("less command header")
	}

	vv, _ := strconv.Atoi(strs[0])
	p.Header.Version = uint(vv)
	p.Header.IdStr = strs[1]
	vv, _ = strconv.Atoi(strs[2])
	p.Header.Direction = uint8(vv)
	vv, _ = strconv.Atoi(strs[3])
	p.Header.Cmd = uint(vv)
	vv, _ = strconv.Atoi(strs[4])
	p.Header.SubCmd = uint(vv)
	vv, _ = strconv.Atoi(strs[5])
	p.Header.Transaction = uint64(vv)
	vv, _ = strconv.Atoi(strs[6])
	p.Header.Sequence = uint16(vv)
	copy(p.Header.Sign[:], []byte(strs[7]))
	vv, _ = strconv.Atoi(strs[8])
	p.Header.Length = uint16(vv)
	vv, _ = strconv.Atoi(strs[9])
	if vv == 1 {
		p.Header.Ret = true
	} else {
		p.Header.Ret = false
	}
	vv, _ = strconv.Atoi(strs[10])
	p.Header.Code = vv

	jklog.L().Debugf("version: %d, id: %s, direction: %d, cmd: %d, subcmd: %d, transaction: %d "+
		", sequence: %d, sign: %s, length: %d, ret: %v, retcode: %d\n", p.Header.Version, string(p.Header.Id[:]),
		p.Header.Direction, p.Header.Cmd, p.Header.SubCmd, p.Header.Transaction, p.Header.Sequence,
		string(p.Header.Sign[:]), p.Header.Length, p.Header.Ret, p.Header.Code)

	p.Body.Data = make([]byte, p.Header.Length)
	copy(p.Body.Data[:], data[p.Header.hlen:uint16(p.Header.hlen)+p.Header.Length])

	return p, nil
}

func KFProtocolParse(data []byte) (*KFProtocol, error) {
	p := NewKFProtocol()
	jklog.Lfile().Debugf("start parse data len: %d\n", len(data))
	if len(data) < mixlen {
		return nil, errors.New("data too less, not enough data")
	}
	p.Header.hlen = uint(mixlen)

	// Version
	v := data[0]
	var vi uint
	var i uint
	for i = 0; i < 4; i++ {
		j := v & (1 << (4 + i))
		if j != 0 {
			vi |= (1 << i)
		}
	}

	p.Header.Version = vi

	jklog.Lfile().Debugln("start direction...")

	// Direction
	j := v & (1 << 3)
	if j != 0 {
		p.Header.Direction = 1
	}

	jklog.Lfile().Debugln("start command ...")

	// Cmd
	v = data[1]
	vi = 0
	for i = 0; i < 4; i++ {
		j := v & (1 << (4 + i))
		if j != 0 {
			vi |= (1 << i)
		}
	}

	p.Header.Cmd = vi

	jklog.Lfile().Debugln("start subcommand ...")
	// SubCmd
	vi = 0
	for i = 0; i < 4; i++ {
		j := v & (1 << i)
		if j != 0 {
			vi |= (1 << (i + 8))
		}
	}

	// SubCmd also
	v = data[2]
	for i = 0; i < 8; i++ {
		j := v & (1 << i)
		if j != 0 {
			vi |= (1 << i)
		}
	}

	p.Header.SubCmd = vi

	p.Header.Ret = false
	v = data[3]
	rett := v & (1 << 7)
	if rett != 0 {
		p.Header.Ret = true
	}
	// 1 byte code for response
	for i = 0; i < 7; i++ {
		j := v & (1 << i)
		if j != 0 {
			p.Header.Code |= (1 << i)
		}
	}

	// id
	jklog.Lfile().Debugln("start id")
	dindex := 4
	id := data[dindex : dindex+6]
	copy(p.Header.Id[:], id)

	// 2 bytes reserved

	jklog.Lfile().Debugln("start transaction ...")
	dindex += 8
	// transaction (8 bytes)

	nv := data[dindex : dindex+8]
	p.Header.Transaction = binary.LittleEndian.Uint64(nv)

	jklog.Lfile().Debugln("start sequence ...")
	dindex += 8
	// sequence (2 bytes)
	ns := data[dindex : dindex+2]
	p.Header.Sequence = binary.LittleEndian.Uint16(ns)

	jklog.Lfile().Debugln("start data ...")
	dindex += 2
	// length (2 bytes)
	nl := data[dindex : dindex+2]
	p.Header.Length = binary.LittleEndian.Uint16(nl)
	if p.Header.Length > 1<<16-1 {
		return nil, errors.New("Error of data length too long")
	}

	jklog.Lfile().Debugln("start sign ...")
	dindex += 2
	// sign (16bytes)
	copy(p.Header.Sign[:], data[dindex:dindex+16])

	jklog.Lfile().Debugln("start data ...")
	dindex += 16
	p.Body.Data = make([]byte, p.Header.Length)
	jklog.L().Debugf("dindex: %d, len data: %d, data: %d\n", dindex, cap(p.Body.Data), len(data[dindex:]))

	lendata := len(data)
	if lendata > int(p.Header.Length)+dindex {
		lendata = int(p.Header.Length) + dindex
	}
	copy(p.Body.Data[:], data[dindex:lendata])

	jklog.L().Debugln("start check valid ...")
	ret := p.checkValid()
	if !ret {
		return nil, errors.New("Error format")
	}

	jklog.L().Debugln("start compare sign ...")
	nsign := p.generateSign()
	if bytes.Compare(nsign[:], p.Header.Sign[:]) != 0 {
		return nil, errors.New("Sign check invalid.")
	}

	return p, nil
}
