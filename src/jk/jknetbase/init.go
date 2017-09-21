package jknetbase

import (
	l4g "github.com/alecthomas/log4go"
)

func InitLog(logfile string, logsize int) {
	lw := l4g.NewFileLogWriter(logfile, false)
	if lw != nil {
		lw.SetRotateSize(logsize)
		l4g.AddFilter("file", l4g.FINE, lw)
	}
}
