package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	//"fmt"
	"github.com/alecthomas/log4go"
	//"github.com/anaskhan96/soup"
	"io/ioutil"
	"net/http"
)

var (
	url  = flag.String("url", "", "url to catch")
	path = flag.String("path", ".", "where to save")
)

func http_catch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	buf, err := ioutil.ReadAll(resp.Body)
	return buf, err
}

var files_list []string

func find_data(data string) []string {
	i := strings.Index(data, "key")
	if i > 0 {
		tm := strings.Split(data[i:80+i], ":")
		if len(tm) > 1 {
			files_list = append(files_list, tm[1])
			//fmt.Println(tm[1])
		}
		find_data(data[i+10:])
	}
	return nil
}

var need_lists []string

func parse_files_list() []string {
	for _, v := range files_list {
		tm := strings.Split(v, "\"")
		if len(tm) > 1 {
			//fmt.Println(tm[1])
			need_lists = append(need_lists, tm[1])
		}
	}
	return need_lists
}

func catch_data(url string) {
	log4go.Debug("catch url [%s]", url)
	resp, err := http_catch(url)
	log4go.Debug("soup get done len [%d]", len(resp))
	if err != nil {
		log4go.Error("error: ", err)
		return
	}
	find_data(string(resp))
	parse_files_list()
}

func save_data(path string) {
	for _, v := range need_lists {
		fullpath := path + "/" + v + ".jpg"
		fullurl := "http://img.hb.aicdn.com/" + v
		buf, err := http_catch(fullurl)
		if err != nil {
			fmt.Println("error ", err)
			continue
		}
		fmt.Println("Catch ", fullurl, " with len ", len(buf))
		ioutil.WriteFile(fullpath, buf, os.ModePerm)
	}
}

func main() {
	flag.Parse()

	log4go.Debug("url : [%s]", *url)
	catch_data(*url)
	save_data(*path)
}
