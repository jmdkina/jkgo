package jkeasycrypto

import (
	"strings"
	"testing"
)

func TestJKCrypt(t *testing.T) {
	oldstr := "This is a test for the data ''? fined why. okaes"
	key := "32bitstringtofor"
	newstr := JKAESEncrypt([]byte(key), []byte(oldstr))
	t.Log("new string is : ", newstr)

	// newstr = "LqJmqpy/n1z+tzIDZZKC5HwmoY0uiRI69fzio7hZmGWSu3m+PHUlc8ulB9HP4TPJGTjbxNvaoLzUOMOwTz9FaQ=="
	// newstr = "YWv+OsFABP6NXfCC/AoVl+ulBSnuGeM+2ggMQAupjapazcMhlw0wB/lx2MW4BNVkHNRkW9M+UzFiz75vOTgB7A=="
	// newstr = "wP3KMRLcCZkZG3ObibCvjhJuYXJWhgO3MQzIkikgbu19PDecsiIa64WxS8/zLRBLdOM1a5CBmCXMTF/PEl7g6A=="
	resStr := JKAESDecrypt([]byte(key), newstr)
	t.Log("restore string is : ", resStr)

	if strings.Compare(oldstr, resStr) != 0 {
		t.Errorf("Error with newstr: %s, restore string: %s", newstr, resStr)
	}
	t.Fatal("error")
}
