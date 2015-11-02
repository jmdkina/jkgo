// Package jkcommon defines some common varients.
package jkcommon

import (
	"bytes"
	"encoding/binary"
	"io"
	"jk/jklog"
	"os"
	// "strconv"
	"path/filepath"
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

	lendata := 0
	data := make([]byte, 2<<20)
	for {
		tdata := make([]byte, 2<<10)
		n, err := f.Read(tdata)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		if n <= 0 {
			continue
		}
		// jklog.L().Debugln("read out data of len : ", n)

		// jklog.L().Debugln("n: ", n, ", data: ", len(data), ", cap: ", cap(data))
		if lendata+n > cap(data) {
			// jklog.L().Debugln("need more length: ", lendata+n)
			newData := make([]byte, lendata*2)
			copy(newData, data)
			data = newData
		}
		copy(data[lendata:], tdata[:n])
		lendata += n
	}

	return string(data[0:lendata]), nil
}

func JKSaveDataToFile(filename string, data []byte, clear bool) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	if clear {
		f.Truncate(0)
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
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
		jklog.L().Errorln("failed parse: ", err)
		return -1
	}
	return n
}

func Int32ToBytes(val int32) []byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, uint32(val))
	return buf
}

func BytesToInt32(buf []byte) int32 {
	n := int32(binary.LittleEndian.Uint32(buf))
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

// Give a filepath return all the directory or files
// @dir: if get directory
// @file: if get files.
func JKFileLists(path string, dir, file bool) ([]string, error) {
	var flists []string

	err := filepath.Walk(path, func(inpath string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if strings.Compare(inpath, path) == 0 {
			return nil
		}
		if !dir && f.IsDir() {
			return nil
		}
		if strings.HasPrefix(f.Name(), ".") {
			return nil
		}
		if !file && !f.IsDir() {
			return nil
		}
		flists = append(flists, f.Name())
		return nil
	})
	if err != nil {
		return nil, err
	}
	return flists, nil
}
