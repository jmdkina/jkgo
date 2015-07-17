package main

import (
	"flag"
	"jk/jklog"
	. "jk/jkparsedoc"
)

var (
	filename = flag.String("filename", ".", "what file/dir to parse")
)

func main() {

	flag.Parse()
	jklog.L().Infoln("jk parse doc start of ", *filename)
	// jklog.L().SetLevel(jklog.LEVEL_IMPORTANT)
	JKParseDocStart(*filename)

}
