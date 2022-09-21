package common

import (
	"fmt"
	"kuke_logger/common/request"
	"log"
)

type Common struct {
}

func (c Common) CreateMapping(msgType string) (string, error) {
	indexName := esMapping.GetIndexName(msgType, "")
	b := esMapping.IndexExits(indexName)
	var err error
	if !b {
		err = esMapping.InitMapping(indexName, "_doc", esMapping.LogMappingTmp())
	}
	return indexName, err

}

func (c Common) CreateIndex(indexName string, logs request.LogRequest) {
	id := esMapping.InES(indexName, logs)
	log.Printf("create success id is:" + id)
}

func (c Common) CreateMappingTime() {
	logType := [4]string{"error", "info", "waring", "notice"}
	for _, value := range logType {
		indexName := esMapping.GetIndexName(value, "+1")
		fmt.Println(indexName)
		b := esMapping.IndexExits(indexName)
		if !b {
			esMapping.InitMapping(indexName, "_doc", esMapping.LogMappingTmp())
		}
	}
}
