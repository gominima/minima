package minima

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
@property {[]Handler} [minmiddleware] The standard minima handler stack
@property {[]http.HandlerFunc} [rawmiddleware] The raw net/http minima handler stack
@property {map[string]interface{}} [properties] The properties for the server instance
@property {*Config} [Config] The core config file for middlewares and router instances
@property {*time.Duration} [drain] The router's drain time
*/
type minima struct {
	server        *http.Server
	started       bool
	Timeout       time.Duration
	router        *Router
	minmiddleware []Handler
	testfun       func(http.Handler) http.Handler
	rawmiddleware []http.HandlerFunc
	properties    map[string]interface{}
	Config        *Config
	drain         time.Duration
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
		router: NewRouter(),
		Config: NewConfig(),
		drain:  0,
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
	if match {
		if err := r.ParseForm(); err != nil {
			log.Printf("Error parsing form: %s", err)
			return
		}
		m.testfun(m).ServeHTTP(w,r)
		res := response(w, r, &m.properties)
		req := request(r)
		req.Params = params
               
		m.ServeMiddleware(res, req)
		f(res, req)
	} else {
		res := response(w, r, &m.properties)
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
@info Injects middlewares and routers directly to core instance
@param {*Config} [config] The config instance
@returns {*minima}
*/
func (m *minima) UseConfig(config *Config) *minima {
	m.minmiddleware = append(m.minmiddleware, config.Middleware...)
	m.rawmiddleware = append(m.rawmiddleware, config.HttpHandler...)
	for _, router := range config.Router {
		m.UseRouter(router)
	}
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
@info Injects minima middleware to the stack
@param {...Handler} [handler] The handler stack to append
@returns {}
*/
func (m *minima) Use(handler ...Handler) {
	m.minmiddleware = append(m.minmiddleware, handler...)
}

/**
@info Injects net/http middleware to the stack
@param {...http.HandlerFunc} [handler] The handler stack to append
@returns {}
*/
func (m *minima) UseRaw(handler ...http.HandlerFunc) {
	m.rawmiddleware = append(m.rawmiddleware, handler...)
}

/**
@info Serves and injects the middlewares to minima logic
@param {Response} [res] The minima response instance
@param {Request} [req] The minima req instance
@returns {}
*/
func (m *minima) ServeMiddleware(res *Response, req *Request) {
	if len(m.rawmiddleware) == 0 {
		return
	}
	for _, raw := range m.rawmiddleware {
		raw(res.ref, req.ref)
	}
	if len(m.minmiddleware) == 0 {
		return
	}
	for _, min := range m.minmiddleware {
		min(res, req)
	}
}

func (m*minima) Test(handler func(http.Handler) http.Handler) {
	
	m.testfun = handler
}