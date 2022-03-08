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
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

/**
 * @info The response instance structure
 * @property {http.ResponseWriter} [ref] The net/http response instance
 * @property {string} [url] The route url
 * @property {string} [method] The route http method
 * @property {OutgoingHeader} [header] The response header instance
 * @property {string} [host] The minima host
 * @property {bool} [HasEnded] Whether the response has ended
 */
type Response struct {
	ref      http.ResponseWriter
	url      string
	method   string
	ended    bool
	header   *OutgoingHeader
	host     string
	HasEnded bool
}

/**
 * @info Make a new default response instance
 * @param {http.Request} [req] The net/http request instance
 * @param {http.ResponseWriter} [rs] The net/http response instance
 * @returns {Response}
 */
func response(rw http.ResponseWriter, req *http.Request) *Response {
	return &Response{
		ref:    rw,
		header: NewResHeader(rw, req),
		url:    req.URL.Path,
		method: req.Method,
		host:   req.Host,
	}
}

/**
 * @info Gets header from response
 * @param {string} [key] Key of the header
 * @returns {string}
 */
func (res *Response) GetHeader(key string) string {
	return res.header.Get(key)
}

/**
 * @info Sets headers for response
 * @param {string} [key] Key of the header
 * @param {string} [value] Value of the header
 * @returns {string}
 */
func (res *Response) SetHeader(key string, value string) *Response {
	res.header.Set(key, value)
	return res
}

/**
 * @info Gets header from response
 * @param {string} [key] Key of the header
 * @returns {string}
 */
func (res *Response) DelHeader(key string) *Response {
	res.header.Del(key)
	return res
}

/**
 * @info Clones all header from response
 * @returns {http.Header}
 */
func (res *Response) CloneHeaders() http.Header {
	return res.header.Clone()
}

/**
 * @info Sets length of the response body
 * @param {string} [len] length value of the header
 * @returns {*Response}
 */
func (res *Response) Setlength(len string) *Response {
	res.header.Setlength(len)
	return res
}

/**
 * @info Sets a good stack of base headers for response
 * @returns {*Response}
 */
func (res *Response) SetBaseHeaders() *Response {
	res.header.BaseHeaders()
	return res
}

/**
 * @info Flushes headers to the response body
 * @returns {*Response}
 */
func (res *Response) FlushHeaders() *Response {
	res.header.Flush()
	return res
}

/**
 * @info Sends string to the route
 * @param {string} [content] The content to write
 * @returns {Response}
 */
func (res *Response) Send(content string) *Response {
	if res.header.Get("Content-Type") == "" {
		res.header.Set("Content-Type", "text/html;charset=utf-8")
	}

	res.WriteBytes([]byte(content))
	return res
}

/**
 * @info Writes bytes to the route
 * @param {[]bytes} [bytes] The bytes to write
 * @returns {Response}
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

func (res *Response) setContent(contentType string) {
	res.header.Set("Content-Type", contentType)
}

/**
 * @info Writes json content to the route
 * @param {interface{}} [content] The json struct to write to the page
 * @returns {Response}
 */
func (res *Response) JSON(content interface{}) *Response {
	output, err := json.Marshal(content)
	if err != nil {
		output = []byte("")
	}

	res.sendContent("application/json", output)
	return res
}

/**
 * @info Writes xml content to the route
 * @param {interface{}} [content] The xml content to write to the page
 * @param {string}  [indent] The indentation of the content
 * @returns {error}
 */
func (res *Response) XML(content interface{}, indent string) error {
	res.setContent("application/json; charset=utf-8")
	enc := xml.NewEncoder(res.ref)
	if indent != "" {
		enc.Indent("", indent)
	}
	if _, err := res.ref.Write([]byte(xml.Header)); err != nil {
		return err
	}
	return enc.Encode(content)
}

/**
 * @info Streams content to the route
 * @param {string} [contentType] The content type to stream
 * @param {io.Reader} [read]  The io.Reader instance
 * @returns {error}
 */
func (res *Response) Stream(contentType string, read io.Reader) error {
	res.setContent(contentType)
	_, err := io.Copy(res.ref, read)
	return err
}

/**
 * @info Sets page's content to none
 * @param {int} [code] The status code
 * @returns {error}
 */
func (res *Response) NoContent(code int) error {
	res.Status(code)
	res.CloseConn()
	return nil
}

/**
 * @info Returns error to the route
 * @param {int} [status] The status code of the error
 * @param {string} [err] The error to write
 * @returns {Response}
 */
