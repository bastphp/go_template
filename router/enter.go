package router

import (
	"kuke_logger/router/log"
	"kuke_logger/router/xxl"
)

type RouterGroup struct {
	Log log.RouterGroup
	Xxl xxl.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
