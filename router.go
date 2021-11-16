package fiable

import (
	"net/http"
	"regexp"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

type router struct {
	routes []route
}

func Router() *router {
	return &router{}
}

func (r *router) newRoute(method, pattern string, handler http.HandlerFunc) route {
	route := route{}
	route.handler = handler
	route.regex = regexp.MustCompile("^" + pattern + "$")
	route.method = method
	r.routes = append(r.routes, route)
	return route
}

func (r *router) Get(pattern string, handler http.HandlerFunc) {
	r.newRoute("GET", pattern, handler)
}
