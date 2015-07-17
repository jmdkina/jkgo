package jkserver

import (
	"bveth"
	"jk/jklog"
	"time"
)

func jk_receive_command(buf []byte) (*jkCommandHeader, []byte) {
	data := string(buf)

	header := jk_parse_command_header(data)
	if header == nil {
		jklog.L().Errorln("header is error")
		return nil, []byte("")
	}

	// Parse header success

	// job: get/set/report,request/response,last/continue
	switch header.name_type(header.eachjob[2]) {
	case JK_SEND_LAST:

		switch header.name_type(header.eachjob[1]) {
		case JK_SEND_DOING:

			switch header.name_type(header.eachjob[0]) {
			case JK_TYPE_GET:

				switch header.name_type(header.eachcommand[0]) {
				case JK_BVWORK:
					switch header.name_type(header.eachcommand[1]) {
					case JK_BVETHERNET:

						switch header.name_type(header.eachcommand[2]) {
						case JK_DEVICELIST:
							str := jk_bv_eth_get_devicelist()
							header.GenJKCommandHeaderOneResponse(JK_BVWORK, JK_BVETHERNET, JK_DEVICELIST)
							return header, str

						}
					}
				}
			case JK_TYPE_SET:
			case JK_TYPE_REPORT:

			}

		case JK_SEND_RESPONSE:
		case JK_SEND_FAIL:

		}

	case JK_SEND_CONTINUE:
		// save data

	}

	// jklog.L().Infoln(str)
	return nil, []byte("")
}

func jk_bv_eth_get_devicelist() []byte {
	bveth.JKStartBroadCast()
	time.Sleep(1000 * time.Millisecond)
	str, _ := bveth.JK_selfresponse_serialize(bveth.GlobalDeviceList.Device)
	// str := bveth.JKBVEthToString()
	bveth.JK_close_listen_udp()
	return str
}
