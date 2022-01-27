package minima

import (
	"context"
	"net/http"
)

type Minima interface {
	//Minima interface is built over net/http so every middleware is compatible with it

	//Initializes net/http server with address
	Listen(address string) error

	//Handler interface
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

type Res interface {
	//Response interface is built over http.ResponseWriter for easy and better utility

	//Returns minima.OutgoingHeader interface
	Header() *OutgoingHeader

	//Utility functions for easier usage
	Send(content string) *Response      //send content
	WriteBytes(bytes []byte) error      //writes bytes to the page
	Json(content interface{}) *Response //sends data in json format
	Error(status int, str string) *Response

	//This functions return http.ResponseWriter instace that means you could use any of your alrady written middlewares in minima!!
	Raw() http.ResponseWriter

	//Renders a html file with data to the page
	Render(path string, data interface{}) *Response

	//Redirects to given url
	Redirect(url string) *Response

	//Sets Header status
	Status(code int) *Response
}

type Req interface {
	//Minima request interface is built on http.Request

	//Returns param from route url
	GetParam(name string) string

	//Returns path url from the route
	GetPathURL() string

	//Returns raw request body
	Body() map[string][]string

	//Finds given key value from body and returns it
	GetBodyValue(key string) []string

	//Returns instance of minima.IncomingHeader for incoming header requests
	Header() *IncomingHeader

	//Returns route method ex.get,post
	Method() string

	//Gets query params from route and returns it
	GetQuery(key string) string
}
