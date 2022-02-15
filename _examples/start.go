package main

import (
	"github.com/gominima/minima"
)

func main() {
	app := minima.New()
	router := minima.NewRouter()
	app.Get("/test/:name/user/:id", func(response *minima.Response, request *minima.Request) {
		p := request.GetParam("name")
		v := request.GetParam("id")
		response.NotFound().Send("Hello").Send(p).Send(v)
		response.CloseConn()

	})
	router.Get("/hi", func(res *minima.Response, req *minima.Request) {
		res.OK().Send("Hello World").Send(req.GetQuery("name"))
	})
	app.NotFound(func(res *minima.Response, req *minima.Request) {
		res.NotFound().Send("Not found handler")
	})
	app.Mount("/api", router)
	app.Listen(":3000")

}
