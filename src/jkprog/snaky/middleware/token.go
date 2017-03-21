package middleware

import (
	"github.com/henrylee2cn/faygo"
)

/*
Token
*/
var Token = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	ctx.Log().Debugf("[ware] token:%q", ctx.QueryParam("token"))
	return nil
})
