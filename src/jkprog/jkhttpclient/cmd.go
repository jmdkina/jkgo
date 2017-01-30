package main

import (
	"net/http"
	"jk/jklog"
	"strings"
	"time"
	"io/ioutil"
	"encoding/json"
)

// Conf of route could set by file
type RouteConf struct {
	Url       string     `json:url`
	User      string     `json:user`
	Pass      string     `json:pass`
	Interval  int        `json:interval`
	Users     []string   `json:users`
	Passes    []string   `json:passes`
}

// Get Conf of route from @filename
func RouteConfFromFile(filename string) *RouteConf {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		jklog.L().Errorf("Read file error: %v\n", err)
		return nil
	}
	var rc RouteConf
	err = json.Unmarshal(data, &rc)
	if err != nil {
		jklog.L().Errorf("Parse data failed %v\n", err)
		return nil
	}
	return &rc
}

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
	//req := make_request_get("http://192.168.1.1")

	rc := RouteConfFromFile("./etc/router.conf")
	if rc == nil {
		return
	}
	jklog.L().Debugln("router: ", rc)

	//auth_users := []string{"root", "admin" }
	//auth_passwds := []string{ "zhuhaiJIAJIA11072", "admin", "root", "123456", "12345678" }

	find := false
	for _, key := range rc.Users {
		for _, value := range rc.Passes {
			jklog.L().Debugf("check with [%s, %s]\n", key, value)
			req := make_request_get("http://" + rc.Url)
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

			time.Sleep(time.Duration(rc.Interval)*time.Millisecond)
		}
		if find {
			break
		}
	}
}
