package minima

import "net/http"

func Build(h Handler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			resp := response(w, req)
			reqs := request(req)
			h(resp, reqs)
			next.ServeHTTP(w, req)
		})
	}
}
