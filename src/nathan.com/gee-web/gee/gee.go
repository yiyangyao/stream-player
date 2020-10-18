package gee

import (
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)

type Engine struct {
	router *Router
	groups []*RouterGroup
	*RouterGroup
}

func Default() *Engine {
	engine := New()
	engine.Apply(Logger(), Recovery())
	return engine
}

func New() *Engine {
	engine := &Engine{router: NewRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// request coming to excuse the method
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.meddlers...)
		}
	}
	c := newContext(w, r)
	c.handlers = middlewares
	engine.router.handler(c)
}
