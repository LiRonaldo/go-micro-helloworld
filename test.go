package main

import (
	"fmt"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"log"
	"time"
)

func main() {
	consulReg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
	for {
		getService, err := consulReg.GetService("prodServices")
		if err != nil {
			log.Fatal(err)
		}
		next := selector.Random(getService)
		node, err := next()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(node.Address)
		time.Sleep(time.Second)
	}

}
