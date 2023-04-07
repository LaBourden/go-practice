package gee

import (
	"log"
	"net/http"
)

type HandleFunc func(ctx *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) AddRoute(method string, path string, handler HandleFunc) {
	e.router.addRoute(method, path, handler)
}

func (e *Engine) GET(path string, handler HandleFunc) {
	e.AddRoute("GET", path, handler)
}

func (e *Engine) POST(path string, handler HandleFunc) {
	e.AddRoute("POST", path, handler)
}

func (e *Engine) PUT(path string, handler HandleFunc) {
	e.AddRoute("PUT", path, handler)
}

func (e *Engine) PATCH(path string, handler HandleFunc) {
	e.AddRoute("PATCH", path, handler)
}

func (e *Engine) DELETE(path string, handler HandleFunc) {
	e.AddRoute("DELETE", path, handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := newContext(w, req)
	e.router.handle(ctx)
}

func (e *Engine) RUN(addr string) {
	log.Fatal(http.ListenAndServe(addr, e))
}
