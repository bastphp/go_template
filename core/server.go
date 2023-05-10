package core

import (
	"template/initialize"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	//initialize.Kafka("fluent-bit")
	Router := initialize.Routers()
	s := initServer(Router)
	s.ListenAndServe()
}
