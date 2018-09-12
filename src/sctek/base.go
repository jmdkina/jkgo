package sctek

import (
	"fmt"
	"jk/jklog"
)

type SctekDiscover struct {
	broadNet *SctekBroadNet
	DevList  []SctekDeviceList
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
	sd.broadNet, err = NewSctekBroadNet("0.0.0.0", 12306)
	if err != nil {
		return nil, err
	}
	return sd, nil
}

func (sd *SctekDiscover) Discover(duration int) ([]SctekDeviceList, error) {
	count := 0
	for {
		if count > 2 {
			break
		}
		buf, err := sd.broadNet.Recv()
		if err != nil {
			jklog.L().Errorln(err)
			return nil, err
		}
		jklog.L().Debugln("recv data ", string(buf))
		count = count + 1
	}
	return sd.DevList, nil
}

func (sd *SctekDiscover) Clear() {

}

func (sd *SctekDiscover) DebugDevicePrint() {
	for k, v := range sd.DevList {
		fmt.Printf("%d : %s, %s, %s\n", k, v.Version, v.IP, v.MAC)
	}
}
