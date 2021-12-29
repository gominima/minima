package minima

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type Handler func(response *Response, request *Request)

type router struct {
	NotFound http.HandlerFunc
	routes   map[string][]*mux
}

func NewRouter() *router {
	router := &router{
		routes: map[string][]*mux{
			"GET":     make([]*mux, 0),
			"POST":    make([]*mux, 0),
			"PUT":     make([]*mux, 0),
			"DELETE":  make([]*mux, 0),
			"PATCH":   make([]*mux, 0),
			"OPTIONS": make([]*mux, 0),
			"HEAD":    make([]*mux, 0),
		},
	}
	return router
}
func RegexPath(path string) (string, []string) {
	var items []string
	var Params []string
	var regexPath string
	parts := strings.Split(path, "/")
	for _, part := range parts {
		if strings.HasPrefix(part, ":") {
			name := strings.Trim(part, ":")
			Params = append(Params, name)
			items = append(items, `([^\/]+)`)

		} else {
			items = append(items, part)
		}

	}
	regexPath = "^" + strings.Join(items, `\/`) + "$"
	fmt.Print(regexPath)
	return regexPath, Params
}

func (r *router) Register(method string, path string, handlers ...Handler) *mux {
	reg, Params := RegexPath(path)
	var newroute = &mux{
		Path:     path,
		Handlers: handlers,
		Regex:    regexp.MustCompile(reg),
		Params:   Params,
	}
	r.routes[method] = append(r.routes[method], newroute)
	return newroute
}

func (r *router) Get(path string, handlers ...Handler) {
	r.Register("GET", path, handlers...)
}

func (r *router) GetRouterRoutes() map[string][]*mux {
	return r.routes
}

func (r *router) UseRouter(router *router) {
	routes := router.GetRouterRoutes()
	for routeType, list := range routes {
		r.routes[routeType] = append(r.routes[routeType], list...)
	}
}

func (r *router) Next(p map[string]string, next Handler, res *Response, req *Request) {
	Path := req.GetPathURl()
	for k, v := range p {

		addParam := &Param{
			Path:  Path,
			key:   k,
			value: v,
		}
		req.Params = append(req.Params, addParam)
	}

	next(res, req)
}
