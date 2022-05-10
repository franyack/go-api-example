package service

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-api-example/internal/business/domain"
	"go-api-example/internal/business/gateway"
	apierrors "go-api-example/internal/infraestructure/delivery/webapi/utils"
	"io/ioutil"
	"net/http"
	"os"
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
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		return nil, apierrors.NewInternalServerApiError("error reading response body", err)
	}

	var item *domain.Item
	if errUnmarshal := json.Unmarshal(responseData, &item); errUnmarshal != nil {
		return nil, apierrors.NewInternalServerApiError("error unmarshalling response", errUnmarshal)
	}

	return item, nil
}