package minima

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

type minima struct {
	server     *http.Server
	started    bool
	Timeout    time.Duration
	router     *router
	properties map[string]interface{}
	Config     *Config
	errorPath  string
	errorData  interface{}
	Middleware *Plugins
}

func New() *minima{
	var router *router = NewRouter()
	var plugin *Plugins = use()
	var Config *Config = NewConfig()
	var minima *minima= &minima{router: router}
	minima.Middleware = plugin
	minima.Config = Config
	minima.errorPath = "../assets/404.html"
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
		path := m.errorPath
		t, err := template.New("404.html").ParseFiles(path)
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, m.errorData)
	}
}

func (m *minima) Get(path string, handler ...Handler) {
	m.router.Get(path, handler...)
}

func (m *minima) Set404(path string, data interface{}) *minima{
	m.errorPath = path
	m.errorData = data
	return m
}
func (m *minima) Use(handler Handler) {
	m.Middleware.AddPlugin(handler)
}
func (m *minima) UseRouter(router *router) {
	m.router.UseRouter(router)

}

func (m *minima) UseConfig(config *Config) {
	for _, v := range config.Middleware {
		m.Middleware.plugin = append(m.Middleware.plugin, &Middleware{handler: v})
	}
	m.Config.Logger = config.Logger
	m.router.UseRouter(config.Router)
}