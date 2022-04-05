package app

import (
	"forum/internal/adapters/repository"
	"forum/internal/service"
	"forum/internal/service/category"
)

func (a *App) category(repo repository.Repo) service.Category {
	return category.NewService(repo)
}
