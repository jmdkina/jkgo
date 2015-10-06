package jknewserver

import (
	// "jk/jklog"
	"bytes"
	"testing"
)

func TestAddItem(t *testing.T) {
	item := &JKServerProcessItem{}
	// item.data = []byte("it is")
	item.id = "123456"
	item.remoteAddr = "192.168.31.111"

	proc := JKServerProcess{}
	proc.addItem(item)

	item.data = []byte("it is")

	if len(proc.Item) != 1 {
		t.Fatalf("need len 1, but real is %d\n", len(proc.Item))
	}

	for k, v := range proc.Item {
		if k == 0 {
			if v.id != "123456" {
				t.Fatalf("need id 123456, but real is %d\n", v.id)
			}
			if v.remoteAddr != "192.168.31.111" {
				t.Fatalf("need remoteAddr is 192.168.31.111, but real is %d\n", v.remoteAddr)
			}

			if bytes.Compare([]byte("it is"), v.data) != 0 {
				t.Fatalf("need data is %v, buf real is %v\n", []byte("it is"), v.data)
			}
			// jklog.L().Infoln(k, ", id: ", v.id, ", remoteaddr: ", v.remoteAddr, ", data: ", v.data)
		}
	}

	item2 := &JKServerProcessItem{}
	item2.id = "1233"
	item2.remoteAddr = "192.168.31.111" // save remote addr
	proc.addItem(item2)

	item2.data = []byte("ok,find")

	if len(proc.Item) != 1 {
		t.Fatalf("need len 1, but real is %d\n", len(proc.Item))
	}
	for k, v := range proc.Item {
		if k == 0 {
			if bytes.Compare([]byte("ok,find"), v.data) != 0 {
				t.Fatalf("need data is %v, buf real is %v\n", []byte("it is"), v.data)
			}
		}
	}

	// third test
	item3 := &JKServerProcessItem{}
	item3.id = "3333"
	item3.remoteAddr = "192.168.31.112"
	proc.addItem(item3)

	item3.data = []byte("this is item3")

	if len(proc.Item) != 2 {
		t.Fatalf("need len 2 , but real is %d\n", len(proc.Item))
	}
	for k, v := range proc.Item {
		if k == 1 {
			if bytes.Compare([]byte("this is item3"), v.data) != 0 {
				t.Fatalf("need data is %v, buf real is %v\n", []byte("it is"), v.data)
			}
		}
	}

}
