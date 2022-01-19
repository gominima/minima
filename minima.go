package minima

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
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
	router     *router
	properties map[string]interface{}
	Config     *Config
	errorPath  string
	errorData  interface{}
	Middleware *Plugins
}
/**
	@info Make a new minima instance
	@returns {minima} Minima instance
*/
func New() *minima {
	var router *router = NewRouter()
	var plugin *Plugins = use()
	var Config *Config = NewConfig()
	var minima *minima= &minima{router: router}
	minima.Middleware = plugin
	minima.Config = Config
	minima.errorPath = "../assets/404.html"
	return minima

}
/**
	@info Listen to the port
	@param {string} [addr] The port
	@returns {error}
*/
func (m *minima) Listen(addr string) error {
	server := &http.Server{Addr: addr, Handler: m}
	if m.started {
		fmt.Errorf("Server is already running", m)
	}
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
		path := m.errorPath
		t, err := template.New("404.html").ParseFiles(path)
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, m.errorData)
	}
}
/**
	@info Handle a GET request
	@param {string} [path] The path
	@param {...Handler} [handler] The handler
*/
func (m *minima) Get(path string, handler ...Handler) {
	m.router.Get(path, handler...)
}
/**
	@info Set 404 page
	@param {string} [path] The path
	@param {interface} [data] The data
	@returns {minima}
*/
func (m *minima) Set404(path string, data interface{}) *minima{
	m.errorPath = path
	m.errorData = data
	return m
}
/**
	@info Use a minima plugin
	@param {Handler} [handler] The handler
*/
func (m *minima) Use(handler Handler) {
	m.Middleware.AddPlugin(handler)
}
/**
	@info Use a router
	@param {router} [router] The router
*/
func (m *minima) UseRouter(router *router) {
	m.router.UseRouter(router)

}
/**
	@info Use a config
	@param {Config} [config] The config
*/
func (m *minima) UseConfig(config *Config) {
	for _, v := range config.Middleware {
		m.Middleware.plugin = append(m.Middleware.plugin, &Middleware{handler: v})
	}
	m.Config.Logger = config.Logger
	m.router.UseRouter(config.Router)
}