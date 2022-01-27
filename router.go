package minima

import (
	"regexp"
	"strings"
)

/**
@info The Handler func structure
@property {Response} [res] The response instance
@property {Request} [req] The request instance
*/
type Handler func(res *Response, req *Request)

/**
@info The router structure
@property {map[string][]*mux} [route] The mux routes
*/
type Router struct {
	routes map[string][]*mux
}

/**
@info Make new default router interface
return {Router}
*/
func NewRouter() *Router {
	return &Router{
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
}

/**
@info Compiles path to regex
@param {string} [path] The route path
return {string, []string}
*/
func RegexPath(path string) (string, []string) {
	var items []string
	var Params []string

	for _, part := range strings.Split(path, "/") {
		if strings.HasPrefix(part, ":") {
			Params = append(Params, strings.Trim(part, ":"))
			items = append(items, `([^\/]+)`)
		} else {
			items = append(items, part)
		}
	}

	return "^" + strings.Join(items, `\/`) + "$", Params
}

/**
@info Registers a new route to router interface
@param {string} [path] The route path
return {string, []string}
*/
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

/**
@info Adds route with Get method
@param {string} [path] The route path
@param {...Handler} [handler] The handler for the given route
@returns {*Router}
*/
func (r *Router) Get(path string, handlers ...Handler) *Router {
	r.Register("GET", path, handlers...)
	return r
}

/**
@info Adds route with Post method
@param {string} [path] The route path
@param {...Handler} [handler] The handler for the given route
@returns {*Router}
*/
func (r *Router) Post(path string, handlers ...Handler) *Router {
	r.Register("POST", path, handlers...)
	return r
}

/**
@info Adds route with Put method
@param {string} [path] The route path
@param {...Handler} [handler] The handler for the given route
@returns {*Router}
*/
func (r *Router) Put(path string, handlers ...Handler) *Router {
	r.Register("PUT", path, handlers...)
	return r
}

/**
@info Adds route with Patch method
@param {string} [path] The route path
@param {...Handler} [handler] The handler for the given route
@returns {*Router}
*/
func (r *Router) Patch(path string, handlers ...Handler) {
	r.Register("PATCH", path, handlers...)
}

/**
@info Adds route with Options method
@param {string} [path] The route path
@param {...Handler} [handler] The handler for the given route
@returns {*Router}
*/
func (r *Router) Options(path string, handlers ...Handler) *Router {
	r.Register("OPTIONS", path, handlers...)
	return r
}

/**
@info Adds route with Head method
@param {string} [path] The route path
@param {...Handler} [handler] The handler for the given route
@returns {*Router}
*/
func (r *Router) Head(path string, handlers ...Handler) *Router {
	r.Register("HEAD", path, handlers...)
	return r
}

/**
@info Adds route with Delete method
@param {string} [path] The route path
@param {...Handler} [handler] The handler for the given route
@returns {*Router}
*/
func (r *Router) Delete(path string, handlers ...Handler) *Router {
	r.Register("DELETE", path, handlers...)
	return r
}

/**
@info Returns all the routes in router
@returns {map[string][]*mux}
*/
func (r *Router) GetRouterRoutes() map[string][]*mux {
	return r.routes
}

/**
@info Appends all routes to core router instance
@param {Router} [Router] The router instance to append
@returns {Router}
*/
func (r *Router) UseRouter(Router *Router) *Router {
	for routeType, list := range Router.GetRouterRoutes() {
		r.routes[routeType] = append(r.routes[routeType], list...)
	}

	return r
}

/**
@info Mounts all routes to a specific path
@param {string} [basepath] The prefix route path
@param {Router} [Router] The router instance to append
@returns {Router}
*/
func (r *Router) Mount(basepath string, Router *Router) *Router {
	for routeType, list := range Router.GetRouterRoutes() {
		for _, v := range list {
			v.Path = basepath + v.Path
			r.Register(routeType, v.Path, v.Handlers...)
		}
	}

	return r
}

func (r *Router) next(p map[string]string, next Handler, res *Response, req *Request) {
	Path := req.GetPathURl()
	for key, value := range p {
		req.Params = append(req.Params, &Param{Path, key, value})
	}

	next(res, req)
}
