package jkcommon

import (
	"bytes"
	"strings"
	"testing"
	"io/ioutil"
	"os"
)

func TestJKFileLists(t *testing.T) {
	filepath := "/home/v/kflogs"

	files, err := JKFileLists(filepath, false, true)
	if err != nil {
		t.Fatalf("failed get file lists ", err)
	}

	if len(files) <= 0 {
		t.Fatalf("files less: ", len(files))
	}
	for _, v := range files {
		t.Logf("files: ", v)
	}
}

func TestJKWriteFileData(t *testing.T) {
	filename := "/tmp/test1"
	ioutil.WriteFile(filename, []byte("test1test2test3test4test5test6\n"), os.ModeType)
	err := JKSaveDataToFile(filename, []byte("test1test2test3test4test5test6\n"), true)
	if err != nil {
		t.Fatalf("error :", err)
	}
	filename = "/tmp/test2"
	ioutil.WriteFile(filename, []byte("test1\ntest2\n"), os.ModeType)
	err = JKSaveDataToFile(filename, []byte("test1\ntest2\n"), true)
	if err != nil {
		t.Fatalf("error:", err)
	}
}

func TestJKReadFileData(t *testing.T) {
	filename := "/tmp/test1"
	str, err := JKReadFileData(filename)
	strcmp := "test1test2test3test4test5test6\n"

	if err != nil || strings.Compare(str, strcmp) != 0 {
		t.Fatalf("need string %s, but real is %s, (%d)?=(%d)\n", strcmp, str, len(str), len(strcmp))
	}

	filename1 := "/tmp/test2"
	str, _ = JKReadFileData(filename1)
	if strings.Compare(str, "test1\ntest2\n") != 0 {
		t.Fatalf("need string test1\\ntest2, but real is %s\n", str)
	}
}

func TestInt32ToBytes(t *testing.T) {
	v := 12
	b := Int32ToBytes(int32(v))
	bold := []byte{0xc, 0, 0x0, 0}
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

	buf = []byte{0xc, 0x0, 0x0, 0x0}
	v1 := BytesToInt(buf)
	if int(v1) != v {
		t.Fatalf("need %v, but real is %v\n", v, v1)
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
