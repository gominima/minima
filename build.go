package minima

import "net/http"

func build(h Handler, params map[string]string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			resp := response(w, req)
			reqs := request(req)
			h(resp, reqs)
			next.ServeHTTP(w, req)
		})
	}
}

func buildHandler(h Handler, params map[string]string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		resp := response(w, req)
		reqs := request(req)
		reqs.Params = params
		h(resp, reqs)

	})
}
