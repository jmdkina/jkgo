package main

import (
	"flag"
	"jk/jklog"
	"jkbase"
	"jkservice/jkstatus/jkstatus"
	"time"

	l4g "github.com/alecthomas/log4go"
)

var (
	address = flag.String("address", "0.0.0.0", "Listen local address")
	port    = flag.Int("port", 20101, "Listen local port")
	bg      = flag.Bool("bg", false, "true|false")
	logfile = flag.String("logfile", "/tmp/jkstatus.log", "log file")
	logsize = flag.Int("logsize", 1024*1024*1024, "log size")
	cmd     = flag.String("cmd", "run", "run/install/remove service")
)

func start() {
	jkbase.InitLog(*logfile, *logsize)

	l4g.Debug("jkstatus start")
	l4g.Info("Listen with [%s:%d]", *address, *port)

	jkbase.InitDeamon(*bg)

	jkstatus.Start(*address, *port, true)
	for {
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	flag.Parse()

	prog := &jkbase.Program{
		Name:        "jkstatus",
		DisplayName: "jk status",
		Desc:        "jk status program",
	}

	prog.Runner = start
	err := prog.CreateService()
	if err != nil {
		jklog.L().Errorln("create service failed ", err)
		return
	}
	err = prog.Ctrl(*cmd)
	if err == nil {
		jklog.L().Infof("Do %s success\n", *cmd)
	}
}
