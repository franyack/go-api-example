package middleware

import (
	"github.com/gin-gonic/gin"
	"go-api-example/internal/business/gateway"
	"go-api-example/internal/infraestructure/service"
)

// IoC Register all your useCases, repositories and service in this middleware.
// This is the first middleware. Show webapi.Start()
func IoC() gin.HandlerFunc {
	return func(c *gin.Context) {
		gateway.RegisterItemsService(service.NewMeliApiItemsService())
	}
}
