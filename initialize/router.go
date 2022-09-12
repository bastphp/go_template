package initialize

import (
	"github.com/gin-gonic/gin"
	"kuke_logger/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	logerRouter := router.RouterGroupApp.Log

	publicGroup := Router.Group("")

	logerRouter.InitApiRouter(publicGroup)

	return Router
}
