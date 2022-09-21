package V1

import (
	"kuke_logger/api/V1/log"
	"kuke_logger/api/V1/xxl"
)

type ApiGroup struct {
	LogApiGroup log.ApiGroup
	XxlApiGroup xxl.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
