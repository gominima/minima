package minima

/**
* Minima is a free and open source software under Mit license

Copyright (c) 2021 gominima

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

* Authors @apoorvcodes @megatank58
* Maintainers @Panquesito7 @savioxavier @Shubhaankar-Sharma @apoorvcodes @megatank58
* Thank you for showing interest in minima and for this beautiful community
*/

import (
	"fmt"
	"net/http"
)

type Handler func(res *Response, req *Request)

/**
 * @info The router structure
 * @property {map[string][]*Routes} [routes] The mux routes
 * @property {Handler} [notfound] The handler for the non matching routes
 * @property {[]Handler} [minmiddleware] The minima handler middleware stack
 * @property {[]func(http.Handler)http.Handler} [middleware] The http.Handler middleware stack
 * @property {http.Handler} [handler] The single http.Handler built on chaining the whole middleware stack
 */
type Router struct {
	notfound      Handler
	handler       http.Handler
	middlewares   []func(http.Handler) http.Handler
	routes        map[string]*Routes
}

/**
@info Make new default router interface
return {Router}
*/
func NewRouter() *Router {
	return &Router{
		routes: map[string]*Routes{
			"GET":     NewRoutes(),
			"POST":    NewRoutes(),
			"PUT":     NewRoutes(),
			"DELETE":  NewRoutes(),
			"PATCH":   NewRoutes(),
			"OPTIONS": NewRoutes(),
			"HEAD":    NewRoutes(),
		},
	}
}

/**
@info Registers a new route to router interface
@param {string} [path] The route path
return {string, []string}
*/
func (r *Router) Register(method string, path string, handler Handler) error {
	if r.handler == nil {
		r.buildHandler()
	}
	routes, ok := r.routes[method]
	if !ok {
		return fmt.Errorf("method %s not valid", method)
	}

	routes.Add(path, handler)
	return nil
}

func (r *Router) NotFound(handler Handler) *Router {
	r.notfound = handler
	return r
}

/**
@info Adds route with Get method
@param {string} [path] The route path
@param {...Handler} [handler] The handler for the given route
@returns {*Router}
*/
func (r *Router) Get(path string, handler Handler) *Router {
	r.Register("GET", path, handler)
	return r
}

/**
@info Adds route with Post method
@param {string} [path] The route path
@param {...Handler} [handler] The handler for the given route
@returns {*Router}
*/
func (r *Router) Post(path string, handler Handler) *Router {
	r.Register("POST", path, handler)
	return r
}

/**
@info Adds route with Put method
@param {string} [path] The route path
@param {...Handler} [handler] The handler for the given route
@returns {*Router}
*/
func (r *Router) Put(path string, handler Handler) *Router {
	r.Register("PUT", path, handler)
	return r
}

/**
@info Adds route with Patch method
@param {string} [path] The route path
@param {...Handler} [handler] The handler for the given route
@returns {*Router}
*/
func (r *Router) Patch(path string, handler Handler) {
	r.Register("PATCH", path, handler)
}

/**
@info Adds route with Options method
@param {string} [path] The route path
@param {...Handler} [handler] The handler for the given route
@returns {*Router}
*/
func (r *Router) Options(path string, handler Handler) *Router {
	r.Register("OPTIONS", path, handler)
	return r
}

/**
@info Adds route with Head method
@param {string} [path] The route path
@param {...Handler} [handler] The handler for the given route
@returns {*Router}
*/
func (r *Router) Head(path string, handler Handler) *Router {
	r.Register("HEAD", path, handler)
	return r
}

/**
@info Adds route with Delete method
@param {string} [path] The route path
@param {...Handler} [handler] The handler for the given route
@returns {*Router}
*/
func (r *Router) Delete(path string, handler Handler) *Router {
	r.Register("DELETE", path, handler)
	return r
}

/**
@info Returns all the routes in router
@returns {map[string][]*mux}
*/
func (r *Router) GetRouterRoutes() map[string]*Routes {
	return r.routes
}

/**
@info Appends all routes to core router instance
@param {Router} [Router] The router instance to append
@returns {Router}
*/
func (r *Router) UseRouter(Router *Router) *Router {
	for t, v := range Router.GetRouterRoutes() {
		for i, vl := range v.roots {
			for _, handle := range vl {
				r.Register(t, i, handle.function)
			}
		}
	}
	return r
}

/**
@info Mounts router to a specific path
@param {string} [path] The route path
@param {*Router} [router] Minima router instance
@returns {*Router}
*/
func (r *Router) Mount(path string, Router *Router) *Router {
	for t, v := range Router.GetRouterRoutes() {
		for i, vl := range v.roots {
			for _, handle := range vl {
				r.Register(t, path+i, handle.function)
			}
		}
	}
	return r
}



/**
 * @info Injects net/http middleware to the stack
 * @param {...func(http.Handler)http.Handler} [handler] The handler stack to append
 * @returns {}
 */
func (r *Router) use(handler ...func(http.Handler) http.Handler) {
	if r.handler != nil {
		panic("Minima: Middlewares can't go after the routes are mounted")
	}
	r.middlewares = append(r.middlewares, handler...)
}

//A dummy function that runs at the end of the middleware stack
func (r *Router) middlewareHTTP(w http.ResponseWriter, rq *http.Request) {}

/**
 * @info Builds whole middleware stack chain into single handler
 */
func (r *Router) buildHandler() {
	r.handler = chain(r.middlewares, http.HandlerFunc(r.middlewareHTTP))
}
