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

type CatchHuaban struct {
	files_list []string
	need_lists []string
	url        []string
	path       string // where to save
}

func NewCatchHuaban(path string) (*CatchHuaban, error) {
	ch := &CatchHuaban{
		path: path,
	}
	ch.url = append(ch.url, "http://huaban.com/favorite/beauty/")
	ch.url = append(ch.url, "http://huaban.com/boards/24116838/?md=newbn&beauty")
	ch.url = append(ch.url, "http://huaban.com/boards/15729161/?md=newbn&beauty")
	ch.url = append(ch.url, "http://huaban.com/boards/19241298/?md=newbn&beauty")
	ch.url = append(ch.url, "http://huaban.com/boards/18398025?md=newbn&beauty")
	return ch, nil
}

func (ch *CatchHuaban) http_catch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	buf, err := ioutil.ReadAll(resp.Body)
	return buf, err
}

func (ch *CatchHuaban) find_data(data string) []string {
	i := strings.Index(data, "key")
	if i > 0 {
		tm := strings.Split(data[i:80+i], ":")
		if len(tm) > 1 {
			ch.files_list = append(ch.files_list, tm[1])
			//fmt.Println(tm[1])
		}
		ch.find_data(data[i+10:])
	}
	return nil
}

func (ch *CatchHuaban) parse_files_list() []string {
	for _, v := range ch.files_list {
		tm := strings.Split(v, "\"")
		if len(tm) > 1 {
			//fmt.Println(tm[1])
			ch.need_lists = append(ch.need_lists, tm[1])
		}
	}
	return ch.need_lists
}

func (ch *CatchHuaban) catch_data(url string) {
	log4go.Debug("catch url [%s]", url)
	resp, err := ch.http_catch(url)
	log4go.Debug("soup get done len [%d]", len(resp))
	if err != nil {
		log4go.Error("error: ", err)
		return
	}
	ch.find_data(string(resp))
	ch.parse_files_list()
}

func (ch *CatchHuaban) save_data(path string) {
	for _, v := range ch.need_lists {
		fullpath := path + "/" + v + ".jpg"
		fullurl := "http://img.hb.aicdn.com/" + v
		buf, err := ch.http_catch(fullurl)
		if err != nil {
			fmt.Println("error ", err)
			continue
		}
		fmt.Println("Catch ", fullurl, " with len ", len(buf))
		ioutil.WriteFile(fullpath, buf, os.ModePerm)
	}
}

func catch_huaban() {
	ch, _ := NewCatchHuaban(*path)
	for _, u := range ch.url {
		ch.catch_data(u)
		ch.save_data(*path)
	}
}

var (
	option = flag.String("option", "", "which mode to run")
	path   = flag.String("path", ".", "where to save")
)

func main() {
	flag.Parse()

	log4go.Debug("option: ", *option)
	switch *option {
	case "huaban":
		catch_huaban()
	}
	log4go.Debug("Everythin done.")
}
