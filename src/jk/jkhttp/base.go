package jkhttp

import (
	"errors"
	"io/ioutil"
	"jk/jklog"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func JKPostDataExtern(ip string, port int, page string, data []byte, timeout int) ([]byte, error) {
	url := "http://" + ip + ":" + strconv.Itoa(port) + "/" + page
	jklog.Lfile().Debugln("post url: ", url)
	client := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	resp, err := client.Post(url, "text/plain", strings.NewReader(string(data)))
	if err != nil {
		jklog.Lfile().Errorln("http post failed: ", err)
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func JKPostData(ip string, port int, page string, data []byte, timeout int) ([]byte, error) {
	ips := strings.Split(ip, ",")
	for _, ipitem := range ips {
		data, err := JKPostDataExtern(ipitem, port, page, data, timeout)
		if err == nil {
			return data, nil
		}
	}
	return nil, errors.New("All post data failed")
}
