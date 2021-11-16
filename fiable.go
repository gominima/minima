package fiable

import (
	"context"
	"fmt"
	"strings"

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
	var router *router = Router()
	var fiable *fiable = &fiable{router: router}
	return fiable
}

func (f *fiable) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var allow []string
	for _, route := range f.router.routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
			route.handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.NotFound(w, r)
	fmt.Fprint(w, "Uh")
}

func (f *fiable) Listen(addr string) error {
	server := &http.Server{Addr: addr}
	if f.started {
		fmt.Errorf("Server is already running", f)
	}
	f.server = server
	f.started = true

	return f.server.ListenAndServe()
}

func (f *fiable) Get(path string, handler http.HandlerFunc) {
	f.router.Get(path, handler)
}
