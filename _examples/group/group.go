package group

import "github.com/gominima/minima"

func RouteGroup() *minima.Group {
	grp := minima.NewGroup("/v1")
	grp.Get("/auth", func(res *minima.Response, req *minima.Request) {
		res.Send("auth")
	})
	return grp
}