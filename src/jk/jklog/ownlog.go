package jklog

import (
	"helper"
	"os"
	"strings"
	"time"
)

var l = New(os.Stdout, "[JK]", LstdFlags)

// default is stderr for if init failed.
// and also called this printf
var lfile = New(os.Stderr, "[JK]", LstdFlags)

var logfilename string

func L() *Logger {
	return lfile
}

func Lfile() *Logger {
	return lfile
}

func LLogFileName() string {
	return logfilename
}

func InitLog(name string) (*Logger, error) {

	nms := strings.Split(name, ".")
	tmstr := helper.FormUnixTimeToStringDate(time.Now().Unix())
	logname := name
	if len(nms) == 2 {
		logname = nms[0] + "-" + tmstr + "." + nms[1]
	}

	if len(logname) > 0 {
		lfile, err := SetLogFileName(logname)
		if err != nil {
			return nil, err
		}
		logfilename = logname
		return lfile, nil

	}
	return nil, nil
}

func SetLogFileName(name string) (*Logger, error) {
	filefd, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil, err
	}
	// defer filefd.Close()
	lfile = New(filefd, "[JK]", LstdFlags)
	return lfile, nil
}

/*
 * name what you want to set in prefix
 * flag: -1 for defulat
 */
func Lext(name string, flag int) *Logger {
	if flag == -1 {
		l = New(os.Stdout, name, LstdFlags)
	} else {
		l = New(os.Stdout, name, flag)
	}
	return l
}

func Lown(name string, flag int) *Logger {
	if flag == -1 {
		return New(os.Stdout, name, LstdFlags)
	} else {
		return New(os.Stdout, name, flag)
	}
	return nil
}
