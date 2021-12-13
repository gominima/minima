package main

import (

	"github.com/gofiable/fiable"
)


func main(){
 app := fiable.New()
 router:=fiable.NewRouter()
 app.Get("/:name", func(response *fiable.Response, request *fiable.Request) {
    p := request.GetParam("name")
    response.Send(300, p)
 })
 router.Get("/hello", func(response *fiable.Response, request *fiable.Request) {
     response.Send(300, "Hello World")
 })


 app.UseConfig(&fiable.Config{
     Logger: false,
     Middleware: []fiable.Handler{},
     Router: router,
 })
 app.Listen(":3000")
}