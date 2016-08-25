package main

import (
	"errors"
	"fmt"
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"jk/jklog"
	"time"
)

type MqttOptionPub struct {
	Client *MQTT.Client
	Exit   chan bool
	Data   chan [2]string
}

func (pub *MqttOptionPub) Init(opt *MQTT.ClientOptions, maxCnts, qos int) error {
	pub.Exit = make(chan bool)
	pub.Client = MQTT.NewClient(opt)
	i := 0
	// connect
	for {
		if i > maxCnts {
			break
		}
		i = i + 1
		if token := pub.Client.Connect(); token.WaitTimeout(time.Millisecond*500) && token.Error() != nil {
			jklog.L().Warnf("trying times %d\n", i)
			time.Sleep(time.Millisecond * 2000)
			continue
		}
		break
	}
	if i > maxCnts {
		return errors.New("Fail Max counts")
	}

	go func() {
		pub.Data = make(chan [2]string)
		jklog.L().Debugln("start to wait data to publish")
		for {
			data := <-pub.Data
			if len(data) >= 2 {
				ptoken := pub.Client.Publish(data[0], byte(qos), false, data[1])
				ptoken.WaitTimeout(time.Millisecond * 2000)
				if ptoken.Error() != nil {
					jklog.L().Errorln("publish failed. ", ptoken.Error())
					break
				}
			}
		}
	}()

	// Create thread for wait if exit.
	go func() {
		jklog.L().Debugln("Start to wait exit for publish")
		for {
			exit := <-pub.Exit
			if exit {
				pub.Client.Disconnect(200)
				break
			}
		}
	}()
	return nil
}

type MqttOptionSub struct {
	Client *MQTT.Client
	Exit   chan bool
	Data   chan [2]string
}

func (sub *MqttOptionSub) Init(opt *MQTT.ClientOptions, mp *mqtt_param, maxCnt int, topic string) error {
	sub.Data = make(chan [2]string)
	sub.Exit = make(chan bool)

	opt.SetDefaultPublishHandler(func(client *MQTT.Client, msg MQTT.Message) {
		sub.Data <- [2]string{msg.Topic(), string(msg.Payload())}
	})

	sub.Client = MQTT.NewClient(opt)

	i := 0
	for {
		if i > maxCnt {
			break
		}
		i = i + 1
		jklog.L().Infoln("Start to connect with subscribe")
		if token := sub.Client.Connect(); token.WaitTimeout(time.Millisecond*500) && token.Error() != nil {
			jklog.L().Warnf("trying times %d\n", i)
			time.Sleep(time.Millisecond * 500)
			continue
		}
		break
	}

	if i > maxCnt {
		return errors.New("subscribe connect max counts")
	}

	if token := sub.Client.Subscribe(topic, byte(mp.Qos), nil); token.Wait() && token.Error() != nil {
		jklog.L().Errorln("subscribe error : ", token.Error())
		sub.Client.Disconnect(200)
		return token.Error()
	}

	// recevier data
	go func() {
		jklog.L().Debugln("Start to recevier data of subscribe")
		for {
			incoming := <-sub.Data
			fmt.Printf("RECEIVED TOPIC: %s MESSAGE: %s\n", incoming[0], incoming[1])
			sub.Client.Publish("res/api", byte(mp.Qos), false, "a response of api")
		}
	}()

	// exit
	go func() {
		jklog.L().Debugln("Wait exit of subscribe")
		exit := <-sub.Exit
		if exit {
			jklog.L().Warnln("exit of subscribe now")
			sub.Client.Disconnect(250)
		}
	}()
	return nil
}

type MqttOption struct {
	MP  *mqtt_param
	Opt *MQTT.ClientOptions
	Pub *MqttOptionPub
}

func InitOption(mp *mqtt_param) (*MqttOption, error) {
	mo := &MqttOption{
		MP: mp,
	}
	mo.Opt = MQTT.NewClientOptions()
	mo.Opt.AddBroker(mp.Broker)
	mo.Opt.SetClientID(mp.Id)
	mo.Opt.SetUsername(mp.User)
	mo.Opt.SetPassword(mp.Password)
	mo.Opt.SetCleanSession(mp.Cleansess)

	mo.Pub = &MqttOptionPub{}

	return mo, nil
}

func (mo *MqttOption) Init() error {
	err := mo.Pub.Init(mo.Opt, 5, mo.MP.Qos)
	if err != nil {
		return err
	}
	return nil
}
