package es

import (
	"context"
	"fmt"
	"kuke_logger/common/request"
	"kuke_logger/global"
	"log"
	"strings"
)

type CreateTool struct {
}

func (ct *CreateTool) InES(indexName string, logs request.LogRequest) string {
	fmt.Println(logs.Time)
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
		LocalTime:    strings.TrimSuffix(logs.Messages.LocalTime, " +08:00"),
		MsgType:      logs.Messages.MsgType,
		Path:         logs.Messages.Path,
		ResponseTime: logs.Messages.ResponseTime,
		Tag:          logs.Messages.Tag,
		UserId:       logs.Messages.UserId,
		TimeStamp:    logs.Time,
	}
	res, err := global.GVA_ES.Index().Index(indexName).BodyJson(msg).Do(context.Background())
	if err != nil {
		log.Panicln("create index error:" + err.Error())
	}
	return res.Id
}
