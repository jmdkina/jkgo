package main

import (
	"errors"
	"io"
	jkmng "jk/jkclimng"
	"jk/jklog"
	jkp "jk/jkprotocol"
	"net"
	"strconv"
)

type JKMNG struct {
	Conf   *jkmng.JKClientConfig
	CliCtl *jkmng.JKClientCtrl
}

func JKMNGInit(useMongo bool, upInterval int64) (*JKMNG, error) {
	mng := JKMNG{}

	ctlConf := jkmng.JKClientConfig{
		UseMongo:       useMongo,
		UpdateInterval: upInterval,
	}
	mng.Conf = &ctlConf
	mng.CliCtl, _ = jkmng.JKClientCtrlNew(ctlConf)
	return &mng, nil
}

func (mng *JKMNG) Listen(addr string, port int) error {
	if port < 0 || port > 65535 {
		return errors.New("error port")
	}
	resolv, err := net.ResolveTCPAddr("tcp", addr+":"+strconv.Itoa(port))
	if err != nil {
		return err
	}

	lis, err := net.ListenTCP("tcp", resolv)
	if err != nil {
		return err
	}

	go func() {
		for {
			c, err := lis.Accept()
			if err != nil {
				jklog.L().Errorln("accept error: ", err)
				break
			}

			go func() {
				for {
					data := make([]byte, 1<<12)
					jklog.L().Debugln("return for read again...")

					n, err := c.Read(data)
					if err != nil {
						if err == io.EOF {
							return
						}
					}

					// jklog.L().Infoln("Get data of length: ", n)

					pUP, _ := jkp.JKProtoUpNew(jkp.JK_PROTOCOL_VERSION_4, "")
					err = pUP.JKProtoUpParse(data[:n])
					if err != nil {
						jklog.L().Errorln("Parse upper protocol failed ", err)
						continue
					}

					// Parse out data and set id.
					gen := jkmng.AIGeneral{}
					gen.Parse(pUP.Proto.Body.Data)
					pUP.Info.Id = gen.Header.ID

					// Find from items
					// Need Id need remove, use another method to get it.
					findItem := jkmng.JKClientItem{
						Id: pUP.Info.Id,
					}
					ret, fItem := mng.CliCtl.ItemExist(findItem)
					if !ret {
						// First comming add it
						jklog.L().Debugf("New dev [ %s ] come in", pUP.Info.Id)
						newItem, _ := jkmng.JKClientItemNew(pUP.Info.Id)
						mng.CliCtl.ItemAdd(*newItem)
					} else {
						// Has exist update it
						jklog.L().Debugf("Dev [ %s ] update", fItem.Id)
						fItem.UpdateUp()
					}

					jklog.L().Debugln("Start to parse command and give response")
					ai := jkmng.ActionInfo{}
					err, rdata := ai.Action(pUP.Proto.Body.Data)
					if rdata == nil {
						jklog.L().Warnln("unsupported command")
						continue
					}
					if err != nil {
						jklog.L().Errorln("action command failed.")
					} else {
						pRet, _ := jkp.JKProtoUpNew(jkp.JK_PROTOCOL_VERSION_4, "")
						rsdata, err := pRet.JKProtoUpInit(true, uint32(len(rdata)), rdata)
						if err != nil {
							jklog.L().Errorln("Generate data fail")
						} else {
							//jklog.L().Infoln("Start to give response ", string(rsdata[:20]))
							c.Write(rsdata)
						}
					}
				}
			}()
		}
	}() // For accept

	return nil
}
