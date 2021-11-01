package fiable

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/url"
)

type request struct {
	ref        *http.Request
	fileReader *multipart.Reader
	query      map[string][]string
	body       map[string][]string
	method     string
	url        string
	_url       *url.URL
	
	query      map[string][]string
	body       map[string][]string

	json       *json.Decoder
	props      *map[string]interface{}
}

func Request(httRequest *http.Request, props *map[string]interface{}) *request{
 req := &request{}
 req.ref = httRequest
 req.fileReader = nil
 req.method = 

}