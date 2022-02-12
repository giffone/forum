package service

import (
	"context"
	"forum/internal/model"
)

type StoragePost interface { // methods from storage
	Create(ctx context.Context, post *model.CreatePostDTO) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]*model.Post, error)
	GetByID(ctx context.Context, id int) (*model.Post, error)
}

type StorageUser interface { // methods from storage
	Create(ctx context.Context, user *model.CreateUserDTO) error
	Delete(ctx context.Context, id int) error
	GetByID(ctx context.Context, id int) (*model.User, error)
	GetByLogin(ctx context.Context, login string) (*model.User, error)
}

type StorageCategory interface { // methods from storage
	Create(ctx context.Context, dto *model.CreateCategoryDTO) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]*model.Category, error)
}

type StorageSession interface { // methods from storage
	Create(ctx context.Context, dto *model.CreateSessionDTO) error
	DeleteByLogin(ctx context.Context, login string) error
	GetByLogin(ctx context.Context, login string) (*model.Session, error)
}
