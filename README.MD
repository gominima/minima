<p align="center">
  <a href="https://gominima.org">
  <img alt="Fiber" src="./assets/logo.png">
</a>
</p>

<h4 align="center">
Minima 🦄 is a reliable and lightweight framework for <a href="https://www.golang.org" target="_blank">Go</a> to carve the web 💻. Developed with core <a href="https://pkg.go.dev/net/http" target="_blank">net/http</a>🔌and other native packages, and with 0 dependencies
</h4>

<br>
<p align="center">
<a href="https://goreportcard.com/badge/github.com/gominima/minima"> <img src="https://goreportcard.com/badge/github.com/gominima/minima"> </a>
<a href="https://img.shields.io/github/go-mod/go-version/gominima/minima"> <img src="https://img.shields.io/github/go-mod/go-version/gominima/minima"></a>
<a href="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"> <img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"></a>
<a href="https://discord.gg/gRyCr5APmg"> <img src="https://img.shields.io/discord/916969864512548904"></a>
<img src="https://img.shields.io/tokei/lines/github/gominima/minima">
<img src="https://img.shields.io/github/languages/code-size/gominima/minima">
<a href="https://gominima.org"> 
<img src="https://img.shields.io/badge/Minima-Docs-blue"></a>
</p>


<h1>⚙️ Setup</h1>
<p>Kindy check if you have <a href="https://go.dev/">Go</a> version 1.15 or higher</p>



```go
mkdir <project-name> && cd  <project-name>

go mod init github.com/<user-name>/<repo-name>

go get github.com/gominima/minima

go run main.go
```
<br>

<h1>🦄 Quickstart </h1>

```go
package main

import "github.com/gominima/minima"

func main() {
	app := minima.New()

	app.Get("/", func(res *minima.Response, req *minima.Request) {
		res.Send(200, "Hello World")
	})

	app.Listen(":3000")
}

```
<br>
<h1>🪶 Features</h1>
<ul>
<li><b>Reliable -</b> Great modular api for building great server side applications</li>
<li><b>Compatible with net/http</b>- use your plain old middlewares written in plain old net/http</li>
<li><b>Lightweight -</b> clocked in ~ 1000 loc</li>
<li><b>No Dependency -</b> Just your plain old go standard libraries</li>
<li><b>Great Documentation -</b> Best in class precise <a href="https://gominima.org/">documentation</a></li>
<li><b>Auto Docs -</b> Docgen for generating all you routing docs from router to json or markdown files</li>
</ul>

<br>

<h1>❓Why minima </h1>

<p>
Minima's name is inspired by the word minimal and is the motivation for building this framework. As a go dev, I was struggling to learn it in my early days due to the steeper learning curve while using net/http.
</p>

<p>Also during checking out some other alternative frameworks, I found out that something like fiber wasn't compatible to net/http modules like gqlgen and other middlewares.</p>

<p>Minima solves this problem as it has a very narrow learning curve as well as a robust structure that supports all net/http modules and other middleware without compromising on performance.</p>
<br>
<h1>🍵 Examples</h1>

<h4>Here are some basic examples related to routing and params</h4>

<h2>📑 Routing & Router</h2>

```go
func UserGetRouter() *minima.Router {
	//router instance which would be used by the main router
	router := minima.NewRouter()
	return router.Get("/user/:id/?", func(response *minima.Response, request *minima.Request) {
		//getting id param from route
		id := request.GetParam("id")

		//as query params are not part of the request path u need to add a ? to initialize them
		username := request.GetQuery("name")

		//get user from db
		userdata, err := db.FindUser(id, username)

		if err != nil {
			panic(err)
			//check for errors
			response.Status(404).Send("No user found with particular id")
		}
		//sending user
		response.Json(userdata).Status(200)
	})
}

func main() {
	//main minima instance
	app := minima.New()
	//UseRouter method takes minima.router as param
	//It appends all the routes used in that specific router to the main instance
	app.UseRouter(UserGetRouter())

	//running the app
	app.Listen(":3000")
}
```

<h2>📑 Params</h2>

```go
func main() {
	app := minima.New()

	app.Get("/getuser/:id", func(response *minima.Response, request *minima.Request) {
		userid := request.GetParam("id")
		// check if user id is available
		if userid == "" {
			response.Error(404, "No user found")
			panic("No user id found in request")
		}
		fmt.Print(userid)
		//Will print 20048 from router /getuser/200048
	})
}
```

<h2>📑 Query Params</h2>

```go
func main() {
	app := minima.New()

	app.Get("/getuser/?", func(response *minima.Response, request *minima.Request) {
		//query params work a bit different instead of adding a param in route ur u just need to add a ? and fetch the param
		userid := request.GetQuery("id")
		
		if userid == "" {
			response.Error(404, "No user found")
			panic("No user id found in request")
		}
		fmt.Print(userid)
		//Will print 20048 from router /getuser/200048
	})
}
```
<h2>📑 Minima interface</h2>

<h4>Minima is based on a looping system which loops through routes and matches the regex of requested route. The router itsekf is fully compatible with <a href="https://pkg.go.dev/net/http" target="_blank">net/http</a></h4>

<h4>Here is minima's interface</h4>

```go
type Minima interface {
	//Minima interface is built over net/http so every middleware is compatible with it

	//initializes net/http server with address
	Listen(address string) error

	//handler interface
	ServeHTTP(w http.ResponseWriter, q *http.Request)

	//Router methods
	Get(path string, handler ...Handler) *minima
	Patch(path string, handler ...Handler) *minima
	Post(path string, handler ...Handler) *minima
	Put(path string, handler ...Handler) *minima
	Options(path string, handler ...Handler) *minima
	Head(path string, handler ...Handler) *minima
	Delete(path string, handler ...Handler) *minima

	//Takes middlewares as a param and adds them to routes
	//middlewares initializes before route handler is mounted
	Use(handler Handler) *minima

	//Takes minima.Router as param and adds the routes from router to main instance
	UseRouter(router *Router) *minima

	//Works as a config for minima, you can add multiple middlewares and routers at once
	UseConfig(config *Config) *minima

	//Shutdowns the net/http server
	Shutdown(ctx context.Context) error

	//Prop methods
	SetProp(key string, value interface{}) *minima
	GetProp(key string) interface{}
}
```