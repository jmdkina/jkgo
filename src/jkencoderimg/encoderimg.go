package main

import (
	"flag"
	"jk/jklog"
	img "jkencoderimg/en"
	"strconv"
)

var (
	scanpath = flag.String("path", ".", "dir to scan")
	quality  = flag.Int("quality", 50, "compress quality")
	scale    = flag.Int("scale", 100, "scale size")
	savepos  = flag.String("savepos", "save", "where to save")
)

func main() {
	flag.Parse()

	jklog.L().Infoln("Scan dir: ", *scanpath)
	savepath := "q" + strconv.Itoa(*quality)
	if *savepos != "save" {
		savepath = *savepos
	}
	en := img.JKEncoderImage{
		ScanPath: *scanpath,
		SavePos:  savepath,
		Quality:  *quality,
		Scale:    *scale,
	}
	en.JK_convertWithPath()
}
