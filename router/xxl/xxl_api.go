package xxl

import (
	"github.com/gin-gonic/gin"
	"template/api/V1"
)

type ApiRouter struct {
}

func (receiver ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("")
	apiRouterXxl := V1.ApiGroupApp.XxlApiGroup.XxlApi
	{
		apiRouter.GET("ping", apiRouterXxl.Ping)
		apiRouter.POST("xxl/createMapping", apiRouterXxl.CreateMapping)
	}
}
