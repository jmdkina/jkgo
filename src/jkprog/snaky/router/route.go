package router

import (
	"github.com/henrylee2cn/faygo"
	"jkprog/snaky/handler"
	"jkprog/snaky/middleware"
)

// Route register router in a tree style.
func Route(frame *faygo.Framework) {
	frame.Route(
		frame.NewNamedAPI("Index", "GET", "/index", handler.Index),
		frame.NewNamedAPI("Snaky", "GET", "/", handler.Snaky),
		frame.NewNamedAPI("test struct handler", "POST", "/test", &handler.Test{}).Use(middleware.Token),
	)
}
