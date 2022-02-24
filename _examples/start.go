package main

import (
	"fmt"
	"net/http"
	"github.com/gominima/minima"
)

func main() {
	app := minima.New()
	app.Get("/", func(res *minima.Response, req *minima.Request) {
		res.Send("Hello")
	})
	app.UseRaw(SimpleTest())
	app.Listen(":3000")
}

func SimpleTest() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte("Works"))
			next.ServeHTTP(w, req)
		})
	}
}


