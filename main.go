package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mumushuiding/go-logging/routers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", routers.Index)

	// 启动服务
	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", Config.Port),
		Handler:        mux,
		ReadTimeout:    time.Duration(Config.ReadTimeout * int(time.Second)),
		WriteTimeout:   time.Duration(Config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("the application start up at port%s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
