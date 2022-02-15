<!-- markdownlint-disable no-hard-tabs no-inline-html -->

# Minima

<!-- Inline HTML inside center aligned paragraphs aren't subjected to markdown refactoring -->

<p align="center">
  <a href="https://gominima.studio">
  <img alt="Minima" src="https://raw.githubusercontent.com/gominima/minima/main/assets/logo.png" />
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
<a href="https://gominima.studio">
<img src="https://img.shields.io/badge/Minima-Docs-blue" /></a>
</p>

## ⚙️ Setup

Please make sure you have [Go](https://go.dev) version 1.15 or higher

```go
mkdir <project-name> && cd  <project-name>

go mod init github.com/<user-name>/<repo-name>

go get github.com/gominima/minima

go run main.go
```

## 🦄 Quickstart

```go
package main

import "github.com/gominima/minima"

func main() {
	app := minima.New()

	app.Get("/", func(res *minima.Response, req *minima.Request) {
		res.OK().Send("Hello World")
	})

	app.Listen(":3000")
}

```

## 🔮 Features

- **Reliable** - great modular API for building great server side applications</li>
- **Compatible with net/http** - use your plain old middlewares written in plain old `net/http`</li>
- **Lightweight** - clocked in ~1000 loc</li>
- **No Dependency** - just your plain old go standard libraries</li>
- **Great Documentation** - best in class precise [documentation](https://gominima.studio/)
- **Auto Docs** - docgen for generating all of your routing docs from router to json or markdown files

## ❓ Why Minima

Minima's name is inspired by the word minimal and is the motivation for building this framework. As a Golang developer, I was struggling to learn it in my early days due to the steeper learning curve while using `net/http`.

Also while checking out some other alternate frameworks, I found out that something like [fiber](https://github.com/gofiber/fiber) wasn't compatible with `net/http` modules like `gqlgen` and other middlewares.

Minima solves this problem as it has a very narrow learning curve as well as a robust structure that supports all `net/http` modules and other middlewares without compromising performance.

## 🍵 Examples

Here are some basic examples related to routing and params:

### 📑 Routing & Router

```go
func UserGetRouter() *minima.Router {
	// router instance which would be used by the main router
	router := minima.NewRouter()
	return router.Get("/user/:id", func(response *minima.Response, request *minima.Request) {
		// getting the id parameter from route
		id := request.GetParam("id")

		// as query params are not part of the request path, they wont be added to the route
		username := request.GetQuery("name")

		// get user from database
		userdata, err := db.FindUser(id, username)

		if err != nil {
			panic(err)
			// check for errors
			response.NotFound().Send("No user found with particular id")
		}
		// send user data
		response.Json(userdata).OK()
	})
}

func main() {
	// main minima instance
	app := minima.New()
	// UseRouter method takes minima.router as a param
	// it appends all the routes used in that specific router to the main instance
	app.UseRouter(UserGetRouter())

	// running the app at port 3000
	app.Listen(":3000")
}
```

### 📑 Params

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

### 📑 Query Params

```go
func main() {
	app := minima.New()

	app.Get("/getuser", func(response *minima.Response, request *minima.Request) {
		// query params work a bit differently
		// instead of adding a param in route, you just need to fetch the param
		userid := request.GetQuery("id")

		if userid == "" {
			response.Error(404, "No user found")
			panic("No user id found in request")
		}
		fmt.Print(userid)
		// the above will print 20048 from router /getuser?id=20048
	})
}
```

## 📒 Minima Interface

Minima is based on a looping system which loops through routes and matches the regex of requested route. The router itself is fully compatible with [`net/http`](https://pkg.go.dev/net/http)

### 🔖 Minima's Interface

```go
type Minima interface {
	// Minima interface is built over net/http so every middleware is compatible with it

	// initializes net/http server with address
	Listen(address string) error

	// main handler interface
	ServeHTTP(w http.ResponseWriter, q *http.Request)

	// main router methods
	Get(path string, handler ...Handler) *minima
	Patch(path string, handler ...Handler) *minima
	Post(path string, handler ...Handler) *minima
	Put(path string, handler ...Handler) *minima
	Options(path string, handler ...Handler) *minima
	Head(path string, handler ...Handler) *minima
	Delete(path string, handler ...Handler) *minima

	// takes middlewares as a param and adds them to routes
	// middlewares initializes before route handler is mounted
	Use(handler Handler) *minima

	// mounts routes to specific base path
	Mount(basePath string, router *Router) *minima

	// takes minima.Router as param and adds the routes from router to main instance
	UseRouter(router *Router) *minima

	// works as a config for minima, you can add multiple middlewares and routers at once
	UseConfig(config *Config) *minima

	// shutdowns the net/http server
	Shutdown(ctx context.Context) error

	// prop methods
	SetProp(key string, value interface{}) *minima
	GetProp(key string) interface{}
}
```

### 🔖 Response and Request Interfaces

Both response and request interfaces of minima are written in `net/http` so you can use any of your old route middlewares in minima out of the box without any hassle.

```go

type Res interface {
	// response interface is built over http.ResponseWriter for easy and better utility

	// returns minima.OutgoingHeader interface
	Header() *OutgoingHeader

	// utility functions for easier usage
	Send(content string) *Response      //send content
	WriteBytes(bytes []byte) error      //writes bytes to the page
	Json(content interface{}) *Response //sends data in json format
	Error(status int, str string) *Response

	// this functions returns http.ResponseWriter instance which means you could use any of your alrady written middlewares in minima!
	Raw() http.ResponseWriter

	// renders an html file with data to the page
	Render(path string, data interface{}) *Response

	// redirects to given url
	Redirect(url string) *Response

	// sets header status
	Status(code int) *Response
}

type Req interface {
	// minima request interface is built on http.Request

	// returns param from route url
	GetParam(name string) string

	// returns path url from the route
	GetPathURL() string

	// returns raw request body
	Body() map[string][]string

	// finds given key value from body and returns it
	GetBodyValue(key string) []string

	// returns instance of minima.IncomingHeader for incoming header requests
	Header() *IncomingHeader

	// returns route method ex.get,post
	Method() string

	// gets query params from route and returns it
	GetQuery(key string) string
}
```

## 🔌 Middlewares

Minima's middlewares are written in its own custom `res` and `req` interfaces in accordance with the standard libraries maintained by Go. You can use `res.Raw()` to get the `http.ResponseWriter` instance and `req.Raw()` to getthe `http.Request` instance, meaning all community written middlewares are compatible with Minima.

Here is an example of standard `net/http` middleware being used with minima:

```go
func MyMiddleWare(res *minima.Response, req *minima.Request) {
	w := res.Raw() // raw http.ResponseWriter instance
	r := req.Raw() // raw http.Request instance

	// your normal net/http middleware
	w.Write([]byte(r.URL.Path))
}
```

## 💫 Contributing

**If you wanna help grow this project or say a thank you!**

1. Give minima a [GitHub star](https://github.com/gominima/minima/stargazers)
2. Fork Minima and Contribute
3. Write a review or blog on Minima
4. Join our [Discord](https://discord.gg/gRyCr5APmg) community

### Contributors

#### Lead Maintainers

- [@apoorvcodes](https://github.com/apoorvcodes)
- [@savioxavier](https://github.com/savioxavier)
- [@megatank58](https://github.com/megatank58)

#### Community Contributors

Thanks to all the contributors, without whom this project would not have been possible:

<a href="https://github.com/gominima/minima/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=gominima/minima" />
</a>

<br>

Be a part of this contribution list by contributing today!

Please note that this project is released with a [Contributor Code of Conduct](CODE_OF_CONDUCT.md). By participating in this project you agree to abide by its terms.

## 🧾 License

**Copyright (c) 2021-present [Apoorv](https://github.com/apoorvcodes) and [Contributors](https://github.com/gominima/minima/graphs/contributors). Minima is a Free and Open Source Software licensed under [MIT License](https://github.com/gominima/minima/blob/main/LICENSE)**

<br />
<br />

<p align="center">
<strong>Happy coding ahead with Minima!</strong>
</p>
