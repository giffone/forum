package app

import (
	"forum/internal/adapters/api"
	post2 "forum/internal/adapters/api/post"
	"forum/internal/adapters/repository"
	"forum/internal/service"
	"forum/internal/service/post"
)

func (a *App) post(repo repository.Repo, srvCategory service.Category,
	srvComment service.Comment, srvRatio service.Ratio,
	apiRatio api.Ratio, ses api.Session) service.Post {
	srv := post.NewService(repo, srvRatio, srvCategory)
	post2.NewHandler(srv, srvCategory, srvComment, apiRatio).Register(a.ctx, a.router, ses)
	return srv
}
