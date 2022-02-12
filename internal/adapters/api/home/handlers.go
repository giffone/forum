package home

import (
	"context"
	"forum/internal/adapters/api"
	"forum/internal/constant"
	"log"
	"net/http"
	"time"
)

type handlerHome struct {
	ctx         context.Context
	serviceHome api.ServiceHome
}

func NewHandler(ctx context.Context, serviceHome api.ServiceHome) api.Handler {
	return &handlerHome{
		ctx:         ctx,
		serviceHome: serviceHome,
	}
}

func (hh *handlerHome) Register(router *http.ServeMux) {
	router.HandleFunc(constant.URLHome, hh.Home)
	router.HandleFunc(constant.URLFavicon, hh.Favicon)
}

func (hh *handlerHome) Home(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, " ", r.URL.Path)
	if r.Method != "GET" {
		api.ErrorsHTTP(w, "", constant.Code405)
		return
	}
	ctx, cancel := context.WithTimeout(hh.ctx, 5*time.Second)
	defer cancel()
	tmpl, err := api.Parse(constant.PathIndex, constant.PathHead, constant.PathFooter, constant.PathPosts, constant.PathCategories)
	if err != nil {
		api.ErrorsHTTP(w, "", constant.Code500)
		return
	}
	data, err := hh.serviceHome.Get(ctx)
	if err != nil {
		api.ErrorsHTTP(w, "", err.Error())
		return
	}
	api.Execute(w, tmpl, "index", data)
}

func (hh *handlerHome) Favicon(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, " ", r.URL.Path)
	w.Header().Set("Content-Type", "image/x-icon")
	w.Header().Set("Cache-Control", "public, max-age=7776000")
	http.ServeFile(w, r, "assets/ico/favicon.ico")
}
