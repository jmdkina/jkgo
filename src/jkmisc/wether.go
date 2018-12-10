package jkmisc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jk/jklog"
	"net/http"
	"net/url"
)

type JKWether struct {
	key   string
	url   string
	param url.Values

	Result  map[string]interface{}
	ResultW JKWetherInfo
}

func JKWetherNew(key string) (*JKWether, error) {
	return &JKWether{
		key:   key,
		url:   "http://v.juhe.cn/weather/index",
		param: url.Values{},
	}, nil
}

// get 网络请求
func (w *JKWether) reqGet(apiURL string, params url.Values) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		fmt.Printf("解析url错误:\r\n%v", err)
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	resp, err := http.Get(Url.String())
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// post 网络请求 ,params 是url.Values类型
func (w *JKWether) reqPost(apiURL string, params url.Values) (rs []byte, err error) {
	resp, err := http.PostForm(apiURL, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
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
		/*
			if w.Result["error_code"].(float64) == 0 {
				jklog.L().Infof("接口返回result字段是:\r\n%v\n", w.Result["result"])
			} else {
				jklog.L().Debugf("result value : \n%v\n", w.Result)
			}
		*/
	}
	return &w.ResultW, nil
}

func (w *JKWether) QueryCity() (*[]JKWetherCity, error) {

	return nil, nil
}
