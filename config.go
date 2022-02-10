package minima

/**
@info The config structure
@property {[]Handler} [Middleware] The plugins to be used
@property {Logger} [bool] Whether logger is enabled or not
@property {Router} [router] The router instance to be used
*/
type Config struct {
	Middleware []Handler
	Router     []*Router
}

/**
@info Make a new default config instance
@returns {Config}
*/
func NewConfig() *Config {
	return &Config{}
}
