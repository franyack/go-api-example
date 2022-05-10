package gateway

import "go-api-example/internal/business/domain"

var itemsService ItemsService

func RegisterItemsService(repository ItemsService) {
	itemsService = repository
}

func NewItemsService() ItemsService {
	if itemsService == nil {
		panic("itemsService not found")
	}
	return itemsService
}

type ItemsService interface {
	GetItemById(itemID string) (*domain.Item, error)
}
