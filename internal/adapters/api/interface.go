package api

import (
	"context"
	"forum/internal/object"
	"net/http"
)

type Handler interface {
	Register(ctx context.Context, router *http.ServeMux, session Session)
}

type Session interface { // middleware for handlers
	Create(ctx context.Context, w http.ResponseWriter, id int, method string) object.Status
	Apply(ctx context.Context, fn func(context.Context, Session,
		http.ResponseWriter, *http.Request)) http.HandlerFunc
	Check(ctx context.Context, fn func(context.Context,
		*object.Cookie, object.Status, http.ResponseWriter,
		*http.Request)) http.HandlerFunc
	End(w http.ResponseWriter) object.Status
}

type Ratio interface {
	Rate(ctx context.Context, ck *object.Cookie, r *http.Request) object.Status
}
