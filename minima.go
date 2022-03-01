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
	"context"
	"log"
	"net/http"
	"time"
)

/**
@info The framework structure
@property {*http.Server} [server] The net/http stock server
@property {bool} [started] Whether the server has started or not
@property {*time.Duration} [Timeout] The router's breathing time
@property {*Router} [router] The core router instance running with the server
@property {map[string]interface{}} [properties] The properties for the server instance
@property {*Config} [Config] The core config file for middlewares and router instances
@property {*time.Duration} [drain] The router's drain time
*/
type minima struct {
	server     *http.Server
	started    bool
	Timeout    time.Duration
	router     *Router
	properties map[string]interface{}
	drain      time.Duration
}

/**
@info Make a new default minima instance
@example `
func main() {
	app := minima.New()

	app.Get("/", func(res *minima.Response, req *minima.Request) {
		res.Status(200).Send("Hello World")
	})

	app.Listen(":3000")
}
`
@returns {minima}
*/
func New() *minima {
	return &minima{
		drain:  0,
		router: NewRouter(),
	}
}

/**
@info Starts the actual http server
@param {string} [addr] The port for the server instance to run on
@returns {error}
*/
func (m *minima) Listen(addr string) error {
	if m.started {
		log.Panicf("Minimia's instance is already running at %s.", m.server.Addr)
	}
	m.server = &http.Server{Addr: addr, Handler: m}
	m.started = true

	return m.server.ListenAndServe()
}

/**
@info Injects the actual minima server logic to http
@param {http.ResponseWriter} [w] The net/http response instance
@param {http.Request} [r] The net/http request instance
@returns {}
*/
func (m *minima) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, params, match := m.router.routes[r.Method].Get(r.URL.Path)
	res := response(w, r)
	req := request(r)
	if match {
		if err := r.ParseForm(); err != nil {
			log.Printf("Error parsing form: %s", err)
			return
		}
		if m.router.handler != nil {
			m.router.handler.ServeHTTP(w, r)
		}
		req.Params = params
		f(res, req)
	} else {
		res := response(w, r)
		req := request(r)
		if m.router.notfound != nil {
			m.router.notfound(res, req)
		} else {
			w.Write([]byte("No matching route found"))
		}
	}
}

/**
@info Adds route with Get method
@param {string} [path] The route path
@param {...Handler} [handler] The handler for the given route
@returns {*minima}
*/
func (m *minima) Get(path string, handler Handler) *minima {

	m.router.Get(path, handler)
	return m
}

/**
@info Adds route with Put method
@param {string} [path] The route path
@param {...Handler} [handler] The handler for the given route
@returns {*minima}
*/
func (m *minima) Put(path string, handler Handler) *minima {
	m.router.Put(path, handler)
	return m
}

/**
@info Adds route with Options method
@param {string} [path] The route path
@param {...Handler} [handler] The handler for the given route
@returns {*minima}
*/
func (m *minima) Options(path string, handler Handler) *minima {
	m.router.Options(path, handler)
	return m
}

/**
@info Adds route with Head method
@param {string} [path] The route path
@param {...Handler} [handler] The handler for the given route
@returns {*minima}
*/
func (m *minima) Head(path string, handler Handler) *minima {
	m.router.Head(path, handler)
	return m
}

/**
@info Adds route with Delete method
@param {string} [path] The route path
@param {...Handler} [handler] The handler for the given route
@returns {*minima}
*/
func (m *minima) Delete(path string, handler Handler) *minima {
	m.router.Delete(path, handler)
	return m
}

/**
@info Adds route with Patch method
@param {string} [path] The route path
@param {...Handler} [handler] The handler for the given route
@returns {*minima}
*/
func (m *minima) Patch(path string, handler Handler) *minima {
	m.router.Patch(path, handler)
	return m
}

/**
@info Adds route with Post method
@param {string} [path] The route path
@param {...Handler} [handler] The handler for the given route
@returns {*minima}
*/
func (m *minima) Post(path string, handler Handler) *minima {
	m.router.Post(path, handler)
	return m
}

/**
@info Injects the NotFound handler to the minima instance
@param {Handler} [handler] Minima handler instance
@returns {*minima}
*/
func (m *minima) NotFound(handler Handler) *minima {
	m.router.NotFound(handler)
	return m
}

/**
@info Injects the routes from the router to core stack
@param {*Router} [router] Minima router instance
@returns {*minima}
*/
func (m *minima) UseRouter(router *Router) *minima {
	m.router.UseRouter(router)
	return m
}

/**
@info Mounts router to a specific path
@param {string} [path] The route path
@param {*Router} [router] Minima router instance
@returns {*minima}
*/
func (m *minima) Mount(path string, router *Router) *minima {
	m.router.Mount(path, router)
	return m
}

/**
@info The drain timeout for the core instance
@param {time.Duration} [time] The time period for drain
@returns {*minima}
*/
func (m *minima) ShutdownTimeout(t time.Duration) *minima {
	m.drain = t
	return m
}

/**
@info Shutdowns the core instance
@param {context.Context} [ctx] The context for shutdown
@returns {error}
*/
func (m *minima) Shutdown(ctx context.Context) error {
	log.Println("Stopping the server")
	return m.server.Shutdown(ctx)
}

/**
@info Declares prop for core properties
@param {string} [key] Key for the prop
@param {interface{}} [value] Value of the prop
@returns {*minima}
*/
func (m *minima) SetProp(key string, value interface{}) *minima {
	m.properties[key] = value
	return m
}

/**
@info Gets prop from core properties
@param {string} [key] Key for the prop
@returns {interface{}}
*/
func (m *minima) GetProp(key string) interface{} {
	return m.properties[key]
}

/**
 * @info Injects Minima middleware to the stack
 * @param {...Handler} [handler] The handler stack to append
 * @returns {}
 */
func (m *minima) Use(handler ...Handler) *minima {
	m.router.use(handler...)
	return m
}

/**
 * @info Injects net/http middleware to the stack
 * @param {...http.HandlerFunc} [handler] The handler stack to append
 * @returns {}
 */
func (m *minima) UseRaw(handler ...func(http.Handler) http.Handler) *minima {
	m.router.useRaw(handler...)
	return m
}
