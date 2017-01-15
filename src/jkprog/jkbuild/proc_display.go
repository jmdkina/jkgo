package main

import (
	"github.com/rthornton128/goncurses"
	"errors"
	"jk/jklog"
	"fmt"
	"strings"
)

type Pos struct {
	x int
	y int
}

type TerminalDisplay struct {
    stdscr *goncurses.Window
	StartPos Pos
	MaxPos   Pos
	Conf   *GlobalConfig
	BArgs   *BuildArgs
	NowRow   int
}

func NewTerminalDisplay() (*TerminalDisplay, error) {
	td := &TerminalDisplay{}
	var err error
	td.stdscr, err = goncurses.Init()
	if err != nil {
		return nil, errors.New("goncurses init failed.")
	}

	td.MaxPos.y, td.MaxPos.x = td.stdscr.MaxYX()
	td.StartPos.x = td.MaxPos.x / 2 - 20
	td.StartPos.y = td.MaxPos.y / 2 - 5

	return td, nil
}

func (td *TerminalDisplay) Terminate() {
	goncurses.End()
}

func (td *TerminalDisplay) Clear() {
	td.stdscr.Clear()
}

func (td *TerminalDisplay) TipsWhichToCompile() {
	tips := []string{
		"Press Enter for next",
		"Choose which you want to compile: ",
	}
	for _, v := range td.Conf.item {
		tips = append(tips, v.Name)
	}
	indx := 0
	for k, v := range tips {
		if k <= 1 {
			td.stdscr.MovePrintln(td.StartPos.y+indx, td.StartPos.x, v)
		} else {
			td.stdscr.MovePrintf(td.StartPos.y+indx, td.StartPos.x, "[X] %s\n", v)
		}
		indx = indx + 1
		td.stdscr.Refresh()
	}
}

func (td *TerminalDisplay) checkPlatform(platform string) bool {
	if len(platform) == 0 {
		td.BArgs.Platform = "x64"
		return true
	}
	switch platform {
	case "x64":
	case "CentOS-x64":
	case "corei7":
	default:
		return false
	}
	td.BArgs.Platform = platform
	return true
}

func (td *TerminalDisplay) checkRelease(release string) bool {
	if len(release) == 0 {
		td.BArgs.Release = "release"
		return true
	}
	switch release {
	case "Release":
	case "Debug":
	case "release":
	case "debug":
	case "RELEASE":
	case "DEBUG":
	default:
		return false
	}
	td.BArgs.Release = release
	return true
}

func (td *TerminalDisplay) checkInstallTo(installto string) bool {
	if len(installto) == 0 {
		return false
	}
	if installto[len(installto) - 1] != '/' {
		// Check directory exist.
		td.BArgs.InstallTo = installto + "/"
	} else {
		td.BArgs.InstallTo = installto
	}
	return true
}

func (td *TerminalDisplay) TipsCompileArgs() error {
	indx := 0 // where is row now.
	x_pos := 10
	tips := "What platform to compile, support x64 (ubuntu 64bit), CentOS-x64 (CentOS), corei7 (crosscompile), x64(default) : "
	td.stdscr.MovePrint(td.StartPos.y+indx, x_pos, tips)
	indx = indx + 1
	platform, err := td.stdscr.GetString(12)
	if err != nil {
		return errors.New("Error platform")
	}
	ret := td.checkPlatform(platform)
	if !ret {
		return errors.New("Invalid platform")
	}

	// For give a line break
	td.stdscr.MovePrintln(td.StartPos.y+indx, x_pos, "")
	indx = indx + 1

	tips = "Compile Debug/Release : "
	td.stdscr.MovePrint(td.StartPos.y+indx, x_pos, tips)
	indx = indx + 1
	release, err := td.stdscr.GetString(8)
	if err != nil {
		return errors.New("Error Release")
	}
	ret = td.checkRelease(release)
	if !ret {
		return errors.New("Invalid Release")
	}

	installto := fmt.Sprintf("%s/../public/lib/%s", td.Conf.baseDir, platform)
	tips = fmt.Sprintf("Install to (%s) : ", installto)
	td.stdscr.MovePrint(td.StartPos.y + indx, x_pos, tips)
	indx = indx + 1
	install, err := td.stdscr.GetString(128)
	if err != nil {
		return errors.New("Error InstallTo")
	}
	// You must give install to because we don't know what the name of
	// the project name in public libs name.
	if len(install) == 0 {
		return errors.New("You must give a Install To")
	} else {
		ret = td.checkInstallTo(install)
		if !ret {
			return errors.New("Invalid InstallTo")
		}
	}

	return nil
}

func (td *TerminalDisplay) DisplayBuildArgs() bool {
	x_pos := 20
	indx := 1
	td.stdscr.MovePrintf(td.StartPos.y + indx, td.StartPos.x - x_pos, "Build platform is : %s\n", td.BArgs.Platform)
	indx = indx + 1
	td.stdscr.MovePrintf(td.StartPos.y + indx, td.StartPos.x - x_pos, "Build Release is : %s\n", td.BArgs.Release)
	indx = indx + 1
	td.stdscr.MovePrintf(td.StartPos.y + indx, td.StartPos.x - x_pos, "Build InstallTo is : %s\n", td.BArgs.InstallTo)
	indx = indx + 1
	td.stdscr.MovePrintln(td.StartPos.y + indx, td.StartPos.x - x_pos, "")
	indx = indx + 1
	td.stdscr.MovePrintln(td.StartPos.y + indx, td.StartPos.x - x_pos, "Please be sure right of them? Enter yes or no, program will exit but enter yes ! ")
	indx = indx + 1
	td.stdscr.MovePrint(td.StartPos.y + indx, td.StartPos.x - x_pos, "Please enter (yes/no): ")

	str, err := td.stdscr.GetString(4)
	if err != nil {
		return false
	}
	if strings.Compare(str, "yes") == 0 {
		return true
	}
	return false
}

func (td *TerminalDisplay) DisplayData(str string) {
	td.stdscr.MovePrintln(td.StartPos.y + td.NowRow, td.StartPos.x - 20, str)
	td.NowRow = td.NowRow + 1
}

func (td *TerminalDisplay) WaitNextForEnter() {
	for {
		// I don't know why the value is 10 when press enter
		if ch := td.stdscr.GetChar(); ch == 10 {
			return
		} else {
			jklog.L().Debugln("key ", ch)
		}
	}
}

func (td *TerminalDisplay) WaitForAnyKey() int {
	return int(td.stdscr.GetChar())
}
