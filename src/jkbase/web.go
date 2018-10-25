package jkbase

// web interface for easy to use in some special only need simple page
// example in demo/jkbase_test

import (
	"encoding/json"
	"errors"
	"jk/jklog"
	"net/http"
	"reflect"
	"strconv"

	"golanger.com/utils"
)

type WebBaseHandle struct {
	path    string
	csspath string
	jspath  string
	imgpath string
	respath string
	addpath string
	port    int
}

type WebBaseInfo struct {
	child interface{}
}

func NewWebBaseHandle(port int, path string) (*WebBaseHandle, error) {
	wbh := &WebBaseHandle{
		port:    port,
		path:    path,
		csspath: path,
		jspath:  path,
		imgpath: path,
		respath: path,
		addpath: path,
	}
	jklog.L().Debugf("Use addon path is [%s], css path [%s]\n",
		wbh.addpath, wbh.csspath)
	http.Handle("/css/", http.FileServer(http.Dir(wbh.csspath)))
	http.Handle("/js/", http.FileServer(http.Dir(wbh.jspath)))
	http.Handle("/imgs/", http.FileServer(http.Dir(wbh.imgpath)))
	http.Handle("/res/", http.FileServer(http.Dir(wbh.respath)))
	http.Handle("/addon/", http.FileServer(http.Dir(wbh.addpath)))
	return wbh, nil
}

// Will Block
func (wbh *WebBaseHandle) Listen() error {
	http.ListenAndServe(":"+strconv.Itoa(wbh.port), nil)
	return errors.New("listen error")
}

func (b *WebBaseInfo) SetFunc(url string, child interface{}) {
	b.child = child
	http.HandleFunc(url, b.serverHttp)
}

func (b *WebBaseInfo) Parse(w http.ResponseWriter, filename string, data interface{}) error {
	sp := SimpleParse{}
	err := sp.Parse(w, filename, data)
	return err
}

func (b *WebBaseInfo) GenerateResponse(w http.ResponseWriter, data interface{}, status int) {
	res := utils.M{
		"Status": status,
	}
	res["Result"] = data
	d, _ := json.Marshal(res)
	w.Write(d)
}

func (b *WebBaseInfo) serverHttp(w http.ResponseWriter, r *http.Request) {
	c := reflect.ValueOf(b.child)
	inputs := make([]reflect.Value, 2)
	inputs[0] = reflect.ValueOf(w)
	inputs[1] = reflect.ValueOf(r)
	jklog.L().Debugf("URL: method: %s, %s, %s, %s\n", r.Method, r.URL.String(), r.RemoteAddr,
		r.UserAgent())

	switch r.Method {
	case "GET":
		method := c.MethodByName("Get")
		if method.IsValid() {
			method.Call(inputs)
		} else {
			jklog.L().Warnln("Undefined GET")
		}
		break
	case "POST":
		method := c.MethodByName("Post")
		if method.IsValid() {
			method.Call(inputs)
		} else {
			jklog.L().Warnln("Undefined POST")
		}
		break

	}
}
