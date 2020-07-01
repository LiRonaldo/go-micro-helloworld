package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"go-micro-helloworld/Services"
	"go-micro-helloworld/WebLib"
)

type logWrapper struct {
	client.Client
}

func (this *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	fmt.Println("调用接口")
	err := this.Client.Call(ctx, req, rsp)
	return err
}
func NewLogWrapper(c client.Client) client.Client {
	return &logWrapper{c}
}

func main() {
	//注册consul
	consulReg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
	myService := micro.NewService(
		micro.Name("prodservice.client"),
		micro.WrapClient(NewLogWrapper),
	)
	//将服务端的服务变为客户端服务
	prodService := Services.NewProdService("prodservice", myService.Client())
	//mirco 生成web服务
	httpServices := web.NewService(
		web.Name("httpprodServices"),
		//注掉是因为要使用server.init（）,同时注册多个服务的，使用init 必须使用命令行 go run prod_main.go --server_address 127.0.0.1:8003 注册多个服务就修改ip或者端口
		web.Address("localhost:8001"),
		web.Handler(WebLib.NewGinRouter(prodService)),
		web.Registry(consulReg),
	)

	//gin框架生成web访问api

	//命令行方式启动必备
	httpServices.Init()
	httpServices.Run()
}
