package error

import (
	"kuke_logger/common/request"
	"kuke_logger/service/common"
	"log"
)

type Error struct {
	commonTools common.Common
}

func (i Error) InES(index string, logs request.LogRequest) {
	indexName, err := i.commonTools.CreateMapping(index)

	if err != nil {
		log.Panicln("mapping created error:" + err.Error())
	}
	i.commonTools.CreateIndex(indexName, logs)
}
