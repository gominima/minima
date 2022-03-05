package minima


const tabSize = 4

type edge struct {
	key string
	n     *Node
}

type Node struct {
	handler Handler
	edges []*edge
	priorty int
	depth int
}

func (n *Node) IsLeaf() bool {
	length := len(n.edges)
	if length == 2 {
		return n.edges[0] == nil && n.edges[1] == nil
	}
	return length == 0
}

func bit(i uint8, c byte) uint8 {
	if 1<<(i-1)&c > 0 {
		return 1
	}
	return 0
}

func (n*Node) insert(key string, handler Handler)  {
	for i := range key {
		for j := uint8(8); j > 0; j-- {
			bbit := bit(j, key[i])
			done := i == len(key)-1 && j == 1
			if e := n.edges[bbit]; e != nil {
				if done {
					e.n.handler = handler
					return 
				}
				goto next
			}
			n.edges[bbit] = &edge{
				n: &Node{
					depth: n.depth + 1,
					edges: make([]*edge, 2),
				},
			}
			if done {
				n.edges[bbit].n.handler = handler
			}
			
		next:
			n = n.edges[bbit].n
		}
	}
}

func (n *Node) delBinary(key string) int {
	var (
		ref *edge
		del int
	)
	for i := range key {
		for j := uint8(8); j > 0; j-- {
			bbit := bit(j, key[i])
			done := i == len(key)-1 && j == 1
			if e := n.edges[bbit]; e != nil {
				del++
				if done && e.n.IsLeaf() { // only delete if node is leaf, otherwise it would break the tree
					ref.n.edges = make([]*edge, 2) // reset edges from the last node that has value
					return del
				}
				ref = e
				n = e.n
				continue
			}
			return 0
		}
	}
	return 0
}

func (n *Node) getBinary(key string) *Node {
	for i := range key {
		for j := uint8(8); j > 0; j-- {
			bbit := bit(j, key[i])
			done := i == len(key)-1 && j == 1
			if e := n.edges[bbit]; e != nil {
				if done {
					return e.n
				}
				n = e.n
				continue
			}
			return nil
		}
	}
	return nil
}