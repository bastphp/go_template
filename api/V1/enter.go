package V1

import "kuke_logger/api/V1/log"

type ApiGroup struct {
	LogApiGroup log.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
