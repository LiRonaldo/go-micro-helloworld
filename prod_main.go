package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"go-micro-helloworld/Helper"
	"go-micro-helloworld/ProdService"
)

func main() {
	//注册consul
	consulReg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
	//gin框架生成web访问api
	ginRouter := gin.Default()
	v1Group := ginRouter.Group("v1")
	{
		v1Group.Handle("POST", "/prods", func(context *gin.Context) {
			var pr Helper.ProdsRequest
			err := context.Bind(&pr)
			if err != nil || pr.Size <= 0 {
				pr = Helper.ProdsRequest{Size: 2}
			}
			context.JSON(200, gin.H{"data": ProdService.NewProdList(pr.Size)})
		})
	}

	//mirco 生成服务
	server := web.NewService(
		web.Name("prodServices"),
		//注掉是因为要使用server.init（）,同时注册多个服务的，使用init 必须使用命令行 go run prod_main.go --server_address 127.0.0.1:8003 注册多个服务就修改ip或者端口
		//web.Address("localhost:8001"),
		web.Handler(ginRouter),
		web.Registry(consulReg),
	)
	//命令行方式启动必备
	server.Init()
	server.Run()
}
