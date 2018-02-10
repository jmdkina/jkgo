package main

import (
	"flag"
	"github.com/alecthomas/log4go"
	"github.com/h2non/filetype"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	dir       = flag.String("dir", ".", "where start search")
	max_level = flag.Int("level", 2, "max dir to search")
)

type FileItem struct {
	Path     string
	Type     string
	TypeFull string
}

var FileLists []FileItem

func debug_out() {
	for _, v := range FileLists {
		log4go.Info("path [%s] type [%s] typefull [%s]",
			v.Path, v.Type, v.TypeFull)
	}
}

var global_level = 0

func do_check(path string) {
	if global_level >= *max_level {
		return
	}
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			global_level = global_level + 1
			do_check(path)
			global_level = global_level - 1
		} else {
			buf, _ := ioutil.ReadFile(path)
			kind, unknow := filetype.Match(buf)
			if unknow != nil {
				log4go.Error("check [%s] unknown [%v]", path, unknow)
				return unknow
			}
			fi := FileItem{
				Path:     path,
				Type:     kind.Extension,
				TypeFull: kind.MIME.Value,
			}
			log4go.Debug("Find one of path [%s]", path)
			FileLists = append(FileLists, fi)
		}
		return nil
	})
}

func main() {
	flag.Parse()
	log4go.AddFilter("stdout", log4go.INFO, log4go.NewConsoleLogWriter())

	do_check(*dir)

	debug_out()
}
