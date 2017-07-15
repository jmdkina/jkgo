package simpleserver

import (
	"net/http"
	"jk/jklog"
	"os"
	"jk/jksys"
)

type Base struct {
	path     string
}

func (b *Base) SetPath(path string) {
	b.path = path
}

type NotFound struct {
	Base
}

func NewNotFound(path string) *NotFound {
	n := &NotFound{}
	n.SetPath(path)
	http.HandleFunc("/", n.ServeHttp)
	return n
}

func (b *NotFound) ServeHttp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/index", http.StatusFound)
	}

	sp := SimpleParse{}
	filename := b.path + "/404.html"
	jklog.L().Debugf("Not found html [%s]\n", filename)

	if _, err := os.Stat(filename); err != nil && !os.IsExist(err) {
		sp.ParseString(w, "The page you request has go to Mars, manager deploy error, please contact to manager", "")
		return
	}

	sp.Parse(w, filename, nil)
}

type Index struct {
	Base
}

type IndexInfo struct {
	Sysinfo      jksys.KFSystemInfo
}

func NewIndex(path string) *Index {
	i := &Index{}
	i.SetPath(path)
	http.HandleFunc("/index", i.ServeHttp)
	return i
}

func (b *Index) ServeHttp(w http.ResponseWriter, r *http.Request) {
	sp := SimpleParse{}
	filename := b.path + "/index.html"
	jklog.L().Debugf("Get html [%s]\n", filename)

	ii := IndexInfo{
		Sysinfo: *jksys.NewSystemInfo(),
	}

	err := sp.Parse(w, filename, ii)
	if err != nil {
		jklog.L().Errorln("Parse error ", err)
	}
}
