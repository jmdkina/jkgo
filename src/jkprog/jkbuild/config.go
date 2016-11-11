package main

import "os"

// May readout from config file
type BuildItem struct {
	Name string
	Type string   // cmake, make, autoconf
	IfInstall bool
}

type GlobalConfig struct {
	baseDir   string
	item      []BuildItem
}

func NewGlobalConfig(conffile string) (*GlobalConfig, error) {
	wd, _ := os.Getwd()
	gc := &GlobalConfig{
		baseDir: wd,
	}

	it := BuildItem{}
	it.Type = "cmake"
	it.Name = "AGS"
	it.IfInstall = true
	gc.item = append(gc.item, it)
	return gc, nil
}

type BuildArgs struct {
	Platform string
	Release string
	InstallTo string
}

// BuildArgs is input from user.

func NewBuildArgs() (*BuildArgs, error) {
	ba := &BuildArgs{}
	return ba, nil
}

