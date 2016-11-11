package main

import (
	"fmt"
	"os"
	"jk/jklog"
	"strings"
	"os/exec"
	"bytes"
)

type BuildCMake struct {
    BaseBuildDir   string   // basedir/[projectname]  where to build
	SrcDir    string   // What source code to build
	BaseName  string   // Project name will to be build
}

func NewBuildCMake(baseDir string, projname string, srcDir string) (*BuildCMake, error) {
    bc := &BuildCMake{}
	bc.BaseName = projname
	bc.SrcDir = srcDir
	bc.BaseBuildDir = fmt.Sprintf("%s/%s", baseDir, projname)
	jklog.Lfile().Debugln("Check dir ", bc.BaseBuildDir, ", and create it.")
	os.RemoveAll(bc.BaseBuildDir)
	bc.readyBaseDir(bc.BaseBuildDir)
    return bc, nil
}

// Create dir if not exist
func (bc *BuildCMake) readyBaseDir(dirname string) error {
	_, err := os.Stat(dirname)
	if err == nil {
		// exist
		return nil
	}
	if os.IsNotExist(err) {
		err := os.MkdirAll(dirname, os.ModeDir|os.ModePerm)
		if err != nil {
			return  err
		}
	}
	return nil
}

func (bc *BuildCMake) Build(args BuildArgs) (string, error) {
	// Enter the directory
	jklog.Lfile().Debugln("Change to dir ", bc.BaseBuildDir)
	err := os.Chdir(bc.BaseBuildDir)
	if err != nil {
		return "Enter Error", err
	}
	cmakecmd := "cmake"

	cmakeargs := []string{}
	cmakeargsStr := ""
	cmakeargs1 := "-DPLATFORM=" + args.Platform

	cmakeargs = append(cmakeargs, cmakeargs1)

	cmakeargs2 := "-DINSTALL_TO=" + args.InstallTo

	cmakeargs = append(cmakeargs, cmakeargs2)

	cmakeargs3 := ""
	if strings.Compare(args.Release, "Debug") == 0 {
		cmakeargs3 = "-DDEBUG_M=ON"
		cmakeargs = append(cmakeargs, cmakeargs3)
	}

	cmakeargs = append(cmakeargs, bc.SrcDir)
	cmakeargsStr = cmakeargs1 + cmakeargs2 + cmakeargs3 + " " + bc.SrcDir

	jklog.Lfile().Debugf("cmake cmd is : [%s]\n", cmakeargsStr)

	cmd := exec.Command(cmakecmd, cmakeargs1, cmakeargs2, cmakeargs3, bc.SrcDir)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		jklog.Lfile().Errorln("error?: ", err)
		return "cmake Error", err
	}

	jklog.Lfile().Debugln("cmake result : ", out.String())

	jklog.Lfile().Debugln("OK, start make")

	cmd1 := exec.Command("make")
	var outmake bytes.Buffer
	cmd1.Stdout = &outmake
	err = cmd1.Run()
	if err != nil {
		jklog.Lfile().Errorln("error ?: ", err)
		return "make Error", err
	}
	jklog.Lfile().Debugln("make result: ", outmake.String())

	return outmake.String(), nil
}

func (bc *BuildCMake) Install() (string, error) {
	pwd, _ := os.Getwd()
	jklog.Lfile().Debugln("current pos : ", pwd)
	cmd := exec.Command("make", "install")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		jklog.Lfile().Errorln("error?: ", err, " stderrr: ", stderr.String())
		return "Make install fail", err
	}

	jklog.Lfile().Debugln("make install result : ", out.String())
	return out.String(), nil
}

func (bc *BuildCMake) Clear() {
	jklog.Lfile().Debugln("Remove dir ", bc.BaseBuildDir)
    os.RemoveAll(bc.BaseBuildDir)
}
