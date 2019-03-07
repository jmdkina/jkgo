package main

import (
	"flag"
	//"io"
	"jk/jklog"
	. "jktools/jkitoy/jkitoy"
)

var (
	filename     = flag.String("f", "", "filename to do")
	saveposition = flag.String("s", "", "Where to save")
	savefilename = flag.String("n", "", "Filename when save")
	revert       = flag.Bool("revert", false, "What style you use")
)

func itoy_do_base() {
	itoy, err := NewItoYItem(*filename)
	if err != nil {
		jklog.L().Errorf("Create error %v\n", err)
		return
	}
	itoy.SetSavePosition(*saveposition, *savefilename)
	err = itoy.DoSomething()
	if err != nil {
		jklog.L().Errorf("do something %v\n", err)
	}
}

func itoy_do_travers() {
	itoy, err := NewItoYItem(*filename)
	if err != nil {
		jklog.L().Errorf("travers create %v\n", err)
		return
	}
	itoy.SetTraversPosition(*saveposition, *savefilename)
	err = itoy.Travers()
	if err != nil {
		jklog.L().Errorf("travers do %v\n", err)
	}
}

func main() {
	flag.Parse()
	if len(*filename) == 0 || len(*saveposition) == 0 || len(*savefilename) == 0 {
		jklog.L().Errorln("Arges need")
		return
	}

	if *revert {
		itoy_do_travers()
	} else {
		itoy_do_base()
	}
	return

	/*
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
	*/
}
