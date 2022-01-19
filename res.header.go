package minima

import (
	"fmt"
	"log"
	"net/http"
)

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

func NewResHeader(res http.ResponseWriter, req *http.Request) *OutgoingHeader {
	h := &OutgoingHeader{}
	h.req = req
	h.res = res
	h.Body = false
	h.Done = false
	return h
}

func (h *OutgoingHeader) Set(key string, value string) {
	h.res.Header().Set(key, value)
}

func (h *OutgoingHeader) Get(key string) string {
	return h.res.Header().Get(key)
}

func (h *OutgoingHeader) Del(key string) {
	h.res.Header().Del(key)
}

func (h *OutgoingHeader) Clone(key string) {
	h.res.Header().Clone()
}

func (h *OutgoingHeader) Setlength(len string) {
	h.Set("Content-lenght", len)
}

func (h *OutgoingHeader) BasicDone() bool {
	return h.Done
}
func (h *OutgoingHeader) Status(code int) {
	h.status = code
}
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
func (h *OutgoingHeader) Flush() bool {
	if h.Body {
		log.Panic("Cannot send OutgoingHeaders in middle of body")
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