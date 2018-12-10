package jkmisc

import (
	"encoding/json"
	"jk/jklog"
	"net/url"
	"time"
)

type JKWether struct {
	JKWetherBase

	timeout     int64
	lastGetTime int64
	param       url.Values
}

func JKWetherNew(key string) (*JKWether, error) {
	w := &JKWether{
		param:       url.Values{},
		lastGetTime: 0,
		timeout:     3600 * 12, // 12 hours update once
	}
	w.key = key
	w.url = "http://v.juhe.cn/weather/index"

	return w, nil
}

func (w *JKWether) generateParam(location string, format string, key string) error {

	//配置请求参数,方法内部已处理urlencode问题,中文参数可以直接传参
	w.param.Set("cityname", location) //城市名或城市ID，如：&quot;苏州&quot;，需要utf8 urlencode
	w.param.Set("dtype", "json")      //返回数据格式：json或xml,默认json
	w.param.Set("format", format)     //未来6天预报(future)两种返回格式，1或2，默认1
	w.param.Set("key", key)           //你申请的key

	return nil
}

func (w *JKWether) Query(location string) (*JKWetherInfo, error) {
	jklog.L().Debugf("wether goto query wether [%s]\n", location)
	// Query only when timeout.
	if time.Now().Unix()-w.lastGetTime < w.timeout {
		return &w.ResultW, nil
	}
	// generate param
	w.generateParam(location, "1", w.key)
	// 发送请求
	data, err := w.reqGet(w.url, w.param)
	if err != nil {
		return nil, err
	} else {
		err = json.Unmarshal(data, &w.ResultW)
		if err != nil {
			return nil, err
		}
		// Set last time only success.
		w.lastGetTime = time.Now().Unix()
		w.ResultW.GetTime = w.lastGetTime
	}
	return &w.ResultW, nil
}

func (w *JKWether) QueryCity() (*[]JKWetherCity, error) {

	return nil, nil
}
