package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
	"go-api-example/internal/business/domain"
	"go-api-example/internal/business/gateway"
	apierrors "go-api-example/internal/infraestructure/delivery/webapi/utils"
)

func NewMeliApiItemsService() gateway.ItemsService {
	return &meliApiItemsService{}
}

type meliApiItemsService struct {
}

func (m *meliApiItemsService) GetItemById(itemID string) (*domain.Item, error) {
	url := fmt.Sprintf("https://api.mercadolibre.com/items/%s", itemID)
	response, err := http.Get(url)
	if err != nil {
		log.Error(err.Error())
		return nil, apierrors.NewInternalServerApiError("error getting api information", err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
		return nil, apierrors.NewInternalServerApiError("error reading response body", err)
	}
	if response.StatusCode != http.StatusOK {
		apiErr, err := apierrors.NewApiErrorFromBytes(responseData)
		if err != nil {
			log.Error(err)
			return nil, apierrors.NewInternalServerApiError("unmarshal error response has failed", err)
		}
		return nil, apierrors.NewApiError(apiErr.Message(), response.Status, response.StatusCode, apiErr.Cause())
	}

	var item *domain.Item
	if errUnmarshal := json.Unmarshal(responseData, &item); errUnmarshal != nil {
		log.Error(errUnmarshal)
		return nil, apierrors.NewInternalServerApiError("error unmarshalling response", errUnmarshal)
	}

	return item, nil
}
