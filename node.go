package minima


/**
 * @info The tree edge structure
 * @property {string} [key] The key of the edge
 * @property {Node} [n] The tree node
*/
type edge struct {
	key string
	n   *Node
}


/**
 * @info The tree Node structure
 * @property {Handler} [handler] The handler to be used
 * @property {[]*edge} [edges] The array of node edges
 * @property {int} [priority] The priority of the node in the tree
 * @property {int} [depth] The depth of the node in the tree
*/
type Node struct {
	handler  Handler
	edges    []*edge
	priority int
	depth    int
}

/**
 * @info Whether the node is a leaf or not
 * @returns {bool}
*/
func (n *Node) IsLeaf() bool {
	length := len(n.edges)
	if length == 2 {
		return n.edges[0] == nil && n.edges[1] == nil
	}
	return length == 0
}

/**
 * @info Clones the current node
 * @returns {*Node}
*/
func (n *Node) clone() *Node {
	c := *n
	c.incrDepth()
	return &c
}


/**
 * @info Increases node's depth in the tree
*/
func (n *Node) incrDepth() {
	n.depth++
	for _, e := range n.edges {
		e.n.incrDepth()
	}
}
