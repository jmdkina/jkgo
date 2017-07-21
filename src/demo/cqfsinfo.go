package main

import (
	"os"
	"strconv"
	"jk/jklog"
)

type CQFSHeader struct {
	Magic       string
	Version     uint32
	Target      string
}

type CQFSInfo struct {
	Header   CQFSHeader
}

func NewCQFS(path string) (*CQFSInfo, error) {
	ci := &CQFSInfo{}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	value := make([]byte, 4096)
	_, err = f.Read(value)
	if err != nil {
		return nil, err
	}
	ci.Header.Magic = string(value[:4])
	ii, _ := strconv.Atoi(string(value[5]))
	ci.Header.Version = uint32(ii)
	ci.Header.Target = string(value[6:132])
	return ci, nil
}

func (ci *CQFSInfo) PrintHeader() {
	jklog.L().Debugf("Header magic [%s] version [%d], Target [%s]\n",
		ci.Header.Magic, ci.Header.Version, ci.Header.Target)
}

func main() {
	ci, err := NewCQFS("/opt/data/more/arges")
	if err != nil {
		jklog.L().Errorln(err)
		return
	}
	ci.PrintHeader()
}
