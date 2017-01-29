package main

import (
	"os"
	"jk/jklog"
	"crypto/md5"
	"fmt"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		jklog.L().Errorf("Usage: %s string\n", os.Args[0])
		return
	}

	retv := md5.Sum([]byte(os.Args[1]))
	for i := 0; i < len(retv); i++ {
	    fmt.Printf("%02x", retv[i])
	}
	fmt.Printf("\n")
	return
}
