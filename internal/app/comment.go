package app

import (
	"forum/internal/adapters/repository"
	"forum/internal/service"
	"forum/internal/service/comment"
)

func (a *App) comment(repo repository.Repo, sLike service.Ratio) service.Comment {
	srv := comment.NewService(repo, sLike)
	return srv
}
