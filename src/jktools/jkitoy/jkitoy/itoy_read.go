package jkitoy

import (
	"os"
	"bytes"
	"errors"
	"io"
	"jk/jklog"
	"fmt"
)

func (itoy *ItoY) check_create_dir(dirname string) error {
	_, err := os.Stat(dirname)
	if os.IsExist(err) == false {
		return os.MkdirAll(dirname, os.ModePerm)
	}
	return nil
}

func (itoy *ItoY) SetTraversPosition(saveposition string, savefilename string) {
	itoy.SaveDir = saveposition
	itoy.SaveFileName = savefilename
}

func (itoy *ItoY) CToGoString(c []byte) string {
	n := -1
	for i, b := range c {
		if b == 0 {
			break
		}
		n = i
	}
	return string(c[:n+1])
}

func (itoy *ItoY) Travers() error {
	fr, err := os.Open(itoy.FileFullName)
	if err != nil {
		return err
	}
	defer fr.Close()

	data := make([]byte, 4)
	n, err := fr.Read(data)
	if err != nil {
		return err
	}
	if bytes.Compare(itoy.Magic, data[0:n]) != 0 {
		return errors.New("Invalid file")
	}

	// Read out ext name
	extname := make([]byte, 10)
	n, err = fr.Read(extname)
	if err != nil {
		return err
	}
	extname_s := itoy.CToGoString(extname)

	dirname := itoy.SaveDir
	err = itoy.check_create_dir(dirname)
	if err != nil {
		return err
	}

	save_file_name := dirname + "/" + itoy.SaveFileName + "." + extname_s
	jklog.L().Debugf("Save position [%s] [%s]\n", save_file_name, extname_s)

	//fw, err := os.OpenFile("/Users/jmdvirus/Documents/item/item/1.jpg", os.O_WRONLY|os.O_CREATE| os.O_TRUNC, os.ModePerm)
	fw, err := os.OpenFile(save_file_name, os.O_WRONLY|os.O_CREATE| os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer fw.Close()

	for {
		data = make([]byte, 4)
		n, err := fr.Read(data)
		if err != nil {
			if io.EOF == err {
				break
			}
			return err
		}
		num := itoy.byteToInt(data[:n])
		data = make([]byte, num)
		nn, err := fr.Read(data)
		if err != nil {
			return err
		}
		if uint32(nn) != num {
			data := fmt.Sprintf("Read data [%d] less then [%d]", nn, num)
			return errors.New(data)
		}
        fw.Write(data[:nn])
	}
	return nil
}