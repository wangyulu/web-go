package middleware

import (
	"github.com/wangyulu/web-go/framework/contract"
	"github.com/wangyulu/web-go/framework/gin"
)

// trace机制，将请求中的traceID保存

func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		tracer := c.MustMake(contract.TraceKey).(contract.Trace)

		traceCtx := tracer.ExtractHTTP(c.Request)

		tracer.WithTrace(c, traceCtx)

		// 使用next执行具体的业务逻辑
		c.Next()
	}
}
