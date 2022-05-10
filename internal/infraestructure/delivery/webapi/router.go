package webapi

import (
	"github.com/gin-gonic/gin"
	"go-api-example/internal/infraestructure/delivery/webapi/handler"
)

// Configure the HealthyChecks application's endpoints.
func registerHealthyCheckEndpoints(diagnostic *gin.RouterGroup) {
	diagnostic.GET("", handler.PingHandler)
}

// In this func you can register the application's endpoints.
func registerApplicationEndpoints(appEndpoints *gin.RouterGroup) {
	appEndpoints.GET("/items/:item_id", handler.GetItemById)
}
