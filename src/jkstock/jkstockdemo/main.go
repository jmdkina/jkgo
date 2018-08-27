package main

import (
	"jk/jklog"
	"flag"
	. "jkstock"
)

var (
	filename = flag.String("filename", "", "filename to read")
)

func main() {
	flag.Parse()
	jklog.L().Debugln("Start stock")

	sp, err := NewStockParse(*filename)
	if err != nil {
		jklog.L().Errorln(err)
		return
	}
	sp.DebugOut()
}