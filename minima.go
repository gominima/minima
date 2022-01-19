package minima

import (
	"context"
	"log"
	"net/http"
	"time"
)
/**
 	@info The minima structure
	@property {http.Server} [server] The server instance
	@property {bool} [started] Whether the server has started yet
	@property {time.Duration} [Timeout] The timeout duration
	@property {router} [router] The router instance
	@property {map[string]interface} [properties] The interal properties of the server
	@property {Config} [Config] The config options
	@property {string} [errorPath] The path for errors
	@property {interface} [errorData] The data for errors
	@property {Plugins} [Middleware] The additional middlewares
*/
type minima struct {
	server     *http.Server
	started    bool
	Timeout    time.Duration
	router     *Router
	properties map[string]interface{}
	Config     *Config
	Middleware *Plugins
	drain      time.Duration
}

func New() *minima {
	var router *Router = NewRouter()
	var plugin *Plugins = use()
	var Config *Config = NewConfig()
	var minima *minima = &minima{router: router}
	minima.Middleware = plugin
	minima.drain = 0
	minima.Config = Config
	return minima

}
/**
	@info Listen to the port
	@param {string} [addr] The port
	@returns {error}
*/
func (m *minima) Listen(addr string) error {
	if m.started {
		panic("Minima server instance is already running")
	}
	server := &http.Server{Addr: addr, Handler: m}
	m.server = server
	m.started = true
	return m.server.ListenAndServe()

}
/**
	@info Server HTTP
	@param {http.ResponseWriter} [w] The response writer
	@param {http.Request} [q] The request
	@returns {error}
*/
func (m *minima) ServeHTTP(w http.ResponseWriter, q *http.Request) {
	match := false

	for _, requestQuery := range m.router.routes[q.Method] {
		if isMatchRoute, Params := requestQuery.matchingPath(q.URL.Path); isMatchRoute {
			match = isMatchRoute
			if err := q.ParseForm(); err != nil {
				log.Printf("Error parsing form: %s", err)
				return
			}

			currentRequest := 0

			res := response(w, q, &m.properties)
			req := request(q, &m.properties)
			m.Middleware.ServePlugin(res, req)

			m.router.Next(Params, requestQuery.Handlers[currentRequest], res, req)
			currentRequest++
			break

		}
	}

	if !match {
		w.Write([]byte("No matching route found"))

	}
}

func (m *minima) Get(path string, handler ...Handler) *minima {
	m.router.Get(path, handler...)
	return m
}

func (m *minima) Put(path string, handler ...Handler) *minima {
	m.router.Put(path, handler...)
	return m
}

func (m *minima) Options(path string, handler ...Handler) *minima {
	m.router.Options(path, handler...)
	return m
}

func (m *minima) Head(path string, handler ...Handler) *minima {
	m.router.Head(path, handler...)
	return m
}

func (m *minima) Delete(path string, handler ...Handler) *minima {
	m.router.Delete(path, handler...)
	return m
}

func (m *minima) Patch(path string, handler ...Handler) *minima {
	m.router.Patch(path, handler...)
	return m
}
func (m *minima) UseRouter(router *router) {
=======
func (m *minima) UseRouter(router *Router) *minima {
>>>>>>> 8c7aafb0132fdea03a58145f8ab9901e321e8614
	m.router.UseRouter(router)
	return m

}

func (m *minima) Mount(path string, router *Router) *minima {
	return m

}
}

	m.drain = t
	return m
}

func (m *minima) Shutdown(ctx context.Context) error {
	log.Println("Stopping the server")
	return m.server.Shutdown(ctx)
}

func (m *minima) SetProp(key string, value interface{}) *minima {
	m.properties[key] = value
	return m

}

func (m *minima) GetProp(key string) interface{} {
	return m.properties[key]
}
