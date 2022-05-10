package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-api-example/internal/business/usecase"
	apierrors "go-api-example/internal/infraestructure/delivery/webapi/utils"
	"net/http"
)

func GetItemById(c *gin.Context) {

	itemID := c.Param("item_id")

	itemByIdUseCase := usecase.NewGetItemByIdUserCase()
	item, err := itemByIdUseCase.GetItemById(itemID)
	if err != nil {
		log.Error("error obtaining item: ", err)
		apiError := err.(apierrors.ApiError)
		c.AbortWithStatusJSON(apiError.Status(), apiError)
		return
	}
	c.JSON(http.StatusOK, item)
}
