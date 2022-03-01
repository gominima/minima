package minima

import "net/http"

type middlewarefunc func(http.Handler) http.Handler

type middleware interface {
	Middleware(handler http.Handler) http.Handler
}

func (mw middlewarefunc) Middleware(handler http.Handler) http.Handler {
	return mw(handler)
}

