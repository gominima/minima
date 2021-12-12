package main

import (
	"fmt"

	"github.com/gofiable/fiable"
)

func fbc(response *fiable.Response, request *fiable.Request){
 fmt.Print(request.GetPathURl())
}

func fyz(response *fiable.Response, request *fiable.Request){
    fmt.Print(request.Params)
   }

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
     Middleware: []fiable.Handler{fbc, fyz},
     Router: router,
 })
 app.Use(func(response *fiable.Response, request *fiable.Request) {
     response.Send(300, "Hello world but middleware")
 })
 app.Listen("3000")
}