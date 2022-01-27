package minima

import "net/http"

type rawHandle func(rw http.ResponseWriter, r *http.Request)

/**
@info The Middleware structure
@property {Handler} [handler] The handler to be used
@property {bool} [israw] Whether the handler is raw net/http or not
@property {rawHandle} [rawHandler] The raw handler to be used
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
	return &Plugins{}
}

/**
@info Add a plugin
@param {Handler} [handler] The handler to add
*/
func (p *Plugins) AddPlugin(handler Handler) {
	p.plugin = append(p.plugin, &Middleware{handler: handler, israw: false})
}

/**
@info Add a raw net/http plugin
@param {rawHandle} [handler] The raw handler to add
*/
func (p *Plugins) AddRawPlugin(handler rawHandle) {
	p.plugin = append(p.plugin, &Middleware{rawHandler: handler, israw: true})
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
