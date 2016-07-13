package jkprotocol

import (
	"errors"
	cm "jk/jkcommon"
	"jk/jklog"
)

type JKProtoV4Header struct {
	Version uint
	Crypt   uint
	ACK     bool
	Length  uint32
}

type JKProtoV4Body struct {
	Data []byte
}

type JKProtoV4 struct {
	Header JKProtoV4Header
	Body   JKProtoV4Body
}

func JKProtoV4New() (*JKProtoV4, error) {
	return &JKProtoV4{}, nil
}

const (
	jk_proto_v4_version       = 1
	jk_proto_v4_header_length = 8 // bytes
)

const (
	JKProtoV4CryptNone = 0
	JKProtoV4CryptAES  = 1
)

// Set header
func (p *JKProtoV4) GenerateHeader(crypt uint, ack bool, length uint32) (*JKProtoV4Header, error) {
	p.Header.Version = jk_proto_v4_version
	p.Header.ACK = ack
	p.Header.Crypt = crypt
	p.Header.Length = length
	return &p.Header, nil
}

// Set Body
func (p *JKProtoV4) GenerateBody(data string) (*JKProtoV4Body, error) {
	p.Header.Length = uint32(len(data))
	p.Body.Data = make([]byte, p.Header.Length)
	p.Body.Data = []byte(data)
	return &p.Body, nil
}

func (p *JKProtoV4) GenerateBodyByte(data []byte) (*JKProtoV4Body, error) {
	p.Header.Length = uint32(len(data))
	p.Body.Data = make([]byte, p.Header.Length)
	p.Body.Data = data
	return &p.Body, nil
}

// Generate to Bytes
func (p *JKProtoV4) ToByte() ([]byte, error) {
	data := make([]byte, jk_proto_v4_header_length+p.Header.Length)
	v := p.Header.Version
	if v <= 0 || v > 1<<4-1 {
		return nil, errors.New("version error")
	}

	// jklog.L().Debugf("version: %d\n", v)
	switch v {
	case 1:
		data[0] |= (1 << 4)
	case 2:
		data[0] |= (1 << 5)
	}

	crypt := p.Header.Crypt
	if crypt < 0 || crypt > 1<<2-1 {
		return nil, errors.New("Crypt error")
	}
	// jklog.L().Debugf("crypt: %d\n", crypt)

	switch crypt {
	case 0:
		data[0] |= (1 << 2)
	case 1:
		data[0] |= (1 << 3)
	}

	ack := p.Header.ACK
	if ack {
		data[0] |= (1 << 1)
	}

	length := p.Header.Length
	if length < 0 || length > 1<<32-1 {
		return nil, errors.New("length error")
	}
	copy(data[4:8], cm.Int32ToBytes(int32(length)))

	copy(data[8:], p.Body.Data[:])

	// jklog.L().Debugf("%v \n", data)

	return data, nil
}

func (p *JKProtoV4) Debug() {
	jklog.L().Debugf("version: %d, crypt: %d, ack: %v, length: %d\n", p.Header.Version,
		p.Header.Crypt, p.Header.ACK, p.Header.Length)
	jklog.L().Debugf("data: %s\n", string(p.Body.Data))
}

func (p *JKProtoV4) Parse(data []byte) error {
	if len(data) < 8 {
		return errors.New("parse data header length less.")
	}

	v := cm.JKBitValueByte(data[0], 8, 4)
	// jklog.L().Debugf("version: %d\n", v)
	if v <= 0 || v >= 1<<4 {
		return errors.New("parse version error")
	}

	p.Header.Version = uint(v)

	crypt := cm.JKBitValueByte(data[0], 4, 2)
	// jklog.L().Debugf("crypte value: %d , %v\n", crypt, data[0])
	if crypt < 0 || crypt >= 1<<2 {
		return errors.New("parse crypt error")
	}
	p.Header.Crypt = uint(crypt)

	ack := data[0] & (1 << 1)
	if ack != 0 {
		p.Header.ACK = true
	}

	// jklog.L().Debugf("data: %v\n", data)

	dataLen := make([]byte, 4)
	copy(dataLen[:], data[4:8])
	length := cm.BytesToUInt32(dataLen)
	// jklog.L().Debugf("length: %d\n", length)
	if length < 0 || length > (1<<32-1) {
		return errors.New("parse length too long")
	}
	p.Header.Length = uint32(length)

	// Data
	p.Body.Data = make([]byte, p.Header.Length)
	copy(p.Body.Data[:], data[8:p.Header.Length+8])

	// p.Debug()

	return nil
}
