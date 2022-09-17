package service

import (
	"kuke_logger/service/log"
)

type ServiceGroup struct {
	LogServiceGroup log.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
