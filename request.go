package minima

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
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

/**
@info The request structure
@property {*http.Request} [ref] The net/http request instance
@property {multipart.Reader} [fileReader] file reader instance
@property {map[string][]string} [body] Value of the request body
@property {string} [method] Request method
@property {[]*Params} [Params] Request path parameters
@property {query} [url.Values] Request path query params
@property {IncomingHeader} [header] Incoming headers of the request
@property {json.Decoder} [json] Json decoder instance
*/
type Request struct {
	ref        *http.Request
	fileReader *multipart.Reader
	body       map[string][]string
	method     string
	Params     map[string]string
	query      url.Values
	header     *IncomingHeader
	json       *json.Decoder
}

/**
@info Make a new default request instance
@param {http.Request} [http.Request] The net/http request instance
@returns {Request}
*/
func request(httpRequest *http.Request) *Request {
	req := &Request{
		ref:        httpRequest,
		header:     &IncomingHeader{},
		fileReader: nil,
		method:     httpRequest.Proto,
		query:      httpRequest.URL.Query(),
	}

	for i, v := range httpRequest.Header {
		req.header.Set(strings.ToLower(i), strings.Join(v, ","))
	}

	if req.header.Get("content-type") == "application/json" {
		req.json = json.NewDecoder(httpRequest.Body)
	} else {
		httpRequest.ParseForm()
	}

	if len(httpRequest.PostForm) > 0 {
		req.body = make(map[string][]string)
		for key, value := range httpRequest.PostForm {
			req.body[key] = value
		}
	}

	return req

}

/**
@info Gets param from route path
@param {string} [key] Key of the route param
@returns {string}
*/
func (r *Request) GetParam(key string) string {
	return r.Params[key]
}

/**
@info Sets param for route path
@param {string} [key] Key of the route param
@param {string} [value] Value of the route param
@returns {Respone}
*/
func (r *Request) SetParam(key string, value string) *Request {
	r.Params[key] = value
	return r
}

/**
@info Gets request path url
@returns {string}
*/
func (r *Request) GetPathURL() string {
	return r.ref.URL.Path
}

/**
@info Gets raw request body
@returns {map[string][]string}
*/
func (r *Request) Body() map[string][]string {
	return r.body
}

/**
@info Gets specified request body
@param {string} [key] Key of the request body
@returns {[]string}
*/
func (r *Request) GetBodyValue(key string) []string {
	return r.body[key]
}

/**
@info Gets raw json decoder instance
@returns {json.Decoder}
*/
func (r *Request) Json() *json.Decoder {
	return r.json
}

/**
@info Gets method of request
@returns {string}
*/
func (r *Request) Method() string {
	return r.method
}

/**
@info Gets raw net/http request instance
@returns {http.Request}
*/
func (r *Request) Raw() *http.Request {
	return r.ref
}

/**
@info Gets request path query
@param {string} [key] key of the request query
@returns {string}
*/
func (r *Request) GetQueryParam(key string) string {
	return r.query[key][0]
}

/**
@info Get all the cookies from the request
@returns {[]*http.Cookie}
*/
func (r *Request) Cookies() []*http.Cookie {
	return r.ref.Cookies()
}

/**
@info Get a paticular cookie by its key
@param {string} [key] key of the cookie
@returns {*http.Cookie}
*/
func (r *Request) GetCookie(key string) *http.Cookie {
	var result *http.Cookie
	for _, cookie := range r.Cookies() {
		if cookie.Name == key {
			result = cookie
		}
	}

	return result
}

/**
@info Set a paticular Header
@param {string} [key] key of the Header
@param {string} [value] value of the Header
@returns {*Request}
*/
func (r *Request) SetHeader(key string, value string) *Request {
	r.header.Set(key, value)
	return r
}

/**
@info Get a paticular Header by its key
@param {string} [key] key of the Header
@returns {string}
*/
func (r *Request) GetHeader(key string) string {
	return r.header.Get(key)
}
