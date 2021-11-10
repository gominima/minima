package fiable

import (
	"net/http"

)


type Middleware func(Request request, Response response)

type NextFunc func(NextFunc)

type Route struct {
	url       string
  method    string
	middleware      Middleware
	hasMiddleware bool
  handler        *http.Handler
    }

type router struct {
  routes []*Route
}

func Router () *router{
  router := router{}
  

  return &router

}

func (router *router) addUrl(method string, hasMiddleware bool, url string,  middleware Middleware){
 r := &Route{}
 r.hasMiddleware = hasMiddleware
 r.middleware = middleware
 r.url = url
 r.method = method
 router.routes = append(router.routes, r)

}

func (r *router) Get(url string, handler Middleware ){
r.addUrl("GET", true, url, handler)
}

