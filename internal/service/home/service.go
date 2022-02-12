package home

import (
	"context"
	"forum/internal/adapters/api"
	"forum/internal/model"
	"log"
)

type serviceHome struct {
	servicePost     api.ServicePost
	serviceCategory api.ServiceCategory
}

func NewService(servicePost api.ServicePost, serviceCategory api.ServiceCategory) api.ServiceHome {
	return &serviceHome{
		servicePost:     servicePost,
		serviceCategory: serviceCategory,
	}
}

func (sh *serviceHome) Get(ctx context.Context) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	modelPost, err := sh.servicePost.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	if len(modelPost) == 0 {
		data["Posts"] = []model.Post{{Text: "no posts are created"}}
	} else {
		data["Posts"] = modelPost
	}

	modelCategory, err := sh.serviceCategory.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	if modelCategory == nil {
		log.Println("!!!")
		data["Category"] = []model.Category{{Name: "no category are created"}}
	} else {
		data["Category"] = modelCategory
	}
	return data, nil
}
