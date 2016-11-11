package main

import "jk/jklog"

type BuildInfo struct {
	Item   []BuildItem
	BaseDir string  // /tmp/jkbuild/
	BaseSrcDir string
	BArgs  BuildArgs
}

func NewBuildInfo(conf *GlobalConfig, args BuildArgs) (*BuildInfo, error) {
	bi := &BuildInfo{}
	bi.BaseDir = "/tmp/jkbuild"
	bi.BaseSrcDir = conf.baseDir
	bi.BArgs = args
	return bi, nil
}

func (bi *BuildInfo) SetBuildItem(item []BuildItem) {
	bi.Item = item
}

func (bi *BuildInfo) AddBuildItem(item BuildItem) {
	bi.Item = append(bi.Item, item)
}

func (bi *BuildInfo) Build() (string, error) {
	for _, v := range bi.Item {
        if v.Type == "cmake" {
			jklog.Lfile().Debugln("Start to build ", v.Name, " with ", v.Type)
			bc, _ := NewBuildCMake(bi.BaseDir, v.Name, bi.BaseSrcDir + "/" + v.Name)
			str, err := bc.Build(bi.BArgs)
			if err != nil {
				return str, err
			}
			if v.IfInstall {
				str, err := bc.Install()
				if err != nil {
					return str, err
				}
			}
		}
	}

	return "", nil
}
