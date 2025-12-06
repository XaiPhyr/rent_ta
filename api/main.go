package main

import (
	"api/middleware"
	"api/router"
	"api/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := utils.InitConfig()
	engine := gin.New()

	middleware.SetLoggers(engine)
	router.InitRouters(engine)
	engine.Run(cfg.Server.Host)
}
