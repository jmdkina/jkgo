package jkcommon

import (
	"bytes"
	"strings"
	"testing"
)

func TestJKReadFileData(t *testing.T) {
	filename := "/tmp/test1"
	str, err := JKReadFileData(filename)
	strcmp := "test1test2test3test4test5test6"

	if err != nil || strings.Compare(str, strcmp) != 0 {
		t.Fatalf("need string %s, but real is %s, (%d)?=(%d)\n", strcmp, str, len(str), len(strcmp))
	}

	filename1 := "/tmp/test2"
	str, _ = JKReadFileData(filename1)
	if strings.Compare(str, "test1\ntest2") != 0 {
		t.Fatalf("need string test1\\ntest2, but real is %s\n", str)
	}
}

func TestBytesInt(t *testing.T) {
	v := 231
	buf := BytesToInt(int64(v), 4)
	obuf := []byte{206, 3, 0, 0}
	if bytes.Compare(buf, obuf) != 0 {
		t.Fatalf("need %v, buf real is %v\n", obuf, buf)
	}

	v1 := IntToBytes(buf)
	if int(v1) != v {
		t.Fatalf("need %d, but real is %d\n", v, v1)
	}
}
