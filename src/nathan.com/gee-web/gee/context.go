package gee

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type H map[string]interface{}

/**
一个web请求->根据*http.Request构造一个http.ResponseWriter进行返回，构造Context对用户的请求和返回进行封装：
- 剥离出请求中的Path和Method，以及请求中携带的Params
- 对返回体的 StatusCode 和 Content-Type 进行封装，实现对不同类型结构的返回
*/
type Context struct {
	Response     http.ResponseWriter
	Request      *http.Request
	Path, Method string
	StatusCode   int
	Params       map[string]string
	handlers     []HandlerFunc // middlewares + handlers
	index        int
	engine       *Engine
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

func (c *Context) SendStringResponse(code int, data string) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	_, _ = c.Response.Write([]byte(data))
}

func (c *Context) SendJsonResponse(code int, object interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Response)
	if err := encoder.Encode(object); err != nil {
		http.Error(c.Response, err.Error(), 500)
	}
}

func (c *Context) SendHTMLResponse(code int, htmlFileName string, data interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Response, htmlFileName, data); err != nil {
		c.Fail(500, err.Error())
	}
}

func (c *Context) SendErrorResponse(code int, errMsg string) {
	c.SendStringResponse(code, errMsg)
}

func (c *Context) SendVideoMP4Response(video *os.File) {
	c.SetHeader("Content-Type", "video/mp4")
	http.ServeContent(c.Response, c.Request, "", time.Now(), video)
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers) - 1
	c.SendJsonResponse(code, H{"message": err})
}

// aspect
func (c *Context) Next() {
	c.index++
	for ; c.index < len(c.handlers); c.index++ {
		c.handlers[c.index](c)
	}
}
