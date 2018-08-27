package jkstock

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type StockInfo struct {
	Area   string  
}

type StockParse struct {
	filename  string
	data      string
	data_json   []interface{}
}

func NewStockParse(filename string) (*StockParse, error) {
	sp := &StockParse{
		filename: filename,
	}
	d, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(d, &sp.data_json)
	if err != nil {
		return nil, err
	}
	
	return sp, nil
}

func (sp *StockParse) ParseArea() error {

	return nil
}

func (sp *StockParse) DebugOut() {
	for i, k := range sp.data_json {
		if i > 10 {
			break
		}
		fmt.Println(k)
	}
}
