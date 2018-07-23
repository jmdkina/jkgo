package simpleserver

import (
	"jk/jklog"
	"net/http"
	"strconv"
	"time"

	"container/list"
	"io/ioutil"
	"strings"

	"golang.org/x/net/websocket"
)

type WSSimpleURLs struct {
	WSURL   string
	HTTPURL string
}

type WSSimple struct {
	datas   list.List
	LisPort int
	LisUrl  string
	Urls    map[string]WSSimpleURLs
	Exit    bool
}

var ws_simple = &WSSimple{}

func GlobalWSSimple() *WSSimple {
	return ws_simple
}

var (
	resunknown = "what you send"
)

func init() {
	ws_simple.Exit = false
	ws_simple.LisUrl = "/message/ws"
	ws_simple.LisPort = 8081
	ws_simple.Urls = map[string]WSSimpleURLs{
		"stable": {"ws://124.43.103.179:8080/message/ws", "http://124.43.103.179/"},
		"test":   {"ws://106.14.61.92:8081/message/ws", "http://124.43.103.179/"},
		"local":  {"ws://127.0.0.1:12345/echo", "http://127.0.0.1:12345"},
		"yt":     {"ws://d3-edu.com:1234/", "http://d3-edu.com"},
	}
}

func (ws *WSSimple) getFileContent(filename string) []byte {
	filefull := "/opt/data/sctek/files-jdh/" + filename
	data, _ := ioutil.ReadFile(filefull)
	return data
}

func (wss *WSSimple) dealWithCmd(ws *websocket.Conn, msg string) {
	jklog.L().Printf("ws read result [%s]\n", msg)
	if strings.Index(msg, "login") > 0 {
		n, err := ws.Write(wss.getFileContent("reslogin"))
		if err != nil {
			jklog.L().Infoln("Write error ", err)
		}
		jklog.L().Printf("Write done n [%d]\n", n)
		//time.Sleep(2 * time.Second)
		n, err = ws.Write(wss.getFileContent("resloginmore"))
		// n, err = ws.Write(wss.getFileContent("resloginnew"))
		if err != nil {
			jklog.L().Infoln("Write error ", err)
		}
		jklog.L().Printf("Write done n [%d]\n", n)
	} else if strings.Index(msg, "initData") > 0 {
		_, err := ws.Write(wss.getFileContent("reschannels"))
		if err != nil {
			jklog.L().Infoln("Write error ", err)
		}
	} else if strings.Index(msg, "getCookSongs") > 0 {
		if strings.Index(msg, "\"channelId\" : 95") > 0 {
			if (strings.Index(msg, "\"cookId\" : 201") > 0) {
			    ws.Write(wss.getFileContent("rescooksongs_1"))
			} else if (strings.Index(msg, "\"cookId\" : 202") > 0) {
			    ws.Write(wss.getFileContent("rescooksongs_2"))
			} else if (strings.Index(msg, "\"cookId\" : 203") > 0) {
			    ws.Write(wss.getFileContent("rescooksongs_5"))
			}
		} else if strings.Index(msg, "\"channelId\" : 96") > 0 {
			if (strings.Index(msg, "\"cookId\" : 201") > 0) {
			    ws.Write(wss.getFileContent("rescooksongs_1"))
			} else if (strings.Index(msg, "\"cookId\" : 202") > 0) {
			    ws.Write(wss.getFileContent("rescooksongs_2"))
			} else if (strings.Index(msg, "\"cookId\" : 203") > 0) {
			    ws.Write(wss.getFileContent("rescooksongs_3"))
			}
		} else if strings.Index(msg, "\"channelId\" : 97") > 0 {
			if (strings.Index(msg, "\"cookId\" : 201") > 0) {
			    ws.Write(wss.getFileContent("rescooksongs_1"))
			} else if (strings.Index(msg, "\"cookId\" : 202") > 0) {
			    ws.Write(wss.getFileContent("rescooksongs_4"))
			} else if (strings.Index(msg, "\"cookId\" : 203") > 0) {
			    ws.Write(wss.getFileContent("rescooksongs_3"))
			}
		}
	} else if strings.Index(msg, "getCallWaitings") > 0 {
		if strings.Index(msg, "\"channelId\" : 95") > 0 ||
			strings.Index(msg, "\"channelId\" : 100") > 0 ||
			strings.Index(msg, "\"channelId\" : 172") > 0 {
			ws.Write(wss.getFileContent("rescallwaitings_1"))
		} else if strings.Index(msg, "\"channelId\" : 96") > 0 ||
			strings.Index(msg, "\"channelId\" : 101") > 0 {
			ws.Write(wss.getFileContent("rescallwaitings_2"))
		} else if strings.Index(msg, "\"channelId\" : 97") > 0 ||
			strings.Index(msg, "\"channelId\" : 102") > 0 {
			ws.Write(wss.getFileContent("rescallwaitings_3"))
		}
	} else if strings.Index(msg, "getAds") > 0 {
		ws.Write(wss.getFileContent("resads"))
	} else {
		msgw := []byte(resunknown)
		jklog.L().Infoln("send unknow response\n")
		ws.Write(msgw)
	}
}

// Echo the data received on the WebSocket.
func (wss *WSSimple) handle_server(ws *websocket.Conn) {
	msg := make([]byte, 1024)
	go func() {
		for {
			e := wss.datas.Front()
			if e != nil {
				senddata := e.Value.(string)
				jklog.L().Infof("will send data \n[%s]\n", senddata)
				ws.Write([]byte(senddata))
				wss.datas.Remove(e)
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()
	for {
		_, err := ws.Read(msg)
		if err != nil {
			jklog.L().Errorln("ws read error ", err)
		} else {
			wss.dealWithCmd(ws, string(msg))
		}
		time.Sleep(100 * time.Millisecond)
		if wss.Exit {
			break
		}
	}
	jklog.L().Infoln("Exit websocket server")
}

func (ws *WSSimple) Start() {
	jklog.L().Infof("start ws server now with port [%d]", ws_simple.LisPort)
	http.Handle(ws_simple.LisUrl, websocket.Handler(ws.handle_server))
	go func() {
		err := http.ListenAndServe(":"+strconv.Itoa(ws_simple.LisPort), nil)
		if err != nil {
			jklog.L().Errorln("ListenAndServe: " + err.Error())
		}
		jklog.L().Infof("exit listen [%s:%d]\n", "localhost", ws_simple.LisPort)
	}()
}

func (ws *WSSimple) Stop() {
	jklog.L().Debugln("stop server")
}

func (ws *WSSimple) CommonCmd(cmd string) {
	jklog.L().Infof("Now do subcmd [%s]\n", cmd)
	switch cmd {
	case "stopPlaying":
		ws.Send(string(ws.getFileContent("reqstopplaying")))
		break
	case "resumePlaying":
		ws.Send(string(ws.getFileContent("reqresumeplaying")))
		break
	case "dataUpdate":
		ws.Send(string(ws.getFileContent("resdataupdate")))
		break
	case "songInsert":
		ws.Send(string(ws.getFileContent("ressonginsert")))
		break
	case "songSkip":
		ws.Send(string(ws.getFileContent("ressongskip")))
		break
	case "logout":
		ws.Send(string(ws.getFileContent("reslogout")))
		break
	default:
		jklog.L().Warnf("Unknow command %s\n", cmd)
		break
	}
}

func (ws *WSSimple) Send(data string) {
	ws.datas.PushBack(data)
}
