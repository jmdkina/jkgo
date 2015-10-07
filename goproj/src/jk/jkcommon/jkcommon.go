// Package jkcommon defines some common varients.
package jkcommon

import (
	"bytes"
	"encoding/binary"
	"io"
	"jk/jklog"
	"os"
	"strings"
)

const (
	JK_RESULT_SUCCESS = 0
	JK_RESULT_E_FAIL  = -100 << iota
	JK_RESULT_E_PARAM_ERROR
	JK_RESULT_E_PARSE_ERROR
	JK_RESULT_E_DATABASE_QUERY_ERROR
	JK_RESULT_E_DATABASE_INSERT_ERROR
	JK_RESULT_E_DATABASE_MOD_ERROR
	JK_RESULT_E_DATABASE_REMOVE_ERROR
	JK_RESULT_E_NOT_EXIST
	JK_RESULT_E_HAS_EXIST
	JK_RESULT_E_DATA_NOT_EXIST
	JK_RESULT_E_CODE_ERROR
	JK_RESULT_E_NET_DIAL_ERROR
	JK_RESULT_E_TIME_FAST
	JK_RESULT_E_NOTSUPPORT
	JK_RESULT_E_NO_PERMISSION
)

const (
	JK_NET_ADDRESS_LOCAL = "0.0.0.0"
	JK_NET_ADDRESS_PORT  = 23888

	JK_SERVER_FILE_POSITION = "kflogs"
)

type ResultStatus struct {
	RS map[string]interface{}
}

func NewResultStatus(v int, desc string) *ResultStatus {
	rs := map[string]interface{}{
		"status": v,
	}
	if len(desc) > 0 {
		rs["desc"] = desc
	}
	rsr := ResultStatus{
		RS: rs,
	}
	return &rsr
}

func (rs *ResultStatus) setStatus(v int, desc string) {
	rs.RS["status"] = v
	rs.RS["desc"] = desc
}

func (rs *ResultStatus) SetCustom(v int, desc string) {
	rs.setStatus(v, desc)
}

func (rs *ResultStatus) SetNoPermission() {
	rs.setStatus(JK_RESULT_E_NO_PERMISSION, "NoPermission")
}

func (rs *ResultStatus) SetInsertFail() {
	rs.setStatus(JK_RESULT_E_DATABASE_INSERT_ERROR, "InsertFail")
}

func (rs *ResultStatus) SetItemExist() {
	rs.setStatus(JK_RESULT_E_HAS_EXIST, "ItemExist")
}

func (rs *ResultStatus) SetItemModFail() {
	rs.setStatus(JK_RESULT_E_DATABASE_MOD_ERROR, "ItemModFail")
}

func (rs *ResultStatus) SetItemDelFail() {
	rs.setStatus(JK_RESULT_E_DATABASE_REMOVE_ERROR, "ItemDelFail")
}

func (rs *ResultStatus) SetItemNotExist() {
	rs.setStatus(JK_RESULT_E_NOT_EXIST, "ItemNotExist")
}

func JKReadFileData(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	data := make([]byte, 2<<12)
	lendata := 0
	for {
		tdata := make([]byte, 2<<10)
		n, err := f.Read(tdata)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		// jklog.L().Debugln("read out data of len : ", n)

		copy(data[lendata:lendata+n-1], tdata[0:n-1])
		lendata += n
	}

	return string(data[0 : lendata-1]), nil
}

// how many bytes to use @cnts
func IntToBytes(v int64, cnts int) []byte {
	buf := make([]byte, cnts)
	binary.PutVarint(buf, v)
	return buf
}

func BytesToInt(buf []byte) int64 {
	nbuf := bytes.NewBuffer(buf)
	n, err := binary.ReadVarint(nbuf)
	if err != nil {
		return -1
	}
	return n
}

func JKSaveFileData(id, filename, data string) bool {
	filepath := "./" + JK_SERVER_FILE_POSITION + "/" + id + "/" + filename
	prefix := os.Getenv("HOME")
	if len(prefix) > 0 {
		filepath = prefix + "/" + JK_SERVER_FILE_POSITION + "/" + id + "/" + filename
	}
	indx := strings.LastIndex(filepath, "/")
	if indx > 0 {
		dirs := filepath[0:indx]
		err := os.MkdirAll(dirs, os.ModePerm)
		if err != nil {
			jklog.L().Errorln("create dir failed: ", err)
			return false
		}
	}

	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		f, err = os.Create(filepath)
		if err != nil {
			jklog.L().Errorf("Open %s failed %v\n", filepath, err)
			return false
		}
	}

	defer f.Close()

	_, err = f.WriteString(data)
	if err != nil {
		jklog.L().Errorln("Write data failed: ", err)
		return false
	}
	return true
}
