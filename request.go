package fiable

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/url"
)

type Request struct {
	ref        *http.Request
	fileReader *multipart.Reader
	query      map[string][]string
	body       map[string][]string
	method     string
	url        string
	_url       *url.URL
	
	

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

func (r*Request) Param(name string) interface{}{
	result := r.ref.Context().Value(name)
	return result
}


func (r * Request) GetPathURl() string {
 return r.url
}