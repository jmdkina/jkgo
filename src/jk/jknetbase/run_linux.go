// +build linux
package jknetbase

import (
	jdaemon "github.com/tyranron/daemonigo"
)

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
