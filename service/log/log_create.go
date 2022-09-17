package log

import (
	"fmt"
	"kuke_logger/common/request"
	logerror "kuke_logger/service/error"
	"kuke_logger/service/info"
	"kuke_logger/service/notice"
	"kuke_logger/service/warning"
)

type CreateService struct {
}

func (lc *CreateService) CreateLog(log []request.LogRequest) (err error) {
	for _, value := range log {
		switch value.Messages.MsgType {
		case "INFO", "info":
			info.InfoService.InES("info", value)
			break
		case "error", "ERROR":
			logerror.ErrorService.InES("error", value)
			break
		case "warning", "WARNING":
			warning.WarningService.InES("warning", value)
			break
		case "notice", "NOTICE":
			notice.NoticeService.InES("notice", value)
			break
		default:
			fmt.Printf("所有日志")
		}
	}
	return nil
}
