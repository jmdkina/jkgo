package simpleserver

import (
	"strconv"

	"code.google.com/p/go.net/websocket"
)

type WSClientSimple struct {
	addr      string
	port      int
	url       string
	fullurl   string
	originurl string
	conn      *websocket.Conn
}

func NewWSClientSimple(addr, url string, port int) (*WSClientSimple, error) {
	ws := &WSClientSimple{}
	ws.addr = addr
	ws.url = url
	ws.port = port
	ws.fullurl = "ws://" + ws.addr + ":" + strconv.Itoa(port) + "/" + ws.url
	ws.originurl = "http://" + ws.addr + ":" + strconv.Itoa(port)
	return ws, nil
}

func (ws *WSClientSimple) Start() error {
	var err error
	ws.conn, err = websocket.Dial(ws.fullurl, "", ws.originurl)
	if err != nil {
		return err
	}
	return nil
}

func (ws *WSClientSimple) Send(data []byte, waitres bool) (int, []byte, error) {
	n, err := ws.conn.Write(data)
	if err != nil {
		return 0, nil, err
	}
	if waitres {
		var msg = make([]byte, 2048)
		n, err = ws.conn.Read(msg)
		if err != nil {
			return 0, nil, err
		}
		return n, msg, nil
	}
	return n, nil, nil
}

func (ws *WSClientSimple) Recv() (int, []byte) {
	var msg = make([]byte, 2048)
	n, err := ws.conn.Read(msg)
	if err != nil {
		return 0, nil
	}
	return n, msg
}

func (ws *WSClientSimple) Close() {
	ws.conn.Close()
}
