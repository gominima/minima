package minima

import "net/http"

/**
@info The config structure
@property {[]Handler} [Middleware] The minima middlewares to be used
@property {[]http.HandlerFunc} [HttpHandler] The net/http middlewares to be used
@property {Router} [router] The router instance to be used
*/
type Config struct {
	Middleware  []Handler
	HttpHandler []http.HandlerFunc
	Router      []*Router
}

/**
@info Make a new default config instance
@returns {Config}
*/
func NewConfig() *Config {
	return &Config{}
}
