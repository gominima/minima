package minima

// This chain implementation is borrowed/ inspired from https://github.com/go-chi/chi, so kindly check that project aswell
import "net/http"

type Middlewares []func(http.Handler) http.Handler

/**
 * @info Create a middleware chain
 */
func Chain(middlewares ...func(http.Handler) http.Handler) Middlewares {
	return Middlewares(middlewares)
}

/**
 * @info Creates a new chain handler instance
 * @param {http.Handler} [h] The handler stack to append
 * @returns {http.Handler}
 */
func (mws Middlewares) Handler(h http.Handler) http.Handler {
	return &ChainHandler{h, chain(mws, h), mws}
}

/**
 * @info Creates a new chain handler instance from handlerfunc
 * @param {http.HandlerFunc} [h] The handlerfunc stack to append
 * @returns {http.Handler}
 */
func (mws Middlewares) HandlerFunc(h http.HandlerFunc) http.Handler {
	return &ChainHandler{h, chain(mws, h), mws}
}

/**
 * @info Create a middleware chain
 * @property {http.Handler} [Endpoint] The endpoint of the chain stack
 * @property {http.Handler} [chain] The actual middleware stack chain
 * @property {Middlewares} [Middlewares] The middleware stack
 */
type ChainHandler struct {
	Endpoint    http.Handler
	chain       http.Handler
	Middlewares Middlewares
}

/**
 * @info Injects the middleware chain to minima instance
 * @param {http.ResponseWriter} [w] The net/http response instance
 * @param {http.Request} [r] The net/http request instance
 */
func (c *ChainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.chain.ServeHTTP(w, r)
}

/**
 * @info Builds the whole chain into one singular http.Handler
 * @param {[]func(http.Handler) http.Handler} [middleware] The array of middleware stack
 * @param {http.Handler} [endpoint] The endpoint of the chain stack
 * @returns {http.Handler}
 */
func chain(middlewares []func(http.Handler) http.Handler, endpoint http.Handler) http.Handler {

	if len(middlewares) == 0 {
		return endpoint
	}
	h := middlewares[len(middlewares)-1](endpoint)
	for i := len(middlewares) - 2; i >= 0; i-- {
		h = middlewares[i](h)
	}

	return h
}
