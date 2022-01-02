package main

import (
	"fmt"

	"github.com/gominima/minima"
)

func main() {
	app := minima.New()
	router := minima.NewRouter()
	app.Get("/test/:name", func(response *minima.Response, request *minima.Request) {
		p := request.GetParam("name")
		response.Send(300, p)
	})
	router.Get("/user/?", func(response *minima.Response, request *minima.Request) {
		type hello struct {
			Name string `json:"name"`
			
		}
		q := request.GetQuery("name")
		fmt.Println(q)
		response.Json(&hello{Name: q})
	})

	app.UseConfig(&minima.Config{
		Logger:     false,
		Middleware: []minima.Handler{},
		Router:     router,
	})
	app.Listen(":3000")
}
