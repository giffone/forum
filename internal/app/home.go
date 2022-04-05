package app

import (
	"forum/internal/adapters/api"
	"forum/internal/adapters/api/home"
	"forum/internal/service"
)

func (a *App) home(srvPost service.Post, srvCategory service.Category, ses api.Session) {
	home.NewHandler(srvPost, srvCategory).Register(a.ctx, a.router, ses)
}
