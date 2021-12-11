package fiable

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type fiable struct {
	server  *http.Server
	started bool
	Timeout time.Duration
	router  *router
        errorRoute bool
	
	properties map[string]interface{}
}


func New() *fiable {
	var router *router = NewRouter()
	var fiable *fiable = &fiable{router: router}
	return fiable
}


func (f *fiable) Listen(addr string) error {
	server := &http.Server{Addr: addr , Handler: f}
	if f.started {
		fmt.Errorf("Server is already running", f)
	}
	f.server = server
	f.started = true
        if !f.errorRoute{
	
	}
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
                        	
			
			f.router.Next(Params, requestQuery.Handlers[currentRequest], res, req)
			currentRequest++
			break
		}
	}

	if !match {
		if f.router.NotFound != nil {
			f.router.NotFound(w, q)
		} else {
			http.NotFound(w, q)
		}
	}
}

func (f *fiable) Get(path string, handler ...Handler) {
	f.router.Get(path, handler...)
}


func (f *fiable) UseRouter(router *router) {
	f.router.UseRouter(router)
}