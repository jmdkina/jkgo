package main

import (
	"jk/jklog"
	"sctek"
)

func main() {
	sd, err := sctek.NewSctekDiscover()
	if err != nil {
		jklog.L().Errorln(err)
		return
	}
	sd.Discover(10)
}
