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
			go info.InfoService.InES("info", value)
			break
		case "error", "ERROR", "ERR":
			go logerror.ErrorService.InES("error", value)
			break
		case "warning", "WARNING", "WARN":
			go warning.WarningService.InES("warning", value)
			break
		case "notice", "NOTICE":
			go notice.NoticeService.InES("notice", value)
			break
		default:
			fmt.Println(value.String())
			if value.Messages.Msg != "" && value.Messages.LocalTime != "" {
				go notice.NoticeService.InES("notice", value)
			}

		}
	}
	return nil
}
