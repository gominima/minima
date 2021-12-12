package fiable


type Middleware struct {
	handler Handler
}

type Plugins struct {
	plugin []*Middleware
}

func use() *Plugins{
 p := &Plugins{}
 return p
}
func (p *Plugins) AddPlugin(handler Handler){
	middleware := &Middleware{handler: handler}
	p.plugin = append(p.plugin, middleware)
}

func (p *Plugins) ServePlugin(res *Response, req *Request) {
	for _, v := range p.plugin {
	 v.handler(res,req)
	}
}