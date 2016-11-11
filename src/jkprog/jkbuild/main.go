package main

import (
	"jk/jklog"
	"os"
	"io/ioutil"
	"strings"
	//"time"
)

// This project for auto build with your project with curses.
// build only support cmake for now.

func main() {
	os.Remove("/tmp/jkbuild.log")
	jklog.SetLogFileName("/tmp/jkbuild.log")

	pd, err := NewTerminalDisplay()
	if err != nil {
		jklog.L().Errorln("failed to init terminal display: ", err)
		return
	}

	jklog.Lfile().Debugln("Start to build.")

	conf, _ := NewGlobalConfig("./Config.jk")
	pd.Conf = conf
	jklog.Lfile().Debugln("Source dir is : ", conf.baseDir)

	ba, _ := NewBuildArgs()
	pd.BArgs = ba

	jklog.Lfile().Debugln("Choose which source to compile")
	pd.TipsWhichToCompile()

	pd.WaitNextForEnter()

	jklog.Lfile().Debugln("Choose compile args")
	pd.Clear()
	ret := pd.TipsCompileArgs()
	if ret != nil {
		pd.Terminate()
		jklog.L().Errorln("Invalid args ", ret)
		return
	}

	pd.Clear()
	jklog.Lfile().Debugln("Sure build args")
	retv := pd.DisplayBuildArgs()
	if !retv {
		pd.Terminate()
		jklog.Lfile().Warnln("User manually exit")
		return
	}

	jklog.Lfile().Debugln("Start to build ...")

	pd.Clear()
	pd.DisplayData("Start to build ...")

	//time.Sleep(time.Millisecond*500)

	// Start to build
	// We didn't transfer pointer of BArgs
	build, _ := NewBuildInfo(conf, *pd.BArgs)
	build.SetBuildItem(conf.item)
	outdata, err := build.Build()
	if err != nil {
		pd.DisplayData(err.Error())
	}
	pd.DisplayData(outdata)
	pd.DisplayData("\n")
	data, err := ioutil.ReadFile("/tmp/jkbuild.log")
	if err == nil {
		str := strings.Split(string(data), "\n")
		for k, v := range str {
			if k > len(str) - 10 {
				pd.DisplayData(v)
			}
		}
	}
	pd.DisplayData("\n")

	pd.DisplayData("Build done! Please any key to exit")
	pd.WaitForAnyKey()

	pd.Terminate()

	jklog.Lfile().Debugln("Program exit")
}