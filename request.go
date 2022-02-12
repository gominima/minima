package minima

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

/**
@info The request param structure
@property {string} [Path] Route path of the param
@property {string} [key] Key for the param
@property {string} [value] Value of the param
*/
type Param struct {
	key   string
	value string
}

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
	Params     []*Param
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
	var value string
	for _, v := range r.Params {
		if v.key == key {
			value = v.value
			break
		}
	}
	return value
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
func (r *Request) GetQuery(key string) string {
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
func (r *Request) Cookie(key string) *http.Cookie {
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
