package initialize

import (
	"github.com/gin-gonic/gin"
	"template/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()

	xxlRouter := router.RouterGroupApp.Xxl

	publicGroup := Router.Group("")

	xxlRouter.InitApiRouter(publicGroup)

	return Router
}
