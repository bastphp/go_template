package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	logRequest "kuke_logger/common/request"
	"kuke_logger/common/response"
)

type LogApi struct {
}

func (l *LogApi) CreateLonger(c *gin.Context) {
	var log []logRequest.LogRequest
	_ = c.ShouldBindJSON(&log)
	fmt.Print(len(log))
	_ = logService.CreateLog(log)
	response.Result(1, "{}", "success", c)
}
