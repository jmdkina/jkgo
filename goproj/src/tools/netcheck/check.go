package main

import (
	"io"
	"jk/jklog"
	"net"
	"os"
	"strconv"
	"time"
)

type NetCheckProgram struct {
	Params *NetCheckParams
	bSend  bool
}

func Init() *NetCheckProgram {
	return &NetCheckProgram{}
}

func (np *NetCheckProgram) InitListen(addr string, port int) error {
	addrto, err := net.ResolveTCPAddr("tcp", addr+":"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	listen, err := net.ListenTCP("tcp", addrto)
	if err != nil {
		return err
	}
	jklog.L().Infoln("Goto listen client")
	var all int64
	var startT int64
	var endT int64
	go func() {
		for {
			if !np.bSend {
				break
			}
			cl, err := listen.Accept()
			if err != nil {
				jklog.L().Errorln("error accept ", err)
				break
			}
			startT = time.Now().Unix()
			for {
				// Receiver
				data := make([]byte, np.Params.Length)
				// cl.SetReadDeadline(time.Now().Add(time.Microsecond * 10))
				n, err := cl.Read(data)
				if err != nil {
					if err == io.EOF {
						jklog.L().Warnln("End with read done")
						endT = time.Now().Unix()
						break
					}
					jklog.L().Errorln("error: ", err)
					break
				}
				all += int64(n)
				jklog.L().Infoln("n: ", n)
			}

			endT = time.Now().Unix()
			ti := endT - startT
			if ti > 0 {
				jklog.L().Infoln("Info: ", all/ti, " B/s")
			} else {
				jklog.L().Errorln("wrong ti: ", ti)
			}
		}
	}()

	return nil
}

func (np *NetCheckProgram) InitCheck(addr string, port int) error {
	addrto, err := net.ResolveTCPAddr("tcp", addr+":"+strconv.Itoa(port))
	if err != nil {
		return err
	}
	cl, err := net.DialTCP("tcp", nil, addrto)
	if err != nil {
		return err
	}

	f, err := os.OpenFile("/tmp/testdata", os.O_RDONLY, os.ModePerm)
	if err != nil {
		jklog.L().Errorln("No file data, ", err)
		return err
	}
	// Start to write data
	jklog.L().Infoln("Goto Write data")
	go func() {
		for {
			if !np.bSend {
				break
			}
			data := make([]byte, np.Params.Length)
			n, err := f.Read(data)
			if err != nil {
				if err != nil {
					if err == io.EOF {
						jklog.L().Warnln("exit with read done")
						// np.bSend = false
						break
					}
					jklog.L().Errorln("read failed ", err)
					break
				}
			}
			jklog.L().Debugln("Will goto write ", n)
			nw, err := cl.Write(data[:n])
			if err != nil {
				jklog.L().Errorln("Write failed ", err)
			} else {
				if nw != n {
					jklog.L().Errorf("Write length not equal %d, %d\n", nw, n)
				}
			}
			time.Sleep(time.Millisecond * 10)
		}
	}()
	return nil
}

func main() {
	np := Init()
	var err error
	np.Params, err = NetCheckParamsInit("/tmp/NetCheckParams.conf")
	if err != nil {
		jklog.L().Errorln("error for check param init ", err)
		return
	}

	np.bSend = true

	port := 20533
	jklog.L().Infoln("params ok, do listen")
	if np.Params.Reverse {
		np.InitListen(np.Params.DstAddr, port)
		np.InitCheck(np.Params.SouceAddr, port)
	} else {
		np.InitListen(np.Params.SouceAddr, port)
		np.InitCheck(np.Params.DstAddr, port)
	}
	for {
		if !np.bSend {
			break
		}
		time.Sleep(time.Millisecond * 10)
	}
}
