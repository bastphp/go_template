package log

import "kuke_logger/service"

type ApiGroup struct {
	LogApi
}

var (
	logService = service.ServiceGroupApp.LogServiceGroup
)
