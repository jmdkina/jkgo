package main

import (
	"net/http"
	"jk/jklog"
	"strings"
	"time"
)

func make_request_get(url string) *http.Request {
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		jklog.L().Errorf("make requlest get failed, %v\n", err)
		return nil
	}
	return req
}


func make_do_req(req *http.Request) (int, *http.Response) {
	h := http.Client{}
	c, err := h.Do(req)
	if err != nil {
		jklog.L().Errorf("make do request failed, %v\n", err)
		return 0, nil
	}
	//jklog.L().Infof("result: %d\n", c.StatusCode)
	return c.StatusCode, c
}

func main() {
	req := make_request_get("http://192.168.1.1")

	auth_users := []string{"root", "admin" }
	auth_passwds := []string{ "zhuhaiJIAJIA11072", "admin", "root", "123456" }

	find := false
	for _, key := range auth_users {
		for _, value := range auth_passwds {
			jklog.L().Debugf("check with [%s, %s]\n", key, value)
			req.SetBasicAuth(key, value)
			code, res := make_do_req(req)
			jklog.L().Infof("code: %d\n", code)
			if res != nil {
                jklog.L().Infof("status: %s\n", res.Status)
			}
			if code == 200 {
				jklog.L().Infof("Find one success, [%s, %s]\n", key, value)
				find = true
				break
			}
			time.Sleep(3000*time.Millisecond)
		}
		if find {
			break
		}
	}
}
