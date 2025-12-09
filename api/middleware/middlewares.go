package middleware

import (
	"api/models"
	"api/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
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

func (m Middleware) CheckPermission(module string, perm ...string) gin.HandlerFunc {
	var UserPermission struct {
		models.User
		HasPermission bool `bun:"has_permissions"`
	}

	return func(ctx *gin.Context) {
		items := []string{}

		for _, v := range perm {
			items = append(items, module+":"+v)
		}

		q := utils.GetPermissions().
			ColumnExpr("u.*, CASE WHEN array_length(ARRAY_AGG(p.name), 1) > 0 THEN TRUE ELSE FALSE END AS has_permissions").
			Where("u.uuid = ?", ctx.GetHeader("UserUUID")).
			Where("p.name IN (?)", bun.In(items)).
			WhereOr("is_admin = true").
			Group("u.id")

		if err := q.Scan(ctx, &UserPermission); err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: You do not have permission."})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
