package netunix

import (
	"testing"
	"time"
)

func TestNetUnixProcess(t *testing.T) {
	nu := &JKNetUnix{}
	err := nu.jk_net_unix_connect("/tmp/kfupload.sock")
	if err != nil {
		t.Fatal("Erorr of unix socket: ", err)
	}
	jsonData := "{\"title\":\"value\", \"nu\":8}"
	n, err := nu.jk_net_unix_send([]byte(jsonData))
	if err != nil {
		t.Fatal("Error of unix socket send: ", err)
	}

	time.Sleep(2000 * time.Millisecond)

	_, err = nu.jk_net_unix_send([]byte(jsonData))
	if err != nil {
		t.Fatal("Error of unix socket send: ", err)
	}
	nu.jk_net_unix_close()
	t.Logf("write done %d\n", n)
}
