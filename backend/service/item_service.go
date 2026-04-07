package service

import (
	"context"
	// "errors"
	"antis/backend/models"
	"antis/backend/repositories"
)

type ItemService struct {
	repo repositories.ItemRepository
}

func NewItemService(repo repositories.ItemRepository) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) CreateItem (ctx context.Context, item *models.Item) (*models.Item, error) {
	
	if err := s.repo.Create(ctx, item); err != nil {
        return nil, err
    }
	 
	
	return item,nil
}

func (s *ItemService) GetItemById (ctx context.Context, id int64) (*models.Item, error) {
	item,err := s.repo.GetByID(ctx, id)

	if err != nil {
		return nil,err
	}

	return item,nil
}

func (s *ItemService) GetAllItems (ctx context.Context) ([]models.Item) {
		items,_ := s.repo.GetAllItems(ctx)

		return items
}