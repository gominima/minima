package fiable

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"strings"
)
type Param struct {
 Path string
 key string
 value string
}


type ReqHeader struct {
 keys map[string]string
}
func (h *ReqHeader) Get(key string) string {
return h.keys[key]
}
      
 func (h *ReqHeader) Set(key string, v string) {
 h.keys[key] = v
 }



type Request struct {
	ref        *http.Request
	fileReader *multipart.Reader
	query      map[string][]string
	body       map[string][]string
	method     string
	url        string
	Params     []*Param
	header      *ReqHeader
	json       *json.Decoder
	props      *map[string]interface{}
}

func request(httRequest *http.Request, props *map[string]interface{}) *Request{
 req := &Request{}
 req.ref = httRequest
 req.fileReader = nil
 req.method = httRequest.Proto
 req.props = props
 for i,v := range httRequest.Header{
	req.header.Set(strings.ToLower(i), strings.Join(v, ","))
 }
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