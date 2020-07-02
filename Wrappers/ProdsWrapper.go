package Wrappers

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
	"go-micro-helloworld/Services"
	"strconv"
)

type ProdsWrapper struct {
	client.Client
}

func (this *ProdsWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	cofigA := hystrix.CommandConfig{
		Timeout: 1000,
	}
	hystrix.ConfigureCommand(req.Service()+"."+req.Endpoint(), cofigA)
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		return this.Client.Call(ctx, req, rsp)
	}, func(err error) error {
		defaultData(rsp)
		return nil
	})
}
func NewLogWrapper(c client.Client) client.Client {
	return &ProdsWrapper{c}
}
func newProd(id int32, pname string) *Services.ProdModel {
	return &Services.ProdModel{ProdID: id, ProdName: pname}
}

func defaultProds(rsp interface{}) {
	var models []*Services.ProdModel = make([]*Services.ProdModel, 0)
	var i int32
	for i = 0; i < 5; i++ {
		models = append(models, newProd(20+i, "prodname"+strconv.Itoa(20+int(i))))
	}
	resp := rsp.(*Services.ProdListResponse)
	resp.Data = models
}

//通用降级方法
func defaultData(resp interface{}) {
	switch t := resp.(type) {
	case *Services.ProdListResponse:
		defaultProds(resp)
	case *Services.ProdDetailResponse:
		t.Data = newProd(10, "测试商品")
	}
}
