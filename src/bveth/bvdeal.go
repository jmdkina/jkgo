package bveth

import (
	"code.google.com/p/goprotobuf/proto"
	"jk/jklog"
	inproto "jk/jkprotobuf"
	// "strconv"
	// "strings"
	"bytes"
	"encoding/binary"
)

const (
	BV_SELFRESPONSE          = 2
	BV_DEVICEINFO            = 10002
	BV_DISABLE_REBOOT        = 5001
	BV_ENABLE_REBOOT         = 5002
	BV_OPEN_MISC_CTRL_DEBUG  = 5003
	BV_CLOSE_MISC_CTRL_DEBUG = 5004
	BV_OPEN_MISC_CTRL_CYCLE  = 5005
	BV_CLOSE_MISC_CTRL_CYCLE = 5006
	BV_OPEN_LIVECAMS_LOG     = 5007
	BV_CLOSE_LIVECAMS_LOG    = 5008
	BV_OPEN_CONSOLE          = 5009
	BV_CLOSE_CONSOLE         = 5010
	BV_DISABLE_READ_SIGNAL   = 5011
	BV_EANBLE_READ_SIGNAL    = 5012
	BV_DISABLE_MESSAGE       = 5013
	BV_ENABLE_MESSAGE        = 5014
)

func jk_deal_response(buf []byte) {
	// var s Selfresponse
	tmpcmd := buf[24:28]
	var cmdtype int32
	binary.Read(bytes.NewReader(tmpcmd), binary.LittleEndian, &cmdtype)
	jklog.L().Debugln("Received cmd -- ", cmdtype)

	tmpcmd = buf[28:32]
	var datalen int32
	binary.Read(bytes.NewReader(tmpcmd), binary.LittleEndian, &datalen)

	jklog.L().Debugln("datalen --- ", datalen)

	jk_deal_cmdtype(cmdtype, buf[32:])
}

func jk_deal_cmdtype(cmdtype int32, buf []byte) {

	switch cmdtype {
	case 2: // Selfresonse
		s := &inproto.Selfresponse{}
		// jklog.L().Infoln("buf out", buf)
		e := proto.Unmarshal(buf, s)
		if e != nil {
			jklog.L().Errorln("unmarshall ", e)
			return
		}
		for _, v := range GlobalDeviceList.Device {
			if v.MacAddr == s.GetMac() {
				return
			}
		}
		di := DeviceInfo{
			MacAddr:     s.GetMac(),
			IpAddr:      s.GetIp(),
			Gateway:     s.GetGateway(),
			Netmask:     s.GetNetmask(),
			Name:        s.GetName(),
			SoftVersion: s.GetSWversion(),
			HardVersion: s.GetHWversion(),
		}
		GlobalDeviceList.Device = append(GlobalDeviceList.Device, di)

		jklog.L().Debugln("get device -- ", di.MacAddr, di.IpAddr, di.Netmask,
			di.Gateway, di.SoftVersion, di.HardVersion, di.Name)

	case 10002: // DeviceInfo

	default:
		jklog.L().Infoln("Sorry, I not support this cmd now ", cmdtype)
	}
}
