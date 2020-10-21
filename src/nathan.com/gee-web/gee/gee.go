package gee

import (
	"net/http"
	"strings"
)

/**
定义函数类型的类型 HandlerFunc，用于接收Context类型的变量
*/
type HandlerFunc func(*Context)

/**
- Engine：web应用的入口，全局唯一
*/
type Engine struct {
	router       *Router        // 该app下的路由规则和前缀树
	groups       []*RouterGroup // 改engine下所包含的所有分组
	*RouterGroup                // 继承RouterGroup，engine作为子类，拥有更多的方法和属性
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

/**
-------main--------
http.ListenAndServe(addr, engine)
- addr: 监听的地址和端口
- engine：Handler接口，需要实现方法 ServeHTTP ，只要传入任何实现了 ServerHTTP 接口的实例，所有的HTTP请求，就都交给了该实例处理了
	type Handler interface {
		ServeHTTP(ResponseWriter, *Request)
	}
*/

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

/**
-------request coming to excuse the method--------
Engine 结构体实现Handler接口的ServeHTTP方法，用于处理传入的http请求
- 1）解析url，遍历engine下所有的分组，匹配到该请求对应的group,子分组属于父分组
- 2）将请求对应的group中已经apply的中间件提取出来
- 3）根据请求创建对应的Context，并将该请求应用的中间件加入到c.handlers中
- 4）执行router的handler，开始正式处理请求
*/
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
