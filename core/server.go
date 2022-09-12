package core

import "kuke_logger/initialize"

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	Router := initialize.Routers()

	address := ":80"
	s := initServer(address, Router)
	s.ListenAndServe()
}
