package app

import (
	"forum/internal/adapters/api"
	ratio2 "forum/internal/adapters/api/ratio"
	"forum/internal/adapters/repository"
	"forum/internal/service"
	"forum/internal/service/ratio"
)

func (a *App) ratio(repo repository.Repo) (service.Ratio, api.Ratio) {
	srvRatio := ratio.NewService(repo)
	return srvRatio, ratio2.NewRatio(srvRatio)
}
