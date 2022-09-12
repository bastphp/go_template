package log

import (
	"github.com/gin-gonic/gin"
	"kuke_logger/api/V1"
)

type ApiRouter struct {
}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("log")
	apiRouterLog := V1.ApiGroupApp.LogApiGroup.LogApi
	{
		apiRouter.POST("createLonger", apiRouterLog.CreateLonger)
	}
}
