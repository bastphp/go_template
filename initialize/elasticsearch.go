package initialize

import (
	"context"
	"github.com/olivere/elastic/v7"
	"kuke_logger/global"
	"log"
)

func Elasticsearch() {
	esCfg := global.GVA_CONFIG.Elasticsearch
	ctx := context.Background()
	client, err := elastic.NewClient(
		elastic.SetURL(esCfg.Host),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.Printf("esclientError:%s\n", err.Error())
	}
	info, code, err := client.Ping(esCfg.Host).Do(ctx)
	if err != nil {
		log.Printf("esPingError:%s\n", err.Error())
	}
	log.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	global.GVA_ES = client
}
