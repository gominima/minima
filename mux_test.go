package minima

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
