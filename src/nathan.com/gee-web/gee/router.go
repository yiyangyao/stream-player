package gee

import (
	"fmt"
	"log"
	"strings"

	"stream-player/src/nathan.com/gee-web/trie"
)

type node = trie.Node

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

func (r *Router) addRouter(method string, pattern string, handler HandlerFunc) {
	log.Printf("Router %4s - %s", method, pattern)
	key := method + "-" + pattern
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].Insert(pattern, parsePattern(pattern), 0)
	r.handlers[key] = handler
}

func (r *Router) getRouter(method string, path string) (*node, map[string]string) {
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	node := root.Search(parsePattern(path), 0)

	if node == nil {
		return nil, nil
	}

	params := make(map[string]string)
	pathPart := parsePattern(path)

	parts := parsePattern(node.Pattern)
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

func (r *Router) handler(c *Context) {
	node, params := r.getRouter(c.Method, c.Path)

	if node != nil {
		c.Params = params
		key := c.Method + "-" + node.Pattern
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			data := fmt.Sprintf("404 NOT FOUND: %s\n", c.Path)
			c.String(404, data)
		})
	}

	c.Next()

}
