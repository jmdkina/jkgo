package simpleserver

import (
	"encoding/json"
	"golanger.com/utils"
	"jk/jklog"
	"net/http"
	"reflect"
)

type Base struct {
	path  string
	child interface{}
}

func (b *Base) SetPath(path string) {
	b.path = path
}

func (b *Base) SetFunc(url string, child interface{}) {
	b.child = child
	http.HandleFunc(url, b.ServerHttp)
}

func (b *Base) WriteSerialData(w http.ResponseWriter, data interface{}, status int) {
	res := utils.M{
		"Status": status,
	}
	res["Result"] = data
	d, _ := json.Marshal(res)
	w.Write(d)
}

func (b *Base) ServerHttp(w http.ResponseWriter, r *http.Request) {
	c := reflect.ValueOf(b.child)
	inputs := make([]reflect.Value, 2)
	inputs[0] = reflect.ValueOf(w)
	inputs[1] = reflect.ValueOf(r)
	jklog.L().Debugf("URL: %s\n", r.URL.String())

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
