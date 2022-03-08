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
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

/**
 * @info The framework structure
 * @property {*http.Server} [server] The net/http stock server
 * @property {bool} [started] Whether the server has started or not
 * @property {*time.Duration} [Timeout] The router's breathing time
 * @property {*Router} [router] The core router instance running with the server
 * @property {map[string]interface{}} [properties] The properties for the server instance
 * @property {*Config} [Config] The core config file for middlewares and router instances
 * @property {*time.Duration} [drain] The router's drain time
 */
type Minima struct {
	server     *http.Server
	started    bool
	Timeout    time.Duration
	router     *Router
	properties map[string]interface{}
	drain      time.Duration
}

var (
	red    = "\u001b[31m"
	green  = "\u001b[32m"
	yellow = "\u001b[33m"
	blue   = "\u001b[34m"
	reset  = "\u001b[0m"
)

const (
	version = "1.1.2"
)

/**
 * @info Make a new default minima instance
 * @example `
func main() {
	app := minima.New()

	app.Get("/", func(res *minima.Response, req *minima.Request) {
		res.Status(200).Send("Hello World")
	})

	app.Listen(":3000")
}
`
 * @returns {minima}
*/
func New() *Minima {
	m := &Minima{
		drain:  0,
		router: NewRouter(),
	}
	m.router.isCache = false
	return m
}

/**
 * @info Starts the actual http server
 * @param {string} [addr] The port for the server instance to run on
 * @returns {error}
 */
func (m *Minima) Listen(addr string) error {
	if m.started {
		log.Panicf("Minima's instance is already running at %s.", m.server.Addr)
	}
	m.server = &http.Server{Addr: addr, Handler: m}
	m.started = true

	banner := fmt.Sprintf(`	
%s  __  __  _  __  _  _  __  __   ____  
%s |  \/  || ||  \| || ||  \/  | / () \ 
%s |_|\/|_||_||_|\__||_||_|\/|_|/__/\__\ %s
%s The Go framework to scale
%s___________________________________________
%s Server started at port %s %v
%s                                                                      
`, green, green, green, version, blue, red, blue, yellow, addr, reset)

	fmt.Println(banner)

	return m.server.ListenAndServe()
}

/**
 * @info Injects the actual minima server logic to http
 * @param {http.ResponseWriter} [w] The net/http response instance
 * @param {http.Request} [r] The net/http request instance
 * @returns {}
 */
func (m *Minima) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, params := m.router.routes[r.Method].GetNode(r.URL.Path)

	if f != nil {
		handler := buildHandler(f.handler, params)
		if err := r.ParseForm(); err != nil {
			log.Printf("Error parsing form: %s", err)
			return
		}
		if m.router.handler != nil {
			m.router.handler.ServeHTTP(w, r)
		}
		handler.ServeHTTP(w, r)
	} else {
		if m.router.notfound != nil {
			buildHandler(m.router.notfound, nil).ServeHTTP(w, r)
		} else {
			w.Write([]byte("No matching route found"))
		}
	}
}

/**
 * @info Adds route with Get method
 * @param {string} [path] The route path
 * @param {...Handler} [handler] The handler for the given route
 * @returns {*minima}
 */
func (m *Minima) Get(path string, handler Handler) *Minima {

	m.router.Get(path, handler)
	return m
}

/**
 * @info Adds route with Put method
 * @param {string} [path] The route path
 * @param {...Handler} [handler] The handler for the given route
 * @returns {*minima}
 */
func (m *Minima) Put(path string, handler Handler) *Minima {
	m.router.Put(path, handler)
	return m
}

/**
 * @info Adds route with Options method
 * @param {string} [path] The route path
 * @param {...Handler} [handler] The handler for the given route
 * @returns {*minima}
 */
func (m *Minima) Options(path string, handler Handler) *Minima {
	m.router.Options(path, handler)
	return m
}

/**
 * @info Adds route with Head method
 * @param {string} [path] The route path
 * @param {...Handler} [handler] The handler for the given route
 * @returns {*minima}
 */
func (m *Minima) Head(path string, handler Handler) *Minima {
	m.router.Head(path, handler)
	return m
}

