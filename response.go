package fiable

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

type Response struct {
	response http.ResponseWriter
	write *bufio.ReadWriter
	connection net.Conn
	url string
	method string
	ended bool
	props *map[string]interface{}
	host string
	HasEnded bool

}

func response(rs http.ResponseWriter, req *http.Request, write *bufio.ReadWriter, connection net.Conn, props *map[string]interface{}) *Response {
 res := &Response{}
 res.response = rs
 res.connection = connection
 res.url = req.URL.Path
 res.method = req.Method
 res.host = req.Host
 res.write = write
 res.props = props
 return res

}


func (r *Response) Header(){
	
}

func (r *Response) Send(status int, content interface{})  {
 fmt.Print(content)
}