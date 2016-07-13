// Pacakge jkprotocol implement jk protocol generate, parse
// and use.
package jkprotocol

import (
	"errors"
	// "jk/jklog"
	"strconv"
	"strings"
)

/**
 * @author: jmdkina@gmail.com
 * @create: 20140914
 *
 * The protocol define
 *
 * |-----------------------------8bytes----------------------------|
 * |JK|version-m(3b)|version-s(4b)|/s/r(1b)|| ? | seq(4byte)      |
 * |     main-command(4byte)      |   slave-command(4byte)        |
 * | length(data)                 |
 * |--------------------------------------------------------------|
 * /s/r(1b)||? (9bits)
 * 2bits: 00(send), 01(response), 10(notify)
 * 1bit: 0 the data over, 1 the data continue(you need save the first data to reconnect)
 *
 * The protocol (string) (return is \r\n)
 * JK
 * (version) main.slave.last
 * (command type) SEND/RESPONSE/NOTIFY
 * (if continued) 0/1
 * (sequence) 28
 * (main command) JK_xxx
 * (slave command) JK_xxx
 * (length) 228
 * (content)
 * xxxxx
 *
 * WARN: This is beta protocol, when you use it, please remember maybe someday
 * it will be changed
 */

// Header const len define
const (
	JK_P_HEADER_LENGTH = 20 // The header length is 20 bytes
)

// Header const define
const (
	JK_P_HEADER_MARK_NAME     = "JK"
	JK_P_HEADER_MAIN_VERSION  = 0
	JK_P_HEADER_SLAVE_VERSION = 1
	JK_P_HEADER_LAST_VERSION  = 0
)

// Header main command type
const (
	JK_P_HEADER_TYPE_SEND     = 0
	JK_P_HEADER_TYPE_RESPONSE = 1
	JK_P_HEADER_TYPE_NOTIFY   = 2
)

const (
	JK_P_HEADER_SINGLE   = 0
	JK_P_HEADER_CONTINUE = 1
)

// Header main command
const (
	JK_P_CMD_DEVICE_CONNECT = 1
)

// Header slave command depends on main command
// Device Connect
const (
	JK_P_CMD_DEVICE_CONNECT_QUERY         = 1
	JK_P_CMD_DEVICE_CONNECT_FILE_TRANSFER = 2
)

type JK_P_Header struct {
	markName     string
	mainVersion  int
	slaveVersion int
	lastVersion  int
	commandType  int
	continued    int
	sequence     int
	mainCommand  int
	slaveCommand int
	length       int
}

type JK_P_DataWithHeader struct {
	header JK_P_Header
	data   []byte
}

// Inner sequence
// increase self.
// Use it if user didn't defined one
var currentSeq = 1 // Init from 1

// GenerateProtocolHeader with args
func JKGenerateProtocolHeader(cmdType, cont, seq int, mCmd, sCmd int, length int) *JK_P_Header {
	header := JK_P_Header{
		markName:     JK_P_HEADER_MARK_NAME,
		mainVersion:  JK_P_HEADER_MAIN_VERSION,
		slaveVersion: JK_P_HEADER_SLAVE_VERSION,
		lastVersion:  JK_P_HEADER_LAST_VERSION,
		commandType:  cmdType,
		continued:    cont,
		mainCommand:  mCmd,
		slaveCommand: sCmd,
		length:       length,
	}
	// Use inner seq if user didn't defined one
	if seq == -1 {
		header.sequence = currentSeq + 1
		currentSeq += 1
	} else {
		header.sequence = seq
	}
	return &header
}

// default version to use
func JKGenerateHeaderSimple(seq, mCmd, sCmd int) *JK_P_Header {
	header := JKGenerateProtocolHeader(JK_P_HEADER_TYPE_SEND, JK_P_HEADER_SINGLE, seq, mCmd, sCmd, 0)
	return header
}

func JKGenerateHeaderSimpleResponse(seq, mCmd, sCmd int) *JK_P_Header {
	header := JKGenerateProtocolHeader(JK_P_HEADER_TYPE_RESPONSE, JK_P_HEADER_SINGLE, seq, mCmd, sCmd, 0)
	return header
}

func JKGenearteHeaderSimpleNotify(seq, mCmd, sCmd int) *JK_P_Header {
	header := JKGenerateProtocolHeader(JK_P_HEADER_TYPE_NOTIFY, JK_P_HEADER_SINGLE, seq, mCmd, sCmd, 0)
	return header
}

// Return a new pointer to the header with data and header
//
func (h *JK_P_Header) SetData(data []byte) *JK_P_DataWithHeader {
	h.SetLength(len(data))
	dataWithHeader := JK_P_DataWithHeader{
		header: *h,
		data:   data,
	}
	return &dataWithHeader
}

// Change the data
func (h *JK_P_DataWithHeader) SetData(data []byte) {
	h.data = data
}

// Get the data
func (h *JK_P_DataWithHeader) GetData() []byte {
	return h.data
}

func (h *JK_P_DataWithHeader) ToByte() []byte {
	buf := string(h.header.ToBytes())
	buf += string(h.data)
	return []byte(buf)
}

