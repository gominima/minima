package minima

type Config struct {
	Middleware []Handler
	Router     *Router
	
}

func NewConfig() *Config {
	return &Config{}
}
