package fiable

type Config struct{
 Middleware []Handler
 Logger   bool
 Router   *router
}

func NewConfig() *Config{
 return &Config{Logger: false}
}