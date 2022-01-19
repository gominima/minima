package minima

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"text/template"
)
/**
	@info The response structure
	@property {http.ResponseWriter} [Ref] The response writer
	@property {string} [url] The url
	@property {string} [method] The method
	@property {bool} [ended] Whether the response has ended
	@property {Header} [header] The response header
	@property {map[]string} [props] The properties
	@property {string} [host] The host
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
	@info Make a new response
	@param {http.ResponseWriter} [Ref] The response writer
	@param {http.Request} [req] The http request
	@param {map[]string} [props] The properties
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
func (res *Response) Header() OutgoingHeader {
	return *res.header
}

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
	@info Write bytes
	@param {[]byte} [bytyes] The bytes to write
	@returns {error}
*/
func (res *Response) WriteBytes(bytes []byte) error {
	var errr error
	_, err := res.ref.Write(bytes)
	if err != nil {
		errr = err
	}
	return errr
}
/**
	@info Send content
	@param {int} [status] The status code
	@param {string} [contentType] The contentType
	@param {[]byte} [content] The content to send
*/
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
	}
	return res
<<<<<<< HEAD
/**
	@info Send an error and log it
	@param {int} [status] The status code
	@param {string} [str] The error to send
*/
func (res *Response) Error(status int, str string) {
=======
	@returns {http.ResponseWriter}
func (res *Response) Raw() http.ResponseWriter {
	return res.ref
}
<<<<<<< HEAD
/**
	@info Render a path
	@param {string} [path] The path
	@param {interface} [data] The data
*/
func (res *Response) Render(path string, data interface{}) {
=======
func (res *Response) Render(path string, data interface{}) *Response {
>>>>>>> 8c7aafb0132fdea03a58145f8ab9901e321e8614
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
	@info Redirect a request
	@param {string} [url] The url
	@returns {Response}
*/
func (res *Response) Redirect(url string) *Response {
	res.header.Status(302)
	res.header.Set("Location", url)
	res.header.Flush()
	res.ended = true
	return res
}

func (res *Response) Status(status int) *Response {
	res.header.Status(status)
	return res
}
