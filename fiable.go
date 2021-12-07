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

	properties map[string]interface{}
}

type ctxKey struct{}

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

	return f.server.ListenAndServe()
}

func (f* fiable) ServeHTTP(w http.ResponseWriter, q *http.Request){
	match := false

	for _, requestedMethod := range f.router.routes[q.Method] {
		if isMatchRoute, namedParams := requestedMethod.matchingPath(q.URL.Path); isMatchRoute {
			match = isMatchRoute
			if err := q.ParseForm(); err != nil {
				log.Printf("Error parsing form: %s", err)
				return
			}
			currentRequest := 0
			
			res := response(w,q, &f.properties)
			req := request(q, &f.properties)

			f.router.NextWithContext(namedParams, requestedMethod.Handlers[currentRequest], res, req)
			currentRequest++
			break
		}
	}

	if !match {

	}
}

func (f *fiable) Get(path string, handler ...Handler) {
	f.router.Get(path, handler...)
}
