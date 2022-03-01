package minima

/**
* Minima is a free and open source software under Mit license

Copyright (c) 2021 gominima

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

* Authors @apoorvcodes @megatank58
* Maintainers @Panquesito7 @savioxavier @Shubhaankar-Sharma @apoorvcodes @megatank58
* Thank you for showing interest in minima and for this beautiful community
*/

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoutes(t *testing.T) {
	tests := []struct {
		route     string
		path      string
		variables map[string]string
	}{
		{
			route:     "/",
			path:      "/",
			variables: map[string]string{},
		},
		{
			route:     "/test",
			path:      "/test",
			variables: map[string]string{},
		},
		{
			route: "/params/:one/:two",
			path:  "/params/one/two",
			variables: map[string]string{
				"one": "one",
				"two": "two",
			},
		},
		{
			route: "/params/:one/fixed/:two",
			path:  "/params/one/fixed/two",
			variables: map[string]string{
				"one": "one",
				"two": "two",
			},
		},
		{
			route: "/params/:one/fixed2/:two",
			path:  "/params/one/fixed2/two",
			variables: map[string]string{
				"one": "one",
				"two": "two",
			},
		},
		{
			route: "/params/:one/:middle/:two",
			path:  "/params/one/middle/two",
			variables: map[string]string{
				"one":    "one",
				"two":    "two",
				"middle": "middle",
			},
		},
	}

	var lastRoute string
	routes := NewRoutes()
	for _, test := range tests {
		test := test
		routes.Add(test.route, func(res *Response, req *Request) {
			lastRoute = test.route
		})
	}

	for _, test := range tests {
		t.Run(test.route, func(t *testing.T) {
			f, params, ok := routes.Get(test.path)
			require.True(t, ok)
			require.Equal(t, test.variables, params)

			f(nil, nil)
			require.Equal(t, test.route, lastRoute)
		})
	}
}
