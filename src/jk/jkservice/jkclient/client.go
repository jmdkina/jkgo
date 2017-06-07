package jkclient

import (
	"jk/jknetbase"
)

type DemoClient struct {
	jknetbase.JKNetBaseClient
}

func NewDemoClient(addr string, port int) (*DemoClient, error) {
	c := &DemoClient{}
	err := c.New(addr, port, 1)
	if err != nil {
		return nil, err
	}
	return c, nil
}
