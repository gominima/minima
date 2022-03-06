package rtr

import (
	"github.com/gominima/minima"
	"github.com/gominima/minima/_examples/rtr/test_routes"
)

func SimpleTest() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
		res.Send("Hi test")
	}
}

func Router() *minima.Router {
	rt := minima.NewRouter()
	rt.Get("/test/one", test.SimpleTests())
	rt.Get("/test/two", test.SimpleTests1())
	rt.Get("/test/three", test.SimpleTests2())
	rt.Get("/test/four", test.SimpleTests3())
	rt.Get("/test/last", test.SimpleTests4())
	return rt
}