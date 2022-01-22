package minima

import (
	"fmt"
	"log"
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
	req    *http.Request
	res    http.ResponseWriter
	Body   bool
	status int
	Done   bool
}

var status = map[int]string{
	200: "OK",
	201: "Created",
	202: "Accepted",
	204: "No Content",
	205: "Reset Content",
	206: "Partial Content",
	301: "Moved Permanently",
	302: "Found",
	304: "Not Modified",
	305: "Use Proxy",
	306: "Switch Proxy",
	307: "Temporary Redirect",
	308: "Permanent Redirect",
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "NOT FOUND",
	405: "Method Not Allowed",
	413: "Payload Too Large",
	414: "URI Too Long",
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailaible",
	504: "Gateway Timeout",
	505: "HTTP Version Not Supported",
}

/**
@info Make a new default request header instance
@param {http.Request} [req] The net/http request instance
@param {http.ResponseWriter} [res] The net/http response instance
@returns {OutgoingHeader}
*/
func NewResHeader(res http.ResponseWriter, req *http.Request) *OutgoingHeader {
	h := &OutgoingHeader{}
	h.req = req
	h.res = res
	h.Body = false
	h.Done = false
	return h
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
@info Sets new header to response
@param {string} [key] Key of the new header
@param {string} [value] Value of the new header
@returns {OutgoingHeader}
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
	h.Set("Content-lenght", len)
	return h
}

func (h *OutgoingHeader) BasicDone() bool {
	return h.Done
}

/**
@info Sets response status
@param {int} [code] The status code for the response
@returns {OutgoingHeader}
*/
func (h *OutgoingHeader) Status(code int) *OutgoingHeader {
	h.status = code
	return h
}

/**
@info Sends good stack of base headers
@returns {}
*/
func (h *OutgoingHeader) SendBaseOutgoingHeaders() {
	if !h.Done && !h.BasicDone() {
		if h.status == 0 {
			h.status = 200
		}
		fmt.Fprintf(h.res, "HTTP/%d.%d %03d %s\r\n", 1, 1, h.status, status[h.status])
		h.Set("transfer-encoding", "chunked")
		h.Set("connection", "keep-alive")
	}
}

/**
@info Flushes and writes header to route
@returns {bool}
*/
func (h *OutgoingHeader) Flush() bool {
	if h.Body {
		log.Panic("Cannot send OutgoingHeaders in middle of Body")
		return false
	}
	if !h.BasicDone() {
		h.SendBaseOutgoingHeaders()
	}
	if h.Get("Content-Type") == "" {
		h.Set("Content-Type", "text/html;charset=utf-8")
	}
	if err := h.res.Header().Write(h.res); err != nil {
		return false
	}
	if f, ok := h.res.(http.Flusher); ok {
		f.Flush()
	}
	return true
}

func (h *OutgoingHeader) CanSend() bool {
	if h.BasicDone() {
		if !h.Body {
			return true
		} else {
			return false
		}
	}
	return true
}
