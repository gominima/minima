package fiable

import (
	"bufio"
	"net"
	"net/http"
)

type Response struct {
	Ref http.ResponseWriter
	write *bufio.ReadWriter
	connection net.Conn
	header     *Header
	url string
	method string
	ended bool
	props *map[string]interface{}
	host string
	HasEnded bool

}

func response(rs http.ResponseWriter, req *http.Request, props *map[string]interface{}, conn net.Conn, writer *bufio.ReadWriter) *Response {
 res := &Response{}
 res.Ref = rs
 res.header = newHeader(rs,req,writer)
 res.url = req.URL.Path
 res.method = req.Method
 res.host = req.Host
 res.write = writer
 res.props = props

 return res

}


func (r *Response) Header(){
	
}

func (res *Response) Send(status int, content string) *Response {
 res.header.SetStatus(status)
 res.write.Writer.Write([]byte(content))
 return res
}
