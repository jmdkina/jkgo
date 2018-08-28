package jkstock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type StockInfo struct {
	Area string
}

type StockParse struct {
	filename  string
	data      string
	data_json []StockInfo

	AreaCollection map[string][]StockInfo
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
	sp.AreaCollection = make(map[string][]StockInfo)

	return sp, nil
}

func (sp *StockParse) ParseArea() error {
	for _, k := range sp.data_json {
		si := k
		area := si.Area
		sp.AreaCollection[area] = append(sp.AreaCollection[area], si)
	}
	return nil
}

func (sp *StockParse) DebugOut() {
	for k, v := range sp.AreaCollection {
		fmt.Printf("area [%s] len [%d]\n", k, len(v))
	}
}
