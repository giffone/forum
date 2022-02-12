package category

import (
	"context"
	"forum/internal/adapters/api"
	"forum/internal/model"
	"forum/internal/service"
)

type serviceCategory struct {
	storageCategory service.StorageCategory
}

func NewService(storage service.StorageCategory) api.ServiceCategory {
	return &serviceCategory{storageCategory: storage}
}

func (sc *serviceCategory) GetAll(ctx context.Context) ([]*model.Category, error) {
	return sc.storageCategory.GetAll(ctx)
}

func (sc *serviceCategory) Create(ctx context.Context, dto *model.CreateCategoryDTO) error {
	return nil
}

func (sc *serviceCategory) Delete(ctx context.Context, id int) error {
	return nil
}
