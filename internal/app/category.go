package app

import (
	"forum/internal/adapters/api"
	category3 "forum/internal/adapters/api/category"
	"forum/internal/adapters/repo/category"
	category2 "forum/internal/service/category"
)

func (a *App) category() api.ServiceCategory {
	str := category.NewStorage(a.db, a.connect, a.query)
	srv := category2.NewService(str)
	hnd := category3.NewHandler(a.ctx, srv)
	hnd.Register(a.router)

	return srv
}
