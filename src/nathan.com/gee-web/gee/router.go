package gee

import (
	"fmt"
	"log"
	"strings"

	"stream-player/src/nathan.com/gee-web/trie"
)

type node = trie.Node

/**
- 定义handlers map存储请求url和处理函数的对应关系
*/
type Router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	parts := make([]string, 0)
	for _, part := range strings.Split(pattern, "/") {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return parts
}

/**
添加路由：
- 通过method找到对应的分支（GET,POST），在前缀树种插入对应节点
- 构建k/v对，放入定义好的handlers map中
	key: method + "-" + pattern
	value: HandlerFunc
*/
func (r *Router) addRouter(method string, pattern string, handler HandlerFunc) {
	log.Printf("Router %4s - %s", method, pattern)
	key := method + "-" + pattern
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].Insert(pattern, parsePattern(pattern), 0)
	r.handlers[key] = handler
}

/**
对url路径进行解析，得到前缀树种对应的节点以及url中携带的动态参数map
*/
func (r *Router) getRouter(method string, path string) (*node, map[string]string) {
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	// 在前缀树中配对对应的路由规则
	node := root.Search(parsePattern(path), 0)

	if node == nil {
		return nil, nil
	}

	params := make(map[string]string)
	pathPart := parsePattern(path)

	parts := parsePattern(node.Pattern)

	// 得到url路径中携带的动态参数，存入params（map）中，（name/nathan）
	for index, part := range parts {
		if part[0] == ':' {
			params[part[1:]] = pathPart[index]
		}
		if part[0] == '*' && len(part) > 1 {
			params[part[1:]] = strings.Join(pathPart[index:], "/")

		}
	}

	return node, params

}

/**
处理用户的请求：
- 解析url，获取url对应的前缀树节点以及url中的动态路径参数/hello/:name
- enrich Context:
	- c.Params(map): 将动态路径参数放入map
	- c.handlers(array): 从node中找到匹配的路径，在router的路径map中找到对应的处理函数，放入context中的handlers数组中
*/
func (r *Router) handler(c *Context) {
	node, params := r.getRouter(c.Method, c.Path)

	if node != nil {
		c.Params = params
		key := c.Method + "-" + node.Pattern
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			data := fmt.Sprintf("404 NOT FOUND: %s\n", c.Path)
			c.SendStringResponse(404, data)
		})
	}

	// 开始根据context中的[]handlers 处理该请求的handler函数
	c.Next()

}
