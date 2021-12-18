package fiable

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"strings"
	// "strings"
)

type Param struct {
	Path  string
	key   string
	value string
}

type HeadInfo struct {
	key   string
	value string
}
type ReqHeader struct {
	headers []*HeadInfo
}

func (h *ReqHeader) Get(key string) string {
	var value string
	for _, v := range h.headers {
		if v.key == key {
			value = v.value
		}
	}
	return value
}

func (h *ReqHeader) Set(key string, v string) {
	h.headers = append(h.headers, &HeadInfo{key: key, value: v})
}

type Request struct {
	Ref        *http.Request
	fileReader *multipart.Reader
	body       map[string][]string
	method     string
	url        string
	Params     []*Param
	header     *ReqHeader
	json       *json.Decoder
	props      *map[string]interface{}
}

func request(httRequest *http.Request, props *map[string]interface{}) *Request {
	req := &Request{}
	req.Ref = httRequest
	req.header = &ReqHeader{}
	req.fileReader = nil
	req.method = httRequest.Proto
	req.props = props
	for i, v := range httRequest.Header {
		req.header.Set(strings.ToLower(i), strings.Join(v, ","))
	}
	if req.header.Get("content-type") == "application/json" {
		req.json = json.NewDecoder(httRequest.Body)
	} else {
		httRequest.ParseForm()
	}
	if len(httRequest.PostForm) > 0 && len(req.body) == 0 {
		req.body = make(map[string][]string)
	}
	for key, value := range httRequest.PostForm {
		req.body[key] = value
	}
	return req

}

func (r *Request) GetParam(name string) string {
	var val string
	for _, v := range r.Params {
		if v.Path == r.GetPathURl() && v.key == name {
			val = v.value
		}
	}
	return val
}

func (r *Request) GetPathURl() string {
	return r.Ref.URL.Path
}

func (r *Request) Body() map[string][]string {
	return r.body
}

func (r *Request) GetBodyValue(key string) []string {
	return r.body[key]
}

func (r *Request) Header() *ReqHeader {
	return r.header
}

func (r *Request) Json() *json.Decoder {
	return r.json
}
