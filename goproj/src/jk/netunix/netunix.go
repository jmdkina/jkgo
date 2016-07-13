package netunix

import (
	"net"
)

type JKNetUnix struct {
	client net.Conn
}

func (nu *JKNetUnix) jk_net_unix_connect(path string) error {
	client, err := net.Dial("unix", path)
	if err != nil {
		return err
	}
	nu.client = client
	return nil
}

func (nu *JKNetUnix) jk_net_unix_close() error {
	nu.client.Close()
	return nil
}

func (nu *JKNetUnix) jk_net_unix_send(data []byte) (int, error) {
	n, err := nu.client.Write(data)
	if err != nil {
		return 0, err
	}
	return n, nil
}
