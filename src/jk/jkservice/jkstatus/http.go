package jkstatus

import (
	"encoding/json"
	"github.com/alecthomas/log4go"
	"net/http"
	"strconv"
)

type StatusHttp struct {
}

func (sh *StatusHttp) WriteSerialData(w http.ResponseWriter, data interface{}, status int) {
	res := map[string]interface{}{
		"Status": status,
	}
	res["Result"] = data
	d, _ := json.Marshal(res)
	w.Write(d)
}

func (sh *StatusHttp) ServerStatus(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log4go.Debug("Get request of cmd %s", r.FormValue("cmd"))
	switch r.FormValue("cmd") {
	case "getstatus":
		sh.WriteSerialData(w, "No data", 200)
		break
	default:
		sh.WriteSerialData(w, "Error", 400)
	}
}

func NewStatusHttp(port int) (*StatusHttp, error) {
	sh := &StatusHttp{}
	log4go.Debug("Start http server of port %d ", port)
	http.HandleFunc("/st", sh.ServerStatus)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
	return sh, nil
}
