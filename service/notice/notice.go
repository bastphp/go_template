package notice

import (
	"kuke_logger/common/request"
	"kuke_logger/service/common"
	"log"
)

type Notice struct {
	commonTools common.Common
}

func (i Notice) InES(index string, logs request.LogRequest) {
	indexName, err := i.commonTools.CreateMapping(index)

	if err != nil {
		log.Panicln("mapping created error:" + err.Error())
	}
	i.commonTools.CreateIndex(indexName, logs)
}
