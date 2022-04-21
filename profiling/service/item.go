package service

import "profiling/entity"

type ItemRepoInterface interface {
	GetByID(id int) (entity.Item, error)
}

type ItemService struct {
	itemRepo ItemRepoInterface
}

func NewItemService(repo ItemRepoInterface) *ItemService {
	return &ItemService{itemRepo: repo}
}

func (s *ItemService) GetItemTitleByID(itemID int) (string, error) {
	item, err := s.itemRepo.GetByID(itemID)
	if err != nil {
		return "", err
	}
	return item.Title, nil
}
