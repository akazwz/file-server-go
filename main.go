package main

import (
	"fmt"
	"github.com/akazwz/file-server/global"
	"github.com/akazwz/file-server/initialize"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	fmt.Println("file server")

	global.VP = initialize.InitViper()
	if global.VP == nil {
		log.Println("init config error")
	}

	gin.SetMode(global.CFG.Server.Mode)

	routers := initialize.Routers()
	addr := fmt.Sprintf(":%d", global.CFG.Server.Addr)
	readTimeOut := global.CFG.Server.ReadTimeout
	writeTimeout := global.CFG.Server.WriteTimeout

	s := &http.Server{
		Addr:         addr,
		Handler:      routers,
		ReadTimeout:  readTimeOut,
		WriteTimeout: writeTimeout,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Println("system sever start error")
	}
}
