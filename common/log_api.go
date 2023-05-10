package common

import "template/common/request"

type LogApi struct {
	Data []request.LogRequest `json:"data"`
}
