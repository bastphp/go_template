package initialize

import (
	"github.com/gin-gonic/gin"
	"kuke_logger/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	logerRouter := router.RouterGroupApp.Log
	xxlRouter := router.RouterGroupApp.Xxl

	publicGroup := Router.Group("")

	logerRouter.InitApiRouter(publicGroup)
	xxlRouter.InitApiRouter(publicGroup)

	return Router
}
