package app

import (
	"forum/internal/adapters/api"
	user3 "forum/internal/adapters/api/user"
	"forum/internal/adapters/repo/user"
	user2 "forum/internal/service/user"
)

func (a *App) user(session api.ServiceSession) {
	str := user.NewStorage(a.db, a.connect, a.query)
	srv := user2.NewService(str)
	hnd := user3.NewHandler(a.ctx, srv, session)
	hnd.Register(a.router)

}
