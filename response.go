package fiable

import (
	"bufio"
	"net"
	"net/http"
)

type response struct {
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

func Response(rs http.ResponseWriter, req *http.Request, write *bufio.ReadWriter, connection net.Conn, props *map[string]interface{}) *response {
 res := &response{}
 res.response = rs
 res.connection = connection
 res.url = req.URL.Path
 res.method = req.Method
 res.host = req.Host
 res.write = write
 res.props = props
 return res

}


func (r *response) Header(){
	
}

func (r *response) Send(status int, content interface{}) error {

}