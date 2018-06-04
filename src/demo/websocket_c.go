package main

import (
	"fmt"
	"net/http"
	"time"

	"encoding/json"

	"golang.org/x/net/websocket"
	"strings"
	"io/ioutil"
)

type URLs struct {
	WSURL   string
	HTTPURL string
}

var (
	urls = map[string]URLs{
		"stable": {"ws://124.43.103.179:8080/message/ws", "http://124.43.103.179/"},
		"test":   {"ws://106.14.61.92:8081/message/ws", "http://124.43.103.179/"},
		"local":  {"ws://127.0.0.1:12345/echo", "http://127.0.0.1:12345"},
	}
)

type RequestSource struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
type RequestData struct {
	MsgData     RequestSource `json:"msgData"`
	Cmd         string        `json:"cmd"`
	MsgType     string        `json:"msgType"`
	TokenId     string        `json:"tokenId"`
	From        string        `json:"from"`
	RequestTime string        `json:"requestTime"`
	ProjectName string        `json:"projectName"`
	SourceType  string        `json:"source_type"`
}

func generate_request_string() (string, error) {
	rs := RequestSource{
		UserName: "00:0c:29:7a:a1:e9",
		Password: "111",
	}
	rd := RequestData{
		MsgData:     rs,
		Cmd:         "login",
		ProjectName: "mbm",
		SourceType:  "box",
		MsgType:     "1",
		TokenId:     "234213415",
		From:        "",
		RequestTime: "1500864060605",
	}
	data, err := json.Marshal(rd)
	if err == nil {
		return string(data), nil
	}
	return "", err
}

var (
	resunknown = "what you send"
)

func GetFileContent(filename string) []byte {
	filefull := "files/" + filename
	data, _ := ioutil.ReadFile(filefull)
	return data
}

func DealWithCmd(ws *websocket.Conn, msg string) {
	fmt.Printf("ws read result [%s]\n", msg)
	if strings.Index(msg, "login") > 0 {
		n, err := ws.Write(GetFileContent("reslogin"))
		if err != nil {
			fmt.Println("Write error ", err)
		}
		fmt.Printf("Write done n [%d]\n", n)
		//time.Sleep(2 * time.Second)
		n, err = ws.Write(GetFileContent("resloginmore"))
		if err != nil {
			fmt.Println("Write error ", err)
		}
		fmt.Printf("Write done n [%d]\n", n)
	} else {
		msgw := []byte(resunknown)
		fmt.Println("send unknow response\n")
		ws.Write(msgw)
	}
}

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	msg := make([]byte, 1024)
	for {
		_, err := ws.Read(msg)
		if err != nil {
			fmt.Println("ws read error ", err)
		} else {
			DealWithCmd(ws, string(msg))
		}
		time.Sleep(100*time.Millisecon)
	}
}

// This example demonstrates a trivial echo server.
func ExampleHandler() {
	http.Handle("/message/ws", websocket.Handler(EchoServer))
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func client_request() error {
	useurl := "test"
	ws, err := websocket.Dial(urls[useurl].WSURL, "", urls[useurl].HTTPURL)
	if err != nil {
		fmt.Println(err)
		return err
	}
	str, err := generate_request_string()
	if err != nil {
		fmt.Println("error generate : ", err)
		return err
	}
	fmt.Printf("Start to write data request [%s]\n", str)
	n, err := ws.Write([]byte(str))
	if err != nil {
		fmt.Println("Write data error ", err)
		return err
	}
	fmt.Printf("Write data success %d\n", n)

	data := make([]byte, 10240)
	n, err = ws.Read(data)
	if err != nil {
		fmt.Println("read error ", err)
		return err
	}
	fmt.Printf("read data %d\n", n)
	fmt.Printf("read out data [%s]\n", string(data))
	return nil
}

func main() {
	go ExampleHandler()
	time.Sleep(2 * time.Second)
	//client_request()

	for {
		time.Sleep(100 * time.Millisecond);
	}
}
