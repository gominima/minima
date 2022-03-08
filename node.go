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

