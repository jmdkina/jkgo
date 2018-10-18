package jklog

import (
	"helper"
	"os"
	"strings"
	"time"
)

type JKLogger struct {
	lstd       *Logger
	lfile      *Logger
	filename   string
	filefd     *os.File
	rotateSize uint64
}

var jkLogger *JKLogger

func init() {
	jkLogger = &JKLogger{
		lstd: New(os.Stdout, "[JK]", LstdFlags|Lmicroseconds),
	}
}

var l = New(os.Stdout, "[JK]", LstdFlags)

// default is stderr for if init failed.
// and also called this printf
var lfile = New(os.Stderr, "[JK]", LstdFlags|Lmicroseconds)

var logfilename string

func L() *JKLogger {
	return jkLogger
}

// func L() *Logger {
// return lfile
// }

func Lfile() *Logger {
	return lfile
}

func LLogFileName() string {
	return logfilename
}

func (l *JKLogger) InitLog(name string) error {
	nms := strings.Split(name, ".")
	tmstr := helper.FormUnixTimeToString(time.Now().Unix())
	logname := name
	if len(nms) == 2 {
		logname = nms[0] + "-" + tmstr + "." + nms[1]
	}

	filefd, err := os.OpenFile(logname, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	l.filename = logname
	l.filefd = filefd
	jkLogger.lfile = New(filefd, "[JK]", LstdFlags)
	return nil
}

func (l *JKLogger) ChangeFile(name string) error {
	if l.filefd != nil {
		l.filefd.Close()
	}
	return l.InitLog(name)
}

func (l *JKLogger) Reopen() error {
	if l.filefd != nil {
		l.filefd.Close()
	}
	return l.InitLog(l.filename)
}

func (l *JKLogger) Debugln(v ...interface{}) {
	if l.lstd != nil {
		l.lstd.Debugln(v...)
	}
	if l.lfile != nil {
		l.lfile.Debugln(v...)
	}
}

func (l *JKLogger) Errorln(v ...interface{}) {
	if l.lstd != nil {
		l.lstd.Errorln(v...)
	}
	if l.lfile != nil {
		l.lfile.Errorln(v...)
	}
}

func (l *JKLogger) Warnln(v ...interface{}) {
	if l.lstd != nil {
		l.lstd.Warnln(v...)
	}
	if l.lfile != nil {
		l.lfile.Warnln(v...)
	}
}

func (l *JKLogger) Infoln(v ...interface{}) {
	if l.lstd != nil {
		l.lstd.Infoln(v...)
	}
	if l.lfile != nil {
		l.lfile.Infoln(v...)
	}
}

func (l *JKLogger) Debug(v ...interface{}) {
	if l.lstd != nil {
		l.lstd.Debug(v...)
	}
	if l.lfile != nil {
		l.lfile.Debug(v...)
	}
}

func (l *JKLogger) Error(v ...interface{}) {
	if l.lstd != nil {
		l.lstd.Error(v...)
	}
	if l.lfile != nil {
		l.lfile.Error(v...)
	}
}

func (l *JKLogger) Warn(v ...interface{}) {
	if l.lstd != nil {
		l.lstd.Warn(v...)
	}
	if l.lfile != nil {
		l.lfile.Warn(v...)
	}
}

func (l *JKLogger) Info(v ...interface{}) {
	if l.lstd != nil {
		l.lstd.Info(v...)
	}
	if l.lfile != nil {
		l.lfile.Info(v...)
	}
}

func (l *JKLogger) Debugf(format string, v ...interface{}) {
	if l.lstd != nil {
		l.lstd.Debugf(format, v...)
	}
	if l.lfile != nil {
		l.lfile.Debugf(format, v...)
	}
}

func (l *JKLogger) Errorf(format string, v ...interface{}) {
	if l.lstd != nil {
		l.lstd.Errorf(format, v...)
	}
	if l.lfile != nil {
		l.lfile.Errorf(format, v...)
	}
}

func (l *JKLogger) Warnf(format string, v ...interface{}) {
	if l.lstd != nil {
		l.lstd.Warnf(format, v...)
	}
	if l.lfile != nil {
		l.lfile.Warnf(format, v...)
	}
}

func (l *JKLogger) Infof(format string, v ...interface{}) {
	if l.lstd != nil {
		l.lstd.Infof(format, v...)
	}
	if l.lfile != nil {
		l.lfile.Infof(format, v...)
	}
}

/*
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
*/

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
