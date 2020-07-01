package WebLib

import (
	"github.com/gin-gonic/gin"
	"go-micro-helloworld/Services"
)

func NewGinRouter(service Services.ProdService) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(InitMiddleware(service))
	v1Group := ginRouter.Group("v1")
	{
		v1Group.Handle("POST", "/prods", GetProdsList)
	}
	return ginRouter
}
