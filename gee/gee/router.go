package gee

import (
	"fmt"
	"log"
)

type router struct {
	handler map[string]HandleFunc
}

func newRouter() *router {
	return &router{handler: make(map[string]HandleFunc)}
}

func (r *router) addRoute(method string, path string, handler HandleFunc) {
	log.Printf("Route %4s - %s", method, path)
	key := method + "-" + path
	r.handler[key] = handler
}

func (r *router) handle(ctx *Context) {
	key := ctx.Method + "-" + ctx.Path
	if handler, ok := r.handler[key]; ok {
		handler(ctx.Writer, ctx.Request)
	} else {
		fmt.Fprintf(ctx.Writer, "404 Not Found %s \n", ctx.Request.URL.String())
	}
}
