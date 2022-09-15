package global

import (
	"github.com/elastic/go-elasticsearch/v7"
	"kuke_logger/configs"
)

var (
	GVA_CONFIG configs.Server
	GVA_ES     *elasticsearch.Client
)
