package minima

import (
	"encoding/json"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	// "strings"
)
/**
	@info The param structure
	@property {string} [Path] The path
	@property {string} [key] The key
	@property {string} [value] The value
*/
type Param struct {
	Path  string
	key   string
	value string
}
/**
	@info The HeadInfo structure
	@property {string} [key] The key
	@property {string} [value] The value
*/
type HeadInfo struct {
	key   string
	value string
}
/**
	@info The ReqHeader structure
	@property {[]HeadInfo} [headers] The headers
*/
type ReqHeader struct {
	headers []*HeadInfo
}
/**
	@info Get a ReqHeader
	@param {string} [key] The key
	@returns {string} The value
*/
func (h *ReqHeader) Get(key string) string {
	var value string
	for _, v := range h.headers {
		if v.key == key {
			value = v.value
		}
	}
	return value
}
/**
	@info Set a ReqHeader
	@param {string} [key] The key
	@param {string} [value] The value
*/
func (h *ReqHeader) Set(key string, v string) {
	h.headers = append(h.headers, &HeadInfo{key: key, value: v})
}
/**
	@info The request structure
	@property {http.Request} [Ref] The request
	@property {multipart.Reader} [fileReader] The file reader
	@property {map[]string[]string} [body] The body
	@property {string} [method] The method
	@property {string} [url] The url
	@property {[]Param} [Params] The params
	@property {url.Values} [query] The query
	@property {ReqHeader} [header] The request header
	@property {json.Decoder} [json] The json decoder
	@property {map[]string} [props] The properties
*/
type Request struct {
	Ref        *http.Request
	fileReader *multipart.Reader
	body       map[string][]string
	method     string
	url        string
	Params     []*Param
	query      url.Values
	header     *ReqHeader
	json       *json.Decoder
	props      *map[string]interface{}
}
/**
	@info Make a new request
	@param {http.Request} [httpRequest] The http request
	@param {map[]string} [props] The properties
	@returns {Request}
*/
func request(httRequest *http.Request, props *map[string]interface{}) *Request {
	req := &Request{}
	req.Ref = httRequest
	req.header = &ReqHeader{}
	req.fileReader = nil
	req.method = httRequest.Proto
	req.props = props
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
	@info Get a param
	@param {string} [name] The name
	@returns {string}
*/
func (r *Request) GetParam(name string) string {
	var val string
	for _, v := range r.Params {
		if v.Path == r.GetPathURl() && v.key == name {
			val = v.value
		}
	}
	return val
}
/**
	@info Get the Path URL
	@returns {string}
*/
func (r *Request) GetPathURl() string {
	return r.Ref.URL.Path
}
/**
	@info Get the Body
	@returns {map[]string[]string}
*/
func (r *Request) Body() map[string][]string {
	return r.body
}
/**
	@info Get a Body value
	@param {string} [key] The key
	@returns {[]string}
*/
func (r *Request) GetBodyValue(key string) []string {
	return r.body[key]
}
/**
	@info Get the Header
	@returns {ReqHeader}
*/
func (r *Request) Header() *ReqHeader {
	return r.header
}
/**
	@info Get the json decoder
	@returns {json.Decoder}
*/
func (r *Request) Json() *json.Decoder {
	return r.json
}
/**
	@info Get the method
	@returns {string}
*/
func (r *Request) Method() string {
	return r.method
}
/**
	@info Get a query
	@param {string} [key] The key
	@returns {string}
*/
func (r *Request) GetQuery(key string) string {
	if r.query[key][0] == "" {
		log.Panic("No query param found with given key")
	}
	return r.query[key][0]
}
