package jkcommon

import (
	"strings"
	"testing"
)

func TestJKReadFileData(t *testing.T) {
	filename := "/tmp/test1"
	str, _ := JKReadFileData(filename)
	strcmp := "test1"
	if strings.Compare(str, strcmp) != 0 {
		t.Fatalf("need string %s, but real is %s\n", strcmp, str)
	}

	filename1 := "/tmp/test2"
	str, _ = JKReadFileData(filename1)
	if strings.Compare(str, "test1\ntest2") != 0 {
		t.Fatalf("need string test1\\ntest2, but real is %s\n", str)
	}
}
