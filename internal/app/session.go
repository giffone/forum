package app

import (
	"forum/internal/adapters/api"
	"forum/internal/adapters/api/middleware"
	"forum/internal/adapters/repository"
	"forum/internal/service/session"
)

func (a *App) session(repo repository.Repo) api.Session {
	srv := session.NewService(repo)
	return middleware.NewSession(srv)
}
