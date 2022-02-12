package category

import (
	"context"
	"forum/internal/adapters/api"
	"forum/internal/constant"
	"forum/internal/model"
	"net/http"
)

type handlerCategory struct {
	ctx             context.Context
	serviceCategory api.ServiceCategory
}

func NewHandler(ctx context.Context, serviceCategory api.ServiceCategory) api.Handler {
	return &handlerCategory{
		ctx:             ctx,
		serviceCategory: serviceCategory,
	}
}

func (hc *handlerCategory) Register(router *http.ServeMux) {
	router.HandleFunc(constant.URLCategory, hc.Create)
}

func (hc *handlerCategory) Create(w http.ResponseWriter, r *http.Request) {
	dto := model.CreateCategoryDTO{}

	err := hc.serviceCategory.Create(context.Background(), &dto)
	if err != nil {
		// w.WriteHeader()
	}

}
