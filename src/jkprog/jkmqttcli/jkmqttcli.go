package main

import (
	"flag"
	"jk/jklog"
	"time"
)

/*
{
  "broker":"tcp://localhost:1883",
  "user":"",
  "password":"",
  "
}
*/
var (
	conffile = flag.String("conf", "./config.json", "Configure file")
)

func main() {
	flag.Parse()

	conf := init_config(*conffile)
	if conf == nil {
		jklog.L().Errorln("Configure read failed, please check.")
		return
	}

	jklog.L().Debugln(conf)

	mq, err := InitOption(conf)
	if err != nil {
		jklog.L().Errorln("init option failed ", err)
		return
	}

	jklog.L().Debugln("Start to init subscribe")
	// connect for subscribe
	sub_api := &MqttOptionSub{}
	err = sub_api.Init(mq.Opt, conf, 5, "api")
	if err != nil {
		jklog.L().Errorln("failed subscript ", err)
	}

	// if we want subscribe exit
	// sub_api.Exit = <-true

	//
	// Publish
	//
	/*
		jklog.L().Debugln("Start to init publish")
		// When we need publish
		err = mq.Init()
		if err != nil {
			jklog.L().Errorln("publish failed ", err)
		}
	*/

	/*
		for {
			time.Sleep(time.Millisecond * 4000)
			jklog.L().Debugln("Send publish")
			// we can send data now if ok
			publish := [2]string{
				"api", "this is test",
			}
			mq.Pub.Data <- publish
		}
	*/

	for {
		time.Sleep(time.Millisecond * 500)
	}
}
