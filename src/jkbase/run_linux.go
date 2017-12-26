// +build linux
package jkbase

import (
	jdaemon "github.com/tyranron/daemonigo"
	"jk/jklog"
)

func InitDeamon(backrun bool) {
	if backrun {
		// Daemonizing echo server application.
		switch isDaemon, err := jdaemon.Daemonize("start"); {
		case !isDaemon:
			return
		case err != nil:
			jklog.L().Errorln("daemon start failed : ", err.Error())
		}
	}
}
