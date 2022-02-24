package minima

import "net/http"

type middlewarefunc func(http.Handler) http.Handler

type middleware interface {
	Middleware(handler http.Handler) http.Handler
}

func (mw middlewarefunc) Middleware(handler http.Handler) http.Handler {
	return mw(handler)
}

/**
@info Injects net/http middleware to the stack
@param {...http.HandlerFunc} [handler] The handler stack to append
@returns {}
*/
func (m *minima) UseRaw(handler ...middlewarefunc) {
	for _, fnc := range handler {
		m.rawmiddleware = append(m.rawmiddleware, fnc)
	}
}
