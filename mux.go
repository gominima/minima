package minima

import (
	"regexp"
	"strings"
)

/**
@info The mux structure
@property {string} [Path] The path to match
@property {[]string} [Params] The params to use
@property {regexp.Regexp} [Regex] The regex to use
@property {[]Handler} [Handlers] The handlers to use
*/
type mux struct {
	Path     string
	Params   []string
	Regex    *regexp.Regexp
	Handlers []Handler
}

/**
@info Match a path
@param {string} [path] The path to match
@returns {bool, map[]string[]string}
*/
func (m *mux) matchingPath(path string) (bool, map[string]string) {
	routeParams := make(map[string]string)
	matchingRegex := m.Regex.FindAllStringSubmatch(path, -1)
	isMatch := len(matchingRegex) != 0

	if isMatch {
		for i, param := range m.Params {
			routeParams[param] = matchingRegex[0][i+1]
		}
	}

	return isMatch, routeParams
}

type Route struct {
	prefix    string
	partNames []string
	function  Handler
}

type Routes struct {
	roots map[string][]Route
}

func NewRoutes() *Routes {
	return &Routes{
		roots: make(map[string][]Route),
	}
}

func (r *Routes) Add(path string, f Handler) {
	parts := strings.Split(path, "/")
	var rootParts []string
	var varParts []string
	for _, p := range parts {
		if strings.HasPrefix(p, ":") {
			varParts = append(varParts, strings.TrimPrefix(p, ":"))
		} else {
			rootParts = append(rootParts, p)
		}
	}

	root := strings.Join(rootParts, "/")

	r.roots[root] = append(r.roots[root], Route{
		prefix:    root,
		partNames: varParts,
		function:  f,
	})
}

func (r *Routes) Run(path string, res *Response, req *Request) bool {
	var routes []Route
	remaining := path
	for {
		var ok bool
		routes, ok = r.roots[remaining]
		if ok {
			return matchRoutes(path, routes, res, req)

		}

		if len(remaining) < 2 {
			return false
		}

		index := strings.LastIndex(remaining, "/")
		if index < 0 {
			return false
		}

		if index > 0 {
			remaining = remaining[:index]
		} else {
			remaining = "/"
		}
	}
}

func matchRoutes(
	path string,
	routes []Route,
	res *Response,
	req *Request,
) bool {
	for _, r := range routes {
		params := strings.Split(
			strings.TrimPrefix(
				strings.TrimPrefix(path, r.prefix),
				"/"),
			"/")
		valid := cleanArray(params)

		if len(valid) == len(r.partNames) {
			paramNames := make(map[string]string)
			for i, n := range r.partNames {
				paramNames[n] = params[i]
			}
			req.Params = paramNames
			r.function(res, req)
			return true
		}
	}
	return false
}

func cleanArray(a []string) []string {
	var valid []string
	for _, s := range a {
		if s != "" {
			valid = append(valid, s)
		}
	}
	return valid
}
