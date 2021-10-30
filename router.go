package fiable

import ("regexp")


type Middleware func()


type Route struct {
	url        *regexp.Regexp
	middleware      Middleware
	hasMiddleware bool
      }

type router struct {
  routes  map[string][]*Route
}

func Router () *router{
  router := router{}
  router.routes = make(map[string][]*Route)
  router.routes["get"] = []*Route{}
  router.routes["post"] = []*Route{}
  router.routes["delete"] = []*Route{}
  router.routes["put"] = []*Route{}
  router.routes["patch"] = []*Route{}

  return &router

}

func (router *router) addUrl(method string, hasMiddleware bool, url *regexp.Regexp, middleware Middleware){
 r := &Route{}
 r.hasMiddleware = hasMiddleware
 r.middleware = middleware
 r.url = url
 router.routes[method] = append(router.routes[method], r)

}

