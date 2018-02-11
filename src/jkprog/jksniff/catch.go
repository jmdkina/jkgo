package main

import (
	"flag"
	"github.com/google/gopacket/pcap"
	"jk/jklog"
)

type DoGoPacket struct {
}

func NewGoPacket() (*DoGoPacket, error) {
	dgp := &DoGoPacket{}

	return dgp, nil
}

func do_gopacket() {
	NewGoPacket()
	ifs, err := pcap.FindAllDevs()
	if err != nil {
		jklog.L().Errorln("find devs fail ", err)
		return
	}

	jklog.L().Debugln("Find Devices len ", len(ifs))
	for _, v := range ifs {
		jklog.L().Infoln("Name : ", v.Name)
		jklog.L().Infoln("Description: ", v.Description)

		for _, addr := range v.Addresses {
			jklog.L().Infoln("IP: ", addr.IP)
		}
	}
}

var (
	mode = flag.String("mode", "gopacket", "which function to do")
)

func main() {
	flag.Parse()

	switch *mode {
	case "gopacket":
		do_gopacket()
	}
}
