//
// Author: jmdvirus@roamter.com
//
// Client of mqtt
//
package main

import (
	"flag"
	"github.com/surgemq/message"
	"github.com/surgemq/surgemq/service"
	"jk/jklog"
	"time"
)

var (
	addr = flag.String("addr", "127.0.0.1", "Connect to server")
	port = flag.String("port", "23444", "Connect to server with port")
)

func onPublishFunc(msg *message.PublishMessage) error {
	jklog.L().Infoln("Recevied on publish func")
	jklog.L().Infoln("name: ", msg.Name())
	jklog.L().Infoln("payload: ", string(msg.Payload()))
	return nil
}

func onCompleteFunc(msg, ack message.Message, err error) error {
	jklog.L().Infoln("Received complete func")
	jklog.L().Infoln("name: ", msg.Name())
	return nil
}

func main() {

	flag.Parse()

	c := service.Client{}

	msg := message.NewConnectMessage()
	msg.SetWillQos(1)
	msg.SetVersion(4)
	msg.SetCleanSession(true)
	msg.SetClientId([]byte("surgemq"))
	msg.SetKeepAlive(10)
	msg.SetWillTopic([]byte("will"))
	msg.SetWillMessage([]byte("send me home"))
	msg.SetUsername([]byte("surgemq"))
	msg.SetPassword([]byte("verysecret"))

	// Connects to the remote server at 127.0.0.1 port 1883
	str := *addr + ":" + *port
	c.Connect("tcp://"+str, msg)

	// Creates a new SUBSCRIBE message to subscribe to topic "abc"
	submsg := message.NewSubscribeMessage()
	submsg.AddTopic([]byte("abc"), 0)

	// Subscribes to the topic by sending the message. The first nil in the function
	// call is a OnCompleteFunc that should handle the SUBACK message from the server.
	// Nil means we are ignoring the SUBACK messages. The second nil should be a
	// OnPublishFunc that handles any messages send to the client because of this
	// subscription. Nil means we are ignoring any PUBLISH messages for this topic.
	c.Subscribe(submsg, onCompleteFunc, onPublishFunc)

	// Creates a new PUBLISH message with the appropriate contents for publishing
	pubmsg := message.NewPublishMessage()
	pubmsg.SetTopic([]byte("abc"))
	pubmsg.SetPayload(make([]byte, 1024))
	pubmsg.SetQoS(0)

	// Publishes to the server by sending the message
	c.Publish(pubmsg, onCompleteFunc)

	// Disconnects from the server
	jklog.L().Infoln("disconnected.")
	for {
		time.Sleep(5000 * time.Millisecond)
		// c.Publish(pubmsg, onCompleteFunc)
	}

	c.Disconnect()
}
