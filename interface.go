package minima

import (
	"context"
	"net/http"
)

type Minima interface {
	//Minima interface is built over net/http so every middleware is compatible with it

	//initializes net/http server with address
	Listen(address string) error

	//handler interface
	ServeHTTP(w http.ResponseWriter, q *http.Request)

	//Router methods
	Get(path string, handler ...Handler) *minima
	Patch(path string, handler ...Handler) *minima
	Post(path string, handler ...Handler) *minima
	Put(path string, handler ...Handler) *minima
	Options(path string, handler ...Handler) *minima
	Head(path string, handler ...Handler) *minima
	Delete(path string, handler ...Handler) *minima

	//Takes middlewares as a param and adds them to routes
	//middlewares initializes before route handler is mounted
	Use(handler Handler) *minima

	//Takes minima.Router as param and adds the routes from router to main instance
	UseRouter(router *Router) *minima

	//Works as a config for minima, you can add multiple middlewares and routers at once
	UseConfig(config *Config) *minima

	//Shutdowns the net/http server
	Shutdown(ctx context.Context) error

	//Prop methods
	SetProp(key string, value interface{}) *minima
	GetProp(key string) interface{}
}
