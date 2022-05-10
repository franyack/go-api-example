package usecase

import (
	"go-api-example/internal/business/domain"
	"go-api-example/internal/business/gateway"
)

type (
	GetItemByIdInterface interface {
		GetItemById(itemID string) (item *domain.Item, err error)
	}

	getItemByIdInterface struct{}
)

func NewGetItemByIdUserCase() GetItemByIdInterface {
	return &getItemByIdInterface{}
}

func (getItems *getItemByIdInterface) GetItemById(itemID string) (item *domain.Item, err error) {
	itemsRepository := gateway.NewItemsService()
	item, err = itemsRepository.GetItemById(itemID)
	if err != nil {
		return nil, err
	}
	return item, err
}
