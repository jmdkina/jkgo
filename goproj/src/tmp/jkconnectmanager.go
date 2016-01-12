package main

import (
	"io"
	"jk/jklog"
	"net"
	"strconv"
)

const (
	JK_CONNECT_TO_PORT = 12001
	JK_LISTEN_PORT     = 13001
	JK_READ_BLOCK_SIZE = 2 << 10
)

// Manager Connection
// How to do
type JKConnectManager struct {
	// Use channel to receiver data
	ReceiverDataUDP chan string
	FromAddr        string
}

func NewJKConnectManager() *JKConnectManager {
	jkcm := JKConnectManager{}
	jkcm.ReceiverDataUDP = make(chan string)
	return &jkcm
}

func (cm *JKConnectManager) JKConnectManagerRelease() {
	close(cm.ReceiverDataUDP)
}

func (cm *JKConnectManager) JKSendBroadcastQuery(data string) error {
	return cm.JKSendUDPCommand("255.255.255.255", data)
}

func (cm *JKConnectManager) JKSendUDPCommand(addr string, data string) error {
	if len(addr) <= 0 {
		addr = cm.FromAddr
	}
	ln, err := net.ResolveUDPAddr("udp", addr+":"+strconv.Itoa(JK_LISTEN_PORT))
	if err != nil {
		jklog.L().Errorln("resolve udp fail of addr ", addr, " error :", err)
		return err
	}

	cn, err := net.DialUDP("udp", nil, ln)
	if err != nil {
		jklog.L().Errorln("Dial udp fail of addr ", addr, " error : ", err)
		return err
	}
	defer cn.Close()
	n, err := cn.Write([]byte(data))
	if err != nil {
		jklog.L().Errorln("write to addr ", addr, " fail ", err)
		return err
	}
	jklog.L().Debugln("send out to ", addr, " data of len ", n)
	jklog.L().Debugln("send out data are : ", data)
	return nil
}

func (cm *JKConnectManager) JKWaitResponseDevices() {
	ln, err := net.ResolveUDPAddr("udp", ":"+strconv.Itoa(JK_LISTEN_PORT))
	if err != nil {
		jklog.L().Errorln("resolve udp failed ", err)
		return
	}
	cn, err := net.ListenUDP("udp", ln)
	if err != nil {
		jklog.L().Errorln("listen udp failed ", err)
		return
	}
	defer cn.Close()
	jklog.L().Infoln("Start to listen here : ", strconv.Itoa(JK_LISTEN_PORT))

	buf := make([]byte, JK_READ_BLOCK_SIZE)
	for {

		n, from, err := cn.ReadFromUDP(buf)
		if err != nil {
			jklog.L().Errorln("read from udp failed ", err)
			break
		}

		jklog.L().Debugln("read from udp result ", n, " from ", from)
		jklog.L().Debugln("out string ", string(buf[0:n]))

		cm.FromAddr = from.IP.String()
		cm.ReceiverDataUDP <- string(buf)
	}
}

func jk_listen_local_with_tcp() {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(JK_LISTEN_PORT))
	if err != nil {
		jklog.L().Errorln("Listen local fail ", err)
		return
	}

	jklog.L().Infoln("Start to listen with port ", JK_LISTEN_PORT)
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				jklog.L().Errorln("Accpet error : ", err)
				break
			}
			jk_receive_data(conn)
		}
	}()
}

func jk_receive_data(c net.Conn) {
	buf := make([]byte, JK_READ_BLOCK_SIZE)

	for {
		// inbuf := make([]byte, JK_READ_BLOCK_SIZE)
		n, err := c.Read(buf)
		if err == io.EOF {
			jklog.L().Infoln("read done, out")
			break
		}
		if err != nil {
			jklog.L().Errorln("Read failed : ", err)
			break
		}

		jklog.L().Infoln("read once done length ", n)
	}

	jklog.L().Infoln("Read data length ", len(buf))
}

func jk_send_query(addr string) {
	jklog.L().Infoln("Start to connect to ", addr)
	_, err := net.Dial("tcp", addr+":"+strconv.Itoa(JK_CONNECT_TO_PORT))
	if err != nil {
		jklog.L().Errorln("Dial fail of ", addr+":"+strconv.Itoa(JK_CONNECT_TO_PORT), " ", err)
		return
	}
}
