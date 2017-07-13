package simpleserver

import (
	"net/http"
	"fmt"
	"golang.org/x/net/html"
)

type Base struct {
}

func (b *Base) ServeHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Base: %q", html.EscapeString(r.URL.Path))
}
