package minima

import (
	"net/http"
	"regexp"
	"strings"
)
/**
	@info The Handler
*/
type Handler func(response *Response, request *Request)
/**
	@info Regex's the path
	@param {string} [path] The path
	@returns {string, []string}
*/
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
>>>>>>> 8c7aafb0132fdea03a58145f8ab9901e321e8614
	for routeType, list := range routes {
		r.routes[routeType] = append(r.routes[routeType], list...)
	}
	return r
}

func (r *Router) Mount(basepath string, Router *Router) *Router {
	routes := Router.GetRouterRoutes()
		for _, v := range list {
			v.Path = basepath + v.Path
			r.Register(routeType, v.Path, v.Handlers...)
		}

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
