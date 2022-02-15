package minima

import (
	"strings"
)

/**
@info The Param structure
@property {string} [name] The name of the param
@property {bool} [fixed] Whether the param is fixed or not
*/
type param struct {
	name  string
	fixed bool
}

/**
@info The Route structure
@property {string} [prefix] The prefix of the route
@property {[]param} [partnames] The route paths split into parts
@property {Handler} [function] The handler to be used
*/
type Route struct {
	prefix    string
	partNames []param
	function  Handler
}

/**
@info The Routes root structure
@property {map[string][]Route} [roots] The map of array routes
*/
type Routes struct {
	roots map[string][]Route
}

/**
@info Makes a new Routes instance
@returns {*Routes}
*/
func NewRoutes() *Routes {
	return &Routes{
		roots: make(map[string][]Route),
	}
}

/**
@info Adds a new route to the routes table
@param {string} [path] Path of the route
@param {Handler} [handler] Handler of the route
*/
func (r *Routes) Add(p string, f Handler) {
	var path string
	path = p
	if p[len(p)-1:] == "/" && p != "/" {
		path = strings.TrimSuffix(p, p[len(p)-1:])
	}
	parts := strings.Split(path, "/")
	var rootParts []string
	var varParts []param
	var paramsFound bool
	for _, p := range parts {
		if strings.HasPrefix(p, ":") {
			paramsFound = true
		}

		if paramsFound {
			if strings.HasPrefix(p, ":") {
				varParts = append(varParts, param{
					name:  strings.TrimPrefix(p, ":"),
					fixed: false,
				})
			} else {
				varParts = append(varParts, param{
					name:  p,
					fixed: true,
				})
			}
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

/**
@info Gets handler and params from the routes table
@param {string} [path] Path of the route to find
@returns {Handler, map[string]string, bool}
*/
func (r *Routes) Get(path string) (Handler, map[string]string, bool) {
	var routes []Route
	remaining := path
	for {
		var ok bool
		routes, ok = r.roots[remaining]
		if ok {
			return matchRoutes(path, routes)

		}

		if len(remaining) < 2 {
			return nil, nil, false
		}

		index := strings.LastIndex(remaining, "/")
		if index < 0 {
			return nil, nil, false
		}

		if index > 0 {
			remaining = remaining[:index]
		} else {
			remaining = "/"
		}
	}
}

/**
@info Matches routes to the request
@param {string} [path] Path of the request route to find
@param {[]Route} [routes] The array of routes to match
@returns {Handler, map[string]string, bool}
*/
func matchRoutes(path string, routes []Route) (Handler, map[string]string, bool) {
outer:
	for _, r := range routes {
		params := strings.Split(
			strings.TrimPrefix(
				strings.TrimPrefix(path, r.prefix),
				"/"),
			"/")
		valid := cleanArray(params)

		if len(valid) == len(r.partNames) {
			paramNames := make(map[string]string)
			for i, p := range r.partNames {
				if p.fixed {
					if params[i] != p.name {
						continue outer
					} else {
						continue
					}
				}
				paramNames[p.name] = params[i]
			}
			return r.function, paramNames, true
		}
	}
	return nil, nil, false
}


/**
@info Cleans the array and finds non nill values
@param {string} [path] The array of string to slice and clean
@returns {[]string}
*/
func cleanArray(a []string) []string {
	var valid []string
	for _, s := range a {
		if s != "" {
			valid = append(valid, s)
		}
	}
	return valid
}
