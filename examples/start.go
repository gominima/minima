package main

import (

	"github.com/gofiable/fiable"
)

func main(){
    
    app := fiable.New()
    router := fiable.NewRouter()
    router.Get("/hello/:name/:id", func(response *fiable.Response, request *fiable.Request) {
        param := request.GetParam("name")
        p := request.GetParam("id")
        r := "The user name is " + param + "and his id is " + p
        response.Send(300,r)
  
    })
    app.Get("/bye/:id/", func(response *fiable.Response, request *fiable.Request) {
        param := request.GetParam("id")
        response.Send(300, param)  
    })
   
    app.UseRouter(router)

    app.Listen(":3000")
}
