package main

import (
	"net"
	"jk/jklog"
	"strconv"
	"strings"
	"errors"
	"time"
)

type KFProvideCommandItem struct {
	version        string
	id             string
	cmd            int
	seq            int
	length         int
	ret            int
	header         string
	data           string
}

type KFProvideItem struct {
	inCmd      []byte
	outCmd     []byte
	inAddr     string
	mac        string
	sendData   string
	respData   string
	waitDevice chan bool
	waitClient chan bool

	cmdItem    KFProvideCommandItem
}

type KFProvide struct {
    lis      *net.TCPListener
	item     KFProvideItem
}

func NewKFProvide(port int) (*KFProvide, error) {
    // start listen to port
	l, err := net.ResolveTCPAddr("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return nil, err
	}
	jklog.Lfile().Debugln("start to listen to port ", port)
	lis, err := net.ListenTCP("tcp", l)
	if err != nil {
		return nil, err
	}
	p := &KFProvide{
		lis: lis,
	}

	return p, nil
}

const (
    KF_PROVIDE_CMD_ONLINE_COUNTS = iota  // all online devices
    KF_PROVIDE_CMD_ONLINE  // device is online?
	KF_PROVIDE_CMD_DEVICE_KEEPALIVE // check if device online
	KF_PROVIDE_CMD_CMD   // send command to device
	KF_PROVIDE_CMD_FILE  // get device file info
)

// version-id-cmd-seq-len-ret \r\n retdata
// version: 01
// cmd: KF_PROVIDE_CMD
// seq: sequence
// len: length of data
// ret: return code

func (p *KFProvideCommandItem) parseCommandHeader(data string) error {
	headers := strings.Split(data, "-")
	if len(headers) < 6 {
		return errors.New("not enough args")
	}
	p.version = headers[0]
	p.id = headers[1]
	p.cmd, _ = strconv.Atoi(headers[2])
	p.seq, _ = strconv.Atoi(headers[3])
	p.length, _ = strconv.Atoi(headers[4])
	p.ret, _ = strconv.Atoi(headers[5])
	return nil
}

func (p *KFProvide) onlineDeviceLists(s *KFServer) ([]byte, error) {
	str := ""
	for _, v := range s.Item {
		now := time.Now().Unix()
		if (now - v.lastOnline < 90) {
			str += v.Id + "-"
		}
	}

	return []byte(str), nil
}

func (p *KFProvide) parseDealCommand(s *KFServer) ([]byte, error) {
	switch p.item.cmdItem.cmd {
	case KF_PROVIDE_CMD_DEVICE_KEEPALIVE:
		break
	case KF_PROVIDE_CMD_CMD:
		break
	case KF_PROVIDE_CMD_FILE:
		break
	case KF_PROVIDE_CMD_ONLINE_COUNTS:
		return p.item.cmdItem.header + "\r\n" + p.onlineDeviceLists(s)
		break
	}
	return nil, errors.New("unknow command")
}

func (p *KFProvide) parseInCmd(s *KFServer, data []byte) error {
	// parses out mac and command need to do
	datas := strings.Split(string(data), "\r\n")
	if len(datas) < 1 {
		return errors.New("not header")
	}
	p.item.cmdItem.header = datas[0]
	err := p.item.cmdItem.parseCommandHeader(datas[0])
	if err != nil {
		return err
	}
	p.item.cmdItem.data = datas[1]
	// query  from s.Item find which need to do
	// then send command.
	return nil
}

func (p *KFProvide) waitResponse(item *KFProvideItem) error {
	// wait command response
//	ret := <- item.waitDevice

	// set response data to item.outCmd
	return nil
}

func (p *KFProvide) ListenConnStart(s *KFServer) ([]byte, error) {
	ac, err := p.lis.AcceptTCP()
	if err != nil {
		return nil, err
	}
//	go func(ac) {
		item := &KFProvideItem{}
		item.inCmd = make([]byte, 1024)
		item.inAddr = ac.RemoteAddr().String()
		n, err := ac.Read(item.inCmd)
		if err != nil {
			return nil, nil
		}
		jklog.Lfile().Debugf("%s: read data of len %d\n ", item.inAddr, n)

	    rr := strings.Split(item.inAddr, ":")
	    if len(rr) > 0 {
			if strings.Compare(rr[0], "127.0.0.1") != 0 ||
      			strings.Compare(rr[0], "0.0.0.0") != 0 {
				jklog.Lfile().Debugln("I don't support outside network")
				return nil, nil
			}
		}
		// send data out
		err = p.parseInCmd(s, item.inCmd)

	    if err != nil {
			// switch which command to deal
			jklog.Lfile().Errorln("error parse commadn ", err)
		}
	    data, err := p.parseDealCommand(s)

	    if err != nil {
			ac.Write(data)
		}

		// wait for data response
		err = p.waitResponse(item)

//	}()
	return data, nil
}