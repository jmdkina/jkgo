package jklog


// For now unuse
// log4go do better

import (
	l4g "github.com/alecthomas/log4go"
)

type LogWrap struct {
	Log *l4g.FileLogWriter
}

var Lw LogWrap

func LW() *LogWrap {
	return &Lw
}

func (lw *LogWrap) Init(filename string) error {
	lw.Log = l4g.NewFileLogWriter(filename, true)
	return nil
}