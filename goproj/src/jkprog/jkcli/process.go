package main

import (
	"errors"
	"flag"
	"jk/jklog"
	p "jk/jkprotocol"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

type ProcConfig struct {
	Proto        *p.JKProtoUp
	Id           string
	Remote       ProcRemote
	Conn         *net.TCPConn
	KeepInterval int
}

type ProcRemote struct {
	Address string
	Port    int
}

func (pc *ProcConfig) Init(address string, port int) error {
	var err error
	pc.Proto, err = p.JKProtoUpNew(p.JK_PROTOCOL_VERSION_4, pc.Id)
	if err != nil {
		return err
	}
	pc.Remote.Address = address
	pc.Remote.Port = port
	return nil
}

func (pc *ProcConfig) Register() ([]byte, error) {
	return pc.Proto.JKProtoUpRegister()
}

func (pc *ProcConfig) Keepalive() ([]byte, error) {
	return pc.Proto.JKProtoUpKeepalive()
}

func (pr *ProcRemote) remoteString() string {
	return pr.Address + ":" + strconv.Itoa(pr.Port)
}

func (pc *ProcConfig) NetInit() error {
	resolv, err := net.ResolveTCPAddr("tcp", pc.Remote.remoteString())
	if err != nil {
		jklog.L().Errorln("error resolve: ", err)
		return err
	}
	pc.Conn, err = net.DialTCP("tcp", nil, resolv)
	if err != nil {
		jklog.L().Errorln("dial error : ", err)
		return err
	}

	return nil
}

func (pc *ProcConfig) NetDeinit() error {
	if pc.Conn != nil {
		pc.Conn.Close()
	}
	return nil
}

func (pc *ProcConfig) NetSend(data []byte) (int, error) {
	if pc.Conn == nil {
		return 0, errors.New("Uninit")
	}
	return pc.Conn.Write(data)
}

func (pc *ProcConfig) NetConnect() {
	for {
		err := pc.NetInit()
		if err != nil {
			jklog.L().Errorln("Connect to server failed will again: ", err)
			time.Sleep(time.Millisecond * 2000)
			continue
		}

		jklog.L().Debugln("Start to register : ", pc.Id)
		data, _ := pc.Register()
		_, err = pc.NetSend(data)

		if err != nil {
			jklog.L().Errorln("failed register, will do later again :", err)
			time.Sleep(time.Millisecond * 2000)
		} else {
			break
		}
	}
}

var (
	address = flag.String("remote_address", "127.0.0.1", "Connect to Remote address")
	port    = flag.Int("remote_port", 24433, "Connect to remote port")
	id      = flag.String("id", "", "Unique ID")
)

func main() {
	flag.Parse()

	// Set to program name is not exist
	if len(*id) == 0 {
		strs := strings.Split(os.Args[0], "/")
		n := len(strs)
		*id = strs[n-1]
	}

	pc := &ProcConfig{
		Id:           *id,
		KeepInterval: 10,
	}
	pc.Init(*address, *port)

	jklog.L().Debugln("Start to connect : ", pc.Remote.remoteString(), " with id: ", *id)

	now := time.Now().Unix()
	last_keep := int64(0)

	go func() {
	reconn:
		// Connect until success.
		pc.NetConnect()
		jklog.L().Infoln("Start do keep alive")
		for {
			now = time.Now().Unix()
			if last_keep == 0 {
			} else {
				// jklog.L().Debugln("time: ", now, ", ", last_keep)
				if now-last_keep <= int64(pc.KeepInterval) {
					continue
				}
			}
			last_keep = now
			jklog.L().Debugln("need to keepalive : ", pc.Id)
			data, _ := pc.Keepalive()
			_, err := pc.NetSend(data)
			if err != nil {
				jklog.L().Errorln("Send keepalive failed : ", err)
				// should go to connect status
				pc.NetDeinit()
				goto reconn
			}
		}
	}()

	for {
		time.Sleep(time.Millisecond * 500)
	}
}
