package jkmng

import (
	"jk/jklog"
	"time"
)

type ProcConfig struct {
	Address        string
	Port           int
	UseMongo       bool
	UpdateInterval int64
}

func Process() bool {
	jklog.L().Infoln("Program start here")

	conf := &ProcConfig{
		Address:        "",
		Port:           24433,
		UseMongo:       true,
		UpdateInterval: 60,
	}

	mng, err := JKMNGInit(conf.UseMongo, conf.UpdateInterval)
	if err != nil {
		jklog.L().Errorln("error: ", err)
		return false
	}

	jklog.L().Infoln("Start listen")

	mng.Listen(conf.Address, conf.Port)

	go func() {
		// Check if item has offline, will remove it from lists
		last := time.Now().Unix()
		for {
			now := time.Now().Unix()
			if now-last > 30 {
				last = now
				mng.CliCtl.ItemCheckAndRemove()
			}
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	for {
		time.Sleep(500 * time.Millisecond)
	}

	return false
}
