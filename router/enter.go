package router

import (
	"template/router/xxl"
)

type RouterGroup struct {
	Xxl xxl.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
