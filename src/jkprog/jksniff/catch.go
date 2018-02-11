package main

import (
	"flag"
	"github.com/alecthomas/log4go"
	"github.com/google/gopacket/pcap"
)

type DoGoPacket struct {
}

func NewGoPacket() (*DoGoPacket, error) {
	dgp := &DoGoPacket{}

	return dgp, nil
}

func do_gopacket() {
	dgp, _ := NewGoPacket()
	ifs, err := pcap.FindAllDevs()
	if err != nil {
		log4go.Error("find devs fail ", err)
		return
	}
	for _, v := range ifs {
		log4go.Info("Name : ", v.Name)
		log4go.Info("Description: ", v.Description)

		for _, addr := range v.Addresses {
			log4go.Info("IP: ", addr.IP)
		}
	}
}

var (
	mode = flag.String("mode", "goopacket", "which function to do")
)

func main() {
	flag.Parse()

	switch *mode {
	case "gopacket":
		do_gopacket()
	}
}
