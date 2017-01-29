package main

import (
	"os"
	"fmt"
	"jk/kfmd5"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("Usage: %s string\n", os.Args[0])
		return
	}

	retv := kfmd5.Sum([]byte(os.Args[1]))
	for i := 0; i < len(retv); i++ {
		fmt.Printf("%02x", retv[i])
	}
	fmt.Printf("\n")
}
