package jkitoy

import (
	"errors"
	"fmt"
	"io"
	"jk/jklog"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type ItoY struct {
	SaveFilePrefix string // will add it to the new filename as prefix.
	ExtName        string // Will replace old extern name
	SaveDirPrefix  string // Prefix of the saved filename, only use when give dir
}

// init it

func NewItoY(savefileprefix, extname, savedirprefix string) *ItoY {
	var base ItoY
	if len(savefileprefix) == 0 {
		base.SaveFilePrefix = "unknown"
	} else {
		base.SaveFilePrefix = savefileprefix
	}
	if len(extname) == 0 {
		base.ExtName = "unknown"
	} else {
		base.ExtName = extname
	}
	if len(savedirprefix) == 0 {
		base.SaveDirPrefix = "unknown"
	} else {
		base.SaveDirPrefix = savedirprefix
	}

	return &base
}

// @param filename Absolute name include the path
func (iy *ItoY) OpenAndSaveToNew(filename string, savelocation string) error {
	if filename[0] == '.' {
		return nil
	}
	// Open file and prepare to close
	fin, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer fin.Close()

	var writefilename = savelocation
	_, err = os.Stat(writefilename)
	var fout *os.File
	// Append to the file if file has exist, create it otherwise
	if os.IsNotExist(err) {
		fout, err = os.Create(writefilename)
	} else {
		fout, err = os.OpenFile(writefilename, os.O_APPEND|os.O_RDWR, os.ModePerm)
	}
	if err != nil {
		jklog.L().Errorln("open error ", err)
		return err
	}
	defer fout.Close()

	err = iy.readAndWrite(fin, fout)
	if err != nil {
		return err
	}

	return nil
}

// read data and write out
func (iy *ItoY) readAndWrite(fin, fout *os.File) error {
	if fin == nil || fout == nil {
		return nil
	}

	f, _ := fin.Stat()
	lendata := f.Size()
	filename := fin.Name()
	// remove first / and last .
	i := strings.LastIndex(filename, "/")
	if i >= 0 {
		filename = filename[i+1 : len(filename)]
	} else {
		// windows
		i = strings.LastIndex(filename, "\\")
		filename = filename[i+1 : len(filename)]
	}
	i = strings.LastIndex(filename, ".")
	if i >= 0 {
		filename = filename[0:i]
	}

	if len(filename) == 0 {
		return nil
	}
	if filename[0] == '.' {
		return nil
	}

	writestring := make([]byte, 128)
	genstring := fmt.Sprintf("vs%dsvend%svvssend", lendata, filename)
	writestring = []byte(genstring)

	_, err := fout.Write([]byte(writestring))
	if err != nil {
		jklog.L().Errorln("write error ", err)
		return err
	}

	buf := make([]byte, 2<<10)
	for {
		n, err := fin.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		_, err = fout.Write(buf[0:n])
		if err != nil {
			return err
		}
	}

	return nil
}

func parseOutFromHeader(header string) (int64, string, int) {
	var filename string
	var lenstring string
	i := strings.Index(header, "svend")
	if i >= 0 {
		lenstring = header[2:i]
	} else {
		return 0, "", 0
	}
	lendata, _ := strconv.ParseInt(lenstring, 10, 64)

	j := strings.Index(header, "vvssend")
	if j >= 0 {
		filename = header[i+5 : j]
	} else {
		return 0, "", 0
	}
	// jklog.L().Infoln("len ", lendata, ", name ", filename)
	return lendata, filename, j + 7
}

func (iy *ItoY) readAndWriteRevers(fin *os.File, savepos string) error {
	if fin == nil {
		return errors.New("NO file handle")
	}

	for {
		bufheader := make([]byte, 64)

		var fout *os.File

		// read header
		n, err := fin.Read(bufheader)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// jklog.L().Infoln("get header ", string(bufheader), ", len ", n)
		lendata, filename, lenheader := parseOutFromHeader(string(bufheader))

		// jklog.L().Infoln("lendata ", lendata, ", lenheader ", lenheader)
		if lendata > 0 {
			// Create file
			fout, err = os.Create(savepos + "/" + filename + ".jpeg")
			if err != nil {
				jklog.L().Errorln("open file error :", err)
				return err
			}
			fout.Write(bufheader[lenheader:n])
		} else {
			break
		}

		lentoread := lendata - ((int64)(n) - (int64)(lenheader))
		// read data
		buf := make([]byte, lentoread)
		n, err = fin.Read(buf)

		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		// jklog.L().Infoln("now read data ", lentoread, ", really read ", n)
		fout.Write(buf)
		fout.Close()
	}

	return nil
}

// The we convert revert
func (iy *ItoY) PrepareToReadWrite(oldfile, saveposition string) error {
	// Open File to Read
	fin, err := os.OpenFile(oldfile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer fin.Close()

	// Open File to Write
	var realfilename string

	i := strings.LastIndex(oldfile, "/")
	if i >= 0 {
		realfilename = oldfile[i+1 : len(oldfile)]
	} else {
		realfilename = oldfile
	}
	if realfilename[0] == '.' {
		return nil
	}

	err = iy.readAndWriteRevers(fin, saveposition)
	return err
}

func (iy *ItoY) scanDirToSave(dirname, saveposition string) error {
	jklog.L().Infoln("scan dir ", dirname)
	filepath.Walk(dirname, func(path string, fi os.FileInfo, err error) error {

		if fi.IsDir() {
			return nil
			err := iy.scanDirToSave(dirname, path)
			if err != nil {
				return err
			}
		} else {
			// jklog.L().Infof("file %s ==> %s\n", path, saveposition)
			err := iy.OpenAndSaveToNew(path, saveposition)
			if err != nil {
				jklog.L().Errorln("scan dir err of ", path, " with err ", err)
				return err
			}
		}
		return nil
	})
	return nil
}

func (iy *ItoY) scanDirToSaveRevert(dirname, saveposition string) error {
	jklog.L().Infoln("scan dir ", dirname)
	filepath.Walk(dirname, func(path string, fi os.FileInfo, err error) error {

		if fi.IsDir() {
			return nil
			err := iy.scanDirToSave(dirname, path)
			if err != nil {
				return err
			}
		} else {
			// jklog.L().Infof("file %s ==> %s\n", path, saveposition)
			err := iy.PrepareToReadWrite(path, saveposition)
			if err != nil {
				jklog.L().Errorln("scan dir err of ", path, " with err ", err)
				return err
			}
		}
		return nil
	})
	return nil
}

// Next we give some func to do with dir
// It will read from @dirname and file the files and do it then save to saveposition
func (iy *ItoY) NewSaveWithDir(dirname, saveposition string) error {
	fi, err := os.Stat(dirname)
	if err != nil {
		return err
	}

	// /xxx/xxx/file
	// need get /xxx/xxx directory, and create it if not exist
	dirposindex := strings.LastIndex(saveposition, "/")
	dirpos := saveposition
	if dirposindex > 0 {
		dirpos = saveposition[0:dirposindex]
	}

	_, err = os.Stat(dirpos)
	if err != nil {
		// return err
	}
	if os.IsExist(err) == false {
		// Create it

		err := os.MkdirAll(dirpos, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// fi can be directory and only one file
	if !fi.IsDir() {
		err = iy.OpenAndSaveToNew(dirname, saveposition+"/"+dirname)
		if err != nil {
			return err
		}
	} else {
		err := iy.scanDirToSave(dirname, saveposition+"/"+dirname)
		if err != nil {
			return err
		}
	}

	return nil
}

func (iy *ItoY) NewSaveWithDirRevert(dirname, saveposition string) error {
	fi, err := os.Stat(dirname)
	if err != nil {
		return err
	}

	_, err = os.Stat(saveposition)
	if err != nil {
		// return err
	}
	if os.IsExist(err) == false {
		// Create it

		err := os.MkdirAll(saveposition, os.ModeDir)
		if err != nil {
			return err
		}
	}

	if !fi.IsDir() {
		err = iy.PrepareToReadWrite(dirname, saveposition)
		if err != nil {
			return err
		}
	} else {
		err := iy.scanDirToSaveRevert(dirname, saveposition)
		if err != nil {
			return err
		}
	}

	return nil
}
