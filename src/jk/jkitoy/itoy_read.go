package jkitoy

import (
	"os"
	"bytes"
	"errors"
	"io"
	"jk/jklog"
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
	extname_s := string(extname[:n])

	dirname := itoy.SaveDir
	err = itoy.check_create_dir(dirname)
	if err != nil {
		return err
	}

	save_position := dirname + "/" + itoy.SaveFileName + "." + extname_s
	jklog.L().Debugf("Save position %s\n", save_position)

	fw, err := os.OpenFile(save_position, os.O_CREATE | os.O_WRONLY | os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer fw.Close()

	for {
		n, err := fr.Read(data)
		if err != nil {
			if io.EOF == err {
				break
			}
			return err
		}
		data = make([]byte, n)
		nn, err := fr.Read(data)
		if err != nil {
			return err
		}
		if nn != n {
			return errors.New("Read data less then need")
		}
        fw.Write(data[:nn])
	}
	return nil
}