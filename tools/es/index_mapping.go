package es

import (
	"context"
	"fmt"
	"kuke_logger/global"
	"log"
	"time"
)

type MappingTool struct {
}

func (mt *MappingTool) LogMappingTmp() string {
	return `{"mappings":{
				"properties":{
					"kk-request-id": {"type":"keyword"},
					"local-time": {"type":"date","format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis||strict_date_optional_time"},
					"msg":  {"type":"text"},
					"msg-info":{
						"properties":{
							"data": {"type":"text"},
							"full-path": {"type":"keyword"},
   							"header": {"type":"text"},
							"type": {"type":"keyword"}
					  }
					  
					},
					"msg-type":{"type":"keyword"},
					"path": {"type":"keyword"},
					"response_time":{"type": "float"},
					"tag": {"type":"keyword"},
					"user-id": {"type":"keyword"},
					"kubernetes": {
					  "properties": {
					    "container_name":{"type":"keyword"},
						  "namespace_name":{"type":"keyword"},
						  "pod_name":{"type":"text"},
						  "container_image":{"type":"text"},
						  "host":{"type":"text"},
						  "labels":{"type":"flattened"},
						  "docker_id":{"type":"text"},
						  "annotations":{"type":"flattened"}
					  }
          }
    }
    
  }
  
}`
}

func (mt *MappingTool) GetIndexName(msgInfo string) string {
	timeObj := time.Now()
	return "k8s-log-" + msgInfo + "-" + timeObj.Format("01-02")
}

func (mt *MappingTool) InitMapping(esIndexName string, esTypeName string, typeMapping string) error {
	fmt.Println(esIndexName)
	_, err := global.GVA_ES.CreateIndex(esIndexName).BodyString(typeMapping).Do(context.Background())
	if err != nil {
		log.Panicln("createMappingError:" + err.Error())
		return err
	}
	return nil
}

func (mt MappingTool) IndexExits(indexName string) bool {
	exists, err := global.GVA_ES.IndexExists(indexName).Do(context.Background())
	if err != nil {
		log.Panicln("IndexExists" + err.Error())
	}
	if exists {
		return true
	}
	return false
}
