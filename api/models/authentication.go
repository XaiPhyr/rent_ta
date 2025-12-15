package models

import (
	"api/utils"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

type Authentication struct{}

func (m Authentication) Login(ctx *gin.Context, username string) (user User, err error) {
	var userPerm struct {
		User
		Permissions []string `bun:"permissions" json:"permissions"`
	}

	utils.GetPermissions(func(sq *bun.SelectQuery) *bun.SelectQuery {
		return sq.Where("username = ?", username).
			WhereOr("email = ?", username)
	}, ctx, &userPerm)

	user = userPerm.User
	user.Permissions = userPerm.Permissions

	return user, err
}
