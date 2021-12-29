package main

import (
	"fmt"

	"github.com/gominima/minima"
)

func main() {
	app := minima.New()
	router := minima.NewRouter()
	app.Get("/o/:name", func(response *minima.Response, request *minima.Request) {
		p := request.GetParam("name")
		response.Send(300, p)
	})
	router.Get("/hello/?", func(response *minima.Response, request *minima.Request) {
		type hello struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		q := request.GetQuery("name")
		fmt.Println(q)
		response.Json(&hello{Name: "totu", Age: 15})
	})

	app.UseConfig(&minima.Config{
		Logger:     false,
		Middleware: []minima.Handler{},
		Router:     router,
	})
	app.Listen(":3000")
}
