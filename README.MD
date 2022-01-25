<!-- markdownlint-disable no-hard-tabs -->

# Minima

<p align="center">
  <a href="https://gominima.org">
  <img alt="Minima" src="https://github.com/gominima/minima/blob/master/assets/logo.png?raw=true" />
</a>
</p>

<p align="center" style="font-weight: 500">
Minima 🦄 is a reliable and lightweight framework for <a href="https://www.golang.org" target="_blank">Go</a> to carve the web 💻. Developed with core <a href="https://pkg.go.dev/net/http" target="_blank">net/http</a>🔌and other native packages, and with 0 dependencies
</p>

<p align="center">
<a href="https://goreportcard.com/badge/github.com/gominima/minima"> <img src="https://goreportcard.com/badge/github.com/gominima/minima" /> </a>
<a href="https://img.shields.io/github/go-mod/go-version/gominima/minima"> <img src="https://img.shields.io/github/go-mod/go-version/gominima/minima" /></a>
<a href="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"> <img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat" /></a>
<a href="https://discord.gg/gRyCr5APmg"> <img src="https://img.shields.io/discord/916969864512548904" /></a>
<img src="https://img.shields.io/tokei/lines/github/gominima/minima" />
<img src="https://img.shields.io/github/languages/code-size/gominima/minima" />
<a href="https://gominima.org">
<img src="https://img.shields.io/badge/Minima-Docs-blue" /></a>
</p>

# ⚙️ Setup

Please make sure you have [Go](https://go.dev) version 1.15 or higher

```go
mkdir <project-name> && cd  <project-name>

go mod init github.com/<user-name>/<repo-name>

go get github.com/gominima/minima

go run main.go
```

# 🦄 Quickstart

```go
package main

import "github.com/gominima/minima"

func main() {
	app := minima.New()

	app.Get("/", func(res *minima.Response, req *minima.Request) {
		res.Status(200).Send("Hello World")
	})

	app.Listen(":3000")
}

```

# 🪶 Features

- **Reliable** - Great modular api for building great server side applications</li>
- **Compatible with net/http** - use your plain old middlewares written in plain old net/http</li>
- **Lightweight** - clocked in ~ 1000 loc</li>
- **No Dependency** - Just your plain old go standard libraries</li>
- **Great Documentation** - Best in class precise [documentation](https://gominima.org/)
- **Auto Docs** - Docgen for generating all you routing docs from router to json or markdown files

# ❓ Why Minima

Minima's name is inspired by the word minimal and is the motivation for building this framework. As a Golang developer, I was struggling to learn it in my early days due to the steeper learning curve while using net/http.

Also during checking out some other alternative frameworks, I found out that something like fiber wasn't compatible to net/http modules like gqlgen and other middlewares.

Minima solves this problem as it has a very narrow learning curve as well as a robust structure that supports all net/http modules and other middleware without compromising on performance.

# 🍵 Examples

**Here are some basic examples related to routing and params:**

## 📑 Routing & Router

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

## 📑 Params

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

## 📑 Query Params

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
		//Will print 20048 from router /getuser?id=20048
	})
}
```

## 📑 Minima Interface

**Minima is based on a looping system which loops through routes and matches the regex of requested route. The router itself is fully compatible with [net/http](https://pkg.go.dev/net/http)**

### Minima's interface

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
        
	//Mounts routes to specific base path
	Mount(basePath string, router *Router) *minima

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

## 📑 Response and Request interface

**Both response and request interfaces of minima are written in net/http so you can use any of your old route middlewares in minima out of the box without any hustle.**

```go

type Res interface {
	//response interface is built over http.ResponseWriter for easy and better utility

	//returns minima.OutgoingHeader interface
	Header() *OutgoingHeader

	//Utility functions for easier usage
	Send(content string) *Response      //send content
	WriteBytes(bytes []byte) error      //writes bytes to the page
	Json(content interface{}) *Response //sends data in json format
	Error(status int, str string) *Response

	//This functions return http.ResponseWriter instace that means you could use any of your alrady written middlewares in minima!!
	Raw() http.ResponseWriter

	//renders a html file with data to the page
	Render(path string, data interface{}) *Response

	//Redirects to given url
	Redirect(url string) *Response

	//Sets Header status
	Status(code int) *Response
}

type Req interface {
	//Minima request interface is built on http.Request

	//returns param from route url
	GetParam(name string) string

	//returns path url from the route
	GetPathURl() string

	//returns raw request body
	Body() map[string][]string

	//finds given key value from body and returns it
	GetBodyValue(key string) []string

	//returns instance of minima.IncomingHeader for incoming header requests
	Header() *IncomingHeader

	//returns route method ex.get,post
	Method() string

	//Gets query params from route and returns it
	GetQuery(key string) string
}
```

## 📑 Middlewares

**Minima's middleware are written in its own custom `res` and `req` interfaces in mind with the standard libs maintained by golang, you can use `res.Raw()` to get `http.ResponseWriter` instance and `req.Raw()` to get `http.Request` instance, meaning all community written middlewares are compatible with Minima.**

**Here is an example of standard net/http middleware being used with minima:**

```go
func MyMiddleWare(res *minima.Response, req *minima.Request) {
	w := res.Raw() //raw http.ResponseWriter instance
	r := req.Raw() //raw http.Request instance

	//your normal net/http middleware
	w.Write([]byte(r.URL.Path))
}
```

## 💫 Contributing

**If you wanna help grow this project or say a thank you!**

1. Give minima a [GitHub star](https://github.com/gominima/minima/stargazers)
2. Fork minima and Contribute
3. Write a review or blog on minima
4. Join our [Discord](https://discord.gg/gRyCr5APmg) community

# 🧾 License

**Copyright (c) 2021-present [Apoorv](https://github.com/apoorvcodes) and [Contributors](https://github.com/gominima/minima/graphs/contributors"). Minima is a free and Open source software licensed under [MIT License](https://github.com/gominima/minima/blob/master/LICENSE)**

<br />
<br />

<p align="center">
<strong>Happy coding ahead with Minima!</strong>
</p>