func (res *Response) Error(status int, err string) *Response {
	res.Status(status)
	res.sendContent("text/html", []byte(err))
	res.CloseConn()
	return res
}

/**
 * @info Returns raw http.ResponseWriter instance
 * @returns {http.ResponseWriter}
 */
func (res *Response) Raw() http.ResponseWriter {
	return res.ref
}

/**
 * @info Renders a html page with payload data to the route
 * @param {string} [path] The dir path of the html page
 * @param {interface{}} [data] The payload data to pass in html page
 * @returns {Response}
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
 * @info Ends connection to the route page
 * @returns {error}
 */
func (res *Response) CloseConn() error {
	var returnErr error
	if w, ok := res.ref.(io.Closer); ok {
		err := w.Close()
		returnErr = err
	}
	return returnErr
}

/**
 * @info Sends a file to the server 
 * @returns {error}
 */
func (res *Response) File(dir string) (error) {
	f, err := os.Open(dir)
	if err != nil {
	  log.Panicf("Minima: file %v wasn't found", dir)
	}
	defer f.Close()

	fi, _ := f.Stat()
	if fi.IsDir() {
		dir = filepath.Join(dir, "")
		f, err = os.Open(dir)
		if err != nil {
			log.Panicf("Minima: file %v wasn't found", dir)
		}
		defer f.Close()
		if fi, err = f.Stat(); err != nil {
			return err
		}
	}
	http.ServeContent(res.ref, res.header.req, fi.Name(), fi.ModTime(), f)
	return nil
}
/**
 * @info Redirects to a different route
 * @param {string} [url] The url of the route to redirect
 * @returns {Response}
 */
func (res *Response) Redirect(url string) *Response {
	http.Redirect(res.Raw(), res.header.req, url, http.StatusTemporaryRedirect)
	res.ended = true
	return res
}

/**
 * @info Sets response status
 * @param {int} [status] The status code for the response
 * @returns {Response}
 */
func (res *Response) Status(status int) *Response {
	res.header.Status(status)
	return res
}

/**
 * @info Set a cookie
 * @param {*http.Cookie} [cookie]
 * @returns {Response}
 */
func (res *Response) SetCookie(cookie *http.Cookie) *Response {
	http.SetCookie(res.ref, cookie)
	return res
}

/**
 * @info Clear a cookie
 * @param {*http.Cookie} [cookie]
 * @returns {Response}
 */
func (res *Response) ClearCookie(cookie *http.Cookie) *Response {
	cookie.MaxAge = -1
	http.SetCookie(res.ref, cookie)
	return res
}

/**
 * @info Set status code as 200
 * @returns {Response}
 */
func (res *Response) OK() *Response {
	res.Status(statusCodes["OK"])
	return res
}

/**
 * @info Set status code as 301
 * @returns {Response}
 */
func (res *Response) MovedPermanently() *Response {
	res.Status(statusCodes["Moved Permanently"])
	return res
}

/**
 * @info Set status code as 307
 * @returns {Response}
 */
func (res *Response) TemporaryRedirect() *Response {
	res.Status(statusCodes["Temporary Redirect"])
	return res
}

/**
 * @info Set status code as 400
 * @returns {Response}
 */
func (res *Response) BadRequest() *Response {
	res.Status(statusCodes["Bad Request"])
	return res
}

/**
 * @info Set status code as 401
 * @returns {Response}
 */
func (res *Response) Unauthorized() *Response {
	res.Status(statusCodes["Unauthorized"])
	return res
}

/**
 * @info Set status code as 403
 * @returns {Response}
 */
func (res *Response) Forbidden() *Response {
	res.Status(statusCodes["Forbidden"])
	return res
}

/**
 * @info Set status code as 404
 * @returns {Response}
 */
func (res *Response) NotFound() *Response {
	res.Status(statusCodes["NOT FOUND"])
	return res
}

/**
 * @info Set status code as 500
 * @returns {Response}
 */
func (res *Response) InternalServerError() *Response {
	res.Status(statusCodes["Internal Server Error"])
	return res
}

/**
 * @info Set status code as 502
 * @returns {Response}
 */
func (res *Response) BadGateway() *Response {
	res.Status(statusCodes["Bad Gateway"])
	return res
}

/**
 * @info Set status code as 503
 * @returns {Response}
 */
func (res *Response) ServiceUnavailable() *Response {
	res.Status(statusCodes["Service Unavailable"])
	return res
}
