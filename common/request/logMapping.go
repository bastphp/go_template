package request

type KubernetesLog struct {
	Annotations    interface{} `json:"annotations"`
	ContainerImage string      `json:"container_image"`
	ContainerName  string      `json:"container_name"`
	DockerId       string      `json:"docker_id"`
	Host           string      `json:"host"`
	Labels         interface{} `json:"labels"`
	NamespaceName  string      `json:"namespace_name"`
	PodName        string      `json:"pod_name"`
}

type MsgInfoLog struct {
	Data     string `json:"data"`
	FullPath string `json:"fullPath"`
	Header   string `json:"header"`
	Type     string `json:"type"`
}

type LogMapping struct {
	KkRequestId  string        `json:"kk-request-id"`
	Kubernetes   KubernetesLog `json:"kubernetes"`
	MsgInfo      MsgInfoLog    `json:"msg-info"`
	LocalTime    string        `json:"local-time"`
	MsgType      string        `json:"msg-type"`
	Path         string        `json:"path"`
	ResponseTime float64       `json:"response_time"`
	Tag          string        `json:"tag"`
	UserId       string        `json:"user-id"`
	TimeStamp    string        `json:"@timestamp"`
}
