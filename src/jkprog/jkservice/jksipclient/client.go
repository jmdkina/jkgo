package main

import (
	"jk/jksip"
	"time"
)

func main() {
	sc, _ := jksip.NewJKSipClient()

	sc.Send()

	for {
		time.Sleep(500 * time.Microsecond)
	}
}
