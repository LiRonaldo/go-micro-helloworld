package WebLib

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-micro-helloworld/Services"
)

func GetProdsList(gincontext *gin.Context) {
	prodService := gincontext.Keys["service"].(Services.ProdService)
	var prodReq Services.ProdsRequest
	err := gincontext.Bind(&prodReq)
	if err != nil {
		gincontext.JSON(500, gin.H{"message:": err.Error()})
	} else {
		prodResp, _ := prodService.GetProdsList(context.Background(), &prodReq)
		gincontext.JSON(200, gin.H{"data:": prodResp.Data})
	}

}
