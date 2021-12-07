package fiable

import (
	"bufio"
	"fmt"

	"net"
	"net/http"
)

type Response struct {
	Ref http.ResponseWriter
	write *bufio.ReadWriter
	connection net.Conn
	url string
	method string
	ended bool
	props *map[string]interface{}
	host string
	HasEnded bool

}

func response(rs http.ResponseWriter, req *http.Request, props *map[string]interface{}) *Response {
 res := &Response{}
 res.Ref = rs

 res.url = req.URL.Path
 res.method = req.Method
 res.host = req.Host

 res.props = props

 return res

}


func (r *Response) Header(){
	
}

func (res *Response) Send(ty int, content string) *Response {
 fmt.Fprint(res.Ref, content)
 return res
}