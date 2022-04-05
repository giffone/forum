package app

import (
	"forum/internal/adapters/api"
	"forum/internal/adapters/api/account"
	"forum/internal/service"
)

func (a *App) account(srvPost service.Post, srvCategory service.Category,
	srvComment service.Comment, apiRatio api.Ratio, ses api.Session) {
	account.NewHandler(srvPost, srvCategory, srvComment, apiRatio).Register(a.ctx, a.router, ses)
}
