package main

import (
	"fmt"
	"net/http"

	"github.com/gominima/minima"
)

func main() {
	app := minima.New()
	app.Use(SimpleTest())
	app.Get("/", func(res *minima.Response, req *minima.Request) {
		res.Send("Hello")
	})
	
	app.Listen(":3000")
}

func SimpleTest() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			fmt.Print(req.Method + "called on" + req.URL.Path)
			next.ServeHTTP(w, req)
		})
	}
}


