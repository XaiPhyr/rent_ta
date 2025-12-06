package router

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouters(router *gin.Engine) {
	var user = controllers.UserController{}
	user.InitUserController(router)
}
