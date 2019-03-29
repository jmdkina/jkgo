package jkstatus

import (
	"encoding/json"
	"github.com/alecthomas/log4go"
	"net/http"
	"strconv"
)

type StatusHttp struct {
	stlink *ServiceStatus
}

func (sh *StatusHttp) AddLinkStatus(stlink *ServiceStatus) {
	sh.stlink = stlink
}

func (sh *StatusHttp) WriteSerialData(w http.ResponseWriter, data interface{}, status int) {
	res := map[string]interface{}{
		"Status": status,
	}
	res["Result"] = data
	d, _ := json.Marshal(res)
	w.Write(d)
}

func (sh *StatusHttp) makeAllStatus() string {
	ris := sh.stlink.RemoteInstances()

	data, err := json.Marshal(ris)
	if err != nil {
		log4go.Error("generate json data failed ", err)
		return ""
	} else {
		return string(data)
	}
}

func (sh *StatusHttp) ServerStatus(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log4go.Debug("http server get request of cmd %s", r.FormValue("cmd"))
	switch r.FormValue("cmd") {
	case "getstatus":
		sh.WriteSerialData(w, sh.makeAllStatus(), 200)
		break
	}
}

func NewStatusHttp(port int) (*StatusHttp, error) {
	sh := &StatusHttp{}
	log4go.Debug("Start http server of port %d ", port)
	http.HandleFunc("/st", sh.ServerStatus)
	go func() {
		http.ListenAndServe(":"+strconv.Itoa(port), nil)
	}()
	return sh, nil
}
