package notice

type ServiceGroup struct {
	Notice
}

var NoticeService = new(ServiceGroup)
