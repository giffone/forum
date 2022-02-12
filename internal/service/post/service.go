package post

import (
	"context"
	"forum/internal/adapters/api"
	"forum/internal/model"
	"forum/internal/service"
)

type servicePost struct {
	storage service.StoragePost
	//userService ServiceUser // for check user exist before create post
}

func NewService(storage service.StoragePost) api.ServicePost {
	return &servicePost{storage: storage}
}

func (sp *servicePost) GetByID(ctx context.Context, id int) (*model.Post, error) {
	return sp.storage.GetByID(ctx, id)
}

func (sp *servicePost) GetAll(ctx context.Context) ([]*model.Post, error) {
	return sp.storage.GetAll(ctx)
}

func (sp *servicePost) Create(ctx context.Context, dto *model.CreatePostDTO) error {

	return nil
}

func (sp *servicePost) Delete(ctx context.Context, id int) error {

	return nil
}
