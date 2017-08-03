package simpleserver

import (
	"net/http"
	"jk/jklog"
)

type WebGL struct {
	Base
}

type WebGLInfo struct {

}

func NewWebGL(path string) *WebGL {
	i := &WebGL{}
	i.SetPath(path)
	http.HandleFunc("/webgl", i.ServeHttp)
	return i
}

func (b *WebGL) ServeHttp(w http.ResponseWriter, r *http.Request) {
	sp := SimpleParse{}
	filename := b.path + "/webgl/start.html"
	jklog.L().Debugf("Get html [%s]\n", filename)

	err := sp.Parse(w, filename, "")
	if err != nil {
		jklog.L().Errorln("Parse error ", err)
	}
}