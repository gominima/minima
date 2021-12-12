package fiable

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"


)

type fiable struct {
	server  *http.Server
	started bool
	Timeout time.Duration
	router  *router
	properties map[string]interface{}
	errorPath  string
	errorData interface{}
	Middleware *Plugins
}


func New() *fiable {
	var router *router = NewRouter()
	var plugin *Plugins = use()
	var fiable *fiable = &fiable{router: router }
	fiable.Middleware = plugin
	fiable.errorPath =  "../assets/404.html"
	return fiable

}


func (f *fiable) Listen(addr string) error {
	server := &http.Server{Addr: addr , Handler: f}
	if f.started {
		fmt.Errorf("Server is already running", f)
	}
	f.server = server
	f.started = true
	return f.server.ListenAndServe()
}

func (f* fiable) ServeHTTP(w http.ResponseWriter, q *http.Request){
	match := false

	for _, requestQuery := range f.router.routes[q.Method] {
		if isMatchRoute,Params := requestQuery.matchingPath(q.URL.Path); isMatchRoute {
			match = isMatchRoute
			if err := q.ParseForm(); err != nil {
				log.Printf("Error parsing form: %s", err)
				return
			}
			
			currentRequest := 0
			
			res := response(w,q, &f.properties)
			req := request(q, &f.properties)
                        f.Middleware.ServePlugin(res, req)	
			
			f.router.Next(Params, requestQuery.Handlers[currentRequest], res, req)
			currentRequest++
			break
		}
	}

	if !match {
	path := f.errorPath
	t, err := template.New("404.html").ParseFiles(path)
	if err != nil {
	 fmt.Println(err)
	}
	t.Execute(w, f.errorData)
	}
}

func (f *fiable) Get(path string, handler ...Handler) {
	f.router.Get(path, handler...)
}

func (f*fiable) Set404(path string, data interface{}) *fiable{
 f.errorPath = path
 f.errorData = data
 return f
}
func (f*fiable) Use(handler Handler){
 f.Middleware.AddPlugin(handler)
}
func (f *fiable) UseRouter(router *router) {
	f.router.UseRouter(router)
	
}