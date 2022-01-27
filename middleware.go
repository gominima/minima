package minima

import "net/http"

type rawHandle func(rw http.ResponseWriter, r *http.Request)

/**
@info The Middleware structure
@property {Handler} [handler] The handler to be used
*/
type Middleware struct {
	handler    Handler
	israw      bool
	rawHandler rawHandle
}

/**
@info The Plugins structure
@property {Middleware} [plugin] The plugin
*/
type Plugins struct {
	plugin []*Middleware
}

/**
@info Initialise the plugins interface
*/
func use() *Plugins {
	p := &Plugins{}
	return p
}

/**
@info Add a plugin
@param {Handler} [handler] The handler to add
*/
func (p *Plugins) AddPlugin(handler Handler) {
	middleware := &Middleware{handler: handler, israw: false}
	p.plugin = append(p.plugin, middleware)
}

func (p *Plugins) AddRawPlugin(handler rawHandle) {
	middleware := &Middleware{rawHandler: handler, israw: true}
	p.plugin = append(p.plugin, middleware)
}

/**
@info Serve a plugin
@param {Response} [res] The response
@param {Request} [req] The request
*/
func (p *Plugins) ServePlugin(res *Response, req *Request) {
	for _, v := range p.plugin {
		if v.israw {
			v.rawHandler(res.Raw(), req.Raw())
		} else {
			v.handler(res, req)
		}
	}
}
