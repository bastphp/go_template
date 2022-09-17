package info

import (
	"kuke_logger/common/request"
	"kuke_logger/service/common"
	"log"
)

type Info struct {
	commonTools common.Common
}

func (i Info) InES(index string, logs request.LogRequest) {
	indexName, err := i.commonTools.CreateMapping(index)

	if err != nil {
		log.Panicln("mapping created error:" + err.Error())
	}
	i.commonTools.CreateIndex(indexName, logs)
}
