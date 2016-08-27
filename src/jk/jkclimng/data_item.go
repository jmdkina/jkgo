package jkclimng

// Get data from database

import (
	"golanger.com/middleware"
	"golanger.com/utils"
	"helper"
	"labix.org/v2/mgo"
	"errors"
)

type DataControl struct {
	ColItem *mgo.Collection
}

var DC *DataControl

func (dc *DataControl) CItem(name string) *mgo.Collection {
	if dc.ColItem != nil {
		return dc.ColItem
	}
	m := middleware.Middleware.Get("mrzy").(*helper.Mongo)
	dc.ColItem = m.C(utils.M{"name": name})
	return dc.ColItem
}

func GlobalDataControl() *DataControl {
	if DC != nil {
		return DC
	}
	mServer := helper.NewMongo("mongodb://127.0.0.1/mryz")
	//defer mServer.Close()
	middleware.Middleware.Add("mrzy", mServer)
	DC = &DataControl{}
	return DC
}

func (dc *DataControl) GetItemGen(colname string, count int, m interface{}) error {
	if ok := dc.CItem(colname).Find(nil).Limit(count).All(m); ok != nil {
		return errors.New("find fail")
	}
	return nil
}

func (dc *DataControl) GetItem(count int) *[]MRYZItem {
	ret := &[]MRYZItem{}
	err := dc.GetItemGen("item", count, ret)
	if err != nil {
		return nil
	}
	return ret
}