// Convert between header struct and []bytes
func (h *JK_P_Header) ToBytes() []byte {
	var buf string

	buf = h.markName + "\r\n"
	buf += strconv.Itoa(h.mainVersion) + "." + strconv.Itoa(h.slaveVersion) + "." + strconv.Itoa(h.lastVersion) + "\r\n"
	buf += "Type:" + strconv.Itoa(h.commandType) + "\r\n"
	buf += "Continued:" + strconv.Itoa(h.continued) + "\r\n"
	buf += "MainCommand:" + strconv.Itoa(h.mainCommand) + "\r\n"
	buf += "SlaveCommand:" + strconv.Itoa(h.slaveCommand) + "\r\n"
	buf += "Sequence:" + strconv.Itoa(h.sequence) + "\r\n"
	buf += "Length:" + strconv.Itoa(h.length) + "\r\n"

	return []byte(buf)
}

// Parse buf out to DataWithHeader
func JKFromBytes(buf []byte) (*JK_P_DataWithHeader, error) {
	dataWithHeader := &JK_P_DataWithHeader{
		header: JK_P_Header{},
	}
	err := dataWithHeader.FromBytes(buf)
	if err != nil {
		return nil, err
	}
	return dataWithHeader, nil
}

// Get header from dataWithHeader
func (header *JK_P_DataWithHeader) Header() *JK_P_Header {
	return &header.header
}

// Get data from dataWithHeader
func (header *JK_P_DataWithHeader) Data() []byte {
	return header.data
}

func (header *JK_P_DataWithHeader) FromBytes(buf []byte) error {
	h := &header.header
	each := strings.Split(string(buf), "\r\n")
	if len(each) < 9 {
		return errors.New("Command less then 8")
	}

	// parse each one
	// 1. mark name
	if each[0] != JK_P_HEADER_MARK_NAME {
		return errors.New("Error mark")
	}

	// 2. version
	vers := strings.Split(each[1], ".")
	if len(vers) < 3 {
		return errors.New("Error version")
	}

	// 3. type
	ts := strings.Split(each[2], ":")
	if len(ts) < 2 || ts[0] != "Type" {
		return errors.New("Error types")
	}

	// 4. Continued mark
	cont := strings.Split(each[3], ":")
	if len(cont) < 2 || cont[0] != "Continued" {
		return errors.New("Error Continued")
	}

	// 5. MainCommand
	mc := strings.Split(each[4], ":")
	if len(mc) < 2 || mc[0] != "MainCommand" {
		return errors.New("Error Main Command")
	}

	// 6. slave command
	sc := strings.Split(each[5], ":")
	if len(sc) < 2 || sc[0] != "SlaveCommand" {
		return errors.New("Error Slave Command")
	}

	// 7. Sequence
	seq := strings.Split(each[6], ":")
	if len(seq) < 2 || seq[0] != "Sequence" {
		return errors.New("Error Sequence")
	}

	// 8. Length
	length := strings.Split(each[7], ":")
	if len(length) < 2 || length[0] != "Length" {
		return errors.New("Error length")
	}

	// Save them
	h.markName = each[0]
	h.mainVersion, _ = strconv.Atoi(vers[0])
	h.slaveVersion, _ = strconv.Atoi(vers[1])
	h.lastVersion, _ = strconv.Atoi(vers[2])
	h.commandType, _ = strconv.Atoi(ts[1])
	h.continued, _ = strconv.Atoi(cont[1])
	h.mainCommand, _ = strconv.Atoi(mc[1])
	h.slaveCommand, _ = strconv.Atoi(sc[1])
	h.sequence, _ = strconv.Atoi(seq[1])
	h.length, _ = strconv.Atoi(length[1])

	// 9. data
	if len(each) >= 9 {
		header.data = []byte(each[8])
	}

	return nil
}

func (h *JK_P_Header) GetVersions() string {
	return strconv.Itoa(h.mainVersion) + "." + strconv.Itoa(h.slaveVersion) + "." + strconv.Itoa(h.lastVersion)
}

func (h *JK_P_Header) SetVersions(m, s, l int) error {
	h.mainVersion = m
	h.slaveVersion = s
	h.lastVersion = l
	return nil
}

func (h *JK_P_Header) GetType() int {
	return h.commandType
}

func (h *JK_P_Header) SetType(t int) error {
	h.commandType = t
	return nil
}

func (h *JK_P_Header) GetContinued() int {
	return h.continued
}

func (h *JK_P_Header) SetContinued(c int) error {
	h.continued = c
	return nil
}

func (h *JK_P_Header) GetMainCommand() int {
	return h.mainCommand
}

func (h *JK_P_Header) SetMainCommand(mc int) error {
	h.mainCommand = mc
	return nil
}

func (h *JK_P_Header) GetSlaveCommand() int {
	return h.slaveCommand
}

func (h *JK_P_Header) SetSlaveCommand(sc int) error {
	h.slaveCommand = sc
	return nil
}

func (h *JK_P_Header) GetSequence() int {
	return h.sequence
}

func (h *JK_P_Header) SetSequence(seq int) error {
	h.sequence = seq
	return nil
}

func (h *JK_P_Header) GetLength() int {
	return h.length
}

func (h *JK_P_Header) SetLength(length int) error {
	h.length = length
	return nil
}
