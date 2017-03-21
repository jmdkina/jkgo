package main

import (
	"github.com/henrylee2cn/faygo"
	"jkprog/snaky/router"
)

func main() {
	router.Route(faygo.New("snaky"))
	faygo.Run()
}
