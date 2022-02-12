package minima

import (
	"bytes"
	"encoding/json"
	"io"
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
@info Gets header from response
@param {string} [key] Key of the header
@returns {string}
*/
func (res *Response) GetHeader(key string) string {
	return res.header.Get(key)
}

/**
@info Sets headers for response
@param {string} [key] Key of the header
@param {string} [value] Value of the header
@returns {string}
*/
func (res *Response) SetHeader(key string, value string) *Response {
	res.header.Set(key, value)
	return res
}

/**
@info Gets header from response
@param {string} [key] Key of the header
@returns {string}
*/
func (res *Response) DelHeader(key string) *Response {
	res.header.Del(key)
	return res
}

/**
@info Clones all header from response
@returns {http.Header}
*/
func (res *Response) CloneHeaders() http.Header {
	return res.header.Clone()

}

/**
@info Sets length of the response body
@param {string} [len] length value of the header
@returns {*Response}
*/
func (res *Response) Setlength(len string) *Response {
	res.header.Setlength(len)
	return res
}

/**
@info Sets a good stack of base headers for response
@returns {*Response}
*/
func (res *Response) SetBaseHeaders() *Response {
	res.header.BaseHeaders()
	return res
}

/**
@info Flushes headers to the response body
@returns {*Response}
*/
func (res *Response) FlushHeaders() *Response {
	res.header.Flush()
	return res
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
@info Ends connection to the route page
@returns {error}
*/
func (res *Response) CloseConn() error {
	var returnerr error
	if w, ok := res.ref.(io.Closer); ok {
		err := w.Close()
		returnerr = err
	}
	return returnerr
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

/**
@info Set a cookie
@param {*http.Cookie} [cookie]
@returns {Response}
*/
func (res *Response) SetCookie(cookie *http.Cookie) *Response {
	http.SetCookie(res.ref, cookie)
	return res
}

/**
@info Clear a cookie
@param {*http.Cookie} [cookie]
@returns {Response}
*/
func (res *Response) ClearCookie(cookie *http.Cookie) *Response {
	cookie.MaxAge = -1
	http.SetCookie(res.ref, cookie)
	return res
}

/**
@info Set status code as 200
@returns {Response}
*/
func (res *Response) OK() *Response {
	res.Status(statusCodes["OK"])
	return res
}

/**
@info Set status code as 301
@returns {Response}
*/
func (res *Response) MovedPermanently() *Response {
	res.Status(statusCodes["Moved Permanently"])
	return res
}

/**
@info Set status code as 307
@returns {Response}
*/
func (res *Response) TemporaryRedirect() *Response {
	res.Status(statusCodes["Temporary Redirect"])
	return res
}

/**
@info Set status code as 400
@returns {Response}
*/
func (res *Response) BadRequest() *Response {
	res.Status(statusCodes["Bad Request"])
	return res
}

/**
@info Set status code as 401
@returns {Response}
*/
func (res *Response) Unauthorized() *Response {
	res.Status(statusCodes["Unauthorized"])
	return res
}

/**
@info Set status code as 403
@returns {Response}
*/
func (res *Response) Forbidden() *Response {
	res.Status(statusCodes["Forbidden"])
	return res
}

/**
@info Set status code as 404
@returns {Response}
*/
func (res *Response) NotFound() *Response {
	res.Status(statusCodes["NOT FOUND"])
	return res
}

/**
@info Set status code as 500
@returns {Response}
*/
func (res *Response) InternalServerError() *Response {
	res.Status(statusCodes["Internal Server Error"])
	return res
}

/**
@info Set status code as 502
@returns {Response}
*/
func (res *Response) BadGateway() *Response {
	res.Status(statusCodes["Bad Gateway"])
	return res
}

/**
@info Set status code as 503
@returns {Response}
*/
func (res *Response) ServiceUnavailaible() *Response {
	res.Status(statusCodes["Service Unavailaible"])
	return res
}
