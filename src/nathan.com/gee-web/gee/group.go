package gee

type RouterGroup struct {
	prefix   string
	meddlers []HandlerFunc
	engine   *Engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRouterRule("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRouterRule("POST", pattern, handler)
}

func (group *RouterGroup) addRouterRule(method string, part string, handler HandlerFunc) {
	pattern := group.prefix + part
	group.engine.router.addRouter(method, pattern, handler)

}

func (group *RouterGroup) Apply(middlewares ...HandlerFunc) {
	group.meddlers = append(group.meddlers, middlewares...)
}
