package main

import (
	"flag"
	"fmt"
	"strings"

	//"fmt"
	"github.com/alecthomas/log4go"
	//"github.com/anaskhan96/soup"
	"io/ioutil"
	"net/http"
)

var (
	url = flag.String("url", "", "url to catch")
)

func http_catch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	buf, err := ioutil.ReadAll(resp.Body)
	return string(buf), err
}

func find_data(data string) []string {
	i := strings.Index(data, "key")
	if i > 0 {
		tm := strings.Split(data[i:80+i], ":")
		if len(tm) > 1 {
			fmt.Println(tm[1])
		}
		find_data(data[i+10:])
	}
	return nil
}

func catch_data(url string) {
	log4go.Debug("catch url [%s]", url)
	resp, err := http_catch(url)
	log4go.Debug("soup get done len [%d]", len(resp))
	if err != nil {
		log4go.Error("error: ", err)
		return
	}
	find_data(resp)
}

func main() {
	flag.Parse()

	log4go.Debug("url : [%s]", *url)
	catch_data(*url)
}
