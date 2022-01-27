package minima

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"text/template"
)

/**
@info The response instance structure
@property {http.ResponseWriter} [ref] The net/http response instance
@property {string} [url] The route url
@property {string} [method] The route http method
@property {OutgoingHeader} [header] The response header instance
@property {*map[string]interface{}} [props] The minima instance props
@property {string} [host] The minima host
@property {bool} [HasEnded] Whether the response has ended
*/
type Response struct {
	ref      http.ResponseWriter
	url      string
	method   string
	ended    bool
	header   *OutgoingHeader
	props    *map[string]interface{}
	host     string
	HasEnded bool
}

/**
@info Make a new default response instance
@param {http.Request} [req] The net/http request instance
@param {http.ResponseWriter} [rs] The net/http response instance
@param {map[string]interface{}} [props] The net/http response instance
@returns {Response}
*/
func response(rw http.ResponseWriter, req *http.Request, props *map[string]interface{}) *Response {
	return &Response{
		ref:    rw,
		header: NewResHeader(rw, req),
		url:    req.URL.Path,
		method: req.Method,
		host:   req.Host,
		props:  props,
	}
}

/**
@info Returns a outgoing header instance
@returns {OutgoingHeader}
*/
func (res *Response) Header() *OutgoingHeader {
	return res.header
}

/**
@info Sends string to the route
@param {string} [content] The content to write
@returns {Response}
*/
func (res *Response) Send(content string) *Response {
	if res.header.Get("Content-Type") == "" {
		res.header.Set("Content-Type", "text/html;charset=utf-8")
	}

	res.WriteBytes([]byte(content))
	return res
}

/**
@info Writes bytes to the route
@param {[]bytes} [bytes] The bytes to write
@returns {Response}
*/
func (res *Response) WriteBytes(bytes []byte) error {
	var err error
	if _, writeErr := res.ref.Write(bytes); writeErr != nil {
		err = writeErr
	}

	return err
}

func (res *Response) sendContent(contentType string, content []byte) {
	res.header.Set("Content-Type", contentType)
	if err := res.WriteBytes(content); err != nil {
		log.Panicf("Failed to flush the buffer: %v", err)
		return
	}
}

/**
@info Writes json content to the route
@param {interface{}} [content] The json struct to write to the page
@returns {Response}
*/
func (res *Response) Json(content interface{}) *Response {
	output, err := json.Marshal(content)
	if err != nil {
		output = []byte("")
	}

	res.sendContent("application/json", output)
	return res
}

/**
@info Returns error to the route
@param {int} [status] The status code of the error
@param {string} [err] The error to write
@returns {Response}
*/
func (res *Response) Error(status int, err string) *Response {
	res.sendContent("text/html", []byte(err))
	log.Panic(err)
	return res
}

/**
@info Returns raw http.ResponseWriter instance
@returns {http.ResponseWriter}
*/
func (res *Response) Raw() http.ResponseWriter {
	return res.ref
}

/**
@info Renders a html page with payload data to the route
@param {string} [path] The dir path of the html page
@param {interface{}} [data] The payload data to pass in html page
@returns {Response}
*/
func (res *Response) Render(path string, data interface{}) *Response {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		log.Panic("Path "+path+" was not found!", err)
		res.header.Status(500)
	}

	var bytes bytes.Buffer
	if err = tmpl.Execute(&bytes, data); err != nil {
		log.Print("Template render failed: ", err)
		res.header.Status(500)
	}

	res.WriteBytes(bytes.Bytes())
	return res

}

/**
@info Redirects to a different route
@param {string} [url] The url of the route to redirect
@returns {Response}
*/
func (res *Response) Redirect(url string) *Response {
	res.header.Status(302)
	res.header.Set("Location", url)
	res.ended = true
	return res
}

/**
@info Sets response status
@param {int} [status] The status code for the response
@returns {Response}
*/
func (res *Response) Status(status int) *Response {
	res.header.Status(status)
	return res
}
