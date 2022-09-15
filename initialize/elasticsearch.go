package initialize

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"kuke_logger/global"
)

func Elasticsearch() {
	esCfg := global.GVA_CONFIG.Elasticsearch
	config := elasticsearch.Config{
		Addresses: esCfg.Addr,
		Username:  esCfg.UserName,
		Password:  esCfg.Password,
	}
	client, err := elasticsearch.NewClient(config)
	if err != nil {
		fmt.Printf(err.Error())
	}
	global.GVA_ES = client
}
