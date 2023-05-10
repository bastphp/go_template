package core

import (
	"fmt"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"template/global"
	"time"
)

func initServer(router *gin.Engine) server {
	system := global.GVA_CONFIG.System
	address := fmt.Sprintf(":%d", system.Addr)
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = time.Duration(system.ReaderTimeOut) * time.Second
	s.WriteTimeout = time.Duration(system.WriterTimeOut) * time.Second
	s.MaxHeaderBytes = system.MaxParams * 1024 * 1024
	return s
}
