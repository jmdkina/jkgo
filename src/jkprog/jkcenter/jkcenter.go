package main

import (
	"flag"
	"jk/jkcenter"
	"jk/jklog"
	"time"
)

var (
	address = flag.String("local_address", "0.0.0.0", "Listen local address")
	port    = flag.Int("local_port", 24433, "Listen local port")
)

func main() {
	flag.Parse()

	cc, err := jkcenter.InitCenter(*address, *port, 0)
	if err != nil {
		jklog.L().Errorln("Error init ", err)
		return
	}

	// thread doing
	err = cc.DoCycle()
	if err != nil {
		jklog.L().Errorln("do cycle failed ", err)
		return
	}

	// thread doing
	err = cc.Recv()
	if err != nil {
		jklog.L().Errorln("recv failed ", err)
		return
	}

	for {
		time.Sleep(500*time.Millisecond)
	}
}
