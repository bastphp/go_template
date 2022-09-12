package log

import (
	"github.com/gin-gonic/gin"
	"kuke_logger/common/response"
)

type LogApi struct {
}

func (l *LogApi) CreateLonger(c *gin.Context) {
	c.Params.Get("test")
	response.Result(1, "{}", "success", c)
}
