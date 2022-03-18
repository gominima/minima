package minima

/**
* Minima is a free and open source software under Mit license

Copyright (c) 2021 gominima

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

* Authors @apoorvcodes @megatank58
* Maintainers @Panquesito7 @savioxavier @Shubhaankar-Sharma @apoorvcodes @megatank58
* Thank you for showing interest in minima and for this beautiful community
*/

// This chain implementation is borrowed/ inspired from https://github.com/go-chi/chi, so kindly check that project aswell
import "net/http"

type Middlewares []func(http.Handler) http.Handler

/**
 n* @info Create a middleware chain
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
 * @param {[]func(http.Handler)http.Handler} [middleware] The array of middleware stack
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
