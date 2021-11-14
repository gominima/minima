package fiable

import (
	"net/http"
	"regexp"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}
