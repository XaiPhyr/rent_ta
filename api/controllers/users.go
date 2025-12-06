package controllers

import (
	"api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun/driver/pgdriver"
)

type UserController struct {
	AppController
	u models.User
}

func (c UserController) InitUserController(router *gin.Engine) {
	r := router.Group(fmt.Sprintf("/%s/user", apiVersion))

	r.POST("", c.mw.Authenticate, c.Upsert)
	r.GET("/:uuid", c.mw.Authenticate, c.Read)
	r.DELETE("/:uuid", c.mw.Authenticate, c.Delete)
}

func (c UserController) Upsert(ctx *gin.Context) {
	form := c.u.ParseUser(ctx)

	if err := ctx.ShouldBindJSON(form); err == nil {
		res, err := c.u.Upsert(ctx, *form)

		if err != nil {
			c.handleError(ctx, err, c.cleanErr(err))
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"data": res})
	} else {
		c.handleError(ctx, err, c.cleanErr(err))
	}
}

func (c UserController) Read(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	if res, err := c.u.Read(c.sanitizeCtx(ctx)); err == nil {
		data := gin.H{"total": res.Count, "data": res.Users}

		if uuid != "all" {
			data = gin.H{"data": res.User}
		}

		ctx.JSON(http.StatusOK, data)
	} else {
		c.handleError(ctx, err, c.cleanErr(err))
	}
}

func (c UserController) Delete(ctx *gin.Context) {
	uuid := ctx.Param("uuid")

	if user, err := c.u.Delete(ctx, uuid); err == nil {
		msg := "User deleted successfully"

		if user.DeletedAt.IsZero() {
			msg = "User restored successfully"
		}

		ctx.JSON(http.StatusOK, gin.H{"deleted_at": user.DeletedAt, "message": msg})
	} else {
		c.handleError(ctx, err, c.cleanErr(err))
	}
}

func (c UserController) cleanErr(err error) string {
	if pgErr, ok := err.(pgdriver.Error); ok {
		return fmt.Sprintf("Postgres Error: %s", pgErr.Field('M'))
	}

	return fmt.Sprintf("Unknown Error: %s", err.Error())
}
