package es

import (
	"context"
	"fmt"
	"kuke_logger/common/request"
	"kuke_logger/global"
	"log"
	"time"
)

type CreateTool struct {
}

func (ct *CreateTool) InES(indexName string, logs request.LogRequest) string {
	timeStamp := time.Now().Format("2006-01-02 15:01:05")
	fmt.Println(timeStamp)
	msg := request.LogMapping{
		KkRequestId: logs.Messages.KkRequestId,
		Kubernetes: request.KubernetesLog{
			Annotations:    logs.Kubernetes.Annotations,
			ContainerImage: logs.Kubernetes.ContainerImage,
			ContainerName:  logs.Kubernetes.ContainerName,
			DockerId:       logs.Kubernetes.DockerId,
			Host:           logs.Kubernetes.Host,
			Labels:         logs.Kubernetes.Labels,
			NamespaceName:  logs.Kubernetes.NamespaceName,
			PodName:        logs.Kubernetes.PodName,
		},
		MsgInfo: request.MsgInfoLog{
			Data:     logs.Messages.MsgInfo.Data,
			FullPath: logs.Messages.MsgInfo.FullPath,
			Header:   logs.Messages.MsgInfo.Header,
			Type:     logs.Messages.MsgInfo.Type,
		},
		LocalTime:    logs.Messages.LocalTime,
		MsgType:      logs.Messages.MsgType,
		Path:         logs.Messages.Path,
		ResponseTime: logs.Messages.ResponseTime,
		Tag:          logs.Messages.Tag,
		UserId:       logs.Messages.UserId,
		TimeStamp:    timeStamp,
	}
	res, err := global.GVA_ES.Index().Index(indexName).BodyJson(msg).Do(context.Background())
	if err != nil {
		log.Panicln("create index error:" + err.Error())
	}
	return res.Id
}
