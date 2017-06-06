package jkclient

import "jk/jknetbase"

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

func (c *DemoClient) Send(data string) int {
	return c.Send(data)
}

func (c *DemoClient) Recv() (string, error) {
	return c.Recv()
}