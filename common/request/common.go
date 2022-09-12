package request

type labels struct {
}

type kubernetes struct {
	containerName  string `json:"container_name"`
	namespaceName  string `json:"namespace_name"`
	podName        string `json:"pod_name"`
	containerImage string `json:"container_image"`
	host           string `json:"host"`
	labels         labels `json:"labels"`
}

type msgInfo struct {
	data     string `json:"data"`
	fullPath string `json:"full_path" from:"full-path"`
	header   string `json:"header"`
	msgType  string `json:"msg_type" from:"type"`
}

type LogRequest struct {
	time         string     `json:"time"`
	kubernetes   kubernetes `json:"kubernetes"`
	kkRequestId  string     `json:"kk_request_id" from:"kk-request-id"`
	localTime    string     `json:"local_time" from:"local-time"`
	msg          string     `json:"msg"`
	msgInfo      msgInfo    `json:"msg_info" from:"msg-info"`
	msgType      string     `json:"msg_type" from:"msg-type"`
	path         string     `json:"path"`
	responseTime string     `json:"response_time"`
	tag          string     `json:"tag"`
	userId       string     `json:"user_id" from:"user-id"`
}
