package fiable

import (
	"context"
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

func NewRouter() *router{
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
func RegexPath(path string) (string, []string){
	 var items []string
	 var Params []string
	var regexPath string
	 parts := strings.Split(path, "/")
	 for _, part := range parts {
		 if strings.HasPrefix(part, ":"){
			 name := strings.Trim(part, ":")
			 Params = append(Params, name)
			 items = append(items, `([^\/]+)`)

		 } else {
			 items = append(items, part)
		 }

	 }
	regexPath = "^" +strings.Join(items, `\/`) +  `/?` + "$"
	return regexPath, Params
}



func (r *router) Register(method string, path string, handlers ...Handler) *mux{
 reg, Params := RegexPath(path)
 var newroute = &mux{
	 Path: path,
	 Handlers: handlers,
	 Regex: regexp.MustCompile(reg),
	 Params: Params,
 }
 r.routes[method] = append(r.routes[method], newroute)
	return newroute
}

func (r*router) Get(path string, handlers ...Handler){
 r.Register("GET", path, handlers...)
}



func (r *router) Next(p map[string]string, next Handler, res *Response, req *Request) {
	ctx := context.Background()
	for k, v := range p {
		ctx = context.WithValue(ctx, k, v)
		fmt.Println(ctx)
		addParam := &Param{
		 path: req.ref.URL.Path,
		 value: v,
		 ctx: ctx,
		}
		req.Params = append(req.Params, addParam)
	}
	
	
        next(res, req)
}

