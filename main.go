// Copyright (c) 2019 数据中心入口文件

package main

import (
	"dc/service"
	"dc/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hprose/hprose-golang/rpc"
	"net/http"
)

func InitHttpServer()  {

	//开启 hprose的
	server := rpc.NewHTTPService()

	router := gin.New()

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	gin.SetMode(util.Cfg.RunMode)
	server.AddInvokeHandler(service.EntryInvokeHandler)
	//server.AddInstanceMethods(&service.HttpService{}, rpc.Options{NameSpace: "DC"})

	router.Any("/path", func(c *gin.Context) {
		fmt.Println(1111111)
		server.ServeHTTP(c.Writer, c.Request)
	})
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", util.Cfg.HttpServer.Port),
		Handler:        router,
		ReadTimeout:    util.Cfg.HttpServer.ReadTime,
		WriteTimeout:   util.Cfg.HttpServer.WriteTime,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
//RPC 方式调用启动程序
func InitRpcServer()  {

	server := rpc.NewTCPServer("tcp4://"+util.Cfg.RpcServer.Host+":" + util.Cfg.RpcServer.Port + "/")
	server.AddInvokeHandler(service.EntryInvokeHandler)
	//注册func
	//server.AddFunction("Hello", hello)

	//注册struct，命名空间是DCService
	server.AddInstanceMethods(&service.DCService{}, rpc.Options{NameSpace: "DC"})
	err := server.Start()
	if err != nil {
		fmt.Printf("start server fail, err:%v\n", err)
		return
	}
}

func main() {
	//初始化配置
	err := util.InitConfig("conf/app.ini")

	if err != nil {
		fmt.Printf("load config file fail, err:%v\n", err)
		return
	}
	//InitHttpServer()
	InitRpcServer()
}