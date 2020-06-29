package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	myhttp "github.com/micro/go-plugins/client/http"
	"github.com/micro/go-plugins/registry/consul"
	"log"
)

func CallApi2(s selector.Selector) {
	myClient := myhttp.NewClient(
		client.Selector(s),
		client.ContentType("application/json"),
	)
	req := myClient.NewRequest("prodServices", "/v1/prods", map[string]string{})
	var rep map[string]interface{}
	err := myClient.Call(context.Background(), req, &rep)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rep["data"])
}
func main() {
	//获取一个consul
	consulReg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
	//随机轮训
	mySelector := selector.NewSelector(
		selector.Registry(consulReg),
		selector.SetStrategy(selector.RoundRobin),
	)
	CallApi2(mySelector)
}
