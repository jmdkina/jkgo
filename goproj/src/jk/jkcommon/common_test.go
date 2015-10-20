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

func TestInt32ToBytes(t *testing.T) {
	v := 12
	b := Int32ToBytes(int32(v))
	bold := []byte{0xc, 0, 0, 0}
	if bytes.Compare(bold, b) != 0 {
		t.Fatalf("need %v, but real is %v\n", bold, b)
	}
}

func TestBytesToInt32(t *testing.T) {

	buf := []byte{12, 0, 0, 0}
	v1 := BytesToInt32(buf)
	if v1 != 12 {
		t.Fatalf("need %d, but real is %d\n", 11, v1)
	}
}

func TestBytesInt(t *testing.T) {
	v := 6
	buf := IntToBytes(int64(v), 4)
	obuf := []byte{12, 0, 0, 0}
	if bytes.Compare(buf, obuf) != 0 {
		t.Fatalf("need %v, buf real is %v\n", obuf, buf)
	}

	buf = []byte{22, 0, 0, 0}
	v1 := BytesToInt(buf)
	if int(v1) != v {
		t.Fatalf("need %d, but real is %d\n", v, v1)
	}
}

func TestWriteData(t *testing.T) {
	filename := "test12"
	data := "test12test12"
	ret := JKSaveFileData("123", filename, data)
	if !ret {
		t.Fatalf("save data failed. \n")
	}
}
