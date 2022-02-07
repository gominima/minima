package main

import (
	"fmt"
	"github.com/gominima/minima"
	"net/http"
)

func main() {
	app := minima.New()
	router := minima.NewRouter()
	app.Get("/test/:name", func(response *minima.Response, request *minima.Request) {
		p := request.GetParam("name")

		response.Status(404).Send(p)
		response.CloseConn()

	})
	router.Get("/user/?", func(response *minima.Response, request *minima.Request) {
		type hello struct {
			Name string `json:"name"`
		}
		q := request.GetQuery("name")
		fmt.Println(q)
		response.Json(&hello{Name: q})
	})

	app.Mount("/v1", router)
	app.UseRaw(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("j"))
		fmt.Print(r.URL.Path)
	})
	app.Listen(":3000")

}

// func UserGetRouter() *minima.Router {
// 	//router instance which would be used by the main router
// 	router := minima.NewRouter()
// 	return router.Get("/user/:id/?", func(response *minima.response, request *minima.Request) {
// 		//getting id param from route
// 		id := request.GetParam("id")

// 		//as query params are not part of the request path u need to add a ? to initialize them
// 		username := request.GetQuery("name")

// 		//get user from db
// 		userdata, err := db.FindUser(id, username)

// 		if err != nil {
// 			panic(err)
// 			//check for errors
// 			response.Status(404).Send("No user found with particular id")
// 		}
// 		//sending user
// 		response.Json(userdata).Status(200)
// 	})
// }
