package middleware

import (
	"api/models"
	"api/utils"
	"fmt"
	"net/http"
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
	// @todo JWT LOGIC
	ctx.Set("userId", 1)
	ctx.Next()
}

func (m Middleware) CheckPermission(module string, perm ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		items := make([]string, len(perm))
		for i, v := range perm {
			items[i] = module + ":" + v
		}

		var userPerm struct {
			models.User
			Permissions []string `bun:"permissions" json:"permissions"`
		}

		err := utils.GetPermissions(ctx.GetHeader("UserUUID"), "u.*", "u.id", ctx, &userPerm)

		if err != nil {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "User does not exist.", "details": err.Error()})
			ctx.Abort()
			return
		}

		if !m.checkPerm(items, userPerm.Permissions, userPerm.IsAdmin) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: You do not have permission.", "user": userPerm})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func (m Middleware) checkPerm(items, permissions []string, isAdmin bool) bool {
	if isAdmin {
		return true
	}

	permissionMap := make(map[string]struct{})
	for _, p := range permissions {
		permissionMap[p] = struct{}{}
	}

	for _, item := range items {
		if _, exists := permissionMap[item]; exists {
			return true
		}
	}

	return false
}
