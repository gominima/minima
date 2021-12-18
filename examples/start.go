package main

import (
	"fmt"

	"github.com/gofiable/fiable"
)

func main() {
	app := fiable.New()
	router := fiable.NewRouter()
	app.Get("/o/:name", func(response *fiable.Response, request *fiable.Request) {
		p := request.GetParam("name")
		response.Send(300, p)
	})
	router.Get("/hello/?", func(response *fiable.Response, request *fiable.Request) {
		type hello struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		q := request.GetQuery("name")
		fmt.Println(q)
		response.Json(&hello{Name: "totu", Age: 15})
	})

	app.UseConfig(&fiable.Config{
		Logger:     false,
		Middleware: []fiable.Handler{},
		Router:     router,
	})
	app.Listen(":3000")
}
