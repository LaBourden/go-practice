package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL: %v", r.URL.String())
	case "/hello":
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q]=%q", k, v)
		}
	default:
		fmt.Fprintf(w, "404 Not Found %v", r.URL.String())
	}
}

func main() {
	e := new(Engine)
	log.Fatal(http.ListenAndServe(":9000", e))
}
