package controllers

import (
	"api/middleware"
	"api/utils"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun/driver/pgdriver"
)

type AppController struct {
	mw middleware.Middleware
}

var (
	apiVersion = utils.InitConfig().Server.Endpoint
)

func (c *AppController) sanitizeCtx(ctx *gin.Context) (ginCtx *gin.Context, uuid, filter, filterExt, sort, status string, limitInt, pageInt int) {
	ginCtx = ctx
	uuid = ctx.Param("uuid")
	sort = ctx.Query("sort")
	limit := ctx.Query("limit")
	page := ctx.Query("page")
	filter = ctx.Query("filter")
	status = ctx.Query("status")

	// QueryUnescape for filterExt
	qUnescape, err := url.QueryUnescape(ctx.Query("filterExt"))
	if err != nil {
		log.Printf("Error Query Unescape: %s", err)
		filterExt = ""
	} else {
		filterExt = qUnescape
	}

	limitInt, err = strconv.Atoi(limit)
	if err != nil || limitInt <= 0 {
		limitInt = 10
	}

	pageInt, err = strconv.Atoi(page)
	if err != nil || pageInt <= 0 {
		pageInt = 1
	}

	return
}

func (c *AppController) cleanErr(err error) string {
	if pgErr, ok := err.(pgdriver.Error); ok {
		return fmt.Sprintf("Postgres Error: %s", pgErr.Field('M'))
	}

	return fmt.Sprintf("Unknown Error: %s", err.Error())
}

func (c *AppController) handleError(ctx *gin.Context, err error, message string) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": message, "details": err.Error()})
}
