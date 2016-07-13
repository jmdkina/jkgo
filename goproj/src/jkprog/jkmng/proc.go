package main

import (
	"flag"
	"jk/jklog"
	"time"
)

type ProcConfig struct {
	Address        string
	Port           int
	UseMongo       bool
	UpdateInterval int64
}

func (pc *ProcConfig) Process() bool {
	jklog.L().Infoln("Program start here")

	conf := pc

	mng, err := JKMNGInit(conf.UseMongo, conf.UpdateInterval)
	if err != nil {
		jklog.L().Errorln("error: ", err)
		return false
	}

	jklog.L().Infof("Start listen [%s:%d]\n", conf.Address, conf.Port)

	mng.Listen(conf.Address, conf.Port)

	go func() {
		// Check if item has offline, will remove it from lists
		last := time.Now().Unix()
		for {
			now := time.Now().Unix()
			if now-last > pc.UpdateInterval {
				last = now
				jklog.L().Debugln("Check if need remove")
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

var (
	address = flag.String("local_address", "0.0.0.0", "Listen local address")
	port    = flag.Int("local_port", 24433, "Listen local port")
)

func main() {
	flag.Parse()

	pc := &ProcConfig{
		Address:        *address,
		Port:           *port,
		UseMongo:       true,
		UpdateInterval: 10,
	}

	pc.Process()
}
