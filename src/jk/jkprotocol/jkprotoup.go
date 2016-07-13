package jkprotocol

import (
	"encoding/json"
)

const (
	JK_PROTOCOL_VERSION_1 = 1 << iota
	JK_PROTOCOL_VERSION_2
	JK_PROTOCOL_VERSION_3
	JK_PROTOCOL_VERSION_4
)

type JKProtoUp struct {
	PType int // proto type
	Proto *JKProtoV4

	Info JKProtoInfo
}

type JKProtoInfo struct {
	Cmd string
	Seq uint32
	Id  string
}

var (
	jk_p_seq = 1
)

func JKProtoUpNew(t int, Id string) (*JKProtoUp, error) {
	p4, _ := JKProtoV4New()
	pUP := JKProtoUp{
		PType: t,
		Proto: p4,
	}
	pUP.Info.Id = Id
	return &pUP, nil
}

// Parse bytes
func (pup *JKProtoUp) JKProtoUpParse(data []byte) error {
	err := pup.Proto.Parse(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(pup.Proto.Body.Data, &pup.Info)
	if err != nil {
		return err
	}

	return nil
}

// Return bytes
// Generate command with protocol v4
func (pup *JKProtoUp) JKProtoUpInit(ack bool, length uint32, content []byte) ([]byte, error) {
	pup.Proto.GenerateHeader(0, ack, length)
	pup.Proto.GenerateBodyByte(content)
	data, err := pup.Proto.ToByte()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Command of register,...
func (pup *JKProtoUp) JKProtoUpCommon(cmd string) ([]byte, error) {
	info := JKProtoInfo{
		Cmd: cmd,
		Seq: uint32(jk_p_seq),
		Id:  pup.Info.Id,
	}

	d, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	pup.JKProtoUpSeqAdd(-1)

	return pup.JKProtoUpInit(false, uint32(len(d)), d)
}

// Auto add sequence, or set sequence
func (pup *JKProtoUp) JKProtoUpSeqAdd(v int32) error {
	if v == -1 {
		jk_p_seq = jk_p_seq + 1
		return nil
	}
	jk_p_seq = int(v)
	if jk_p_seq >= 2<<32 {
		jk_p_seq = 1
	}

	return nil
}

// Get last command's sequence
func (pup *JKProtoUp) JKProtoUpSeq() int {
	seq := jk_p_seq - 1
	return seq
}

func (pup *JKProtoUp) JKProtoUpRegister() ([]byte, error) {
	return pup.JKProtoUpCommon("Register")
}

func (pup *JKProtoUp) JKProtoUpKeepalive() ([]byte, error) {
	return pup.JKProtoUpCommon("Keepalive")
}
