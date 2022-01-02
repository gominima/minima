package minima

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"text/template"
)

type Response struct {
	Ref      http.ResponseWriter
	url      string
	method   string
	ended    bool
	header   *OutgoingHeader
	props    *map[string]interface{}
	host     string
	HasEnded bool
}

func response(rs http.ResponseWriter, req *http.Request, props *map[string]interface{}) *Response {
	res := &Response{}
	res.Ref = rs
	res.header = NewResHeader(rs, req)
	res.url = req.URL.Path
	res.method = req.Method
	res.host = req.Host
	res.props = props
	return res

}
func (res *Response) Header() OutgoingHeader {
	return *res.header
}

func (res *Response) Send(status int, content string) *Response {
	res.header.status = status
	if !res.header.BasicDone() && res.header.CanSend() {
		if res.header.Flush() {
			log.Print("Failed to push headers")
		}
		res.header.Done = true
		res.header.Body = true
	}
	var bytes = []byte(content)
	res.WriteBytes(bytes)
	return res
}

func (res *Response) WriteBytes(bytes []byte) error {
	var errr error
	_, err := res.Ref.Write(bytes)
	if err != nil {
		errr = err
	}
	return errr
}

func (res *Response) sendContent(status int, contentType string, content []byte) {
	if res.header.BasicDone() {
		res.header.Status(status)
	}
	if res.header.CanSend() {
		res.header.Set("Content-Type", contentType)
		if Done := res.header.Flush(); !Done {
			log.Print("Failed to write headers")
			res.header.Done = true
			res.header.Body = true
			return
		}
	}
	err := res.WriteBytes(content)
	if err != nil {
		log.Panicf("Failed to flush the buffer, error: %v", err)
		return
	}

}

func (res *Response) Json(content interface{}) *Response {
	output, err := json.Marshal(content)
	if err != nil {
		res.sendContent(500, "application/json", []byte(""))
	} else {
		res.sendContent(200, "application/json", output)
	}
	return res
}
func (res *Response) Error(status int, str string) {
	res.sendContent(status, "text/html", []byte(str))
	log.Panic(str)
}

func (res *Response) Raw() http.ResponseWriter {
	return res.Ref
}
func (res *Response) Render(path string, data interface{}) *Response {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		log.Panicf("Given path was not found", err)
		res.header.Status(500)
		res.header.Flush()

	}
	var byt bytes.Buffer
	err = tmpl.Execute(&byt, data)
	if err != nil {
		log.Print("Template render failed ", err)
		res.header.Status(500)
		res.header.Flush()
	}
	res.WriteBytes(byt.Bytes())
	return res

}

func (res *Response) Redirect(url string) *Response {
	res.header.Status(302)
	res.header.Set("Location", url)
	res.header.Flush()
	res.ended = true
	return res
}

func (res *Response) Status(status int) *Response {
 res.header.Status(status)
 return res
}