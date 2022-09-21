package xxl

import "github.com/gin-gonic/gin"

type XxlApi struct {
}

func (a XxlApi) Ping(c *gin.Context) {
	c.JSON(200, "pong")
}

func (a XxlApi) CreateMapping(c *gin.Context) {

}
