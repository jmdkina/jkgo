package main

import (
	"jk/jksip"
	"time"
)

func main() {
	ss, _ := jksip.NewJKSipServer()
	ss.Recv()

	for {
		time.Sleep(500 * time.Microsecond)
	}
}
