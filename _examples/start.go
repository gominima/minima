package main

import (

	"github.com/gominima/minima"

)

func main() {
	app := minima.New()
	router := minima.NewRouter()
	app.Get("/test/:name/ok", func(response *minima.Response, request *minima.Request) {
		p := request.GetParam("name")
		v := request.GetQuery("age")
		response.NotFound().Send("Hello").Send(p).Send(v)
		response.CloseConn()

	})
	router.Get("/", func(res *minima.Response, req *minima.Request) {
		res.OK().Send("Hello World")
	})
	app.UseRouter(router)
	app.Listen(":3000")

}