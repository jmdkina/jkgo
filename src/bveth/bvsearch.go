package bveth

import (
	"code.google.com/p/goprotobuf/proto"
	. "jk/jkcommon"
	"jk/jklog"
	. "jk/jkprotobuf"
	"time"
)

type DeviceInfo struct {
	MacAddr     string
	IpAddr      string
	Netmask     string
	Gateway     string
	SoftVersion string
	HardVersion string
	Name        string
}

type DeviceList struct {
	Device []DeviceInfo
}

var GlobalDeviceList DeviceList

const (
	jk_dial_out_port = 9921
	jk_listen_port   = 9920
)

func clearDeviceList() {
	GlobalDeviceList.Device = GlobalDeviceList.Device[:0]
}

func JKStartBroadCast() int {

	JK_close_listen_udp()
	time.Sleep(100 * time.Millisecond)

	err := jk_dial_udp(jk_dial_out_port)
	if err != nil {
		jklog.L().Errorln("jk_dial_udp ", err)
		return JK_RESULT_E_FAIL
	}

	clearDeviceList()

	err = jk_listen_udp_msg(jk_listen_port)
	if err != nil {
		jklog.L().Errorln("jk_listen_udp_msg ", err)
		return JK_RESULT_E_FAIL
	}

	return JK_RESULT_SUCCESS
}

func JKControlCommon(mac, ip string, cmdType int) int {
	return JKControlCommonExtString(mac, ip, cmdType, 0, "")
}

func JKControlCommonExtString(mac, ip string, cmdType int, dlen int, dstring string) int {
	err := jk_dial_udp_ext_data(mac, ip, jk_dial_out_port, cmdType, dlen, dstring)
	if err != nil {
		jklog.L().Errorln("jk_dial_udp_ext ", err)
		return JK_RESULT_E_FAIL
	}

	JK_close_listen_udp()
	time.Sleep(100 * time.Millisecond)

	err = jk_listen_udp_msg(jk_listen_port)
	if err != nil {
		jklog.L().Errorln("jk_list_udp_msg ", err)
		return JK_RESULT_E_FAIL
	}
	return JK_RESULT_SUCCESS
}

func JKFindItem(mac string) *DeviceInfo {
	for _, v := range GlobalDeviceList.Device {
		if v.MacAddr == mac {
			return &v
		}
	}
	return nil
}

func JKBVEthToString() string {
	var str string
	for _, v := range GlobalDeviceList.Device {
		str += "{MacAddr:" + v.MacAddr +
			";Ip:" + v.IpAddr +
			";Netmask:" + v.Netmask +
			";Gateway:" + v.Gateway +
			";Name:" + v.Name +
			";Software:" + v.SoftVersion +
			";Hardware:" + v.HardVersion +
			"}"
	}
	return str
}

func JK_selfresponse_serialize(sSelfresponse []DeviceInfo) ([]byte, error) {
	pSelfresponse := &SelfresponseAll{}

	for i, v := range sSelfresponse {
		*pSelfresponse.SzSelfresponse[i].Ip = v.IpAddr
		*pSelfresponse.SzSelfresponse[i].Mac = v.MacAddr
		*pSelfresponse.SzSelfresponse[i].Name = v.Name
		*pSelfresponse.SzSelfresponse[i].Netmask = v.Netmask
		*pSelfresponse.SzSelfresponse[i].Gateway = v.Gateway
		*pSelfresponse.SzSelfresponse[i].Firmware = v.SoftVersion
		*pSelfresponse.SzSelfresponse[i].HWversion = v.HardVersion
	}
	ret, err := proto.Marshal(pSelfresponse)
	return ret, err
}

func JK_selfresponse_parse(buf []byte) []DeviceInfo {
	pSelfresponse := &SelfresponseAll{}

	err := proto.Unmarshal(buf, pSelfresponse)
	if err != nil {
		return nil
	}

	var di []DeviceInfo
	for _, v := range pSelfresponse.SzSelfresponse {
		indi := &DeviceInfo{}
		indi.IpAddr = v.GetIp()
		indi.Gateway = v.GetGateway()
		indi.Netmask = v.GetNetmask()
		indi.HardVersion = v.GetHWversion()
		indi.SoftVersion = v.GetSWversion()
		indi.MacAddr = v.GetMac()
		indi.Name = v.GetName()
		di = append(di, *indi)
	}

	return di
}
