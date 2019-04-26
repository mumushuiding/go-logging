package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mumushuiding/go-logging/config"
	"github.com/mumushuiding/go-logging/model"
	"github.com/mumushuiding/go-logging/routers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", routers.Index)
	// -------------------- 日志 -------------
	mux.HandleFunc("/log/save", routers.SaveLogdata)
	mux.HandleFunc("/log/deleteById", routers.DeleteLogdataByID)
	mux.HandleFunc("/log/findAll", routers.PostHandler(routers.FindLogdatas))
	// 配置
	var config = *config.Config
	// 启动数据库连接
	model.Setup()
	// 启动服务
	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", config.Port),
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("the application start up at port%s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
