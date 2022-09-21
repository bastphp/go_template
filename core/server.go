package core

import (
	"fmt"
	"kuke_logger/global"
	"kuke_logger/initialize"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	initialize.Elasticsearch()
	Router := initialize.Routers()
	initialize.Xxl(Router)
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	fmt.Printf(address)
	//address := ":80"
	s := initServer(address, Router)
	s.ListenAndServe()
}
