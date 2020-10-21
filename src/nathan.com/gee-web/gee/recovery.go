package gee

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

/**
print stack trace for debug，用来获取触发 panic 的堆栈信息
*/
func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

/**
Recovery() 作为全局的中间件，预防系统触发了某些异常，例如数组越界，空指针等，导致系统宕机
recover函数可以避免因为panic发生而导致整个程序终止，recover函数只在defer中生效
*/
func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()

		c.Next()
	}
}
