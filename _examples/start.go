package main

import (
	"github.com/gominima/minima"
	"net/http"
)

func main() {
	app := minima.New()
	app.Get("/",func(res *minima.Response, req *minima.Request) {
		res.Send("Hello")
	})
	app.Test(test())
	app.Listen(":3000")
}


func test() func(next http.Handler) http.Handler {
    
	return func(next http.Handler) http.Handler {
	    fn := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
		
		next.ServeHTTP(w, r)
	    }
	    return http.HandlerFunc(fn)
	}
    }