package minima

import "net/http"


/**
 * @info Converts minima handler into middleware chain handler
 * @param {Handler} [h] The handler to convert
 * @param {map[string]string} [params] The handler params
 * @return {func(http.Handler) http.Handler}
 */
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

/**
 * @info Converts minima handler into net/http handler func
 * @param {Handler} [h] The handler to convert
 * @param {map[string]string} [params] The handler params
 * @return {http.Handler}
 */
func buildHandler(h Handler, params map[string]string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		resp := response(w, req)
		reqs := request(req)
		reqs.Params = params
		h(resp, reqs)

	})
}
