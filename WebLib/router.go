package WebLib

import (
	"github.com/gin-gonic/gin"
	"go-micro-helloworld/Services"
)

func NewGinRouter(service Services.ProdService) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(InitMiddleware(service), ErrorMiddleware())
	v1Group := ginRouter.Group("v1")
	{
		v1Group.Handle("POST", "/prods", GetProdsList)
		v1Group.Handle("GET", "/prods/:pid", GetProdDetail)
	}
	return ginRouter
}
