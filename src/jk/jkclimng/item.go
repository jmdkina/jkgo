package jkclimng

import (
	"time"
)

type JKClientItem struct {
	Id          string
	OnlineTime  int64
	OfflineTime int64
	UpdateTime  int64
}

func JKClientItemNew(id string) (*JKClientItem, error) {
	return &JKClientItem{
		Id:          id,
		OnlineTime:  time.Now().Unix(),
		OfflineTime: 0,
		UpdateTime:  time.Now().Unix(),
	}, nil
}

func (item *JKClientItem) UpdateOnline() error {
	item.OnlineTime = time.Now().Unix()
	return nil
}

func (item *JKClientItem) UpdateOffline() error {
	item.OfflineTime = time.Now().Unix()
	return nil
}

func (item *JKClientItem) UpdateUp() error {
	item.UpdateTime = time.Now().Unix()
	return nil
}

// Check if item has offline @interval time.
func (item *JKClientItem) OfflineCheck(interval int64) bool {
	if time.Now().Unix()-item.UpdateTime > interval {
		return true
	}
	return false
}
