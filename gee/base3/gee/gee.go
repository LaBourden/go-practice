package gee

import (
	"fmt"
	"log"
	"net/http"
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandleFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

func (e *Engine) AddRoute(method string, path string, handler HandleFunc) {
	e.router[method+"-"+path] = handler
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
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, ok := e.router[r.Method+"-"+r.URL.Path]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 Not Found %s \n", r.URL.String())
	}
}

func (e *Engine) RUN(addr string) {
	log.Fatal(http.ListenAndServe(addr, e))
}
