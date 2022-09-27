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
		if value.Messages.Path == "/health" {
			continue
		}
		switch value.Messages.MsgType {
		case "INFO", "info":
			info.InfoService.InES("info", value)
			break
		case "error", "ERROR", "ERR":
			logerror.ErrorService.InES("error", value)
			break
		case "warning", "WARNING", "WARN":
			warning.WarningService.InES("warning", value)
			break
		case "notice", "NOTICE":
			notice.NoticeService.InES("notice", value)
			break
		default:
			fmt.Println(value.String())
			if value.Messages.Msg != "" && value.Messages.LocalTime != "" {
				notice.NoticeService.InES("notice", value)
			}

		}
	}
	return nil
}
