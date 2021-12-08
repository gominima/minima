package fiable

import (
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"
)
type Param struct {
 path string
 value string
 ctx   context.Context
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

func (r*Request) Param(name string) (string, context.Context){
 var val string
 var ctx context.Context
 for _, v := range r.Params{
   if v.path == r.GetPathURl(){
   val = v.value
   ctx = v.ctx
   }
  	 
 }
 return val, ctx
}

func (r * Request) GetPathURl() string {
 return r.ref.URL.Path
}