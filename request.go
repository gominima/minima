package minima

import (
	"encoding/json"
	"log"
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
	Path  string
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
func request(httRequest *http.Request) *Request {
	req := &Request{}
	req.ref = httRequest
	req.header = &IncomingHeader{}
	req.fileReader = nil
	req.method = httRequest.Proto
	req.query = httRequest.URL.Query()
	for i, v := range httRequest.Header {
		req.header.Set(strings.ToLower(i), strings.Join(v, ","))
	}
	if req.header.Get("content-type") == "application/json" {
		req.json = json.NewDecoder(httRequest.Body)
	} else {
		httRequest.ParseForm()
	}
	if len(httRequest.PostForm) > 0 && len(req.body) == 0 {
		req.body = make(map[string][]string)
	}
	for key, value := range httRequest.PostForm {
		req.body[key] = value
	}
	return req

}

/**
@info Gets param from route path
@param {string} [key] Key of the route param
@returns {string}
*/
func (r *Request) GetParam(key string) string {
	var val string
	for _, v := range r.Params {
		if v.Path == r.GetPathURl() && v.key == key {
			val = v.value
		}
	}
	return val
}

/**
@info Gets request path url
@returns {string}
*/
func (r *Request) GetPathURl() string {
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
@info Gets raw IncomingHeader instance
@returns {IncomingHeader}
*/
func (r *Request) Header() *IncomingHeader {
	return r.header
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
	if r.query[key][0] == "" {
		log.Panic("No query param found with given key")
	}
	return r.query[key][0]
}
