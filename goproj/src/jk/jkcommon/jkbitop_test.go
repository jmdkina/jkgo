package jkcommon

import (
	"testing"
)

func TestJKBitValue(t *testing.T) {
	testvalue := uint(15)
	ret := JKBitValue(testvalue, 5, 3)
	if ret != 5 {
		t.Fatalf("Error ret %d, expect %d\n", ret, 5)
	}
}
