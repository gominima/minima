package minima

import "net/http"

/**
@info The config structure
@property {[]Handler} [Middleware] The plugins to be used
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
