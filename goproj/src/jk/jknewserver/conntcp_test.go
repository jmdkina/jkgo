package jknewserver

import (
	// "jk/jklog"
	"bytes"
	"testing"
)

func TestAddItem(t *testing.T) {
	item := &JKServerProcessItem{}
	// item.data = []byte("it is")
	item.Id = "123456"
	item.RemoteAddr = "192.168.31.111"

	proc := JKServerProcess{}
	proc.addItem(item)

	item.Data = []byte("it is")

	if len(proc.Item) != 1 {
		t.Fatalf("need len 1, but real is %d\n", len(proc.Item))
	}

	for k, v := range proc.Item {
		if k == 0 {
			if v.Id != "123456" {
				t.Fatalf("need id 123456, but real is %d\n", v.Id)
			}
			if v.RemoteAddr != "192.168.31.111" {
				t.Fatalf("need remoteAddr is 192.168.31.111, but real is %d\n", v.RemoteAddr)
			}

			if bytes.Compare([]byte("it is"), v.Data) != 0 {
				t.Fatalf("need data is %v, buf real is %v\n", []byte("it is"), v.Data)
			}
			// jklog.L().Infoln(k, ", id: ", v.id, ", remoteaddr: ", v.remoteAddr, ", data: ", v.data)
		}
	}

	item2 := &JKServerProcessItem{}
	item2.Id = "1233"
	item2.RemoteAddr = "192.168.31.111" // save remote addr
	proc.addItem(item2)

	item2.Data = []byte("ok,find")

	if len(proc.Item) != 1 {
		t.Fatalf("need len 1, but real is %d\n", len(proc.Item))
	}
	for k, v := range proc.Item {
		if k == 0 {
			if bytes.Compare([]byte("ok,find"), v.Data) != 0 {
				t.Fatalf("need data is %v, buf real is %v\n", []byte("it is"), v.Data)
			}
		}
	}

	// third test
	item3 := &JKServerProcessItem{}
	item3.Id = "3333"
	item3.RemoteAddr = "192.168.31.112"
	proc.addItem(item3)

	item3.Data = []byte("this is item3")

	if len(proc.Item) != 2 {
		t.Fatalf("need len 2 , but real is %d\n", len(proc.Item))
	}
	for k, v := range proc.Item {
		if k == 1 {
			if bytes.Compare([]byte("this is item3"), v.Data) != 0 {
				t.Fatalf("need data is %v, buf real is %v\n", []byte("it is"), v.Data)
			}
		}
	}

}
