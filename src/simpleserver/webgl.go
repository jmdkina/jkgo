package simpleserver

import (
	"github.com/gorilla/websocket"
	"jk/jklog"
	"net/http"
)

var upgrader = websocket.Upgrader{} // use default options

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

type WebSocket struct {
	Base
}

type WebSocketInfo struct {
}

func NewWebSocket(path string) *WebSocket {
	i := &WebSocket{}
	i.SetPath(path)
	http.HandleFunc("/websocket", i.ServeHttp)
	return i
}

func (s *WebSocket) ServeHttp(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		jklog.L().Errorln("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			jklog.L().Errorln("read:", err)
			break
		}
		jklog.L().Debugf("recv: %s\n", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			jklog.L().Errorln("write:", err)
			break
		}
	}
}
