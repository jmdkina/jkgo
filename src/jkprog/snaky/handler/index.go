package handler

import (
	"github.com/henrylee2cn/faygo"
)

/*
Index
*/
var Index = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	return ctx.Render(200, faygo.JoinStatic("index.html"), faygo.Map{
		"TITLE":   "faygo",
		"VERSION": faygo.VERSION,
		"CONTENT": "Welcome To Faygo",
		"AUTHOR":  "HenryLee",
	})
})
