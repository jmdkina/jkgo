package jknetwork

import (
	"io"
	"jk/jklog"
	"net"
	"strconv"
	"time"
)

type JKDeviceDebug struct {
	Address string
	Port    int
	conn    net.Conn
}

func New(address string, port int) *JKDeviceDebug {
	debug := &JKDeviceDebug{
		Address: address,
		Port:    port,
	}

	// connect the address.
	connAddr := address + ":" + strconv.Itoa(port)
	conn, err := net.Dial("tcp", connAddr)
	if err != nil {
		jklog.L().Errorln("fail to dial ", connAddr, " with error ", err)
		return nil
	}
	debug.conn = conn

	return debug
}

func (debug *JKDeviceDebug) Close() {
	debug.conn.Close()
}

func (debug *JKDeviceDebug) sendCommandOut(cmd, args string) error {
	sendData := cmd + "[" + args + "]"
	n, err := debug.conn.Write([]byte(sendData))
	if err != nil {
		jklog.L().Errorln("write command failed ", sendData, " with error ", err)
		return err
	} else {
		jklog.L().Infoln("write done data ", sendData, " with len ", n)
	}
	return nil
}

func (debug *JKDeviceDebug) DoPipe(cmd, args string) error {
	err := debug.sendCommandOut(cmd, args)
	if err != nil {
		return err
	}
	return nil
}

func (debug *JKDeviceDebug) DoCommand(cmd, args string) string {
	err := debug.sendCommandOut(cmd, args)
	if err != nil {
		return ""
	}

	// wait the result of the data.
	var resultData string
	//t := time.Now()
	for {
		//now := time.Now()
		//v := now - t
		//if v > 5 {
		//	break
		//}
		recvData := make([]byte, 1024)
		n, err := debug.conn.Read(recvData)
		if io.EOF == err {
			break
		}
		if err != nil {
			jklog.L().Errorln("read error : ", err)
			break
		}
		// resultData = make([]byte, n)
		// resultData = append(resultData, recvData[0:n])
		if n > 0 {
			resultData += string(recvData[0:n])
		}
		jklog.L().Infof("recv data [%s], len [%d]\n", recvData, n)
		break
	}
	return string(resultData)
}

func (debug *JKDeviceDebug) DoFile(cmd, args string) (string, error) {
	err := debug.sendCommandOut(cmd, args)
	if err != nil {
		return "", err
	}

	lListen, err := net.Listen("tcp", debug.conn.LocalAddr().String())
	if err != nil {
		return "", err
	}
	defer lListen.Close()

	c, err := lListen.Accept()
	if err != nil {
		jklog.L().Errorln("Accept error : ", err)
		return "", err
	}

	var resultData string
	for {
		jklog.L().Infoln("Start to recever data from ", debug.conn.LocalAddr().String(), " ...")
		recvdata := make([]byte, 1024)

		// go func() {
		n, err := c.Read(recvdata)
		if io.EOF == err {
			time.Sleep(10 * time.Microsecond)
			break
		}
		if err != nil {
			jklog.L().Errorln("read error : ", err)
			break
		}

		jklog.L().Infoln("get data [", n, "] from out ", string(recvdata))
		// resultData = make([]byte, n+len(resultData))
		// resultData = append(resultData, recvdata[0:n])
		resultData += string(recvdata[0:n])
	}

	return string(resultData), nil
}
