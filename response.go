package minima

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"text/template"
)

type Response struct {
	ref      http.ResponseWriter
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
	res.ref = rs
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

func (res *Response) Send(content string) *Response {
	var bytes = []byte(content)
	res.WriteBytes(bytes)
	return res
}

func (res *Response) WriteBytes(bytes []byte) error {
	var errr error
	_, err := res.ref.Write(bytes)
	if err != nil {
		errr = err
	}
	return errr
}

func (res *Response) sendContent(status int, contentType string, content []byte) {

	res.header.Status(status)

	if res.header.CanSend() {
		res.header.Set("Content-Type", contentType)
		if Done := res.header.FlushHeader(); !Done {
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
func (res *Response) Error(status int, str string) *Response {
	res.sendContent(status, "text/html", []byte(str))
	log.Panic(str)
	return res
}

func (res *Response) Raw() http.ResponseWriter {
	return res.ref
}
func (res *Response) Render(path string, data interface{}) *Response {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		log.Panic("Given path was not found", err)
		res.header.Status(500)

	}
	var byt bytes.Buffer
	err = tmpl.Execute(&byt, data)
	if err != nil {
		log.Print("Template render failed ", err)
		res.header.Status(500)
	}
	res.WriteBytes(byt.Bytes())
	return res

}

func (res *Response) Redirect(url string) *Response {
	res.header.Status(302)
	res.header.Set("Location", url)
	res.ended = true
	return res
}

func (res *Response) Status(status int) *Response {
	res.header.Status(status)
	return res
}

func (res *Response) FlushHeader() *Response {
	done := res.header.FlushHeader()
	if !done {
		panic("Failed to push headers")
	}
	return res
}
