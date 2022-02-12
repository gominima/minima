package minima

import (
	"net/http"
)

/**
@info The Outgoing header structure
@property {http.Request} [req] The net/http request instance
@property {http.ResponseWriter} [res] The net/http response instance
@property {bool} [body] Whether body has been sent or not
@property {int} [status] response status code
*/
type OutgoingHeader struct {
	req *http.Request
	res http.ResponseWriter
}

var statusCodes = map[string]int{
	"OK":                         200,
	"Created":                    201,
	"Accepted":                   202,
	"No Content":                 204,
	"Reset Content":              205,
	"Partial Content":            206,
	"Moved Permanently":          301,
	"Found":                      302,
	"Not Modified":               304,
	"Use Proxy":                  305,
	"Switch Proxy":               306,
	"Temporary Redirect":         307,
	"Permanent Redirect":         308,
	"Bad Request":                400,
	"Unauthorized":               401,
	"Forbidden":                  403,
	"NOT FOUND":                  404,
	"Method Not Allowed":         405,
	"Payload Too Large":          413,
	"URI Too Long":               414,
	"Internal Server Error":      500,
	"Not Implemented":            501,
	"Bad Gateway":                502,
	"Service Unavailaible":       503,
	"Gateway Timeout":            504,
	"HTTP Version Not Supported": 505,
}

/**
@info Make a new default request header instance
@param {http.Request} [req] The net/http request instance
@param {http.ResponseWriter} [res] The net/http response instance
@returns {OutgoingHeader}
*/
func NewResHeader(res http.ResponseWriter, req *http.Request) *OutgoingHeader {
	return &OutgoingHeader{req, res}
}

/**
@info Sets and new header to response
@param {string} [key] Key of the new header
@param {string} [value] Value of the new header
@returns {OutgoingHeader}
*/
func (h *OutgoingHeader) Set(key string, value string) *OutgoingHeader {
	h.res.Header().Set(key, value)
	return h
}

/**
@info Gets the header from response headers
@param {string} [key] Key of the header
@returns {string}
*/
func (h *OutgoingHeader) Get(key string) string {
	return h.res.Header().Get(key)
}

/**
@info Deletes header from respose
@param {string} [key] Key of the header
@returns {OutgoingHeader}
*/
func (h *OutgoingHeader) Del(key string) *OutgoingHeader {
	h.res.Header().Del(key)
	return h
}

/**
@info Clones all headers from response
@returns {OutgoingHeader}
*/
func (h *OutgoingHeader) Clone() http.Header {
	return h.res.Header().Clone()
}

/**
@info Sets content lenght
@param {string} [len] The lenght of the content
@returns {OutgoingHeader}
*/
func (h *OutgoingHeader) Setlength(len string) *OutgoingHeader {
	h.Set("Content-length", len)
	return h
}

/**
@info Sets response status
@param {int} [code] The status code for the response
@returns {OutgoingHeader}
*/
func (h *OutgoingHeader) Status(code int) *OutgoingHeader {
	h.res.WriteHeader(code)
	return h
}

/**
@info Sends good stack of base headers
@returns {}
*/
func (h *OutgoingHeader) BaseHeaders() {
	h.Set("transfer-encoding", "chunked")
	h.Set("connection", "keep-alive")
}

/**
@info Flushes and writes header to route
@returns {bool}
*/
func (h *OutgoingHeader) Flush() bool {
	if h.Get("Content-Type") == "" {
		h.Set("Content-Type", "text/html;charset=utf-8")
	}

	if f, ok := h.res.(http.Flusher); ok {
		f.Flush()
	}

	return true
}
