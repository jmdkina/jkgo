package jkclimng

import (
	"strings"
)

type JKClientConfig struct {
	UseMongo       bool
	UpdateInterval int64
}

type JKClientCtrl struct {
	Conf  JKClientConfig
	Items []JKClientItem
}

func JKClientCtrlNew(conf JKClientConfig) (*JKClientCtrl, error) {
	cc := &JKClientCtrl{}
	cc.Conf = conf
	return cc, nil
}

func (cc *JKClientCtrl) ItemAdd(item JKClientItem) error {
	cc.Items = append(cc.Items, item)
	return nil
}

func (cc *JKClientCtrl) ItemExist(item JKClientItem) (bool, *JKClientItem) {
	for k, v := range cc.Items {
		if strings.Compare(v.Id, item.Id) == 0 {
			return true, &cc.Items[k]
		}
	}
	return false, nil
}

func (cc *JKClientCtrl) ItemRemove(item JKClientItem) bool {
	for k, v := range cc.Items {
		if strings.Compare(v.Id, item.Id) == 0 {
			cc.Items = append(cc.Items[:k-1], cc.Items[k+1:]...)
			return true
		}
	}
	return false
}

func (cc *JKClientCtrl) ItemCounts() int {
	return len(cc.Items)
}

func (cc *JKClientCtrl) ItemCheckAndRemove() {
	for k, v := range cc.Items {
		if v.OfflineCheck(cc.Conf.UpdateInterval) {
			cc.Items = append(cc.Items[:k-1], cc.Items[k+1:]...)
		}
	}
}
