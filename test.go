package main

import (
	"fmt"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"io/ioutil"
	"log"
	"net/http"
)

func CallApi(addr string, path string, method string) (string, error) {
	req, _ := http.NewRequest(method, "http://"+addr+path, nil)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	buf, _ := ioutil.ReadAll(res.Body)
	return string(buf), nil
}

func main() {
	consulReg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))

	getService, err := consulReg.GetService("prodServices")
	if err != nil {
		log.Fatal(err)
	}
	next := selector.Random(getService)
	node, err := next()
	if err != nil {
		log.Fatal(err)
	}
	res, err := CallApi(node.Address, "/v1/prods", "GET")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
