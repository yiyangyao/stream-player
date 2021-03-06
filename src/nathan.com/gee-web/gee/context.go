package gee

import (
	"encoding/json"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Response     http.ResponseWriter
	Request      *http.Request
	Path, Method string
	StatusCode   int
	Params       map[string]string
	handlers     []HandlerFunc // middlewares + handlers
	index        int
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Response: w,
		Request:  r,
		Path:     r.URL.Path,
		Method:   r.Method,
		index:    -1,
	}
}

func (c *Context) GetParamValue(key string) string {
	value, _ := c.Params[key]
	return value
}

func (c *Context) GetPostFormValue(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) GetQueryValue(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Response.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Response.Header().Set(key, value)
}

func (c *Context) String(code int, data string) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Response.Write([]byte(data))
}

func (c *Context) Json(code int, object interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Response)
	if err := encoder.Encode(object); err != nil {
		http.Error(c.Response, err.Error(), 500)
	}
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Response.Write([]byte(html))
}

// aspect
func (c *Context) Next() {
	c.index++
	for ; c.index < len(c.handlers); c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers) - 1
	c.Json(code, H{"message": err})
}
