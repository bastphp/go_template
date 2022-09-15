package log

import (
	"fmt"
	"kuke_logger/common/request"
)

type CreateService struct {
}

func (lc *CreateService) CreateLog(log []request.LogRequest) (err error) {
	for _, value := range log {
		switch value.Messages.MsgType {
		case "INFO", "info":
			fmt.Printf("info日志")
			break
		case "error", "ERROR":
			fmt.Printf("error日志")
			break
		case "warning", "WARNING":
			fmt.Printf("waring日志")
			break
		case "notice", "NOTICE":
			fmt.Printf("notice 日志")
			break
		default:
			fmt.Printf("所有日志")
		}
	}
	return nil
}
