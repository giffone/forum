package app

import (
	"context"
	"database/sql"
	"forum/internal/adapters/repo"
	"forum/internal/config"
	"net/http"
)

type App struct {
	db      *sql.DB
	connect *config.Conn
	query   *config.Query
	router  *http.ServeMux
	ctx     context.Context
}

func NewApp(ctx context.Context) *App {
	return &App{
		ctx:    ctx,
		router: http.NewServeMux(),
	}
}

func (a *App) Run(driver string) (*sql.DB, *http.ServeMux, string) {
	a.db, a.connect, a.query = Switcher(a.ctx, driver)

	post := a.post()
	category := a.category()
	session := a.session()
	a.user(session)
	a.home(post, category)

	dir := http.Dir("internal/web/assets")
	dirHandler := http.StripPrefix("/assets/", http.FileServer(dir))
	a.router.Handle("/assets/", dirHandler)

	// FOR TEST ONLY
	repo.InsertTestUsers(a.db, a.query) // adding random users
	repo.InsertTestPosts(a.db, a.query) // adding random posts
	// FOR TEST ONLY

	return a.db, a.router, a.connect.Port
}
