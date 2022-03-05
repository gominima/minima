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
	"net"
	"net/http"
	"net/url"
	"strings"
)

/**
 * @info The request structure
 * @property {*http.Request} [ref] The net/http request instance
 * @property {multipart.Reader} [fileReader] file reader instance
 * @property {map[string][]string} [body] Value of the request body
 * @property {string} [method] Request method
 * @property {[]*Params} [Params] Request path parameters
 * @property {query} [url.Values] Request path query params
 * @property {IncomingHeader} [header] Incoming headers of the request
 * @property {json.Decoder} [json] Json decoder instance
 */
type Request struct {
	ref        *http.Request
	fileReader *multipart.Reader
	body       map[string][]string
	method     string
	Params     map[string]string
	rawQuery   url.Values
	header     *IncomingHeader
	json       *json.Decoder
}

/**
 * @info Make a new default request instance
 * @param {http.Request} [http.Request] The net/http request instance
 * @returns {Request}
 */
func request(httpRequest *http.Request) *Request {
	req := &Request{
		ref:        httpRequest,
		header:     &IncomingHeader{},
		fileReader: nil,
		method:     httpRequest.Proto,
		rawQuery:   httpRequest.URL.Query(),
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
 * @info Gets param from route path
 * @param {string} [key] Key of the route param
 * @returns {string}
 */
func (r *Request) Param(key string) string {
	return r.Params[key]
}

/**
 * @info Sets param for route path
 * @param {string} [key] Key of the route param
 * @param {string} [value] Value of the route param
 * @returns {Respone}
 */
func (r *Request) SetParam(key string, value string) *Request {
	r.Params[key] = value
	return r
}

/**
 * @info Gets request path url
 * @returns {string}
 */
func (r *Request) Path() string {
	return r.ref.URL.Path
}

/**
 * @info Gets raw request body
 * @returns {map[string][]string}
 */
func (r *Request) Body() map[string][]string {
	return r.body
}

/**
 * @info Gets specified request body
 * @param {string} [key] Key of the request body
 * @returns {[]string}
 */
func (r *Request) BodyValue(key string) []string {
	return r.body[key]
}

/**
 * @info Gets raw json decoder instance
 * @returns {json.Decoder}
 */
func (r *Request) Json() *json.Decoder {
	return r.json
}

/**
 * @info Gets method of request
 * @returns {string}
 */
func (r *Request) Method() string {
	return r.method
}

/**
 * @info Gets raw net/http request instance
 * @returns {http.Request}
 */
func (r *Request) Raw() *http.Request {
	return r.ref
}

/**
 * @info Gets request path query
 * @param {string} [key] key of the request query
 * @returns {string}
 */
func (r *Request) Query(key string) string {
	return r.rawQuery.Get("key")
}

/**
 * @info Gets request path query in a string
 * @returns {string}
 */
func (r *Request) QueryString() string {
	return r.ref.URL.RawQuery
}

/**
 * @info Gets request path query in an array
 * @param {string} [key] key of the request query
 * @returns {string}
 */
func (r *Request) QueryParams() url.Values {
	return r.rawQuery
}

/**
 * @info Gets ip of the request origin
 * @returns {string}
 */
func (r *Request) IP() string {
	if ip := r.ref.Header.Get("X-Forwarded-For"); ip != "" {
		i := strings.IndexAny(ip, ",")
		if i > 0 {
			return strings.TrimSpace(ip[:i])
		}
		return ip
	}
	ra, _, _ := net.SplitHostPort(r.ref.RemoteAddr)
	return ra
}

/**
 * @info Whether the request is TLS or not
 * @returns {bool}
 */
func (r *Request) IsTLS() bool {
	return r.ref.TLS != nil
}

/**
 * @info Whether the request is a websocket or not
 * @returns {bool}
 */
func (r *Request) IsSocket() bool {
	upgrade := r.ref.Header.Get("Upgrade")
	return strings.EqualFold(upgrade, "websocket")
}

/**
 * @info Gets the scheme type of the request body
 * @returns {bool}
 */
func (r *Request) SchemeType() string {
	if r.IsTLS() {
		return "http"
	}
	if scheme := r.ref.Header.Get("X-Forwarded-Proto"); scheme != "" {
		return scheme
	}
	if scheme := r.ref.Header.Get("X-Forwarded-Protocol"); scheme != "" {
		return scheme
	}
	if ssl := r.ref.Header.Get("X-Forwarded-Ssl"); ssl == "on" {
		return "https"
	}
	if scheme := r.ref.Header.Get("X-Forwarded-Scheme"); scheme != "" {
		return scheme
	}
	return "http"
}

/**
 * @info Gets the values from request form
 * @param {string} [key] The key of the value
 * @returns {string}
 */
func (r *Request) FormValue(key string) string {
	return r.ref.FormValue(key)
}

/**
 * @info Gets all the form param values
 * @returns {url.Values, error}
 */
func (r *Request) FormParams() (url.Values, error) {
	if strings.HasPrefix(r.ref.Header.Get("Content-type"), "multipart/form-data") {
		if err := r.ref.ParseMultipartForm(24); err != nil {
			return nil, err
		}
	} else {
		if err := r.ref.ParseForm(); err != nil {
			return nil, err
		}
	}
	return r.ref.Form, nil
}

/**
 * @info Gets a file from request form
 * @returns {multipart.FileHeader, error}
 */
func (r *Request) FormFile(key string) (*multipart.FileHeader, error) {
	f, file, err := r.ref.FormFile(key)
	if err != nil {
		return nil, err
	}
	f.Close()
	return file, nil
}

/**
 * @info Gets a Multi part form from request form
 * @returns {multipart.Form, error}
 */
func (r *Request) MultipartForm() (*multipart.Form, error) {
	err := r.ref.ParseMultipartForm(24)
	return r.ref.MultipartForm, err
}

/**
 * @info Get all the cookies from the request
 * @returns {[]*http.Cookie}
 */
func (r *Request) Cookies() []*http.Cookie {
	return r.ref.Cookies()
}

/**
 * @info Get a paticular cookie by its key
 * @param {string} [key] key of the cookie
 * @returns {*http.Cookie}
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
 * @info Set a paticular Header
 * @param {string} [key] key of the Header
 * @param {string} [value] value of the Header
 * @returns {*Request}
 */
func (r *Request) SetHeader(key string, value string) *Request {
	r.header.Set(key, value)
	return r
}

/**
 * @info Get a paticular Header by its key
 * @param {string} [key] key of the Header
 * @returns {string}
 */
func (r *Request) GetHeader(key string) string {
	return r.header.Get(key)
}
