package jkitoy

import (
	"strings"
	"errors"
	"os"
	"io"
	"encoding/binary"
	"jk/jklog"
)

func NewItoYItem(filename string) (*ItoY, error) {
	itoy := &ItoY{};

	fi, err := os.Stat(filename)
	if err != nil {
		return nil, err
	}

	itoy.FileFullName = filename

	itoy.FileSize = fi.Size()

	splits := strings.Split(filename, ".")
	if len(splits) < 2 {
		return nil, errors.New("Invalid filename")
	}
	itoy.ExtName = splits[1]

	index := strings.LastIndex(splits[0], "/")
	if index <= 0 {
		itoy.FilePath = ""
		itoy.FileName = splits[0]
	} else {
		itoy.FilePath = splits[0][:index]
		itoy.FileName = splits[0][index+1:]
	}
	if len(itoy.FilePath) == 0 {
		itoy.SaveDir = itoy.FilePath
	} else {
		itoy.SaveDir = itoy.FilePath + "/item"
	}
	itoy.SaveFileName = itoy.FileName
	itoy.SaveFileSuffix = "bin"

	itoy.Magic = []byte{ 0xae, 0xbb, 0xee, 0x48 }

	itoy.SegInfo = map[int16]uint32 {
		0: 1<< 10,
		1: 1<< 12,
		2: 1<< 20,
		3: 1<< 20,
		4: 1<< 20,
		5: 1<< 30,
	}

	return itoy, nil
}

func (itoy *ItoY) SetSavePosition(position string, filename string) {
	itoy.SaveDir = position
	itoy.SaveFileName = filename
}

// When converter you will call it
func (itoy *ItoY) setHeader(save_file_name string) error {

	_, err := os.Stat(itoy.SaveDir)
	if os.IsExist(err) == false {
		os.MkdirAll(itoy.SaveDir, os.ModePerm)
	}

	fw, err := os.OpenFile(save_file_name, os.O_CREATE | os.O_WRONLY | os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer fw.Close()
	fw.Write(itoy.Magic)

	extname := []byte(itoy.ExtName)

	if len(extname) > 10 {
		fw.Write(extname[:10])
	} else {
		remain := 10 - len(extname)
		fw.Write(extname)
		remain_data := make([]byte, remain)
		fw.Write(remain_data)
	}

	return nil
}

func (itoy *ItoY) setDataHeader(fi *os.File, size uint32) {
	intbyte := make([]byte, 4)
	binary.LittleEndian.PutUint32(intbyte, size)
	fi.Write(intbyte)
}

func (itoy *ItoY) byteToInt(data []byte) uint32 {
	return binary.LittleEndian.Uint32(data)
}

func (itoy *ItoY) DoSomething() error {
	fhandler, err := os.Open(itoy.FileFullName)
	if err != nil {
		return err
	}
	defer fhandler.Close()

	save_file_name := itoy.SaveDir + "/" + itoy.SaveFileName + ".bin"

	err = itoy.setHeader(save_file_name)
	if err != nil {
		jklog.L().Errorf("set header failed")
		return err
	}

	fw, err := os.OpenFile(save_file_name, os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		jklog.L().Errorf("open file %s error", save_file_name)
		return err
	}
	jklog.L().Debugf("Write to %s\n", save_file_name)
	defer fw.Close()

	var index int16
	index = 0
	for {
		data := make([]byte, itoy.SegInfo[index])
		n, err := fhandler.Read(data)
		if err != nil {
			if io.EOF == err {
				break;
			}
			return err
		}

		// Write header
		itoy.setDataHeader(fw, uint32(n))
		fw.Write(data[:n])
		index = index + 1
	}

	return nil
}