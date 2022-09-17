package common

import (
	"kuke_logger/common/request"
	"log"
)

type Common struct {
}

func (c Common) CreateMapping(msgType string) (string, error) {
	indexName := esMapping.GetIndexName(msgType)
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
