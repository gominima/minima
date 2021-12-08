package fiable

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
)
type Param struct {
 Path string
 key string
 value string

}
type Request struct {
	ref        *http.Request
	fileReader *multipart.Reader
	query      map[string][]string
	body       map[string][]string
	method     string
	url        string
	Params     []*Param
	
	

	json       *json.Decoder
	props      *map[string]interface{}
}

func request(httRequest *http.Request, props *map[string]interface{}) *Request{
 req := &Request{}
 req.ref = httRequest
 req.fileReader = nil
 req.method = httRequest.Proto
 req.props = props
 return req

}

func (r*Request) GetParam(name string) string{
var val string
for _, v:= range r.Params{
 if v.Path == r.GetPathURl() && v.key == name{
   val = v.value
 }
}
 return val
}


func (r * Request) GetPathURl() string {
 return r.ref.URL.Path
}