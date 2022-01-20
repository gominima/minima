package minima

import (
	"regexp"
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
	isMatch := false
	matchingRegex := m.Regex.FindAllStringSubmatch(path, -1)

	if isMatch = len(matchingRegex) != 0; isMatch {
		for i, param := range m.Params {
			routeParams[param] = matchingRegex[0][i+1]
		}
	}
	return isMatch, routeParams
}
