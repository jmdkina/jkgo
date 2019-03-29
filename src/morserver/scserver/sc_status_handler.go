package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"jk/jklog"
	"jkbase"
	"jkdbs"
	"net/http"
	"time"
)

type SCDevStatusHeader struct {
	Version       string
	Cmd           string
	SubCmd        string // login, logout
	Time          int64
	TransactionID string
}

type SCDevStatusBody struct {
	Version  string
	Duration int
	DevType  string
	Mac      string
	IP       string
}

type SCDevStatusInfo struct {
	Header SCDevStatusHeader `json:"header"`
	Body   SCDevStatusBody   `json:"body"`
}

type SCStatusPage struct {
	jkbase.WebBaseInfo
}

var dbhandle *jkdbs.CMMysqlSC

func (st *SCStatusPage) Get(w http.ResponseWriter, r *http.Request) {
	jklog.L().Debugln("sctek st get")
	key := r.FormValue("sctek")
	if key != "sctekkey_get" {
		return
	}
	filename := "./html/sctek/sctek_status.html"
	dbstatus := []*jkdbs.SCDevStatus{}
	if dbhandle != nil {
		dbstatus, _ = dbhandle.QueryDevStatus()
	}
	err := st.Parse(w, filename, dbstatus)
	if err != nil {
		jklog.L().Errorln("error parse ", err)
	}
}

func (st *SCStatusPage) Post(w http.ResponseWriter, r *http.Request) {
	jklog.L().Debugln("sctek st post")
	key := r.FormValue("sctek")
	if key != "sctekkey" {
		return
	}
	buffer, _ := ioutil.ReadAll(r.Body)
	stinfo := SCDevStatusInfo{}
	err := json.Unmarshal(buffer, &stinfo)
	if err != nil {
		jklog.L().Errorln("json parse fail ", err)
		jklog.L().Debugln("json: ", string(buffer))
		st.GenerateResponse(w, "json error", 444)
		return
	}
	err = st.parseSaveStatus(stinfo, r.RemoteAddr)
	if err != nil {
		jklog.L().Errorln("db insert fail ", err)
		st.GenerateResponse(w, "db insert fail", 505)
	} else {
		st.GenerateResponse(w, "OK", 200)
	}
}

func (st *SCStatusPage) checkTransactionID(stinfo SCDevStatusInfo) bool {
	originstr := fmt.Sprintf("%s%s%s%d", stinfo.Header.Version, "sctek", stinfo.Header.SubCmd, stinfo.Header.Time)
	calmd5 := md5.Sum([]byte(originstr))
	calmd5str := hex.EncodeToString(calmd5[:])
	jklog.L().Debugf("check md5 cal [%s] origin [%s]\n",
		calmd5str, stinfo.Header.TransactionID)
	return calmd5str == stinfo.Header.TransactionID
}

func (st *SCStatusPage) parseSaveStatus(stinfo SCDevStatusInfo, remoteip string) error {
	if !st.checkTransactionID(stinfo) {
		return errors.New("Invalid transaction id")
	}
	if stinfo.Body.Mac == "" {
		return errors.New("Mac not exist")
	}
	dbst := jkdbs.SCDevStatus{}
	dbst.Mac = stinfo.Body.Mac
	dbst.IP = stinfo.Body.IP
	dbst.DevType = stinfo.Body.DevType
	dbst.ActionType = stinfo.Header.SubCmd
	dbst.RemoteIp = remoteip
	dbst.Time = time.Now().Unix()
	// dbst.Time = stinfo.Header.Time / 1000 // history status use.
	dbst.Version = stinfo.Body.Version
	if stinfo.Header.SubCmd == "login" {
		dbst.Online = true
	} else {
		dbst.Online = false
	}

	if dbhandle != nil {
		err := dbhandle.AddDevStatus(dbst)
		if err != nil {
			return err
		}
	} else {
		return errors.New("DB not inited.")
	}
	return nil
}
