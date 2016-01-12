package main

import (
	"jk/jklog"
	client "jk/jknewclient"
	// ser "jk/jknewserver"
	pt "jk/jkprotocol"
)

type KFClient struct {
	cli *client.KFClient
}

func (s *KFClient) Send(data []byte) error {
	pro := pt.NewKFProtocol()
	pro.Init()
	pro.SetCmd(pt.KF_CMD_CONTROL, pt.KF_SUBCMD_KEEPALIVE, []byte("jktest"))
	pro.SetData(data)
	d, err := pro.GenerateData(true) // send data
	if err != nil {
		jklog.Lfile().Errorln("failed generate data : ", err)
		return err
	}

	err = s.cli.Send(d, 10)
	if err != nil {
		jklog.Lfile().Errorln("failed send data : ", err)
		return err
	}

	jklog.Lfile().Debugln("start to recv")
	d, err = s.cli.Recv()
	if err != nil {
		jklog.Lfile().Errorln("failed recv data: ", err)
		return err
	}

	npro, err := pt.KFProtocolParse(d)
	if err != nil {
		jklog.Lfile().Errorln("parse failed of response")
		return nil
	}

	jklog.Lfile().Infoln("recv data: ", string(npro.Body.Data))

	return nil
}

func main() {
	s := &KFClient{}
	cli, err := client.KFClientNew("192.168.133.173", 9988)
	if err != nil {
		jklog.Lfile().Errorln("kf client new failed. ", err)
		return
	}
	s.cli = cli

	err = s.Send([]byte("This is a test data"))
	if err != nil {
		jklog.L().Errorln("error send : ", err)
	}
}
