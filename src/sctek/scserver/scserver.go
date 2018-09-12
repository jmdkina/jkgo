package main

import (
	"jk/jklog"
	"sctek"
	"time"
)

func main() {
	sd, err := sctek.NewSctekDiscover()
	if err != nil {
		jklog.L().Errorln(err)
		return
	}
	sd.Discover(10)
	for {
		time.Sleep(time.Millisecond * 500)
	}
}
