package minima

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testServer *httptest.Server

func setupMinima() *minima {
	app := New()
	app.Get("/", func(res *Response, req *Request) {
		res.OK().Send("Hello World")
	})
	return app
}

func setup() {
	minimaApp := setupMinima()
	testServer = httptest.NewServer(minimaApp)
}

func shutdown() {
	testServer.Close()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func TestHelloWorld(t *testing.T) {
	resp, err := http.Get(testServer.URL)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, "Hello World", string(body))
}
