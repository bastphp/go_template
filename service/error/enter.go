package error

type ServiceGroup struct {
	Error
}

var ErrorService = new(ServiceGroup)
