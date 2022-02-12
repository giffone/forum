package post

import (
	"context"
	"forum/internal/adapters/api"
	"forum/internal/constant"
	"net/http"
)

type handlerPost struct {
	ctx         context.Context
	servicePost api.ServicePost
}

func NewHandler(ctx context.Context, servicePost api.ServicePost) api.Handler {
	return &handlerPost{
		ctx:         ctx,
		servicePost: servicePost,
	}
}

func (hp *handlerPost) Register(router *http.ServeMux) {
	router.HandleFunc(constant.URLPost, hp.GetPost)
	//router.HandleFunc(postsURL, hp.GetAllPost)
}

func (hp *handlerPost) GetPost(w http.ResponseWriter, r *http.Request) {
	//w.Write()
}
