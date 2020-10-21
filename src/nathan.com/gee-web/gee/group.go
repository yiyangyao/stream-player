package gee

/**
为路由router提供分组控制，以相同的前缀来区分不同的分组
*/
type RouterGroup struct {
	prefix   string        // 以相同的前缀来区分不同的分组
	meddlers []HandlerFunc // 分组支持中间件
	engine   *Engine       // 所有的RouterGroup共享同一个Engine实例
}

/**
创建分组：
- 将RouterGroup.engine指向全局唯一的Engine，这样可以通过engine间接访问到各种接口
*/
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

/**
添加路由规则：将part和分组的前缀拼接到一起，通过engine间接的调用router.addRouter()方法
*/
func (group *RouterGroup) addRouterRule(method string, part string, handler HandlerFunc) {
	pattern := group.prefix + part
	group.engine.router.addRouter(method, pattern, handler)

}

/**
对该分组应用中间件
*/
func (group *RouterGroup) Apply(middlewares ...HandlerFunc) {
	group.meddlers = append(group.meddlers, middlewares...)
}
