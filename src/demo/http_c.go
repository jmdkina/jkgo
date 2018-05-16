package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

//var url = "http://124.43.103.179:8080/message/ws"
var url = "http://result.eolinker.com/jkgV297298df5e76190ed507cd513e7841251220cdd8cfd?uri=--"

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
		UserName: "08:00:27:AD:4A:91",
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

func client_test() {
	tm := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: tm,
	}

	str, e := generate_request_string()

	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println("ok, next, str : ", str)

	req, err := http.NewRequest("POST", url, strings.NewReader(str))
	if err != nil {
		// handle error
		fmt.Println("request create error ", err)
		return
	}

	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Something wrong of client request ", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println("Read data error ", err)
	}

	fmt.Println(string(body))
}

func main() {
	client_test()
}
