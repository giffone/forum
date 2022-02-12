package app

import (
	"forum/internal/adapters/api"
	post3 "forum/internal/adapters/api/post"
	"forum/internal/adapters/repo/post"
	post2 "forum/internal/service/post"
)

func (a *App) post() api.ServicePost {
	str := post.NewStorage(a.db, a.connect, a.query)
	srv := post2.NewService(str)
	hnd := post3.NewHandler(a.ctx, srv)
	hnd.Register(a.router)

	return srv
}