/**
 * @info Adds route with Delete method
 * @param {string} [path] The route path
 * @param {...Handler} [handler] The handler for the given route
 * @returns {*minima}
 */
func (m *Minima) Delete(path string, handler Handler) *Minima {
	m.router.Delete(path, handler)
	return m
}

/**
 * @info Adds route with Patch method
 * @param {string} [path] The route path
 * @param {...Handler} [handler] The handler for the given route
 * @returns {*minima}
 */
func (m *Minima) Patch(path string, handler Handler) *Minima {
	m.router.Patch(path, handler)
	return m
}

/**
 * @info Adds route with Post method
 * @param {string} [path] The route path
 * @param {...Handler} [handler] The handler for the given route
 * @returns {*minima}
 */
func (m *Minima) Post(path string, handler Handler) *Minima {
	m.router.Post(path, handler)
	return m
}

/**
 * @info Injects the NotFound handler to the minima instance
 * @param {Handler} [handler] Minima handler instance
 * @returns {*minima}
 */
func (m *Minima) NotFound(handler Handler) *Minima {
	m.router.NotFound(handler)
	return m
}

/**
 * @info Injects the routes from the router to core stack
 * @param {*Router} [router] Minima router instance
 * @returns {*minima}
 */
func (m *Minima) UseRouter(router *Router) *Minima {
	m.router.UseRouter(router)
	return m
}

/**
 * @info The drain timeout for the core instance
 * @param {time.Duration} [time] The time period for drain
 * @returns {*minima}
 */
func (m *Minima) ShutdownTimeout(t time.Duration) *Minima {
	m.drain = t
	return m
}

/**
 * @info Shutdowns the core instance
 * @param {context.Context} [ctx] The context for shutdown
 * @returns {error}
 */
func (m *Minima) Shutdown(ctx context.Context) error {
	log.Println("Stopping the server")
	return m.server.Shutdown(ctx)
}

/**
 * @info Declares prop for core properties
 * @param {string} [key] Key for the prop
 * @param {interface{}} [value] Value of the prop
 * @returns {*minima}
 */
func (m *Minima) SetProp(key string, value interface{}) *Minima {
	m.properties[key] = value
	return m
}

/**
 * @info Gets prop from core properties
 * @param {string} [key] Key for the prop
 * @returns {interface{}}
 */
func (m *Minima) GetProp(key string) interface{} {
	return m.properties[key]
}

/**
 * @info Injects net/http middleware to the stack
 * @param {...http.HandlerFunc} [handler] The handler stack to append
 * @returns {}
 */
func (m *Minima) UseRaw(handler ...func(http.Handler) http.Handler) *Minima {
	m.router.use(handler...)
	return m
}

/**
 * @info Injects minima middleware to the stack
 * @param {Handler} [handler] The handler stack to append
 * @returns {}
 */
func (m *Minima) Use(handler Handler) *Minima {
	m.router.use(build(handler, nil))
	return m
}

/**
 * @info Injects minima group to main router stack
 * @param {Group} [grp] The minima group to append
 * @returns {}
 */
func (m *Minima) UseGroup(grp *Group) *Minima {
	for _, v := range grp.GetGroupRoutes() {
		fmt.Print(v.method)
		m.router.routes[v.method].InsertNode(v.path, v.handler)
	}
	return m
}

/**
 * @info Injects a static file to minima instance
 * @param {string} [pth] The route path for static serve
 * @param {string} [dir] The dir of the file
 * @returns {}
 */
func (m *Minima) File(pth string, dir string) {
	m.Get(pth, func(res *Response, req *Request) {
		res.File(dir)
	})
}

/**
 * @info Injects a static directory to minima instance
 * @param {string} [pth] The route path for static serve
 * @param {string} [dir] The dir of the static folder
 * @returns {}
 */
func (m *Minima) Static(pth string, dir string) {
	if dir == "" {
		dir = "./"
	}
	files, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatal(err)
    }
    for _, f := range files {
		path := []string{pth, "/", f.Name()}
		dr := []string{dir, "/", f.Name()}
        m.Get(strings.Join(path, ""), func(res *Response, req *Request) {
			res.File(strings.Join(dr, ""))
		})
    }
}