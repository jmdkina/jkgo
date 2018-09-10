package sctek

import "fmt"

type SctekDiscover struct {
	DevList []SctekDeviceList
}

type SctekDeviceList struct {
	Version string
	IP      string
	MAC     string
}

func NewSctekDiscover() (*SctekDiscover, error) {
	return &SctekDiscover{}, nil
}

func (sd *SctekDiscover) Discover(duration int) ([]SctekDeviceList, error) {
	return sd.DevList, nil
}

func (sd *SctekDiscover) Clear() {

}

func (sd *SctekDiscover) DebugDevicePrint() {
	for k, v := range sd.DevList {
		fmt.Printf("%d : %s, %s, %s\n", k, v.Version, v.IP, v.MAC)
	}
}
