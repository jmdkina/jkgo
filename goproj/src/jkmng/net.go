package jkmng

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

					n, err := c.Read(data)
					if err != nil {
						if err == io.EOF {
							return
						}
					}

					jklog.L().Infoln("Get data of length: ", n)

					pUP, _ := jkp.JKProtoUpNew(jkp.JK_PROTOCOL_VERSION_4, "")
					err = pUP.JKProtoUpParse(data[:n])
					if err != nil {
						jklog.L().Errorln("Parse upper protocol failed ", err)
						continue
					}

					// Find from items
					findItem := jkmng.JKClientItem{
						Id: pUP.Info.Id,
					}
					ret, fItem := mng.CliCtl.ItemExist(findItem)
					if !ret {
						// First comming add it
						newItem, _ := jkmng.JKClientItemNew(pUP.Info.Id)
						mng.CliCtl.ItemAdd(*newItem)
					} else {
						// Has exist update it
						fItem.UpdateUp()
					}
				}
			}()
		}
	}() // For accept

	return nil
}
