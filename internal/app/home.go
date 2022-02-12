package app

import (
	"forum/internal/adapters/api"
	home2 "forum/internal/adapters/api/home"
	"forum/internal/service/home"
)

func (a *App) home(post api.ServicePost, category api.ServiceCategory) {
	srv := home.NewService(post, category)
	hnd := home2.NewHandler(a.ctx, srv)
	hnd.Register(a.router)
}
