package app

import (
	"forum/internal/adapters/api"
	"forum/internal/adapters/repo/session"
	session2 "forum/internal/service/session"
)

func (a *App) session() api.ServiceSession {
	str := session.NewStorage(a.db, a.connect, a.query)
	return session2.NewService(str)
}
