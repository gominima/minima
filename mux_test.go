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
	}

	var result string
	routes := NewRoutes()
	for _, test := range tests {
		test := test
		routes.Add(test.route, func(res *Response, req *Request) {
			require.Equal(t, test.variables, req.Params)
			result = test.route
		})
	}

	for _, test := range tests {
		t.Run(test.route, func(t *testing.T) {
			ok := routes.Run(test.path, nil, &Request{})
			require.True(t, ok)
			require.Equal(t, test.route, result)
		})
	}
}
