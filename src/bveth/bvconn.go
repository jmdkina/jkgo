package bveth

import (
	// "encoding"
	// "encoding/hex"
	// "encoding/binary"
	// "fmt"
	"bytes"
	"encoding/binary"
	"jk/jklog"
	"net"
	"strconv"
	"strings"
	"time"
)

func jk_get_mac_addr() []byte {
	hardaddr := make([]byte, 32)
	ints, _ := net.Interfaces()
	for _, v := range ints {
		hardaddr = v.HardwareAddr
	}

	if len(hardaddr) < 6 {
		for i := 0; i < 6; i++ {
			hardaddr[i] = 0xa1
		}
	}

	return hardaddr
}

func jk_dial_udp_ext(mac, ip string, port int, cmdType int) error {
	return jk_dial_udp_ext_data(mac, ip, port, cmdType, 0, "")
}

func jk_dial_udp_ext_data(mac, ip string, port int, cmdType int, dlen int, dstring string) error {
	// start to send command to everywhere through udp
	ser, err := net.ResolveUDPAddr("udp", ip+":"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	// jklog.L().Infoln("dial udp ", ser.IP, ser.Port, ser.Zone, ser.Network(), ser.String())
	dialout, err := net.DialUDP("udp", nil, ser)
	if err != nil {
		return err
	}

	// version dstmac srcmac sessionid dataid cmdtype datalen data

	// inmac := []byte(mac)
	lm := jk_get_mac_addr()

	jklog.L().Infoln("The mac addr ", lm)

	b_buf := bytes.NewBuffer([]byte{})
	binary.Write(b_buf, binary.LittleEndian, int32(cmdType))
	cmdbuf := b_buf.Bytes()
	// jklog.L().Infoln("cmd type ", cmdType, ", b_buf ", cmdbuf)

	dlenbuf := []byte{0, 0, 0, 0}
	if dlen > 0 {
		binary.Write(b_buf, binary.LittleEndian, int32(dlen))
		dlenbuf = b_buf.Bytes()
	}

	// 00 00 00 00 ff ff ff ff ff ff 00 0d 60 dc c3 d3 74 65 6d 70 00 00 00 80 01 00 00 00 00 00 00 00
	sendstr := []byte{
		0, 0, 0, 0,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		lm[0], lm[1], lm[2], lm[3], lm[4], lm[5],
		0x4, 0x1, 0x2, 0x3,
		0x0, 0x0, 0x0, 0x80,
		cmdbuf[0], cmdbuf[1], cmdbuf[2], cmdbuf[3],
		dlenbuf[0], dlenbuf[1], dlenbuf[2], dlenbuf[3],
	}

	if dlen > 0 {
		dialout.Write(sendstr)
	} else {
		v := []byte(dstring)
		for _, t := range v {
			sendstr = append(sendstr, t)
		}
		dialout.Write(sendstr)
	}

	jklog.L().Debugln("send string to udp ", port, " down!")

	dialout.Close()
	return nil
}

func jk_dial_udp(port int) error {
	// start to send command to everywhere through udp
	ser, err := net.ResolveUDPAddr("udp4", "255.255.255.255:"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	// jklog.L().Infoln("dial udp ", ser.IP, ser.Port, ser.Zone, ser.Network(), ser.String())
	dialout, err := net.DialUDP("udp", nil, ser)
	if err != nil {
		return err
	}

	// version dstmac srcmac sessionid dataid cmdtype datalen data

	lm := jk_get_mac_addr()

	jklog.L().Infoln("The mac addr ", lm)

	sendstr := make([]byte, 32)
	// 00 00 00 00 ff ff ff ff ff ff 00 0d 60 dc c3 d3 74 65 6d 70 00 00 00 80 01 00 00 00 00 00 00 00
	sendstr = []byte{0, 0, 0, 0,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
		lm[0], lm[1], lm[2], lm[3], lm[4], lm[5],
		// 0x1, 0x3, 0x25, 0x44, 0x12, 0x3,
		0x4, 0x1, 0x2, 0x3,
		0x0, 0x0, 0x0, 0x80,
		0x1, 0x0, 0x0, 0x0,
		0x0, 0x0, 0x0, 0x0}

	dialout.Write(sendstr)

	jklog.L().Debugln("send string to udp ", port, " down!")

	dialout.Close()
	return nil
}

var waitresp *net.UDPConn

func jk_listen_udp_msg(port int) error {
	//
	// start to listen response command
	//
	jklog.L().Infoln("listen udp addr " + strconv.Itoa(port))

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return err
	}
	var localAddr string
	for _, t := range addrs {
		if strings.HasPrefix(t.String(), "192.168") {
			localAddr = t.String()
		}
	}

	jklog.L().Debugln("listen to ", localAddr, ":", strconv.Itoa(port))
	service, err := net.ResolveUDPAddr("udp4", ":"+strconv.Itoa(port))
	if err != nil {
		return err
	}

	waitresp, err = net.ListenUDP("udp", service)

	if err != nil {
		return err
	}

	go func() {
		for {
			buf := make([]byte, 512)
			n, _, err := waitresp.ReadFromUDP(buf)
			if err != nil {
				// jklog.L().Errorln("read from udp ", err)
				break
			}
			if n == 0 {
				time.Sleep(500 * time.Microsecond)
			} else {
				// jklog.L().Infoln("out recv -- ", n, addr)
				// waitresp.WriteToUDP([]byte("get message send response"), addr)
				jk_deal_response(buf[:n])
			}
		}
		// jklog.L().Errorln("shouldn't down here")
	}()

	return nil
}

func JK_close_listen_udp() {
	if waitresp != nil {
		waitresp.Close()
		waitresp = nil
	}
}
