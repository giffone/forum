package app

import (
	"forum/internal/adapters/api"
	user2 "forum/internal/adapters/api/user"
	"forum/internal/adapters/repository"
	"forum/internal/service/user"
)

func (a *App) user(repo repository.Repo, ses api.Session) {
	srv := user.NewService(repo)
	user2.NewHandler(srv).Register(a.ctx, a.router, ses)
}
