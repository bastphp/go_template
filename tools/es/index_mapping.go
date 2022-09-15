package es

func LogMappingTmp() string {
	return `{
				"properties":{
					"kk-request-id": {"type":"keyword"},
					"local-time": {"type":"date", "format":"yyyy-MM-dd HH:mm:ss|strict_date_optional_time|epoch_mills"},
					"msg":  {"type":"text"},
					"msg-info":{
						"properties":{
							"data": {"type":"flattened"},
							"full-path": {"type":"keyword"},
   							"header": {"type":"flattened"},
							"type": {"type":"keyword"}
}
					"msg-type": {"type":"keyword"},
					"path": {"type":"keyword"},
					"response_time":{"type": "float"},
					"tag": {"type":"keyword"},
					"user-id": {"type":"keyword"},
					"kubernetes": {
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
}`
}

func GetIndexName(msgInfo string) string {
	return "k8s-log-" + msgInfo
}

func InitMapping(esIndexName string) {

}
