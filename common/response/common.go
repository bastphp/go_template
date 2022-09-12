package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Res struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Res{
		code,
		msg,
		data,
	})
}
