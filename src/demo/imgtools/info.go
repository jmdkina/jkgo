package main

import (
	"flag"
	"github.com/rwcarlsen/goexif/exif"
	"jk/jklog"
	"os"
)

func main() {
	img_path := flag.String("image", "", "image to parse")

	flag.Parse()
	f, _ := os.Open(*img_path)
	x, err := exif.Decode(f)
	if err != nil {
		jklog.L().Errorln("error decode ", err)
		return
	}
	jklog.L().Debugln(x)
}
