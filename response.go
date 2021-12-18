package fiable

import (
	"net/http"
)

type Response struct {
	Ref http.ResponseWriter

	url      string
	method   string
	ended    bool
	props    *map[string]interface{}
	host     string
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

func (r *Response) Header() {

}

func (res *Response) Send(status int, content string) *Response {
	res.Ref.Write([]byte(content))

	return res
}
