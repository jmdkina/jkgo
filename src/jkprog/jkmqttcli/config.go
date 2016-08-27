package main

import (
	"encoding/json"
	"io/ioutil"
)

type mqtt_param struct {
	Broker    string `json:broker`
	User      string `json:user`
	Password  string `json:password`
	Id        string `json:id`
	Cleansess bool   `json:cleansess`
	Qos       int    `json:qos`
}

func init_config(path string) *mqtt_param {
	mp := &mqtt_param{}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(data, mp)
	if err != nil {
		return nil
	}
	return mp
}
