package main

import (
	"github.com/gin-gonic/gin"
	//"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	//"github.com/micro/go-plugins/registry/consul"
)

func main() {

	//gin框架生成web访问api
	ginRouter := gin.Default()
	ginRouter.Handle("GET", "/", func(context *gin.Context) {
		data := make([]interface{}, 0)
		context.JSON(200, gin.H{
			"data": data,
		})
	})
	//mirco 生成服务
	server := web.NewService(
		web.Address("localhost:8000"),
		web.Handler(ginRouter),
	)
	server.Run()
}
