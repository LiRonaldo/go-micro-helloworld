package WebLib

import (
	"github.com/gin-gonic/gin"
	"go-micro-helloworld/Services"
)

func InitMiddleware(service Services.ProdService) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Keys = make(map[string]interface{})
		context.Keys["service"] = service
		context.Next()
	}
}
