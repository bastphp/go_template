package service

import (
	"kuke_logger/service/common"
	"kuke_logger/service/log"
)

type ServiceGroup struct {
	LogServiceGroup log.ServiceGroup
	Common          common.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
