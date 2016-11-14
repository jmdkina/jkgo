package main

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

// May readout from config file
type BuildItem struct {
	Name string
	Type string   // cmake, make, autoconf
	IfInstall bool
}

type GlobalConfig struct {
	confFile string
	baseDir   string
	item      []BuildItem
}

type ConfigInfo struct {
	Items   []BuildItem
}

// @conffile: relative path from execute path
func NewGlobalConfig(conffile string) (*GlobalConfig, error) {
	wd, _ := os.Getwd()
	gc := &GlobalConfig{
		baseDir: wd,
		confFile: wd + "/" + conffile,
	}

	err := gc.Config()
	if err != nil {
		return nil, err
	}

	//it := BuildItem{}
	//it.Type = "cmake"
	//it.Name = "AGS"
	//it.IfInstall = true
	//gc.item = append(gc.item, it)
	return gc, nil
}

func (gc *GlobalConfig) Config() error {
	data, err := ioutil.ReadFile(gc.confFile)
	if err != nil {
		return err
	}
	var ci ConfigInfo
	err = json.Unmarshal(data, &ci)
	if err != nil {
		return err
	}
	gc.item = ci.Items
	return nil
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

