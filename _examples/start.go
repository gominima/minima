package main

import (
	"net/http"
	"github.com/gominima/minima"
)

func main() {
	app := minima.New()
	app.UseRaw(SimpleTest())
	app.UseRaw(SimpleTest2nd())
	app.Get("/", func(res *minima.Response, req *minima.Request) {
		res.Send("Hello")
	})
	
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


func SimpleTest2nd() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte("Works 2nd handler"))
			next.ServeHTTP(w, req)
		})
	}
}


