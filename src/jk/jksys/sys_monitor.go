package jksys

import (
	"bytes"
	"jk/jklog"
	"os/exec"
	"strings"
)

func SysProgramRunning(runshell, progName string) bool {
	cmd := exec.Command(runshell, progName)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		jklog.Lfile().Errorln("execute failed: ", err)
		return false
	}
	jklog.Lfile().Debugln("Check " + progName + ", result: " + out.String())
	if out.String()[:1] == "1" {
		return true
	}
	return false
}

func SysRestartProgram(progName, args string) bool {
	cmd := exec.Command(progName, strings.Split(args, " ")...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		jklog.Lfile().Errorln("execute failed: ", err)
		return false
	}
	jklog.Lfile().Debugln("Restart " + progName + ", result: " + out.String())
	return true
}
