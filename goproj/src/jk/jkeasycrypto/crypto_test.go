package jkeasycrypto

import (
	"strings"
	"testing"
)

func TestJKCrypt(t *testing.T) {
	oldstr := "This is a test for the data ''? fined why. ok"
	key := "32bitstringtofores/-asel"
	newstr := JKAESEncrypt([]byte(key), []byte(oldstr))
	t.Log("new string is : ", newstr)

	resStr := JKAESDecrypt([]byte(key), newstr)
	t.Log("restore string is : ", resStr)

	if strings.Compare(oldstr, resStr) != 0 {
		t.Errorf("Error with newstr: %s, restore string %s", newstr, resStr)
	}
}
