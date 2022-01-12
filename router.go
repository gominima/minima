package minima

import (
	"net/http"
	"regexp"
	"strings"
)

type Handler func(response *Response, request *Request)

type Router struct {
	NotFound http.HandlerFunc
	routes   map[string][]*mux
}

func NewRouter() *Router {
	Router := &Router{
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
	return Router
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
	return regexPath, Params
}

func (r *Router) Register(method string, path string, handlers ...Handler) *mux {
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

func (r *Router) Get(path string, handlers ...Handler) *Router {
	r.Register("GET", path, handlers...)
	return r
}

func (r *Router) Post(path string, handlers ...Handler) *Router {
	r.Register("POST", path, handlers...)
	return r
}

func (r *Router) Put(path string, handlers ...Handler) *Router {
	r.Register("PUT", path, handlers...)
	return r
}

func (r *Router) Patch(path string, handlers ...Handler) {
	r.Register("PATCH", path, handlers...)
}

func (r *Router) Options(path string, handlers ...Handler) *Router {
	r.Register("Options", path, handlers...)
	return r
}

func (r *Router) Head(path string, handlers ...Handler) *Router {
	r.Register("HEAD", path, handlers...)
	return r
}

func (r *Router) Delete(path string, handlers ...Handler) *Router {
	r.Register("DELETE", path, handlers...)
	return r
}

func (r *Router) GetRouterRoutes() map[string][]*mux {
	return r.routes
}

func (r *Router) UseRouter(Router *Router) *Router {
	routes := Router.GetRouterRoutes()
	for routeType, list := range routes {
		r.routes[routeType] = append(r.routes[routeType], list...)
	}
	return r
}

func (r *Router) Next(p map[string]string, next Handler, res *Response, req *Request) {
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
