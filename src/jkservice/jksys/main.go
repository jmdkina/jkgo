package main

import (
	"flag"
	"jk/jklog"
	"jkbase"
	"jkservice/jksys/jksys"
	"time"

	l4g "github.com/alecthomas/log4go"
)

type sysArgs struct {
	Addr       string
	Port       int
	LogFile    string
	LogSize    int
	ClientAddr string
	ClientPort int
}

var sys_args sysArgs

var (
	cmd  = flag.String("cmd", "run", "run/install/remove service")
	conf = flag.String("conf", "c:/jk/jksys.json", "conf file")
)

func start() {

	jkbase.InitLog(sys_args.LogFile, sys_args.LogSize)

	l4g.Debug("jksys start")
	l4g.Info("Listen with [%s:%d]", sys_args.Addr, sys_args.Port)

	_, _, err := jksys.Start(sys_args.Addr, sys_args.Port, sys_args.ClientAddr, sys_args.ClientPort, true, true)
	if err != nil {
		jklog.L().Errorln("start error ", err)
		return
	}
	for {
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {

	flag.Parse()

	err := jkbase.GetConfigInfo(*conf, &sys_args)
	if err != nil {
		jklog.L().Errorf("get config %s failed %v\n", *conf, err)
		return
	}
	jklog.L().Infoln("conf data ", sys_args)

	prog := &jkbase.Program{
		Name:        "jksys",
		DisplayName: "jk system",
		Desc:        "jk system program",
	}

	prog.Runner = start
	err = prog.CreateService()
	if err != nil {
		jklog.L().Errorln("create service failed ", err)
		return
	}
	err = prog.Ctrl(*cmd)
	if err == nil {
		jklog.L().Infof("Do %s success\n", *cmd)
	}
}
