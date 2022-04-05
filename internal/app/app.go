package app

import (
	"context"
	"database/sql"
	"net/http"
)

type App struct {
	router *http.ServeMux
	ctx    context.Context
}

func NewApp(ctx context.Context) *App {
	return &App{
		ctx:    ctx,
		router: http.NewServeMux(),
	}
}

func (a *App) Run(driver string) (*sql.DB, *http.ServeMux, string) {
	repo := switcher(a.ctx, driver)
	db, port, _ := repo.ExportSettings()

	srvCategory := a.category(repo)
	session := a.session(repo)
	srvRatio, apiRatio := a.ratio(repo)
	srvComment := a.comment(repo, srvRatio)
	a.user(repo, session)
	srvPost := a.post(repo, srvCategory, srvComment, srvRatio, apiRatio, session)
	a.home(srvPost, srvCategory, session)
	a.account(srvPost, srvCategory, srvComment, apiRatio, session)

	dir := http.Dir("internal/web/assets")
	dirHandler := http.StripPrefix("/assets/", http.FileServer(dir))
	a.router.Handle("/assets/", dirHandler)

	// FOR TEST ONLY
	//_, _, schema := repo.ExportSettings()
	//repository.NewLoremIpsum().Run(db, schema)
	// FOR TEST ONLY

	return db, a.router, port
}
