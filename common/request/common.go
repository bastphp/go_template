package request

type Labels struct {
	Layer   string `json:"k8s.kuboard.cn/layer"`
	Name    string `json:"k8s.kuboard.cn/name"`
	Logging string `json:"logging"`
	PodHash string `json:"pod-template-hash"`
}

type Annotations struct {
	RestartedAt string `json:"kubectl.kubernetes.io/restartedAt"`
}

type Kubernetes struct {
	ContainerName  string      `json:"container_name"`
	NamespaceName  string      `json:"namespace_name"`
	PodName        string      `json:"pod_name"`
	ContainerImage string      `json:"container_image"`
	Host           string      `json:"host"`
	Labels         Labels      `json:"labels"`
	DockerId       string      `json:"docker_id"`
	Annotations    Annotations `json:"annotations"`
}

type msgInfo struct {
	Data     string `json:"data"`
	FullPath string `json:"full-path"`
	Header   string `json:"header"`
	Type     string `json:"type"`
}

type LogRequest struct {
	Time       string     `json:"time"`
	Messages   Message    `json:"messages"`
	Date       int64      `json:"date"`
	On         int        `json:"on"`
	Kubernetes Kubernetes `json:"kubernetes"`
	Stream     string     `json:"stream"`
}

type Message struct {
	KkRequestId  string  `json:"kk-request-id"`
	LocalTime    string  `json:"local-time"`
	Msg          string  `json:"msg"`
	MsgInfo      msgInfo `json:"msg-info"`
	MsgType      string  `json:"msg-type"`
	Path         string  `json:"path"`
	ResponseTime int     `json:"response_time"`
	Tag          string  `json:"tag"`
	UserId       string  `json:"user-id"`
}
