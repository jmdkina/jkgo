package main

import (
	"flag"
	"io"
	"jk/jklog"
)

var (
	filename     = flag.String("filename", "", "what file to do")
	saveposition = flag.String("saveposition", "./", "Where to save")
	revert       = flag.Bool("revert", false, "What style you use")
)

func main() {
	flag.Parse()
	if len(*filename) == 0 {
		jklog.L().Errorln("I must know what file to do with ...")
		return
	}

	if *revert == false {
		todo := NewItoY("", "", "")
		err := todo.NewSaveWithDir(*filename, *saveposition)
		if err != nil && err != io.EOF {
			jklog.L().Errorln("err: ", err)
			return
		}
	} else {

		todo := NewItoY("", "", "")
		err := todo.NewSaveWithDirRevert(*filename, *saveposition)
		if err != nil && err != io.EOF {
			jklog.L().Errorln("error : ", err)
			return
		}

	}

	jklog.L().Infoln("Do it success ...")
	return
}
