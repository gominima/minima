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
func response(rs http.ResponseWriter, req *http.Request, props *map[string]interface{}) *Response {
	res := &Response{}
	res.ref = rs
	res.header = NewResHeader(rs, req)
	res.url = req.URL.Path
	res.method = req.Method
	res.host = req.Host
	res.props = props
	return res

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

	if !res.header.BasicDone() && res.header.CanSend() {
		if res.header.Flush() {
			log.Print("Failed to push headers")
		}
		res.header.Done = true
		res.header.Body = true
	}
	var bytes = []byte(content)
	res.WriteBytes(bytes)
	return res
}

/**
@info Writes bytes to the route
@param {[]bytes} [bytes] The bytes to write
@returns {Response}
*/
func (res *Response) WriteBytes(bytes []byte) error {
	var errr error
	_, err := res.ref.Write(bytes)
	if err != nil {
		errr = err
	}
	return errr
}

func (res *Response) sendContent(status int, contentType string, content []byte) {
	if res.header.BasicDone() {
		res.header.Status(status)
	}
	if res.header.CanSend() {
		res.header.Set("Content-Type", contentType)
		if Done := res.header.Flush(); !Done {
			log.Print("Failed to write headers")
			res.header.Done = true
			res.header.Body = true
			return
		}
	}
	err := res.WriteBytes(content)
	if err != nil {
		log.Panicf("Failed to flush the buffer, error: %v", err)
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
		res.sendContent(500, "application/json", []byte(""))
	} else {
		res.sendContent(200, "application/json", output)
	}
	return res
}

/**
@info Returns error to the route
@param {int} [status] The status code of the error
@param {string} [err] The error to write
@returns {Response}
*/
func (res *Response) Error(status int, err string) *Response {
	res.sendContent(status, "text/html", []byte(err))
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
		log.Panic("Given path was not found", err)
		res.header.Status(500)
		res.header.Flush()

	}
	var byt bytes.Buffer
	err = tmpl.Execute(&byt, data)
	if err != nil {
		log.Print("Template render failed ", err)
		res.header.Status(500)
		res.header.Flush()
	}
	res.WriteBytes(byt.Bytes())
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
	res.header.Flush()
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
