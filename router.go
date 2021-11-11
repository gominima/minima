package fiable

import (
	"fmt"
	"net/http"
)

type handler func(Request *Request, Response *Response)

type NextFunc func(NextFunc)

type Route struct {
	url       string
  method    string
	handler      handler
	hashandler bool
  http      *http.ServeMux
 
    }

type router struct {
  routes []*Route
}

func Router () *router{
  router := router{}
  

  return &router

}

func (router *router) addUrl(method string, hashandler bool, url string,  handler handler){
 r := &Route{}
 r.hashandler = hashandler
 r.handler = handler
 r.url = url
 r.method = method
 router.routes = append(router.routes, r)

}

func (r *router) Get(url string, handler handler ){
fmt.Printf("Get route added")
r.addUrl("GET", true, url, handler)
}

