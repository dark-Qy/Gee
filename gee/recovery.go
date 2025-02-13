package gee

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

func trace(message string) string {
	var stack [32]uintptr
	// skip为0时，表示runtime.Callers调用者的栈帧，1表示调用者的调用者的栈帧
	length := runtime.Callers(3, stack[:])
	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range stack[:length] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				// 恢复失败
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.String(500, "Internal Server Error")
			}
		}()

		c.Next()
	}
}
