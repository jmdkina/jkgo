package main

import (
	"flag"
	l4g "github.com/alecthomas/log4go"
	"jk/jklog"
	"jk/jkservice/jksys"
	"jkbase"
	"time"
)

var (
	address     = flag.String("address", "0.0.0.0", "Listen local address")
	port        = flag.Int("port", 20103, "Listen local port")
	bg          = flag.Bool("bg", false, "true|false")
	logfile     = flag.String("logfile", "/tmp/jksys.log", "log file")
	logsize     = flag.Int("logsize", 1024*1024*1024, "log size")
	client_addr = flag.String("client_address", "0.0.0.0", "dial client address")
	client_port = flag.Int("client_port", 20101, "dial client port")
	cmd         = flag.String("cmd", "run", "run/install/remove service")
)

func start() {

	jkbase.InitLog(*logfile, *logsize)

	l4g.Debug("jksys start")
	l4g.Info("Listen with [%s:%d]", *address, *port)

	// jkbase.InitDeamon(*bg)
	jksys.Start(*address, *port, *client_addr, *client_port, true, true)
	for {
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {

	flag.Parse()

	prog := &jkbase.Program{
		Name:        "jksys",
		DisplayName: "jk system",
		Desc:        "jk system program",
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
