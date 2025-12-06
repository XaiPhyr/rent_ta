package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type Middleware struct{}

func SetLoggers(router *gin.Engine) {
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.Use(gin.Recovery())
}

func (m Middleware) Authenticate(ctx *gin.Context) {
	ctx.Set("userId", 1)
	ctx.Next()
}
