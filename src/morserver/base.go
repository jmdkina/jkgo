package sctek

import (
	"fmt"
	"jk/jklog"
	"strings"
	"time"
)

type SctekDiscover struct {
	broadNet *SctekBroadNet
	DevList  map[string]*SctekDeviceList
	stop     bool
}

type SctekDeviceList struct {
	LastTime string
	Data     string
	Version  string
	IP       string
	MAC      string
}

func (sd *SctekDeviceList) parse(buffer string) error {
	items := strings.Split(buffer, ",")
	sd.LastTime = time.Now().Format("2006-01-02 15:04:05")
	sd.MAC = items[1][:len(items[1])]
	sd.IP = items[2]
	sd.Version = items[3][:len(items[3])]
	return nil
}

func NewSctekDiscover() (*SctekDiscover, error) {
	sd := &SctekDiscover{}
	sd.DevList = make(map[string]*SctekDeviceList, 1024)
	var err error
	sd.broadNet, err = NewSctekBroadNet("0.0.0.0", 30001)
	if err != nil {
		return nil, err
	}
	return sd, nil
}

func (sd *SctekDiscover) Discover(duration int) (map[string]*SctekDeviceList, error) {
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
		sdc := &SctekDeviceList{}
		sdc.parse(string(buf))
		sd.DevList[sdc.MAC] = sdc
	}
	return sd.DevList, nil
}

func (sd *SctekDiscover) Clear() {
	sd.stop = true
}

func (sd *SctekDiscover) DebugDevicePrint() {
	for k, v := range sd.DevList {
		fmt.Printf("%s : %s, %s, %s\n", k, v.Version, v.IP, v.MAC)
	}
}
