package jknetbase

import (
	l4g "github.com/alecthomas/log4go"
	jdaemon "github.com/tyranron/daemonigo"
)

func InitLog(logfile string, logsize int) {
	lw := l4g.NewFileLogWriter(logfile, false)
	if lw != nil {
		lw.SetRotateSize(logsize)
		l4g.AddFilter("file", l4g.FINE, lw)
	}
}

func InitDeamon(backrun bool) {
	if backrun {
		// Daemonizing echo server application.
		switch isDaemon, err := jdaemon.Daemonize("start"); {
		case !isDaemon:
			return
		case err != nil:
			l4g.Error("daemon start failed : ", err.Error())
		}
	}
}