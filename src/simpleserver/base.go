package simpleserver

import (
	"encoding/json"
	"golanger.com/utils"
	"jk/jklog"
	"jkbase"
	"net/http"
	"reflect"
	. "simpleserver/dbs"
	"time"
)

type BaseConfig struct {
	PicsPath string
	HtmlPath string
	Port     int
	Index    string
	DBType   string
	DBUrl    string
	LogFile  string
}

var bConfig BaseConfig

func GlobalSetConfig(filename string) {
	jklog.L().Infof("Use config of file %s\n", filename)
	jkbase.GetConfigInfo(filename, &bConfig)
	jklog.L().Debugln("Config Info ", bConfig)
}

func GlobalBaseConfig() BaseConfig {
	return bConfig
}

type Base struct {
	path  string
	child interface{}
}

func (b *Base) SetPath(path string) {
	b.path = path
}

func (b *Base) Path() string {
	return b.path
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

func (b *Base) LogToDB(r *http.Request) {
	if GlobalDBS() == nil {
		return
	}
	url := r.URL.String()
	remote := r.RemoteAddr
	useragent := r.UserAgent()
	method := r.Method
	data := utils.M{
		"TimeUnix":   time.Now().Unix(),
		"TimeStr":    time.Now().String(),
		"URL":        url,
		"RemoteAddr": remote,
		"UserAgent":  useragent,
		"method":     method,
	}
	GlobalDBS().Add("proj", "log", data)
}

func (b *Base) ServerHttp(w http.ResponseWriter, r *http.Request) {
	c := reflect.ValueOf(b.child)
	inputs := make([]reflect.Value, 2)
	inputs[0] = reflect.ValueOf(w)
	inputs[1] = reflect.ValueOf(r)
	jklog.L().Debugf("URL: method: %s, %s, %s, %s\n", r.Method, r.URL.String(), r.RemoteAddr,
		r.UserAgent())
	// jklog.L().Debugln(r.Referer(), ",", r.Host, ",", r.RequestURI)

	b.LogToDB(r)

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
