package WebLib

import (
	"fmt"
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
func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				context.JSON(500, gin.H{"status": fmt.Sprint("%s", r)})
				context.Abort()

			}
		}()
		context.Next()
	}
}
