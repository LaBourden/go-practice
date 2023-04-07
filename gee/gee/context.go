package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// original data
	Writer  http.ResponseWriter
	Request *http.Request
	// request
	Method string
	Path   string
	// response
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: req,
		Method:  req.Method,
		Path:    req.URL.Path,
	}
}

func (ctx *Context) PostForm(key string) string {
	return ctx.Request.FormValue(key)
}

func (ctx *Context) Query(key string) string {
	return ctx.Request.URL.Query().Get(key)
}

func (ctx *Context) Status(code int) {
	ctx.StatusCode = code
	ctx.Writer.WriteHeader(code)
}

func (ctx *Context) SetHeader(key string, value string) {
	ctx.Writer.Header().Set(key, value)
}

func (ctx *Context) String(code int, format string, value ...any) {
	ctx.SetHeader("Content-Type", "text/plain")
	ctx.Status(code)
	ctx.Writer.Write([]byte(fmt.Sprintf(format, value...)))
}

func (ctx *Context) JSON(code int, obj interface{}) {
	ctx.SetHeader("Content-Type", "application/json")
	ctx.Status(code)
	encoder := json.NewEncoder(ctx.Writer)
	err := encoder.Encode(obj)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (ctx *Context) Data(code int, data []byte) {
	ctx.Status(code)
	ctx.Writer.Write(data)
}

func (ctx *Context) HTML(code int, html string) {
	ctx.Status(code)
	ctx.SetHeader("Content-Type", "text/html")
	ctx.Writer.Write([]byte(html))
}
