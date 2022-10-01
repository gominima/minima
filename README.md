<!-- markdownlint-disable no-hard-tabs no-inline-html -->

# Minima

<!-- Inline HTML inside center aligned paragraphs aren't subjected to markdown refactoring -->

<p align="center">
  <a href="https://gominima.studio">
  <img alt="Minima" src="https://raw.githubusercontent.com/gominima/minima/main/assets/logo.png" />
</a>
</p>

<p align="center" style="font-weight: 500">
Minima 🦄 is a reliable and lightweight framework for <a href="https://www.golang.org" target="_blank">Go</a> to carve the web 💻. Developed with core <a href="https://pkg.go.dev/net/http" target="_blank">net/http</a>🔌and other native packages, and with 0 dependencies.
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
<a href="https://gominima.studio"><img src="https://api.netlify.com/api/v1/badges/34ade51b-20a8-4fad-8cfd-0d574c9c40ed/deploy-status" /></a>
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

## 🔖 Guide

Official guide: https://guide.gominima.studio
Official docs: https://gominima.studio

## 🔮 Features

- **Reliable** - great modular API for building great server side applications</li>
- **Compatible with net/http** - use your plain old middlewares written in plain old `net/http`</li>
- **Lightweight** - clocked in ~1000 loc</li>
- **No Dependency** - just your plain old go standard libraries</li>
- **Great Documentation** - best in class precise [documentation](https://gominima.studio/)
- **Auto Docs** - docgen for generating all of your routing docs from jsdoc like comments to json or markdown files

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

	return router.Get("/user/:id", func(res *minima.Response, req *minima.Request) {
		// getting the id parameter from route
		id := req.Param("id")

		// instead of adding a param in route, you just need to fetch it

		username := req.Query("name")

		// get user from database
		userdata, err := db.FindUser(id, username)

		if err != nil {
			// check for errors
			res.NotFound().Send("No user found with particular id")
		}
		// send user data
		res.OK().Json(userdata)
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

	app.Get("/getuser/:id", func(res *minima.Response, req *minima.Request) {
		userid := req.Param("id")
		// check if user id is available
		if userid == "" {
			res.Error(404, "No user found")
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
		// instead of adding a param in route, you just need to fetch it

		userid := req.Query("id")

		if userid == "" {
			res.Error(404, "No user found")
			panic("No user id found in request")
		}
		fmt.Print(userid)
		// the above will print 20048 from router /getuser?id=20048
	})
}
```

## 📒 Minima Interface

Minima is based on a custom implementation of radix tree which makes it extremely performant. The router itself is fully compatible with [`net/http`](https://pkg.go.dev/net/http)

### 🔖 Minima's Interface

```go
type Minima interface {
	// Minima interface is built over net/http so every middleware is compatible with it

	// initializes net/http server with address
	Listen(address string) error
    

	//static file serve methods
	File(pth string, dir string) //serves a single static file to route
	Static(pth string, dir string) //serves whole directory with specified pth ex /static/main.html
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

       //Takes http.Handler and appends it to middleware chain
	UseRaw(handler func(http.Handler) http.Handler) *minima

        // an custom handler when route is not matched
	NotFound(handler Handler)*minima

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

	// Header methods
	GetHeader(key string) string // gets a header from response body
        
	SetHeader(key string, value string)  *Response // sets a new header to response body

	DelHeader(key string)  *Response // Deletes a header from response body
        
	CloneHeaders() http.Header // clones all headers of the response body

	Setlenght(len string) *Response // sets content length of the response body

	SetBaseHeaders() *Response // sets a good stack of base headers for response body

	FlushHeaders() // flushes headers

	// utility functions for easier usage
	Send(content string) *Response      //send content
	WriteBytes(bytes []byte) error      //writes bytes to the page
	JSON(content interface{}) *Response //sends data in json format
	XML(content interface{}, indent string) //sends data in xml format
	Stream(contentType string read io.Reader) // streams content to the route
	NoContent(code int)
	Error(status int, str string) *Response

	// this functions returns http.ResponseWriter instance which means you could use any of your alrady written middlewares in minima!
	Raw() http.ResponseWriter

	// renders an html file with data to the page
	Render(path string, data interface{}) *Response
        
	// custom method when there's an error
	Error(content interface{}) *Response
        
	// closes io.Writer connection from the route
	CloseConn() *Response

	// redirects to given url
	Redirect(url string) *Response

	// sets header status
	Status(code int) *Response

	//cookie methods
	SetCookie(cookie *http.Cookie) *Response // sets a new cookie to the response

	
	SetCookie(cookie *http.Cookie) *Response // clears a cookie to the response

}

type Req interface {
	// minima request interface is built on http.Request

	// returns param from route 
	Param(name string) string
        
	// sets a new param to the request instance
	SetParam(key string, value) string
	
	//returns query param from route url
	Query(key string) string
        
	//returns query params to a string
	QueryString() string
        
	//returns raw url.Values
	QueryParams() url.Values
        
	//returns the ip of the request origin
	IP() string

	//returns whether the request is tls or not
	IsTLS() bool

	//returns whether the request is a socket or not
	IsSocket() bool

	//returns scheme type of request body
	SchemeType() string

	//Gets form value from request body 
	FormValue(key string) string

	//Gets all the form param values
	FormParams() (url.Values, error)

	//Gets file from request 
	FormFile(key string) (*multipart.FileHeader, error)

	//Gets request Multipart form
	MultipartForm() (*multipart.Form, error)
	// returns path url from the route
	Path() string

	// returns raw request body
	Body() map[string][]string

	// finds given key value from body and returns it
	BodyValue(key string) []string

	// returns instance of minima.IncomingHeader for incoming header requests
	Header() *IncomingHeader

	// returns route method ex.get,post
	Method() string

	// Header methods
	SetHeader(key string, value string) *Response //sets a new header to request body
        
	//Get a header from request body
	GetHeader(key string) string

	//Cookie methods
        Cookies() []*http.Cookie // gets all cookies from request body

	GetCookie(key string) *http.Cookie // gets a specific cookie from request body

}
```

## 🔌 Middlewares

Minima's middlewares are written in its own custom `res` and `req` interfaces in accordance with the standard libraries maintained by Go. You can use `res.Raw()` to get the `http.ResponseWriter` instance and `req.Raw()` to getthe `http.Request` instance, meaning all community written middlewares are compatible with Minima.

Minima also takes `http.Handler` while using `.UseRaw` function and runs it in a chain.

Here is an example of standard `net/http` middleware being used with minima:

```go
func MyMiddleWare(res *minima.Response, req *minima.Request) {
	w := res.Raw() // raw http.ResponseWriter instance
	r := req.Raw() // raw http.Request instance

	// your normal net/http middleware
	w.Write([]byte(r.URL.Path))
}
app.UseRaw(HttpHandler())
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
- [@megatank58](https://github.com/megatank58)

#### Core Team

- [@apoorvcodes](https://github.com/apoorvcodes)
- [@megatank58](https://github.com/megatank58)
- [@Shubhaankar-Sharma](https://github.com/Shubhaankar-Sharma)
- [@savioxavier](https://github.com/savioxavier)

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
