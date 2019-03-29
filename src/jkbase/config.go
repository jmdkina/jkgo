package jkbase

import (
	"encoding/json"
	"io/ioutil"
)

func GetConfigInfo(filepath string, ci interface{}) error {
	d, e := ioutil.ReadFile(filepath)
	if e != nil {
		return e
	}
	err := json.Unmarshal(d, ci)
	if e != nil {
		return err
	}
	return nil
}

func CMConfigFile(filepath string, ci interface{}) error {
	return GetConfigInfo(filepath, ci)
}
