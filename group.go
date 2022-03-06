package minima

/**
 * @info The minima group structure
 * @property {[]cacheroute} [route] The array of cached routes
 * @property {[string} [prefix] The group prefix
 */
type Group struct {
	route  []*cacheRoute
	prefix string
}

/**
 * @info Creates a new minima group
 * @return {Group}
 */
func NewGroup(prefix string) *Group {
	return &Group{
		route:  make([]*cacheRoute, 0),
		prefix: prefix,
	}
}

func (g *Group) register(method string, path string, handler Handler) {
	g.route = append(g.route, &cacheRoute{
		method:  method,
		path:    g.prefix + path,
		handler: handler,
	})
}

/**
 * @info Adds route with Get method
 * @param {string} [path] The route path
 * @param {...Handler} [handler] The handler for the given route
 * @returns {*Router}
 */
func (g *Group) Get(path string, handler Handler) *Group {
	g.register("GET", path, handler)
	return g
}

/**
 * @info Adds route with Post method
 * @param {string} [path] The route path
 * @param {...Handler} [handler] The handler for the given route
 * @returns {*Group}
 */
func (g *Group) Post(path string, handler Handler) *Group {
	g.register("POST", path, handler)
	return g
}

/**
 * @info Adds route with Put method
 * @param {string} [path] The route path
 * @param {...Handler} [handler] The handler for the given route
 * @returns {*Group}
 */
func (g *Group) Put(path string, handler Handler) *Group {
	g.register("PUT", path, handler)
	return g
}

/**
 * @info Adds route with Patch method
 * @param {string} [path] The route path
 * @param {...Handler} [handler] The handler for the given route
 * @returns {*Group}
 */
func (g *Group) Patch(path string, handler Handler) {
	g.register("PATCH", path, handler)
}

/**
 * @info Adds route with Options method
 * @param {string} [path] The route path
 * @param {...Handler} [handler] The handler for the given route
 * @returns {*Group}
 */
func (g *Group) Options(path string, handler Handler) *Group {
	g.register("OPTIONS", path, handler)
	return g
}

/**
 * @info Adds route with Head method
 * @param {string} [path] The route path
 * @param {...Handler} [handler] The handler for the given route
 * @returns {*Group}
 */
func (g *Group) Head(path string, handler Handler) *Group {
	g.register("HEAD", path, handler)
	return g
}

/**
 * @info Adds route with Delete method
 * @param {string} [path] The route path
 * @param {...Handler} [handler] The handler for the given route
 * @returns {*Group}
 */
func (g *Group) Delete(path string, handler Handler) *Group {
	g.register("DELETE", path, handler)
	return g
}

/**
 * @info Returns all routes for the group
 * @return {[]cachRoute}
 */
func (g *Group) GetGroupRoutes() []*cacheRoute {
	return g.route
}
