package sctek

import (
	"fmt"
	"jk/jklog"
)

type SctekDiscover struct {
	broadNet *SctekBroadNet
	DevList  []SctekDeviceList
	stop     bool
}

type SctekDeviceList struct {
	Data    string
	Version string
	IP      string
	MAC     string
}

func NewSctekDiscover() (*SctekDiscover, error) {
	sd := &SctekDiscover{}
	var err error
	sd.broadNet, err = NewSctekBroadNet("0.0.0.0", 30001)
	if err != nil {
		return nil, err
	}
	return sd, nil
}

func (sd *SctekDiscover) Discover(duration int) ([]SctekDeviceList, error) {
	sd.stop = false
	for {
		if sd.stop {
			break
		}
		jklog.L().Debugln("Start to recv data")
		buf, err := sd.broadNet.Recv()
		if err != nil {
			jklog.L().Errorln(err)
			return nil, err
		}
		jklog.L().Debugln("recv data ", string(buf))
	}
	return sd.DevList, nil
}

func (sd *SctekDiscover) Clear() {
	sd.stop = true
}

func (sd *SctekDiscover) DebugDevicePrint() {
	for k, v := range sd.DevList {
		fmt.Printf("%d : %s, %s, %s\n", k, v.Version, v.IP, v.MAC)
	}
}
