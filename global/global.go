package global

import (
	"github.com/olivere/elastic/v7"
	"kuke_logger/configs"
)

var (
	GVA_CONFIG configs.Server
	GVA_ES     *elastic.Client
)
