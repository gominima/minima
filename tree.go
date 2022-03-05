package minima

import (
	"strings"
	"sync"
)

/**
 * @info the radix tree structure 
 * @property {Node} [roots] The roots of the tree
 * @property {int} [len] The lenght of the tree
 * @property {int} [size] The size of the tree
 * @property {bool} [safe] Whether the mutex is enabled or not
 * @property {byte} [placeholder] The regex byte for params
 * @property {byte} [delim] The regex byte for params
 * @property {sync.Mutex} [mu] The synx.Mutex instance
*/
type tree struct {
	root        *Node
	len         int
	size        int
	safe        bool
	placeholder byte
	delim       byte
	mu          *sync.Mutex
}

/**
 * @info Creates a new radix tree
 */
func NewTree() *tree {
	return &tree{
		root:        &Node{},
		len:         1,
		placeholder: ':',
		delim:       '/',
		mu:          &sync.Mutex{},
		safe:        true,
	}
}

/**
 * @info Inserts a new node in the tree
 * @param {string} [key] The route path used as key
 * @param {Handler} [handler] The handler to be used
 * @returns {}
*/
func (tr *tree) InsertNode(key string, handler Handler) {
	if key == "" || handler == nil {
		return
	}
	if tr.safe {
		defer tr.mu.Unlock()
		tr.mu.Lock()
	}
	n := tr.root

	for {
		var next *edge
		var slice string

		for _, edge := range n.edges {
			var found int
			slice = edge.key
			for i := range slice {
				if i < len(key) && slice[i] == key[i] {
					found++
					continue
				}
				break
			}
			if found > 0 {
				key = key[found:]
				slice = slice[found:]
				next = edge
				break
			}
		}
		if next != nil {
			n = next.n
			n.priority++

			if len(key) == 0 {
				if len(slice) == 0 {
					n.handler = handler
					return
				}
				next.key = next.key[:len(next.key)-len(slice)]
				c := n.clone()
				c.priority--
				n.edges = []*edge{
					&edge{
						key: slice,
						n:   c,
					},
				}
				n.handler = handler
				tr.len++
				return
			}
			if len(slice) > 0 {
				c := n.clone()
				c.priority--
				n.edges = []*edge{
					&edge{ // the suffix that is clone into a new node
						key: slice,
						n:   c,
					},
					&edge{ // the new node
						key: key,
						n: &Node{
							handler:  handler,
							depth:    n.depth + 1,
							priority: 1,
						},
					},
				}
				next.key = next.key[:len(next.key)-len(slice)]
				n.handler = nil
				tr.len += 2
				tr.size += len(key)
				return
			}
			continue
		}
		n.edges = append(n.edges, &edge{
			key: key,
			n: &Node{
				handler:  handler,
				depth:    n.depth + 1,
				priority: 1,
			},
		})
		tr.len++
		tr.size += len(key)
		return
	}
}

/**
 * @info Finds a specific node from the tree
 * @param {string} [key] The route path used as key
 * @returns {*Node, map[string]string}
*/
func (tr *tree) GetNode(key string) (*Node, map[string]string) {
	if key == "" {
		return nil, nil
	}
	if tr.safe {
		defer tr.mu.Unlock()
		tr.mu.Lock()
	}
	n := tr.root
	var params map[string]string
	for n != nil && key != "" {
		var next *edge
	Walk:
		for _, edge := range n.edges {
			slice := edge.key

			for {
				pindex := len(slice)
				if i := strings.IndexByte(slice, tr.placeholder); i >= 0 {
					pindex = i
				}
				prefix := slice[:pindex]
				if !strings.HasPrefix(key, prefix) {
					continue Walk
				}
				key = key[len(prefix):]

				if len(prefix) == len(slice) {
					next = edge
					break Walk
				}
				var delimint int
				slice = slice[pindex:]
				if delimint = strings.IndexByte(slice[1:], tr.delim) + 1; delimint <= 0 {
					delimint = len(slice)
				}
				k := slice[1:delimint]
				slice = slice[delimint:]
				if delimint = strings.IndexByte(key[1:], tr.delim) + 1; delimint <= 0 {
					delimint = len(key)
				}
				if params == nil {
					params = make(map[string]string)
				}
				params[k] = key[:delimint]
				key = key[delimint:]
				if slice == "" && key == "" {
					next = edge
					break Walk
				}

			}
		}
		if next != nil {
			n = next.n
			continue
		}
		n = nil
	}
	return n, params
}

/**
 * @info Turns a radix tree into a hash map
 * @param {tree} [tre] The tree to convert
 * @returns {map[string]Handler}
*/
func ToMap(tre *tree) map[string]Handler {
	ma := make(map[string]Handler, tre.len)
	for _, edge := range tre.root.edges {
		ma[edge.key] = edge.n.handler
	}
	return ma
}

/**
 * @info Inserts a hash map to the tree
 * @param {map[string]Handler} [m] The hash map to insert
 * @returns {}
*/
func (tr *tree) InsertMap(m map[string]Handler) {
	for i, v := range m {
		tr.InsertNode(i, v)
	}
}
