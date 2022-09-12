package router

import "kuke_logger/router/log"

type RouterGroup struct {
	Log log.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
