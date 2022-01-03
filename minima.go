package minima

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

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

func (m *minima) Listen(addr string) error {
	server := &http.Server{Addr: addr, Handler: m}
	if m.started {
		fmt.Errorf("Server is already running", m)
	}
	m.server = server
	m.started = true
	return m.server.ListenAndServe()
}

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

func (m *minima) Post(path string, handler ...Handler) *minima {
	m.router.Post(path, handler...)
	return m
}

func (m *minima) Use(handler Handler) *minima {
	m.Middleware.AddPlugin(handler)
	return m
}
func (m *minima) UseRouter(router *Router) *minima {
	m.router.UseRouter(router)
	return m

}

func (m *minima) UseConfig(config *Config) *minima {
	for _, v := range config.Middleware {
		m.Middleware.plugin = append(m.Middleware.plugin, &Middleware{handler: v})
	}
	m.Config.Logger = config.Logger
	m.router.UseRouter(config.Router)
	return m
}

func (m *minima) ShutdownTimeout(t time.Duration) *minima {
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
