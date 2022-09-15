package common

import "kuke_logger/common/request"

type LogApi struct {
	Data []request.LogRequest `json:"data"`
}
