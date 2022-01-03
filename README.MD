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