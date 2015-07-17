package main

import (
	"flag"
	"jk/jklog"
	. "jk/jkprotocol"
	"time"
)

type JKConnectCenter struct {
	jkcm JKConnectManager
}

var center JKConnectCenter

var bexit = false

var (
	addr = flag.String("addr", "", "Addr to connect")
)

func jk_set_exit() {
	bexit = true
}

func main() {
	flag.Parse()

	jkcm := &center.jkcm

	jkcm = NewJKConnectManager()

	// send command out
	// content is null
	header := JKGenerateHeaderSimple(-1, JK_P_CMD_DEVICE_CONNECT, JK_P_CMD_DEVICE_CONNECT_QUERY)
	jkcm.JKSendBroadcastQuery(string(header.ToBytes()))

	header_notify := JKGenearteHeaderSimpleNotify(-1, JK_P_CMD_DEVICE_CONNECT, JK_P_CMD_DEVICE_CONNECT_QUERY)
	di := JKDeviceInfo{
		Name:  "xiao mac",
		Addr:  "192.168.6.11",
		ID:    18392831,
		Model: "macos-go",
	}
	noti_data := header_notify.SetData(di.ToByte())
	jkcm.JKSendBroadcastQuery(string(noti_data.ToByte()))

	// wait response from outside
	go jkcm.JKWaitResponseDevices()

	for {
		select {
		case v := <-jkcm.ReceiverDataUDP:
			ParseData(v)
			// jklog.L().Infoln("out the value ", v)
		default:
		}

		if bexit {
			break
		}
		time.Sleep(200 * time.Microsecond)
	}

	return

}

func ParseData(data string) {
	header, err := JKFromBytes([]byte(data))
	if err != nil {
		jklog.L().Errorln("data error ", err)
		return
	}
	intype := header.Header().GetType()

	jklog.L().Debugln("response type ", intype)

	switch intype {
	case JK_P_HEADER_TYPE_SEND:
		ParseDataSend(*header.Header(), string(header.GetData()))
	case JK_P_HEADER_TYPE_RESPONSE:
		ParseDataResponse(*header.Header(), string(header.GetData()))
	case JK_P_HEADER_TYPE_NOTIFY:
	}
}

func ParseDataResponse(h JK_P_Header, data string) {
	mCmd := h.GetMainCommand()
	sCmd := h.GetSlaveCommand()

	jklog.L().Debugln("response ", mCmd, "/", sCmd)

	switch mCmd {
	case JK_P_CMD_DEVICE_CONNECT:
		switch sCmd {
		case JK_P_CMD_DEVICE_CONNECT_QUERY:
			jklog.L().Infoln("data is : ", data)
		case JK_P_CMD_DEVICE_CONNECT_FILE_TRANSFER:
		}
	}
}

func ParseDataSend(h JK_P_Header, data string) {
	mCmd := h.GetMainCommand()
	sCmd := h.GetSlaveCommand()

	jklog.L().Debugln("response ", mCmd, "/", sCmd)

	switch mCmd {
	case JK_P_CMD_DEVICE_CONNECT:
		switch sCmd {
		case JK_P_CMD_DEVICE_CONNECT_QUERY:
			di := JKDeviceInfo{
				Name:  "macos",
				Addr:  "192.168.6.155",
				ID:    18392834,
				Model: "windows10",
			}
			didata := di.ToByte()
			h.SetType(JK_P_HEADER_TYPE_RESPONSE)
			outdata := h.SetData(didata)
			center.jkcm.JKSendUDPCommand("192.168.6.151", string(outdata.ToByte()))
		case JK_P_CMD_DEVICE_CONNECT_FILE_TRANSFER:
		}
	}
}
