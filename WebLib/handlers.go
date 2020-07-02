package WebLib

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-micro-helloworld/Services"
	"strconv"
)

func newProd(id int32, pname string) *Services.ProdModel {
	return &Services.ProdModel{ProdID: id, ProdName: pname}
}

func defaultProds() (resp *Services.ProdListResponse, err error) {
	var models []*Services.ProdModel = make([]*Services.ProdModel, 0)
	var i int32
	for i = 0; i < 5; i++ {
		models = append(models, newProd(100+i, "prodname"+strconv.Itoa(100+int(i))))
	}
	resp = &Services.ProdListResponse{}
	resp.Data = models
	return resp, nil
}

func GetProdsList(gincontext *gin.Context) {
	prodService := gincontext.Keys["service"].(Services.ProdService)
	var prodReq Services.ProdsRequest
	err := gincontext.Bind(&prodReq)
	if err != nil {
		gincontext.JSON(500, gin.H{"message:": err.Error()})
	} else {
		prodResp, err := prodService.GetProdsList(context.Background(), &prodReq)
		if err != nil {
			gincontext.JSON(500, gin.H{"message:": err.Error()})
		} else {
			gincontext.JSON(200, gin.H{"data:": prodResp.Data})
		}
		/*//熔断代码改造
		//1.配置config
		cofigA := hystrix.CommandConfig{
			Timeout: 1000,
		}
		//2.配置commond
		hystrix.ConfigureCommand("getprods", cofigA)
		//3.执行do方法
		var prodResp *Services.ProdListResponse
		err:=hystrix.Do("getprods", func() error {

			return err
		}, func(err error) error {
			prodResp,err=defaultProds()
			return  err
		})*/

	}

}

func GetProdDetail(ginCtx *gin.Context) {
	var prodReq Services.ProdsRequest
	PanicIfError(ginCtx.BindUri(&prodReq))
	prodService := ginCtx.Keys["service"].(Services.ProdService)
	resp, err := prodService.GetProdDetail(context.Background(), &prodReq)
	if err != nil {
		ginCtx.JSON(500, gin.H{"message:": err.Error()})
	} else {
		ginCtx.JSON(200, gin.H{"data:": resp.Data})
	}
}
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
