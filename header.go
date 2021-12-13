package fiable

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strconv"

)

var statusCodes = map[int]string{
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

type Header struct {
	res   http.ResponseWriter
	req    *http.Request
	w     *bufio.ReadWriter
	bodyDone   bool
	basicDone bool
	hasLength  bool
	StatusCode int
	ProtoMajor int
	ProtoMinor int
}

func newHeader(res http.ResponseWriter, req *http.Request, w *bufio.ReadWriter) *Header{
	h := &Header{}
	h.res = res
	h.req = req
	h.w= w
	h.bodyDone = false
	h.basicDone = false
	h.ProtoMinor = 1
	h.ProtoMajor = 1
	return h
}


func (h *Header) Set(key string, value string) *Header {
	h.res.Header().Set(key, value)
	return h
}

func (h*Header) Delete(key string) *Header{
 h.res.Header().Del(key)
 return h
}
func (h *Header) Get(key string) string {
	return h.res.Header().Get(key)
}

func (h*Header) GetReqHeaders(key string) []string{
 return h.req.Header[key]
}

func (h *Header) SetLength(len int){
	h.res.Header().Set("Content-Length", strconv.Itoa(len))
	h.hasLength = true
}
func (h *Header) SetStatus(code int) {
	h.StatusCode = code
}
func (h *Header) CanSendHeader() bool {
	if h.basicDone == true {
		if h.bodyDone == false {
			return true
		}
		return false
	}
	return true
}

func (h*Header) Flush() bool{
	if h.bodyDone == true {
		log.Panic("Cannot send headers in middle of body")
		return false
	}
	if h.basicDone == false {
		h.Basics()
	}

	if h.Get("Content-Type") == "" {
		h.Set("Content-Type", "text/html;charset=utf-8")
	}
	if err := h.res.Header().Write(h.w); err != nil {
		return false
	}
	var chunkSize = fmt.Sprintf("%x", 0)
	h.w.WriteString(chunkSize + "\r\n" + "\r\n")
	h.w.Writer.Flush()
 return true
}

func (h *Header) Basics(){
	if h.StatusCode == 0 {h.StatusCode = 200}
	fmt.Fprintf(h.w, "HTTP/%d.%d %03d %s\r\n", h.ProtoMajor, h.ProtoMinor, h.StatusCode, statusCodes[h.StatusCode])
	h.Set("transfer-encoding", "chunked")
	h.Set("connection", "keep-alive")
	h.basicDone = true
}