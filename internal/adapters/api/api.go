package api

import (
	"context"
	"forum/internal/model"
	"net/http"
)

type Handler interface {
	Register(router *http.ServeMux)
}

type ServicePost interface { // use cases
	Create(ctx context.Context, dto *model.CreatePostDTO) error
	GetByID(ctx context.Context, id int) (*model.Post, error)
	GetAll(ctx context.Context) ([]*model.Post, error)
}

type ServiceUser interface { // use cases
	Create(ctx context.Context, dto *model.CreateUserDTO) (error, string)
	GetByID(ctx context.Context, id int) (*model.User, error)
	CheckLoginPassword(ctx context.Context, dto *model.CreateUserDTO) (error, string)
}

type ServiceCategory interface { // use cases
	Create(ctx context.Context, dto *model.CreateCategoryDTO) error
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]*model.Category, error)
}

type ServiceHome interface { // use cases
	Get(ctx context.Context) (map[string]interface{}, error)
}

type ServiceSession interface { // use cases
	Create(ctx context.Context, dto *model.CreateSessionDTO) (error, string)
	DeleteByLogin(ctx context.Context, login string) error
	DeleteExpired(ctx context.Context) error
	GetByLogin(ctx context.Context, login string) (*model.Session, error)
}
