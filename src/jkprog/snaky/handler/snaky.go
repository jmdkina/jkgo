package handler

import "github.com/henrylee2cn/faygo"

var Snaky = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	return ctx.Render(200, faygo.JoinStatic("snaky.html"), faygo.Map{
		"TITLE":   "Snaky",
		"VERSION": faygo.VERSION,
		"CONTENT": "Welcome To Snaky",
		"AUTHOR":  "Jmdvirus",
	})
})