package rtr

import (
	"github.com/gominima/minima"
	"github.com/gominima/minima/_examples/rtr/auth"
)

func SimpleTest() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
		res.Send("Hi test")
	}
}

func Router() *minima.Router {
	rt := minima.NewRouter()
	rt.Get("/test/one", auth.SimpleTests())
	rt.Get("/test/two", auth.SimpleTests1())
	rt.Get("/test/three", auth.SimpleTests2())
	rt.Get("/test/four", auth.SimpleTests3())
	rt.Get("/test/last", auth.SimpleTests4())
	return rt
}