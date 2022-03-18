package main

/**
* Minima is a free and open source software under Mit license

Copyright (c) 2021 gominima

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

* Authors @apoorvcodes @megatank58
* Maintainers @Panquesito7 @savioxavier @Shubhaankar-Sharma @apoorvcodes @megatank58
* Thank you for showing interest in minima and for this beautiful community
*/

import (
	"fmt"
	"github.com/gominima/minima"
	"github.com/gominima/minima/_examples/group"
	"github.com/gominima/minima/_examples/rtr"
	"net/http"
)

func main() {
	app := minima.New()
	app.UseRaw(SimpleTest())
	app.UseRouter(rtr.Router())
	app.UseGroup(group.RouteGroup())
	app.File("/main.html", "./static/main.html")
	app.Static("/static", "./static")
	app.Get("/", func(res *minima.Response, req *minima.Request) {
		res.Send("Hello")
		res.Send(req.Query("name"))
	})
	
	app.Listen(":3000")
}

func SimpleTest() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			fmt.Print(req.Method + "called on" + req.URL.Path)
			next.ServeHTTP(w, req)
		})
	}
}
